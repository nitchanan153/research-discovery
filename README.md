# Bio-RD Research Discovery MVP

MVP starter for the Bio-RD intelligence and discovery workflow. Frontend: Next.js + TypeScript. Backend: Go standard library HTTP API.

## Run locally

Requirements: Node.js 20.9+ with npm, and Go 1.22+.

Terminal 1 — backend:

```powershell
cd backend
go run .
```

The API listens at `http://localhost:8080`.

Terminal 2 — frontend:

```powershell
cd frontend
npm install
npm run dev
```

Open `http://localhost:3000`. The MVP UI currently uses clearly labeled sample data; API endpoints are available at `/api/health`, `/api/v1/dashboard`, and `/api/v1/opportunities?q=...`.

## Current scope

- Research Intelligence dashboard, sample opportunity ranking, technology trend summary and pipeline cards.
- Search/filter interactions and a Research Brief draft modal.
- Go read-only JSON API with sample data and CORS for local development.

## Next implementation steps

1. Connect the frontend to Go API and replace sample data with an approved initial data source.
2. Add PostgreSQL persistence, migrations, authentication/RBAC and audit trail.
3. Add ingestion jobs and source provenance before enabling AI summaries.
4. Add evidence citations, expert review and configurable opportunity scoring.

All displayed research entries are illustrative and are not validated scientific recommendations.
