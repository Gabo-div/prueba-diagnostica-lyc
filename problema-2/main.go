package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

func generatePascalTriangle(n int) []*big.Int {
	if n < 0 {
		return []*big.Int{}
	}

	coefficients := make([]*big.Int, n+1)
	for i := 0; i <= n; i++ {
		coefficients[i] = new(big.Int).Binomial(int64(n), int64(i))
	}
	return coefficients
}

func formatPolynomial(coefficients []*big.Int) string {
	var builder strings.Builder
	n := len(coefficients) - 1
	for i, coeff := range coefficients {
		power := n - i

		if coeff.Cmp(big.NewInt(0)) == 0 {
			continue
		}

		if builder.Len() > 0 {
			builder.WriteString(" + ")
		}

		if coeff.Cmp(big.NewInt(1)) != 0 || power == 0 {
			builder.WriteString(coeff.String())
		}

		if power > 0 {
			builder.WriteString("x")
			if power > 1 {
				builder.WriteString("^" + strconv.Itoa(power))
			}
		}
	}
	return builder.String()
}

func evaluatePolynomial(coefficients []*big.Int, x int64) {
	n := len(coefficients) - 1
	total := new(big.Int)
	var calculationSteps strings.Builder
	var resultSteps strings.Builder
	termResults := make([]*big.Int, n+1)

	fmt.Printf("Calculando para f(%d):\n", x)

	for i, coeff := range coefficients {
		power := n - i
		termVal := new(big.Int).Exp(big.NewInt(x), big.NewInt(int64(power)), nil)
		stepResult := new(big.Int).Mul(coeff, termVal)
		termResults[i] = stepResult
		total.Add(total, stepResult)

		if i > 0 {
			calculationSteps.WriteString(" + ")
		}
		calculationSteps.WriteString(fmt.Sprintf("%s*(%d^%d)", coeff.String(), x, power))
	}
	fmt.Println("f(x) =", calculationSteps.String())

	for i, res := range termResults {
		if i > 0 {
			resultSteps.WriteString(" + ")
		}
		resultSteps.WriteString(res.String())
	}
	fmt.Println("     =", resultSteps.String())
	fmt.Println("     =", total.String())
}

func main() {
	var n int
	var x int64

	fmt.Print("Ingrese el valor de n (entero no negativo): ")
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println("Entrada inválida para n. Debe ser un número entero.")
		return
	}

	if n < 0 {
		fmt.Println("n no puede ser negativo.")
		return
	}

	startTime := time.Now()
	coefficients := generatePascalTriangle(n)
	duration := time.Since(startTime)

	polynomial := formatPolynomial(coefficients)
	fmt.Printf("\nEl polinomio (x+1)^%d es:\n", n)
	fmt.Println(polynomial)

	fmt.Print("\nIngrese el valor de x para evaluar el polinomio: ")
	if _, err := fmt.Scanln(&x); err != nil {
		fmt.Println("Entrada inválida para x. Debe ser un número entero.")
		return
	}
	evaluatePolynomial(coefficients, x)

	fmt.Printf("\n--- Medición de Tiempo para n=%d ---\n", n)
	fmt.Printf("Tiempo de ejecución en Go: %s\n", duration)

	fileName := "resultados.txt"
	fileContent := fmt.Sprintf("Go: n=%d, x=%d, tiempo: %s\n", n, x, duration)

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error al abrir el archivo %s: %v\n", fileName, err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(fileContent); err != nil {
		fmt.Printf("Error al escribir en el archivo %s: %v\n", fileName, err)
	} else {
		fmt.Printf("Resultado de la medición guardado en '%s'\n", fileName)
	}
}
