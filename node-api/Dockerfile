# Dockerfile para API Node.js
# Etapa 1: Builder - Instalar dependencias
FROM node:20-alpine AS builder

# Establecer directorio de trabajo
WORKDIR /app

# Copiar archivos de dependencias
COPY package*.json ./

# Instalar solo dependencias de producción
RUN npm ci --only=production && npm cache clean --force

# Etapa 2: Producción - Imagen final optimizada
FROM node:20-alpine

# Instalar wget para healthchecks
RUN apk --no-cache add wget

# Establecer directorio de trabajo
WORKDIR /app

# Copiar archivos de dependencias desde builder
COPY --from=builder /app/package*.json ./

# Copiar node_modules desde builder (solo producción)
COPY --from=builder /app/node_modules ./node_modules

# Copiar código fuente
COPY . .

# Crear usuario no-root para mayor seguridad
RUN addgroup -g 1001 -S nodejs && \
    adduser -S nodejs -u 1001 && \
    chown -R nodejs:nodejs /app

USER nodejs

# Exponer puerto
EXPOSE 3001

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost:3001/health || exit 1

# Comando para ejecutar la aplicación
CMD ["npm", "start"]

