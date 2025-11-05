package main

import (
	"fmt"
	"os"
	"regexp"
)

var keywords = map[string]string{
	"auto":     "automatico",
	"break":    "romper",
	"case":     "caso",
	"char":     "caracter",
	"const":    "constante",
	"continue": "continuar",
	"default":  "defecto",
	"do":       "hacer",
	"double":   "doble",
	"else":     "si no",
	"enum":     "enumeracion",
	"extern":   "externo", "float": "flotante",
	"for":      "para",
	"goto":     "ir a",
	"if":       "si",
	"int":      "entero",
	"long":     "largo",
	"register": "registro",
	"return":   "retornar",
	"short":    "corto",
	"signed":   "con signo",
	"sizeof":   "tamano de",
	"static":   "estatico",
	"struct":   "estructura",
	"switch":   "cambiar",
	"typedef":  "definir tipo",
	"union":    "union",
	"unsigned": "sin signo",
	"void":     "vacio",
	"volatile": "volatil",
	"while":    "mientras",
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Uso: %s <archivo.c>\n", os.Args[0])
		os.Exit(1)
	}

	filePath := os.Args[1]

	content, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al leer el archivo: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Contenido del archivo:\n%s\n", content)
	fmt.Println("\nPalabras reservadas encontradas y sus traducciones:")

	re := regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`)
	words := re.FindAllString(string(content), -1)

	foundKeywords := make(map[string]bool)

	for _, word := range words {
		translation, isKeyword := keywords[word]
		if isKeyword {
			if _, found := foundKeywords[word]; !found {
				fmt.Printf("%s -> %s\n", word, translation)
				foundKeywords[word] = true
			}
		}
	}
}

