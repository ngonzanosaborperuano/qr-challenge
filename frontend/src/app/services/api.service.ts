import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { AuthService } from './auth.service';
import { LoginRequest, LoginResponse, MatrixRequest, MatrixProcessResponse } from '../models/api.models';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private readonly GO_API_URL = environment.goApiUrl;
  private readonly NODE_API_URL = environment.nodeApiUrl;

  constructor(
    private http: HttpClient,
    private authService: AuthService
  ) {}

  private getHeaders(): HttpHeaders {
    const token = this.authService.getToken();
    let headers = new HttpHeaders().set('Content-Type', 'application/json');
    if (token) {
      headers = headers.set('Authorization', `Bearer ${token}`);
    }
    return headers;
  }

  login(credentials: LoginRequest): Observable<LoginResponse> {
    console.log('API Service - Login URL:', `${this.GO_API_URL}/auth/login`);
    console.log('API Service - Credentials:', credentials);
    
    // Para login no necesitamos el token, así que creamos headers sin Authorization
    const headers = new HttpHeaders().set('Content-Type', 'application/json');
    
    return this.http.post<LoginResponse>(
      `${this.GO_API_URL}/auth/login`,
      credentials,
      { headers }
    );
  }

  processMatrix(matrix: number[][]): Observable<MatrixProcessResponse> {
    const request: MatrixRequest = { matrix };
    return this.http.post<MatrixProcessResponse>(
      `${this.GO_API_URL}/matrix/process`,
      request,
      { headers: this.getHeaders() }
    );
  }

  getHealth(api: 'go' | 'node'): Observable<any> {
    const url = api === 'go' ? this.GO_API_URL : this.NODE_API_URL;
    return this.http.get(`${url}/health`);
  }
}

