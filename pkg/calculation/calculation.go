package calculation

import (
	"strconv"
	"strings"
)

func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func parseNumber(input string) (float64, error) {
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		return 0, ErrInvalidExpression
	}
	return num, nil
}

func Calc(expression string) (float64, error) {
	// Удаляем все пробелы
	expression = strings.ReplaceAll(expression, " ", "")
	if len(expression) == 0 {
		return 0, ErrInvalidExpression
	}

	for _, char := range expression {
		if !isOperator(char) && (char < '0' || char > '9') && char != '(' && char != ')' {
			return 0, ErrInvalidExpression
		}
	}

	return evaluate(expression)
}

func evaluate(expression string) (float64, error) {
	for strings.Contains(expression, "(") {
		start := strings.LastIndex(expression, "(")
		end := strings.Index(expression[start:], ")") + start
		if end <= start {
			return 0, ErrInvalidExpression
		}
		subExpression := expression[start+1 : end]
		result, err := evaluate(subExpression)
		if err != nil {
			return 0, err
		}
		expression = strings.Replace(expression, "("+subExpression+")", strconv.FormatFloat(result, 'f', -1, 64), 1)
	}

	for strings.ContainsAny(expression, "*/") {
		expression = processOperator(expression, "*/")
	}

	for strings.ContainsAny(expression, "+-") {
		expression = processOperator(expression, "+-")
	}

	return parseNumber(expression)
}

func processOperator(expression string, operators string) string {
	for i, char := range expression {
		if strings.ContainsRune(operators, char) {
			left, right := findOperands(expression, i)
			leftNum, _ := parseNumber(left)
			rightNum, _ := parseNumber(right)
			var result float64

			switch char {
			case '*':
				result = leftNum * rightNum
			case '/':
				if rightNum == 0 {
					return "division_by_zero"
				}
				result = leftNum / rightNum
			case '+':
				result = leftNum + rightNum
			case '-':
				result = leftNum - rightNum
			}

			expression = strings.Replace(expression, left+string(char)+right, strconv.FormatFloat(result, 'f', -1, 64), 1)
			break
		}
	}
	return expression
}

func findOperands(expression string, operatorIndex int) (string, string) {
	left := ""
	right := ""
	i := operatorIndex - 1

	for i >= 0 && (expression[i] >= '0' && expression[i] <= '9' || expression[i] == '.') {
		left = string(expression[i]) + left
		i--
	}

	i = operatorIndex + 1
	for i < len(expression) && (expression[i] >= '0' && expression[i] <= '9' || expression[i] == '.') {
		right += string(expression[i])
		i++
	}

	return left, right
}
