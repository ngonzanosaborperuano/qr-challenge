module.exports = {
  testEnvironment: 'node',
  coverageDirectory: 'coverage',
  collectCoverageFrom: [
    'src/**/*.js',
    '!src/index.js', // Excluir archivo principal
  ],
  testMatch: [
    '**/__tests__/**/*.js',
    '**/*.test.js'
  ],
  verbose: true
};

