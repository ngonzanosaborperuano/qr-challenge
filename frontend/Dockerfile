# Dockerfile para Frontend Angular
# Etapa 1: Builder - Compilar la aplicación
FROM node:20-alpine AS builder

WORKDIR /app

# Copiar archivos de dependencias
COPY package*.json ./
RUN npm ci

# Copiar código fuente
COPY . .

# Compilar la aplicación
RUN npm run build

# Normalizar salida de Angular para Nginx:
# Angular 17+ suele generar en dist/browser (o dist/<app>/browser).
# Construimos /app/out con index.html en raíz para que el COPY final siempre reemplace
# el index.html default de Nginx (evita el "Welcome to nginx!").
RUN set -eux; \
  rm -rf /app/out; mkdir -p /app/out; \
  if [ -d /app/dist/browser ]; then \
    cp -R /app/dist/browser/. /app/out/; \
  elif [ -f /app/dist/index.html ]; then \
    cp -R /app/dist/. /app/out/; \
  else \
    idx="$(find /app/dist -maxdepth 6 -type f -name index.html | head -n 1)"; \
    if [ -z "$idx" ]; then echo "No se encontró index.html en /app/dist"; exit 1; fi; \
    dir="$(dirname "$idx")"; \
    cp -R "$dir"/. /app/out/; \
  fi

# Etapa 2: Producción - Servir con nginx
FROM nginx:alpine

# Copiar archivos compilados
COPY --from=builder /app/out /usr/share/nginx/html

# Copiar configuración de nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Exponer puerto
EXPOSE 80

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost/ || exit 1

# Iniciar nginx
CMD ["nginx", "-g", "daemon off;"]

