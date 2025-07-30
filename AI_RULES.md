# AI Rules for Dyad Monorepo

## Structure
- `backend/`: backend REST API in Go 1.24
- `frontend/`: frontend app in React + Tailwind + shadcn/ui

## Frontend
- React w/ TypeScript
- TailwindCSS
- Mantine or Shadcn UI (match the setup)
- Router: React Router
- Build: Vite
- Directory: `react-template/src/pages`, `components`, etc.

## Backend
- Go 1.24
- Cloud Run + Datastore
- YAML/env config system
- JWT & OAuth2 support
- Docker-ready

## Shared Conventions
- All components should be self-contained and testable
- AI should assume top-level context unless scoped by folder
- Config is `.env` based, but fallback to YAML allowed
- You have detailed AI_RULES.md files in frontend & backend folders