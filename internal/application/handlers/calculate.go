package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AlexDillz/Calc_server_yandex/pkg/calculation"
)

type CalcRequest struct {
	Expression string `json:"expression"`
}

type CalcResponse struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req CalcRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(CalcResponse{Error: "Invalid request format"})
		return
	}

	result, err := calculation.Calc(req.Expression)
	if err != nil {
		if err == calculation.ErrInvalidExpression {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(CalcResponse{Error: "Expression is not valid"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(CalcResponse{Error: "Internal server error"})
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CalcResponse{Result: result})
}
