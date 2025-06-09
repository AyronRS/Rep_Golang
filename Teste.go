package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operacao string

	fmt.Println("== Calculadora em Go ==")
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite a operação (+, -, *, /): ")
	fmt.Scanln(&operacao)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	switch operacao {
	case "+":
		fmt.Printf("Resultado: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Resultado: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Resultado: %.2f\n", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Erro: divisão por zero!")
		} else {
			fmt.Printf("Resultado: %.2f\n", num1/num2)
		}
	default:
		fmt.Println("Operação inválida.")
	}
}
