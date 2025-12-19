# Factorización QR - Explicación

## ¿Qué es la Factorización QR?

La **factorización QR** (también llamada descomposición QR) es una técnica fundamental en álgebra lineal que descompone una matriz **A** en el producto de dos matrices:

```
A = Q × R
```

Donde:
- **Q** es una matriz **ortogonal** (sus columnas son vectores ortonormales)
- **R** es una matriz **triangular superior** (todos los elementos debajo de la diagonal principal son 0)

---

## Propiedades de las Matrices Q y R

### Matriz Q (Ortogonal)
- Las columnas de Q son **vectores ortonormales** (perpendiculares entre sí y de longitud 1)
- Q^T × Q = I (matriz identidad)
- **Puede contener valores negativos** - esto es completamente normal
- Ejemplo: Q puede tener valores como `[-0.12, 0.90, 0.41]`

### Matriz R (Triangular Superior)
- Todos los elementos **debajo de la diagonal principal son 0** (o muy cercanos a 0 por errores de punto flotante)
- Los elementos en y arriba de la diagonal pueden ser cualquier número real
- **Puede contener valores negativos** - esto es completamente normal
- Ejemplo: R puede tener valores como `[-8.12, -9.60, -11.08]`

---

## ¿Por qué el Mínimo puede ser Negativo?

### Respuesta Corta
**Es completamente normal y matemáticamente correcto.** Las matrices Q y R pueden contener valores negativos como resultado natural de la factorización QR.

### Explicación Detallada

1. **La factorización QR no garantiza valores positivos:**
   - El algoritmo de factorización QR (usando métodos como Householder o Gram-Schmidt) puede producir valores negativos
   - Esto depende de la estructura y los valores de la matriz original

2. **Ejemplo práctico:**
   ```
   Matriz original: [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
   
   Matriz Q (puede tener negativos):
   [-0.12,  0.90,  0.41]
   [-0.49,  0.30, -0.82]
   [-0.86, -0.30,  0.41]
   
   Matriz R (puede tener negativos):
   [-8.12, -9.60, -11.08]
   [ 0.00,  0.90,   1.81]
   [ 0.00,  0.00,  -0.00]
   ```

3. **Las estadísticas se calculan sobre todas las matrices:**
   - Se combinan todos los valores de Q, R y la matriz rotada
   - Si alguna de estas matrices tiene valores negativos, el mínimo será negativo
   - En el ejemplo anterior, el mínimo es `-11.08` (viene de R)

---

## ¿Es Correcto que Haya Valores Negativos?

**✅ SÍ, es completamente correcto.**

### Verificaciones Matemáticas

1. **Q × R debe igualar la matriz original:**
   ```python
   # Verificación (con tolerancia para errores de punto flotante)
   Q × R ≈ A
   ```

2. **Q debe ser ortogonal:**
   ```python
   Q^T × Q = I  # Matriz identidad
   ```

3. **R debe ser triangular superior:**
   - Elementos debajo de la diagonal ≈ 0

Nuestros tests verifican estas propiedades:
- ✅ R es triangular superior (elementos debajo de diagonal ≈ 0)
- ✅ Q × R ≈ A (con tolerancia numérica)

---

## Aplicaciones de la Factorización QR

1. **Resolución de sistemas de ecuaciones lineales:**
   - Ax = b se convierte en Rx = Q^T × b (más fácil de resolver)

2. **Cálculo de valores propios (eigenvalues):**
   - Algoritmo QR iterativo

3. **Mínimos cuadrados:**
   - Resolución de problemas de regresión lineal

4. **Análisis numérico:**
   - Descomposición estable y numéricamente robusta

---

## Ejemplo Visual

### Matriz Original (A)
```
[1, 2, 3]
[4, 5, 6]
[7, 8, 9]
```

### Factorización QR: A = Q × R

**Matriz Q (Ortogonal):**
```
[-0.12,  0.90,  0.41]
[-0.49,  0.30, -0.82]
[-0.86, -0.30,  0.41]
```
*Nota: Valores negativos son normales*

**Matriz R (Triangular Superior):**
```
[-8.12, -9.60, -11.08]
[ 0.00,  0.90,   1.81]
[ 0.00,  0.00,  -0.00]
```
*Nota: Valores negativos son normales, elementos debajo de diagonal ≈ 0*

### Verificación
```
Q × R ≈ A  ✅
```

---

## Estadísticas Calculadas

Las estadísticas (máximo, mínimo, promedio, suma) se calculan sobre **todas las matrices**:
- Matriz Q (puede tener negativos)
- Matriz R (puede tener negativos)
- Matriz rotada (solo valores de la matriz original)

**Ejemplo:**
- Si R tiene `-11.08`, el mínimo será `-11.08`
- Si Q tiene `-0.86`, también se incluye en el cálculo
- El promedio puede ser positivo o negativo dependiendo de los valores

---

## Conclusión

1. ✅ **Valores negativos en Q y R son normales y correctos**
2. ✅ **La factorización QR puede producir valores negativos**
3. ✅ **El mínimo negativo en las estadísticas es esperado cuando hay valores negativos en Q o R**
4. ✅ **Nuestros tests verifican la correctitud matemática de la factorización**

---

## Referencias

- **Algoritmo usado:** Librería `gonum.org/v1/gonum` (método estándar de factorización QR)
- **Complejidad:** O(n³) para una matriz n×n
- **Precisión:** Usa tolerancia numérica (1e-10) para comparaciones de punto flotante

