import { Component, DoCheck } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../services/api.service';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule, FormsModule],
  template: `
    <div class="login-section">
      <h2>Iniciar Sesión</h2>
      <form (ngSubmit)="onLogin()" *ngIf="!isLoggedIn">
        <div class="form-group">
          <label for="username">Usuario:</label>
          <input
            type="text"
            id="username"
            [(ngModel)]="username"
            name="username"
            required
            placeholder="admin"
          />
        </div>
        <div class="form-group">
          <label for="password">Contraseña:</label>
          <input
            type="password"
            id="password"
            [(ngModel)]="password"
            name="password"
            required
            placeholder="admin"
          />
        </div>
        <button type="submit" [disabled]="loading">
          {{ loading ? 'Iniciando sesión...' : 'Iniciar Sesión' }}
        </button>
      </form>
      <div *ngIf="error" class="error">{{ error }}</div>
      <div *ngIf="isLoggedIn" class="success">
        ✅ Sesión iniciada correctamente. Token guardado.
      </div>
    </div>
  `,
  styles: []
})
export class LoginComponent implements DoCheck {
  username = 'admin';
  password = 'admin';
  loading = false;
  error = '';
  isLoggedIn = false;

  constructor(
    private apiService: ApiService,
    public authService: AuthService
  ) {
    this.updateLoginStatus();
  }

  /**
   * Actualiza el estado de login basado en el token
   */
  private updateLoginStatus(): void {
    this.isLoggedIn = this.authService.isAuthenticated();
  }

  /**
   * Se ejecuta cuando cambia el estado de autenticación
   * Angular detecta cambios en authService.isAuthenticated()
   */
  ngDoCheck(): void {
    const currentStatus = this.authService.isAuthenticated();
    if (currentStatus !== this.isLoggedIn) {
      this.isLoggedIn = currentStatus;
      if (!this.isLoggedIn) {
        // Limpiar campos si se cerró sesión
        this.error = '';
        this.username = 'admin';
        this.password = 'admin';
      }
    }
  }

  onLogin(): void {
    if (!this.username || !this.password) {
      this.error = 'Por favor completa todos los campos';
      return;
    }

    this.loading = true;
    this.error = '';

    console.log('Intentando login con:', { username: this.username });

    this.apiService.login({
      username: this.username,
      password: this.password
    }).subscribe({
      next: (response) => {
        console.log('Respuesta del servidor:', response);
        if (response && response.success && response.token) {
          this.authService.setToken(response.token);
          this.isLoggedIn = true;
          this.loading = false;
          this.error = '';
        } else {
          this.error = response?.message || 'Error al iniciar sesión';
          this.loading = false;
        }
      },
      error: (err) => {
        console.error('Error completo en login:', err);
        console.error('Error status:', err.status);
        console.error('Error message:', err.message);
        console.error('Error error:', err.error);
        
        if (err.status === 0 || err.status === undefined) {
          this.error = 'No se puede conectar con el servidor. Verifica que las APIs estén levantadas y que la configuración del frontend sea correcta.';
        } else if (err.status === 401) {
          this.error = err.error?.message || 'Credenciales incorrectas';
        } else if (err.status === 400) {
          this.error = err.error?.error || err.error?.message || 'Datos inválidos';
        } else if (err.error?.error) {
          this.error = err.error.error;
        } else if (err.error?.message) {
          this.error = err.error.message;
        } else {
          this.error = `Error al conectar con el servidor (${err.status || 'desconocido'})`;
        }
        this.loading = false;
      }
    });
  }
}

