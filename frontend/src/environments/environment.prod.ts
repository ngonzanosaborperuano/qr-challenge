export const environment = {
  production: true,
  // En producción, el frontend se sirve por Nginx y este hace reverse-proxy a las APIs.
  // Así evitamos CORS y evitamos hardcodear host/puertos en el browser.
  goApiUrl: '/api/go',
  nodeApiUrl: '/api/node'
};

