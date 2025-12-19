import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './components/login.component';
import { MatrixProcessorComponent } from './components/matrix-processor.component';
import { AuthService } from './services/auth.service';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, LoginComponent, MatrixProcessorComponent],
  template: `
    <div class="container">
      <h1>🔢 QR Challenge - Matrix Processor</h1>
      <p style="color: #666; margin-bottom: 30px;">
        Procesa matrices: rotación 90°, factorización QR y estadísticas
      </p>

      <app-login></app-login>

      <div class="section-divider" *ngIf="authService.isAuthenticated()"></div>

      <app-matrix-processor *ngIf="authService.isAuthenticated()"></app-matrix-processor>

      <div *ngIf="!authService.isAuthenticated()" style="text-align: center; color: #999; margin-top: 30px;">
        <p>Inicia sesión para comenzar a procesar matrices</p>
      </div>
    </div>
  `,
  styles: []
})
export class AppComponent {
  constructor(public authService: AuthService) {}
}

