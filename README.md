# Real Estate Listings API

A robust REST API built with Go (Gin framework) for managing real estate property listings. This API provides endpoints for creating, retrieving, and managing property listings with features like pagination and realtor authentication.

## 🚀 Features

- Property listing management
- Pagination support
- Image handling for multiple property photos
- Realtor authentication
- Duplicate listing prevention
- Transaction support for data integrity

## 🛠️ Tech Stack

- **Go** - Primary programming language
- **Gin** - Web framework
- **GORM** - ORM for database operations
- **PostgreSQL** - Database (assumed based on the codebase)
- **UUID** - For unique identifier generation

## 📋 API Endpoints

### Listings

#### Get Listings

GET /listings?page=1&limit=10

Query Parameters:

- `page` (optional) - Page number (default: 1)
- `limit` (optional) - Items per page (default: 10)

Response:

```json
{
    "status": "OK",
    "data": [...],
    "total": 100,
    "current_page": 1,
    "total_page": 10,
    "per_page": 10
}
```

#### Create Listing

```http
POST /listings
```

Required fields in request body:

- `title` - Property title
- `address` - Property address
- `price` - Property price
- ... (other fields as defined in CreateListingInput)

## 🔐 Realtor Authentication

### Register

```http
POST /realtors/register
```

Required fields in request body:

- `name` - Realtor's name
- `email` - Realtor's email
- `password` - Realtor's password

### Login Endpoint

**Request Body Parameters:**

- **email**: The user's email address.
- **password**: The user's password.

**Response:**

Upon successful authentication, the API will redirect to `/api/v1/users/me` and return a JSON object with the user's details, for example:

```json
{
  "data": {
    "id": "user_id",
    "email": "user_email",
    "username": "username",
    "full_name": "User's Full Name",
    "start_date": "If realtor",
    "is_staff": false,
    "is_active": true,
    "provider": "email",
    "created_at": "date",
    "updated_at": "date"
  }
}

Note: Session data is stored in the Redis database.
```

## 🏗️ Project Structure

```
app/
├── api/
│   └── listings.go # API handlers
├── config/
│   └── database.go # Database configuration
├── middleware/
│   └── auth.go # Authentication middleware
├── models/
│   └── listing.go # Data models
└── repositories/
    └── listing.repo.go # Database operations
```

# -
