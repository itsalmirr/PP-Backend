# Real Estate Listings API

A robust REST API built with Go (Gin framework) for managing real estate property listings. This API provides endpoints for creating, retrieving, and managing property listings with features like pagination and realtor authentication.

## 🚀 Features

- Property listing management
- Realtor and user management
- Pagination support
- Image handling for multiple property photos
- Realtor authentication
- Duplicate listing prevention
- Transaction support for data integrity
- Database schema generation with Ent

## 🛠️ Tech Stack

- **Go** - Primary programming language
- **Gin** - Web framework
- **Ent** - Entity framework for Go (database operations and schema management)
- **PostgreSQL** - Database
- **Docker** - Containerization
- **UUID** - For unique identifier generation

## 📋 API Endpoints

### Listings
- `GET /listings` - Get all property listings with pagination
- `GET /listings/:id` - Get a specific listing by ID
- `POST /listings` - Create a new property listing
- `PUT /listings/:id` - Update an existing listing
- `DELETE /listings/:id` - Delete a listing

### Realtors
- `GET /realtors` - Get all realtors
- `GET /realtors/:id` - Get a specific realtor
- `POST /realtors` - Register a new realtor
- `PUT /realtors/:id` - Update realtor information
- `DELETE /realtors/:id` - Delete a realtor

### Users
- `GET /users` - Get all users
- `GET /users/:id` - Get a specific user
- `POST /users` - Create a new user
- `PUT /users/:id` - Update user information
- `DELETE /users/:id` - Delete a user

## 🔐 Authentication

The API includes realtor authentication middleware to ensure only authorized realtors can create, update, or delete property listings.

## 🏗️ Project Structure

```
PP-Backend/
├── cmd/
│   ├── main.go          # Application entry point
│   └── server.go        # Server configuration and setup
├── ent/                 # Ent generated code and schema
│   ├── client.go        # Database client
│   ├── listing.go       # Listing entity
│   ├── realtor.go       # Realtor entity
│   ├── user.go          # User entity
│   ├── *_create.go      # Create operations
│   ├── *_query.go       # Query operations
│   ├── *_update.go      # Update operations
│   ├── *_delete.go      # Delete operations
│   └── enttest/         # Testing utilities
├── internal/            # Internal application logic
├── proto/              # Protocol buffer definitions
├── docker-compose.yml  # Docker services configuration
├── Dockerfile         # Container build instructions
├── go.mod            # Go module dependencies
└── go.sum           # Dependency checksums
```

## 🚀 Getting Started

### Prerequisites
- Go 1.19 or higher
- Docker and Docker Compose
- PostgreSQL (if running locally)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/itsalmirr/PP-Backend.git
cd PP-Backend
```

2. Copy environment variables:
```bash
cp .env.example .env
```

3. Edit `.env` file with your database credentials and configuration.

4. Run with Docker Compose:
```bash
docker-compose up --build
```

Or run locally:
```bash
go mod download
go run cmd/main.go
```

### Database Schema Generation

This project uses Ent for database schema management. To generate or update the database schema:

```bash
go generate ./ent
```

## 🧪 Testing

Run the test suite:
```bash
go test ./...
```

## 📝 Environment Variables

Copy `.env.example` to `.env` and configure:

- `DATABASE_URL` - PostgreSQL connection string
- `PORT` - Server port (default: 8080)
- `JWT_SECRET` - Secret key for JWT authentication
- Additional configuration as needed

## 🐳 Docker

The application includes Docker support with multi-stage builds for production deployment:

```bash
# Build and run with Docker Compose
docker-compose up --build

# Or build the Docker image manually
docker build -t real-estate-api .
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📄 License

This project is licensed under the MIT License.