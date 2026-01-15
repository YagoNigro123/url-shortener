# URL Shortener (High Performance)

## Technologies
* **System Core:** Golang (Concurrency handling).
* **SQL Database:** PostgreSQL (Persistent storage).
* **NoSQL / Cache:** Redis (In-memory cache for high-speed retrieval).
* **Infrastructure:** Docker & Docker Compose.
* **Event Processing:** Asynchronous analytics using Go routines.

## Project Journal

### Day 1: Infrastructure as Code
I established the development environment using **Docker Compose** to orchestrate PostgreSQL and Redis services.
* **Persistence:** Configured Docker Volumes to ensure data persists even if containers are destroyed.
* **Networking:** Configured Port Mapping to access the databases from the host machine during development.
* **Caching Strategy:** Set up Redis to act as a look-aside cache, reducing load on the primary SQL database.

### Day 2: Core Domain & Business Logic
I focused on the **Hexagonal Architecture** principles (or Layered Architecture), isolating the core logic from external dependencies.
* **Domain Model:** Defined the `Link` struct and the `LinkStore` interface in the `core` package.
* **Polymorphism:** The `LinkStore` interface decouples the logic from the storage, allowing future implementations of Redis or PostgreSQL without changing the core code.
* **Business Logic:** Implemented the `Service` struct to handle URL shortening using a random ID generator.
* **Dependency Injection:** Applied DI pattern in the `NewService` constructor to inject the storage implementation.

### Day 3: Persistence Layer & Entrypoint
I integrated the PostgreSQL database with the Go application to ensure data persistence.

* **PostgreSQL Adapter:** Implemented the `PostgresStore` struct using the `database/sql` package. Used **Parameterized Queries** ($1, $2) to prevent SQL Injection.
* **Schema Migration:** Created and executed `schema.sql` to define the database structure (`links` table) and indexes directly inside the Docker container.
* **Main Entrypoint:** Orchestrated the application startup. I connected the components by injecting the concrete `PostgresStore` into the `Service` (Dependency Injection).

### Day 4: RESTful API Implementation
Transformed the CLI application into a fully functional Web Service using the **Chi** router.

* **HTTP Layer:** Implemented the `Handler` struct to act as an adapter between the HTTP world (JSON) and the Core Domain.
* **Endpoints:**
  * `POST /api/shorten`: Accepts a JSON payload, validates input, and returns the shortened ID (Status `201 Created`).
  * `GET /{id}`: Retrieves the original URL and performs an HTTP `301 Moved Permanently` redirection.
* **Routing:** Integrated `github.com/go-chi/chi` for lightweight and idiomatic routing, including middleware for logging and recovery.

### Day 5: Configuration & Security
Refactored the application to follow **12-Factor App** principles regarding configuration.

* **Environment Variables:** Removed hardcoded credentials. The application now reads configuration (`DB_URL`, `PORT`) from the system environment.
* **Security:** Implemented `.env` file support using `godotenv` for local development, while ensuring secrets are excluded from version control via `.gitignore`.

curl -X POST http://localhost:8080/api/shorten \
     -H "Content-Type: application/json" \
     -d '{"original_url": "https://www.youtube.com/watch?v=dQw4w9WgXcQ"}'