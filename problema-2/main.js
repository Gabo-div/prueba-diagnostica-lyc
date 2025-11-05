const readline = require('readline');
const fs = require('fs');
const { performance } = require('perf_hooks');

function binomial(n, k) {
  if (k < 0n || k > n) {
    return 0n;
  }

  if (k > n / 2n) {
    k = n - k;
  }

  let res = 1n;

  for (let i = 0n; i < k; i++) {
    res = res * (n - i) / (i + 1n);
  }

  return res;
}

function generatePascalTriangle(n) {
  if (n < 0) {
    return [];
  }
  const coefficients = new Array(n + 1);
  for (let i = 0; i <= n; i++) {
    coefficients[i] = binomial(BigInt(n), BigInt(i));
  }
  return coefficients;
}

function formatPolynomial(coefficients) {
  const n = coefficients.length - 1;
  if (n < 0) return "0";

  const parts = [];
  for (let i = 0; i <= n; i++) {
    const coeff = coefficients[i];
    const power = n - i;

    if (coeff === 0n) {
      continue;
    }

    let term = '';
    // Coeficiente: No mostrar si es 1 y la potencia es > 0.
    if (coeff !== 1n || power === 0) {
      term += coeff.toString();
    }

    // Variable x y su potencia
    if (power > 0) {
      term += 'x';
      if (power > 1) {
        term += '^' + power;
      }
    }
    parts.push(term);
  }
  return parts.join(' + ');
}

/**
 * Muestra el cálculo paso a paso de la evaluación del polinomio.
 * @param {bigint[]} coefficients - Coeficientes del polinomio.
 * @param {number} x - El valor con el que evaluar.
 */
function evaluatePolynomial(coefficients, x) {
  const n = coefficients.length - 1;
  const xBig = BigInt(x);
  let total = 0n;

  const calculationSteps = [];
  const resultSteps = [];

  console.log(`\nCalculando para f(${x}):`);

  // Calcular cada término y construir las cadenas de texto
  for (let i = 0; i <= n; i++) {
    const coeff = coefficients[i];
    const power = n - i;

    const termVal = xBig ** BigInt(power);
    const stepResult = coeff * termVal;
    total += stepResult;

    calculationSteps.push(`${coeff.toString()}*(${x}^${power})`);
    resultSteps.push(stepResult.toString());
  }

  console.log('f(x) =', calculationSteps.join(' + '));
  console.log('     =', resultSteps.join(' + '));
  console.log('     =', total.toString());
}

/**
 * Función principal que maneja la interacción con el usuario.
 */
function main() {
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
  });

  rl.question('Ingrese el valor de n (entero no negativo): ', (nStr) => {
    const n = parseInt(nStr, 10);
    if (isNaN(n) || n < 0 || !Number.isInteger(n)) {
      console.log('Entrada inválida para n. Debe ser un número entero no negativo.');
      rl.close();
      return;
    }

    const startTime = performance.now();
    const coefficients = generatePascalTriangle(n);
    const endTime = performance.now();
    const duration = (endTime - startTime);

    const polynomial = formatPolynomial(coefficients);
    console.log(`\nEl polinomio (x+1)^${n} es:`);
    console.log(polynomial);

    rl.question('\nIngrese el valor de x para evaluar el polinomio: ', (xStr) => {
      const x = parseInt(xStr, 10);
      if (isNaN(x) || !Number.isInteger(x)) {
        console.log('Entrada inválida para x. Debe ser un número entero.');
        rl.close();
        return;
      }

      evaluatePolynomial(coefficients, x);

      console.log(`\n--- Medición de Tiempo para n=${n} ---`);
      console.log(`Tiempo de ejecución en JS: ${duration.toFixed(4)} ms`);

      const fileName = 'resultados.txt';
      const fileContent = `JS: n=${n}, x=${x}, tiempo: ${duration.toFixed(4)}ms\n`;

      fs.appendFile(fileName, fileContent, { encoding: 'utf8' }, (err) => {
        if (err) {
          console.error(`Error al abrir o escribir en el archivo ${fileName}:`, err);
        } else {
          console.log(`Resultado de la medición guardado en '${fileName}'`);
        }
        rl.close();
      });
    });
  });
}

main();
