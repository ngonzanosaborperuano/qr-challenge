import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';

/**
 * Componente reutilizable para mostrar cajas de información
 * Aplica el principio de responsabilidad única (SRP) y reutilización
 */
@Component({
  selector: 'app-info-box',
  standalone: true,
  imports: [CommonModule],
  template: `
    <div 
      class="info-box" 
      [ngClass]="type"
      [style.background]="background"
      [style.border-left-color]="borderColor">
      <ng-content></ng-content>
    </div>
  `,
  styles: [`
    .info-box {
      padding: 10px;
      margin-bottom: 10px;
      border-radius: 5px;
      font-size: 0.9em;
      border-left: 4px solid;
    }
    
    .info-box.info {
      background: #f0f0f0;
      color: #555;
      border-left-color: #6c757d;
    }
    
    .info-box.warning {
      background: #fff3cd;
      color: #856404;
      border-left-color: #ffc107;
    }
  `]
})
export class InfoBoxComponent {
  @Input() type: 'info' | 'warning' = 'info';
  @Input() background?: string;
  @Input() borderColor?: string;
}

