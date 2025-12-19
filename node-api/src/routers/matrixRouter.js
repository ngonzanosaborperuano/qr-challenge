/**
 * Router para endpoints de matrices
 */

const express = require('express');
const router = express.Router();
const matrixController = require('../controllers/matrixController');

// POST /matrix/stats - Calcula estadísticas sobre matrices
router.post('/stats', matrixController.getMatrixStats);

module.exports = router;


