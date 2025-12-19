#!/bin/bash

echo "=== QR Challenge - Flujo Completo ==="
echo ""

# Paso 1: Verificar servicios
echo "1. Verificando servicios..."
curl -s http://localhost:3000/health | jq '.'
curl -s http://localhost:3001/health | jq '.'
echo ""

# Paso 2: Obtener tokens
echo "2. Obteniendo tokens JWT..."
TOKEN_GO=$(curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin"}' \
  -s | jq -r '.token')

TOKEN_NODE=$(curl -X POST http://localhost:3001/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin"}' \
  -s | jq -r '.token')

echo "✅ Tokens obtenidos"
echo ""

# Paso 3: Procesar matriz
echo "3. Procesando matriz..."
RESPONSE=$(curl -X POST http://localhost:3000/matrix/process \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN_GO" \
  -d '{
    "matrix": [
      [1, 2, 3],
      [4, 5, 6],
      [7, 8, 9]
    ]
  }' \
  -s)

echo "$RESPONSE" | jq '.'
echo ""

# Paso 4: Mostrar estadísticas
echo "4. Estadísticas calculadas:"
echo "$RESPONSE" | jq '.nodeStats'
echo ""

# Paso 5: Verificar matriz rotada
echo "5. Matriz rotada:"
echo "$RESPONSE" | jq '.rotated'
echo ""

echo "✅ Flujo completo ejecutado exitosamente"