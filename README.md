# University Academic RESTful API

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Echo Framework](https://img.shields.io/badge/Echo-v4.13-blue?style=flat)](https://echo.labstack.com/)
[![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL-336791?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![GORM](https://img.shields.io/badge/ORM-GORM-red?style=flat)](https://gorm.io/)
[![JWT](https://img.shields.io/badge/Auth-JWT-000000?style=flat&logo=jsonwebtokens)](https://jwt.io/)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?style=flat&logo=docker)](https://www.docker.com/)

A comprehensive RESTful API for university academic management system built with Go, Echo framework, and PostgreSQL. Features student registration, authentication, course enrollment, and academic data management with clean architecture principles.

## 🚀 Features

- **Student Authentication**: JWT-based registration and login system
- **Course Management**: Browse and manage university courses
- **Enrollment System**: Enroll/unenroll students in courses
- **Department Structure**: Organized course categorization by departments
- **Professor Management**: Faculty and teaching assignments
- **Secure API**: JWT authentication with middleware protection
- **Database Relations**: Comprehensive relational database design
- **Input Validation**: Robust request validation using ozzo-validation
- **Error Handling**: Centralized error handling with custom exceptions
- **CORS Support**: Cross-origin resource sharing enabled
- **Swagger Documentation**: Interactive API documentation
- **Docker Support**: Containerized deployment ready

## 📋 Database Schema

### Core Entities

- **Students**: User accounts with authentication
- **Professors**: Faculty members
- **Departments**: Academic departments
- **Courses**: Academic courses with department relations
- **Enrollments**: Student-course relationships
- **Teachings**: Professor-course assignments

### Entity Relationships

```
students ──┐
           ├── enrollments ──── courses ──── departments
           │                      │
professors ────── teachings ───────┘
```

## 🛠 Tech Stack

- **Language**: Go 1.21+
- **Framework**: Echo v4.13
- **Database**: PostgreSQL with UUID extensions
- **ORM**: GORM v1.25
- **Authentication**: JWT tokens
- **Validation**: ozzo-validation
- **Password**: bcrypt hashing
- **Documentation**: OpenAPI 3.0 (Swagger)
- **Deployment**: Docker, Google Cloud Run

## 📁 Project Structure

```
univ-academic/
├── app/
│   └── database.go              # Database configuration
├── controller/
│   ├── student_controller.go    # Student endpoints
│   ├── course_controller.go     # Course endpoints
│   └── enrollment_controller.go # Enrollment endpoints
├── exception/
│   ├── error_handler.go         # Centralized error handling
│   ├── authentication_credential_error.go
│   ├── data_conflict_error.go
│   └── data_not_found_error.go
├── middleware/
│   └── jwt_middleware.go        # JWT authentication middleware
├── repository/
│   ├── student_repository.go    # Student data access
│   ├── course_repository.go     # Course data access
│   └── enrollment_repository.go # Enrollment data access
├── service/
│   ├── student_service.go       # Student business logic
│   ├── course_service.go        # Course business logic
│   └── enrollment_service.go    # Enrollment business logic
├── validator/
│   ├── student_validator.go     # Student validation rules
│   └── enrollment_validator.go  # Enrollment validation rules
├── ddl.sql                      # Database schema
├── query.sql                    # Sample queries
├── Dockerfile                   # Container configuration
├── docker-compose.yml           # Multi-service setup
├── Univ_Academic.swagger.yml    # API documentation
└── main.go                      # Application entry point
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Docker (optional)
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/FathurRohmna/univ-academic.git
   cd univ-academic
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   
   Create a `.env` file in the root directory:
   ```env
   DATABASE_URL=postgres://username:password@localhost:5432/univ_academic?sslmode=disable
   PORT=8080
   JWT_SECRET=your-secret-key-here
   ```

4. **Set up the database**
   
   Create PostgreSQL database and run the schema:
   ```bash
   createdb univ_academic
   psql -d univ_academic -f ddl.sql
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:8080`

## 🐳 Docker Deployment

### Using Docker Compose

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Manual Docker Build

```bash
# Build the API
docker build -t univ-academic-api .

# Run with environment variables
docker run -p 8080:8080 \
  -e DATABASE_URL="your-database-url" \
  -e JWT_SECRET="your-jwt-secret" \
  univ-academic-api
```

## 📚 API Documentation

### Base URL
```
http://localhost:8080/api
```

### Authentication

Most endpoints require JWT authentication. Include the token in the Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

### Core Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/students/register` | Register new student | ❌ |
| POST | `/students/login` | Student login | ❌ |
| GET | `/students/me` | Get current student details | ✅ |
| GET | `/courses` | Get all available courses | ✅ |
| POST | `/enrollments` | Enroll in a course | ✅ |
| DELETE | `/enrollments/{course_id}` | Cancel enrollment | ✅ |

### Request/Response Examples

#### Student Registration
```bash
POST /api/students/register
Content-Type: application/json

{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@university.edu",
  "password": "securePassword123",
  "address": "123 University St, Academic City",
  "birth_date": "2000-05-15"
}
```

**Response:**
```json
{
  "message": "Student registered successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@university.edu",
    "address": "123 University St, Academic City",
    "birth_date": "2000-05-15"
  }
}
```

#### Student Login
```bash
POST /api/students/login
Content-Type: application/json

{
  "email": "john.doe@university.edu",
  "password": "securePassword123"
}
```

**Response:**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### Get Student Details
```bash
GET /api/students/me
Authorization: Bearer <token>
```

**Response:**
```json
{
  "message": "Student details retrieved successfully",
  "data": {
    "student": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "first_name": "John",
      "last_name": "Doe",
      "email": "john.doe@university.edu",
      "address": "123 University St, Academic City",
      "birth_date": "2000-05-15"
    },
    "enrollments": [
      {
        "course_title": "Struktur Data",
        "enrollment_date": "2024-01-15T00:00:00Z"
      }
    ]
  }
}
```

#### Course Enrollment
```bash
POST /api/enrollments
Authorization: Bearer <token>
Content-Type: application/json

{
  "course_id": "550e8400-e29b-41d4-a716-446655440001"
}
```

**Response:**
```json
{
  "message": "Successfully enrolled in course",
  "data": {
    "course": {
      "id": "550e8400-e29b-41d4-a716-446655440001",
      "title": "Kalkulus",
      "description": "Konsep-konsep matematika tingkat lanjut"
    },
    "enrollment_date": "2024-01-16T10:30:00Z"
  }
}
```

## 🗄️ Database Queries

The `query.sql` file contains useful queries for data analysis:

- List all students with their enrolled courses
- List all courses with departments, professors, and enrolled students
- List all professors with their teaching assignments
- Enrollment details with dates and credits
- Department course listings
- Student enrollment statistics
- Average enrollment per department

## 🔐 Security Features

- **Password Hashing**: bcrypt with salt rounds
- **JWT Authentication**: Secure token-based authentication
- **Input Validation**: Comprehensive request validation
- **SQL Injection Prevention**: GORM ORM protection
- **CORS Configuration**: Controlled cross-origin access
- **Error Handling**: Secure error responses without sensitive data exposure

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 📊 Swagger Documentation

Interactive API documentation is available at:
- **Local**: `http://localhost:8080/swagger/`
- **Production**: `https://univ-academic-swagger-70017640279.us-central1.run.app/`

## 🌐 Live Demo

- **API**: `https://univ-academic-restful-api-70017640279.us-central1.run.app/api`
- **Documentation**: `https://univ-academic-swagger-70017640279.us-central1.run.app/`

## 🔧 Environment Variables

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `DATABASE_URL` | PostgreSQL connection string | ✅ | - |
| `PORT` | Server port | ❌ | 8080 |
| `JWT_SECRET` | JWT signing secret | ✅ | - |

## 📈 Performance & Scalability

- Connection pooling configured for optimal database performance
- JWT stateless authentication for horizontal scaling
- Clean architecture for maintainable codebase
- Docker containerization for consistent deployments
- Google Cloud Run deployment for auto-scaling

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go conventions and best practices
- Use `go fmt` to format your code
- Run `go vet` to check for common errors
- Write unit tests for new functionality
- Update documentation for API changes

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **Fathur Rohman** - *Initial work* - [FathurRohmna](https://github.com/FathurRohmna)

## 🙏 Acknowledgments

- Echo framework for excellent HTTP routing
- GORM for powerful ORM capabilities
- PostgreSQL for robust database foundation
- JWT for secure authentication
- Go community for amazing tools and libraries

## 📞 Support

For support and questions:
- 📧 Email: fr081938@gmail.com
- 🐛 Issues: [GitHub Issues](https://github.com/FathurRohmna/univ-academic-restful-api-with-swagger/issues)

---

**Happy Coding! 🎓**