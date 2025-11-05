# Prueba de Compiladores 1

Este repositorio contiene las soluciones a cuatro problemas para la primera prueba de lenguajes y compiladores.

## Problema 1: Validador FEN

Este programa verifica si una cadena dada es una cadena FEN (Forsyth-Edwards Notation) válida.

### Cómo ejecutar

```bash
go run problema-1/main.go
```

## Problema 2: Expansión Polinomial

Este programa genera el polinomio `(x+1)^n` para un `n` dado, y lo evalúa para un `x` dado. Está implementado tanto en Go como en JavaScript, y compara sus tiempos de ejecución.

### Cómo ejecutar

Para ejecutar la versión en Go:

```bash
go run problema-2/main.go
```

Para ejecutar la versión en JavaScript:

```bash
node problema-2/main.js
```

El programa pedirá los valores de `n` y `x`. Los resultados del tiempo de ejecución se guardan en el archivo `resultados.txt`.

## Problema 3: Reconocimiento de Cadenas

Este programa utiliza expresiones regulares para identificar si una cadena dada es un literal de cadena, un número en notación científica, una dirección IP o una dirección de correo electrónico.

### Cómo ejecutar

```bash
go run problema-3/main.go
```

## Problema 4: Traductor de Palabras Clave de C

Este programa lee un archivo fuente en C, encuentra todas las palabras clave de C e imprime sus traducciones al español.

### Cómo ejecutar

```bash
go run problema-4/main.go ruta/a/tu/archivo.c
```

Por ejemplo:

```bash
go run problema-4/main.go problema-4/test.c
```

### Video explicativo

https://youtu.be/lrOsnC0cllg?si=OnZafce07Xu7Uw3A
