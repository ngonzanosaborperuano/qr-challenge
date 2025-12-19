export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  success: boolean;
  token: string;
  message: string;
  expiresIn: string;
}

export interface MatrixRequest {
  matrix: number[][];
}

export interface MatrixStats {
  max: number;
  min: number;
  avg: number;
  sum: number;
  anyDiagonal: boolean;
}

export interface MatrixProcessResponse {
  rotated: number[][];
  q: number[][];
  r: number[][];
  nodeStats?: MatrixStats;
  error?: string;
}

