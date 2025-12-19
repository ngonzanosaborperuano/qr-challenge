const { authenticateToken } = require('./auth');
const jwt = require('jsonwebtoken');

// Mock de Express request, response y next
const createMockReq = (authHeader) => ({
  headers: {
    authorization: authHeader
  }
});

const createMockRes = () => {
  const res = {};
  res.status = jest.fn().mockReturnValue(res);
  res.json = jest.fn().mockReturnValue(res);
  return res;
};

const createMockNext = () => jest.fn();

describe('authenticateToken middleware', () => {
  const originalEnv = process.env.JWT_SECRET;

  beforeEach(() => {
    process.env.JWT_SECRET = 'test-secret-key';
  });

  afterEach(() => {
    process.env.JWT_SECRET = originalEnv;
  });

  test('debe permitir acceso con token válido', () => {
    const token = jwt.sign(
      { username: 'admin', id: 1, role: 'admin' },
      'test-secret-key',
      { expiresIn: '1h' }
    );

    const req = createMockReq(`Bearer ${token}`);
    const res = createMockRes();
    const next = createMockNext();

    authenticateToken(req, res, next);

    expect(next).toHaveBeenCalled();
    expect(res.status).not.toHaveBeenCalled();
    expect(req.user).toBeDefined();
    expect(req.user.username).toBe('admin');
  });

  test('debe rechazar acceso sin token', () => {
    const req = createMockReq(undefined);
    const res = createMockRes();
    const next = createMockNext();

    authenticateToken(req, res, next);

    expect(next).not.toHaveBeenCalled();
    expect(res.status).toHaveBeenCalledWith(401);
    expect(res.json).toHaveBeenCalledWith(
      expect.objectContaining({
        error: 'Token de acceso requerido'
      })
    );
  });

  test('debe rechazar token inválido', () => {
    const req = createMockReq('Bearer invalid-token');
    const res = createMockRes();
    const next = createMockNext();

    authenticateToken(req, res, next);

    expect(next).not.toHaveBeenCalled();
    expect(res.status).toHaveBeenCalledWith(403);
    expect(res.json).toHaveBeenCalledWith(
      expect.objectContaining({
        error: 'Token inválido o expirado'
      })
    );
  });

  test('debe rechazar token con secreto incorrecto', () => {
    const token = jwt.sign(
      { username: 'admin' },
      'wrong-secret',
      { expiresIn: '1h' }
    );

    const req = createMockReq(`Bearer ${token}`);
    const res = createMockRes();
    const next = createMockNext();

    authenticateToken(req, res, next);

    expect(next).not.toHaveBeenCalled();
    expect(res.status).toHaveBeenCalledWith(403);
  });

  test('debe rechazar token expirado', () => {
    const token = jwt.sign(
      { username: 'admin' },
      'test-secret-key',
      { expiresIn: '-1h' } // Token expirado
    );

    const req = createMockReq(`Bearer ${token}`);
    const res = createMockRes();
    const next = createMockNext();

    authenticateToken(req, res, next);

    expect(next).not.toHaveBeenCalled();
    expect(res.status).toHaveBeenCalledWith(403);
  });

  test('debe rechazar formato de header incorrecto', () => {
    const req = createMockReq('InvalidFormat token');
    const res = createMockRes();
    const next = createMockNext();

    authenticateToken(req, res, next);

    expect(next).not.toHaveBeenCalled();
    expect(res.status).toHaveBeenCalledWith(401);
  });
});

