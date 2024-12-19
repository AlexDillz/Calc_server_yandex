# Calc_server_yandex

## Веб-сервис для вычисления арифметических выражений

## Возможности

- Вычисление выражений, отправленных через HTTP
- Обработка ошибок (деление на ноль, неправильно введенные выражения)
- Поддержка арифметических операций (вычитание, умножение, деление, сложение)

My Calc Project — это библиотека на Go, которая позволяет выполнять арифметические вычисления с поддержкой операций сложения, вычитания, умножения, деления и приоритета операций (включая использование скобок)

## Установка

Склонируйте репозиторий:

git clone https://github.com/AlexDillz/Calc_server_yandex.git  

Перейдите в папку проекта:

cd Calc_server_yandex

Инициализируйте модуль Go:

go mod init github.com/AlexDillz//Calc_server_yandex

Зависимости:

go mod tidy

## Пример использования:

package main

import (
	"fmt"
	"log"

	"github.com/AlexDillz/Calc_server_yandex/pkg/calculation"
)

func main() {
	expression := "2 + (3 * 4) - 5 / 2"
	result, err := calculation.Calc(expression)
	if err != nil {
		log.Fatalf("Ошибка при вычислении: %v", err)
	}
	fmt.Printf("Результат выражения '%s': %f\n", expression, result)
}


## Использование curl:

curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'

Response:

{
  "result": 6
}

## Ошибка 422:

curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "1+1*"
}'

Response:

{
  "error": "Expression is not valid"
}

## Ошибка 500:

{
  "error": "Internal server error"
}


## Для запуска тестов:

go test ./...

## Структура проекта:
pkg/calculation/ — основная библиотека
calculation.go — функции для вычисления выражений
errors.go — определение ошибок
pkg/calculation/calculation_test.go — модульные тесты для библиотеки
