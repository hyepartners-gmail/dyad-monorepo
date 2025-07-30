# Dyad Monorepo: Go Backend + React Frontend

This monorepo contains both the Go backend and React frontend templates configured for Dyad-based development and deployment.

---

## 📦 Structure

```

dyad-monorepo/
├── go-template/         # Go REST API service (Cloud Run ready)
├── react-template/      # React + Tailwind frontend (Vite, shadcn/ui)
├── ai\_rules.md          # Shared Dyad instruction set
├── dyad.config.json     # Multi-app routing for Dyad (when supported)

````

---

## ⚙️ Backend (Go)

Located in `backend/`

### 🔹 Run Locally

```bash
cd backend
go mod tidy
go run ./cmd/app
````

### 🔹 Test

```bash
go test ./...
```

### 🔹 Docker

```bash
make docker-build
make docker-run
```

### 🔹 Deploy to Cloud Run

```bash
make deploy
```

---

## 🎨 Frontend (React)

Located in `frontend/`

### 🔹 Run Locally

```bash
cd frontend
npm install
npm run dev
```

---

## 🧠 Dyad Compatibility

This monorepo is designed to support Dyad’s AI tooling across both frontend and backend apps.

* Dyad currently assumes the root contains a `package.json` or `go.mod`
* **To work with the frontend**, open Dyad from the `react-template/` folder:

```bash
dyad react-template/
```

* The included `dyad.config.json` is a **hint file** for future Dyad support of multi-service monorepos.

---

## 🔐 Environment Variables

Sample `.env` is provided for both apps in their respective directories.

* Go config supports both `.env` and `config.yaml`
* React config uses standard `VITE_`-prefixed env vars

---

## 🛠 Make Targets

Top-level Makefile (optional) can be added to support:

```bash
make run-dev      # Start both apps locally
make test-all     # Run Go tests + (optional) React tests
```

---

## 🧪 Contributing

1. Keep all new apps in their own subfolders
2. Use consistent naming: `backend/`, `frontend/`, etc.
3. Update `ai_rules.md` when tech stack or conventions evolve

---

## 💬 Contact

Questions? Ideas? Reach out via the Dyad Hub discussion board or fork this repo and open a PR. Created by HYE Partners.
