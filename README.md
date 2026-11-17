# Prestige Properties API

REST API for a real-estate listings platform, written in Go. It serves the [Next.js client](https://github.com/itsalmirr/Prestige-Properties-Client).

**Stack:** Go 1.24 · Gin · PostgreSQL (pgx) · Ent · Redis · Cloudinary · Docker

## Features

- Listings, realtors and users modeled with Ent; creates run in database transactions, and listings are checked for duplicate titles and addresses.
- Paginated listing queries, with page size capped at 100.
- Session authentication: Redis-backed cookie sessions, email/password sign-in with Argon2id hashing, and Google/GitHub OAuth.
- Multipart listing creation with image uploads to Cloudinary.
- CORS restricted to the configured frontend origin.
- Multi-stage Docker image that runs as a non-root user, with a `/health` check.
- CodeQL scanning on every push and pull request.

## API

| Method | Path | Description |
|---|---|---|
| GET | `/health` | Health check |
| GET | `/auth/:provider`, `/auth/:provider/callback` | OAuth sign-in (Google, GitHub) |
| POST | `/auth/signout` | End the session |
| POST | `/api/v1/users/` | Register |
| POST | `/api/v1/users/signin` | Email/password sign-in |
| GET | `/api/v1/users/me` | Current user (requires a session) |
| GET | `/api/v1/realtors/all`, `/api/v1/realtors/:email` | Realtors |
| POST | `/api/v1/realtors/` | Create a realtor |
| GET | `/api/v1/properties/buy` | Paginated listings |
| POST | `/api/v1/properties/add`, `/api/v1/properties/add-json` | Create a listing (multipart with images, or JSON) |
| PATCH | `/api/v1/properties/update` | Update a listing |
| DELETE | `/api/v1/properties/` | Delete a listing |
| POST | `/upload-images` | Upload images to Cloudinary |

## Run locally

```sh
cp .env.example .env   # database, Redis, OAuth, Cloudinary and frontend settings
docker compose up --build
```

The API listens on port 8080 and runs Ent migrations on startup.

## Layout

```
cmd/                    entry point and server wiring
ent/schema/             data model: listing, realtor, user
internal/api/           HTTP handlers; auth/ holds OAuth, sign-in and sign-out
internal/repositories/  data access and transactions
internal/routers/       routes, CORS, sessions and auth middleware
internal/services/      Cloudinary image service
internal/util/          image upload validation
internal/config/        configuration, database, sessions, OAuth
```

## Tests

```sh
go test ./...
```
