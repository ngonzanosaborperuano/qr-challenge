import { Pipe, PipeTransform } from '@angular/core';

/**
 * Pipe para formatear matrices a string con formato legible
 * Aplica el principio de responsabilidad única (SRP)
 */
@Pipe({
  name: 'matrixFormat',
  standalone: true
})
export class MatrixFormatPipe implements PipeTransform {
  transform(matrix: number[][] | null | undefined, decimals: number = 2): string {
    if (!matrix || matrix.length === 0) {
      return '[]';
    }

    return matrix
      .map(row => 
        '[' + row.map(val => val.toFixed(decimals)).join(', ') + ']'
      )
      .join('\n');
  }
}

