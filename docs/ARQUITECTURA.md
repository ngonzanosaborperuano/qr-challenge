# Arquitectura del proyecto (QR Challenge)

Este documento describe la **arquitectura actual** del repositorio, el **porqué** de las decisiones tomadas para el challenge y **qué cambiaría** (opcionalmente) si el objetivo fuera evolucionarlo a un sistema más “productivo” sin perder simplicidad.

## Contexto y objetivo

El challenge pide (implícita o explícitamente) un flujo distribuido:

- **Go API**: recibe una matriz, valida, rota 90°, calcula factorización QR, y **coordina** la llamada a Node.
- **Node API**: recibe matrices derivadas (Q, R y/o rotada) y devuelve **estadísticas** (max/min/avg/sum) y verificación **diagonal**.
- **Frontend Angular**: interfaz simple para autenticar (JWT) y consumir el flujo.

## Vista general (alto nivel)

```
Cliente (Angular / curl / Postman)
        |
        | HTTP/REST + JWT
        v
Go API (Fiber)
  - valida, rota, calcula QR
  - llama a Node API para stats
        |
        | HTTP/REST + JWT (reuso de token)
        v
Node API (Express)
  - calcula stats y anyDiagonal
        |
        v
Go API retorna respuesta agregada al cliente
```

## Arquitectura actual por servicio

### 1) Go API (`go-api/`)

**Estilo predominante:** arquitectura en capas (Layered) + separación por paquetes (`controllers`, `handlers`, `services`, `middleware`, `models`).

**Estructura:**

- **Entrada (HTTP / delivery)**
  - `cmd/server/main.go`: composición del servidor, rutas, middlewares.
  - `internal/handlers/matrix_handler.go`: endpoint principal `/matrix/process`.
  - `internal/controllers/auth_controller.go`: endpoint `/auth/login`.
- **Dominio / lógica de negocio**
  - `internal/services/validator.go`: validación de matriz (rectangular/no vacía).
  - `internal/services/rotation.go`: rotación 90°.
  - `internal/services/qr_decomposition.go`: QR decomposition (vía `gonum`).
- **Salida (integración con otro servicio)**
  - `internal/services/node_client.go`: cliente HTTP hacia Node API.
- **Cross-cutting**
  - `internal/middleware/auth.go`: JWT middleware.
  - `internal/models/*`: contratos de request/response.

**Nota importante (limitación actual):** el handler `matrix_handler.go` hace tanto de **adaptador de entrada** (HTTP) como de **caso de uso** (orquestación del flujo). Para el challenge esto es totalmente aceptable; para un producto, convendría separar esa orquestación en una capa de “application/usecases”.

### 2) Node API (`node-api/`)

**Estilo predominante:** arquitectura en capas (Layered) típica en Express.

**Estructura:**

- **Entrada**
  - `src/index.js`: servidor, middlewares, rutas.
  - `src/routers/matrixRouter.js`: ruteo del feature.
  - `src/controllers/matrixController.js`: valida el request y delega.
- **Dominio / lógica**
  - `src/services/statsService.js`: cálculo de estadísticas y `isDiagonal`.
- **Cross-cutting**
  - `src/middleware/auth.js`: JWT middleware.

### 3) Frontend Angular (`frontend/`)

**Estilo predominante:** arquitectura basada en componentes (Component-based) + separación por responsabilidades.

**Estructura:**

- `src/app/components/`: UI (login, procesador de matrices, etc.)
- `src/app/services/`: comunicación HTTP y auth (`api.service.ts`, `auth.service.ts`)
- `src/app/models/`: contratos TypeScript
- `src/app/pipes/`: formateos/reutilización de vista
- `src/environments/`: configuración de URLs por entorno

## Por qué se eligió esta arquitectura (para el propósito del challenge)

- **Claridad y rapidez**: la separación por carpetas (handlers/controllers/services) hace el flujo fácil de seguir y explicar.
- **Separación de responsabilidades suficiente**: cada API tiene un objetivo acotado (Go “procesa y orquesta”, Node “calcula stats”).
- **Testabilidad razonable**: las funciones de `services` y `statsService` se testean sin necesidad de UI.
- **Despliegue simple**: con Docker/Compose, cada componente se levanta como servicio independiente.
- **Enfoque en el objetivo**: el challenge prioriza cumplir requerimientos con código entendible más que implementar una arquitectura “enterprise”.

## Relación con Hexagonal (Ports & Adapters) / Clean Architecture

En el repositorio, los conceptos de Hexagonal/Clean están **parcialmente presentes**, pero no formalizados:

- **Adapters de entrada (inbound):**
  - Go: `handlers/` y `controllers/`
  - Node: `controllers/` + `routers/`
- **Adapters de salida (outbound):**
  - Go: `services/node_client.go` (cliente HTTP a Node)
- **Dominio:**
  - Go: `services/validator.go`, `rotation.go`, `qr_decomposition.go`
  - Node: `services/statsService.js`

**Qué falta para decir “está aplicada Hexagonal/Clean” de forma explícita:**

- **Casos de uso (application layer)**
  - Ej.: `ProcessMatrix` como unidad de orquestación fuera del handler HTTP.
- **Puertos (interfaces)**
  - Ej.: `StatsClient` (interfaz) para que el caso de uso dependa de abstracciones.
- **Inversión de dependencias**
  - El handler debería llamar a un use case, y el use case a un puerto; el adaptador HTTP a Node implementa ese puerto.

## ¿Se debería cambiar la arquitectura?

### Para el challenge (propósito actual)
**No es necesario cambiarla.** La arquitectura actual es correcta: es simple, testeable y alineada al alcance.

### Para un contexto más “productivo” (mejoras recomendadas, opcionales)

**Recomendación principal (Go API):** introducir una capa `usecases/` + puertos (interfaces).

Ejemplo conceptual (sin imponer nombres exactos):

- `internal/usecases/process_matrix.go`
  - define `type StatsClient interface { GetMatrixStats(...) (...) }`
  - implementa `ProcessMatrix` orquestando:
    - Validate → Rotate → QR → StatsClient.GetMatrixStats
- `internal/adapters/http/node_client.go`
  - implementa `StatsClient` usando HTTP
- `internal/adapters/http/fiber/matrix_handler.go`
  - parsea request, llama al use case, devuelve response

**Beneficios reales de ese cambio:**

- tests del caso de uso sin Fiber ni HTTP real
- menos acoplamiento a detalles (framework y transporte)
- evolución más sencilla (si Node cambia a gRPC/cola, solo cambia el adaptador)

**Qué NO haría (para no sobre-ingenierizar):**

- No introducir CQRS, DDD “pesado”, ni múltiples bounded contexts.
- No meter una capa extra si no hay cambios futuros esperados.

## Decisión final (resumen)

- **Arquitectura actual**: óptima para el challenge por simplicidad y claridad.
- **Cambio recomendado si el sistema crece**: mover la orquestación del flujo a una capa de **use cases** y definir **ports** para integraciones (Hexagonal/Clean explícita), especialmente en Go.

