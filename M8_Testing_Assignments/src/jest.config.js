module.exports = {
    testEnvironment: 'node',
    verbose: true,
    testMatch: ['**/__tests__/**/*.js'],
    collectCoverage: true,
    coverageDirectory: 'coverage',
    coverageReporters: ['text','lcov']
}