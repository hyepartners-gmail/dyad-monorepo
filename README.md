# Go Template for Dyad

A production-ready Go template with Google Cloud integration, optimized for Dyad AI workflows and modular backend services.

---

## 🔧 Features

- ✅ Google Cloud Datastore integration
- ✅ OAuth2 + JWT authentication ready
- ✅ CORS + structured logging middleware
- ✅ Config from `.env` or `config.yaml`
- ✅ Dockerfile + Cloud Run ready
- ✅ Health check endpoint for GCP
- ✅ Makefile with build/test/deploy tasks

---

## 📂 Project Structure

```

.
├── cmd/app/                 # Main server entrypoint
├── config/                  # Loads YAML or env config
├── internal/
│   ├── handler/             # HTTP handlers
│   └── middleware/          # Logging + CORS
├── Dockerfile
├── cloudrun.yaml
├── .env.sample
├── go.mod / go.sum

````

---

## 🚀 Getting Started

### Install

```bash
go mod tidy
````

### Run Locally

```bash
go run ./cmd/app
```

### Test

```bash
go test ./...
```

---

## 🐳 Docker

```bash
docker build -t go-template .
docker run -p 8080:8080 go-template
```

---

## ☁️ Deploy to Google Cloud Run

```bash
make deploy
```

Or manually:

```bash
gcloud run deploy go-template \
  --source=. \
  --region=us-west1 \
  --platform=managed \
  --allow-unauthenticated
```

---

## 📄 Sample `.env`

```env
GO_ENV=production
PORT=8080
PROJECT_ID=my-gcp-project
JWT_SECRET=dev-secret
```

---

## 🧠 Designed for Dyad

This template works out-of-the-box with Dyad's AI Assist + Component Edit. It's structured for traceability, testability, and AI-guided workflows.
