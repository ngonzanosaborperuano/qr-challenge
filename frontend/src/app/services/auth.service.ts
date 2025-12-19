import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private readonly TOKEN_KEY = 'jwt_token';

  setToken(token: string): void {
    localStorage.setItem(this.TOKEN_KEY, token);
  }

  getToken(): string | null {
    return localStorage.getItem(this.TOKEN_KEY);
  }

  removeToken(): void {
    localStorage.removeItem(this.TOKEN_KEY);
  }

  /**
   * Cierra la sesión del usuario
   * Elimina el token y limpia el almacenamiento local
   */
  logout(): void {
    this.removeToken();
  }

  isAuthenticated(): boolean {
    return this.getToken() !== null;
  }
}

