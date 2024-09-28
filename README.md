# Financial Assistance Scheme API

## Overview

The Financial Assistance Scheme API is a Go-based web service that manages applicants, schemes, and applications for financial assistance programs. It provides a RESTful API interface and includes Swagger documentation for easy integration and testing.

## Features

- CRUD operations for applicants, schemes, and applications
- Pagination support for list endpoints
- Swagger UI for interactive API documentation
- Dockerized for easy deployment and consistency across environments
- Built with Go and uses gorilla/mux for routing

## Prerequisites

- Go 1.18 or higher
- Docker and Docker Compose (for containerized deployment)
- PostgreSQL (for local development without Docker)

## Quick Start

1. Clone the repository:
   ```
   git clone https://github.com/jeredwong/financial-scheme-manager-api.git   
   cd financial-scheme-manager-api
   ```

2. Build and run using Docker Compose:
   ```
   docker-compose up --build
   ```

3. Access the API at `http://localhost:8080`
4. View the Swagger UI documentation at `http://localhost:8080/swagger-ui/`

## Local Development

1. Ensure you have Go 1.18+ installed
2. Install dependencies:
   ```
   go mod download
   ```
3. Set up your local PostgreSQL database
4. Copy `.env.example` to `.env` and update the values
5. Run the application:
   ```
   go run cmd/server/main.go
   ```

## API Endpoints

- `GET /api/health`: Health check endpoint
- `GET /api/applicants`: List all applicants
- `POST /api/applicants`: Create a new applicant
- `GET /api/schemes`: List all schemes
- `GET /api/schemes/eligible1`: List all eligible schemes for an applicant
- `GET /api/applications`: List all applications
- `POST /api/applications`: Create a new application

For detailed API documentation, please refer to the Swagger UI when running the application.

<!-- ## Configuration

The application can be configured using environment variables. See `.env.example` for available options.

## Testing

Run the test suite with:

```
go test ./...
```

## Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for more details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Security

If you discover any security-related issues, please email security@example.com instead of using the issue tracker.

## Support

If you have any questions or need assistance, please open an issue on the GitHub repository. -->
