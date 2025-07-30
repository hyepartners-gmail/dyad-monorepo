# AI Rules for Go Template

## Tech Stack
- **Language**: Go 1.24+
- **Core Dependencies**:
  - Google Cloud Datastore (`cloud.google.com/go/datastore`)
  - JWT Authentication (`github.com/dgrijalva/jwt-go`)
  - OAuth2 Client (`golang.org/x/oauth2`)
  - UUID Generation (`github.com/google/uuid`)
  - YAML parsing (`gopkg.in/yaml.v3`)
  - CORS middleware (`github.com/rs/cors`)

- **Indirect Dependencies** (managed via Go modules):
  - gRPC, IAM, OpenTelemetry, Google Cloud auth/compute helpers

## Architecture
- `cmd/app`: entrypoint/main function
- `config/`: loads app config from `.env` or `config.yaml`
- `internal/handler`: HTTP handlers for core routes
- `internal/middleware`: CORS + logging middleware
- `pkg/`: reserved for reusable logic (currently unused)

## Cloud Infrastructure
- Designed for Google Cloud Run
- Environment variables injected via `.env` or deploy YAML
- Compatible with Secret Manager or Vault

## Dev Workflow
- Run: `go run ./cmd/app`
- Test: `go test ./...`
- Build: `go build ./cmd/app`
- Lint: `go vet ./...`
- Docker: `make docker-build && make docker-run`
- Deploy: `make deploy` (Cloud Run)

## AI Instructions
- Use idiomatic Go style (short funcs, minimal global state)
- All new logic should accept context (`ctx context.Context`)
- Use interfaces where testability is important
- Middleware must be chainable via `http.Handler`
