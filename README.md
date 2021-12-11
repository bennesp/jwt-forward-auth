# Jwt Micro-Service

Simple http server that can receive a JWT in the http request and answer with a 200 if the JWT is valid, 401 otherwise.

It implements both JWKS (with both asymmetric RSA/ECDSA/... keys and symmetric OCT keys) and SECRET modes.

It can read jwt from headers and cookies and it can be extended to read it from anywhere (just write a sources/source.go file).

## Environment variables

### General

- `PORT` (default is `8080`): port where the http server will listen to

### Read JWT from a header

- `HEADER_JWT_SOURCE_ENABLED` (default is `true`): If true, header source is enabled
- `HEADER_JWT_SOURCE_NAME` (default is `Authorization`): Name of the header whose value is the jwt 
- `HEADER_JWT_SOURCE_PREFIX` (default is `Bearer `): If the value of the header is prefixed by a value, specify it with this environment variable so that it will be trimmed. If a value is specified but it is not found in the header, no errors will be thrown and no value will be trimmed.

### Read JWT from a cookie

- `COOKIE_JWT_SOURCE_ENABLED` (default is `false`):  If true, cookie source is enabled
- `COOKIE_JWT_SOURCE_NAME` (default is `token`): Name of the cookie whose value is the jwt

### Validate JWT with JWKS

- `JWKS_ENABLED` (default is `true`): If true, validation with JWKS is enabled
- `JWKS_URL` (default is `""`): URL of the keys of your IdP. For example https://jwks-service.appspot.com/.well-known/jwks.json
- `JWKS_REFRESH_INTERVAL` (default is `1h`): Interval between the refresh of the keys.
- `JWKS_REFRESH_RATE_LIMIT` (default is `5m`): Rate limit for the refresh of the keys.
- `JWKS_REFRESH_TIMEOUT` (default is `5s`): Timeout for the refresh of the keys.
- `JWKS_REFRESH_UNKNOWN_KID` (default is `true`): If true, unknown kid will be refreshed.

### Validate JWT with a secret

- `JWT_SECRET_ENABLED` (default is `false`): If true, validation with a secret is enabled
- `JWT_SECRET` (default is `""`): Secret used to sign and verify the JWT.
