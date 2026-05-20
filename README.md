# Article Management API (Coding Test)

Simple RESTful API built with Go (Golang) for managing articles.

## Tech Stack
- **Language**: Go (Golang)
- **Framework**: Gin Gonic
- **ORM**: GORM
- **Database**: MySQL

## Prerequisites
- Go 1.26+ installed
- MySQL Database running
- Postman (optional, for testing)

## Setup Instructions

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yudhaeric/service-article-sv.git
   cd service-article-sv
   ```

2. **Configure Environment Variables**:
   Create a `.env` file in the root directory and configure your database connection:
   ```env
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_HOST=127.0.0.1
   DB_NAME=article
   ```

3. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

4. **Run the Application**:
   ```bash
   go run main.go
   ```
   The server will start on `http://localhost:8080`.

## API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| **POST** | `/article/` | Create a new article |
| **GET** | `/article/` | Get articles with pagination (`?limit=10&offset=0`) |
| **GET** | `/article/:id` | Get a single article by ID |
| **PUT** | `/article/:id` | Update an article |
| **DELETE** | `/article/:id` | Delete an article |

## Validation Rules
- **Title**: Required, minimum 20 characters.
- **Content**: Required, minimum 200 characters.
- **Category**: Required, minimum 3 characters.
- **Status**: Required, must be one of: `publish`, `draft`, `thrash`.

## Testing with Postman
1. Open Postman.
2. Import `Article_API.postman_collection.json` located in the root folder.
3. You can now test all endpoints with pre-configured requests.
