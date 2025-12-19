import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../services/api.service';
import { AuthService } from '../services/auth.service';
import { MatrixProcessResponse } from '../models/api.models';
import { MatrixFormatPipe } from '../pipes/matrix-format.pipe';
import { InfoBoxComponent } from './info-box.component';

/**
 * Componente para procesar matrices
 * Aplica principios SOLID:
 * - SRP: Solo maneja la presentación y coordinación
 * - DIP: Depende de abstracciones (servicios)
 */
@Component({
  selector: 'app-matrix-processor',
  standalone: true,
  imports: [CommonModule, FormsModule, MatrixFormatPipe, InfoBoxComponent],
  templateUrl: './matrix-processor.component.html',
  styleUrls: ['./matrix-processor.component.css']
})
export class MatrixProcessorComponent {
  matrixInput = '[[1, 2, 3], [4, 5, 6], [7, 8, 9]]';
  result: MatrixProcessResponse | null = null;
  loading = false;
  error = '';

  constructor(
    public apiService: ApiService,
    public authService: AuthService
  ) {}

  /**
   * Procesa la matriz ingresada por el usuario
   * Aplica validación y manejo de errores robusto
   */
  processMatrix(): void {
    if (this.isProcessing()) {
      return;
    }

    this.resetState();
    
    if (!this.isValidInput()) {
      this.error = 'Por favor ingresa una matriz válida';
      return;
    }

    const matrix = this.parseMatrix();
    if (!matrix) {
      return; // Error ya fue establecido en parseMatrix
    }

    this.executeMatrixProcessing(matrix);
  }

  /**
   * Verifica si hay un procesamiento en curso
   */
  private isProcessing(): boolean {
    return this.loading;
  }

  /**
   * Resetea el estado del componente
   */
  private resetState(): void {
    this.loading = true;
    this.error = '';
    this.result = null;
  }

  /**
   * Valida que el input no esté vacío
   */
  private isValidInput(): boolean {
    return this.matrixInput.trim().length > 0;
  }

  /**
   * Parsea y valida la matriz del input
   * Retorna null si hay error (y establece this.error)
   */
  private parseMatrix(): number[][] | null {
    try {
      const matrix = JSON.parse(this.matrixInput);
      
      if (!Array.isArray(matrix) || matrix.length === 0) {
        this.error = 'La matriz debe ser un array no vacío';
        this.loading = false;
        return null;
      }

      return matrix;
    } catch (parseError) {
      this.error = 'Error al parsear la matriz. Asegúrate de usar formato JSON válido.';
      this.loading = false;
      return null;
    }
  }

  /**
   * Ejecuta el procesamiento de la matriz a través del servicio
   */
  private executeMatrixProcessing(matrix: number[][]): void {
    this.apiService.processMatrix(matrix).subscribe({
      next: (response) => this.handleSuccess(response),
      error: (err) => this.handleError(err)
    });
  }

  /**
   * Maneja la respuesta exitosa del procesamiento
   */
  private handleSuccess(response: MatrixProcessResponse): void {
    this.result = response;
    this.loading = false;
    
    if (response.error) {
      this.error = `Se procesó la matriz pero hubo un error al obtener estadísticas: ${response.error}`;
    }
  }

  /**
   * Maneja errores del procesamiento
   */
  private handleError(err: any): void {
    this.error = err.error?.error || 'Error al procesar la matriz';
    this.loading = false;
  }
}

