# 🚀 High-Performance URL Shortener

A scalable, distributed URL shortener service built with Go, focusing on performance, clean architecture, and cloud-native practices.

**Live Demo:** [https://url-shortener-production-38b4.up.railway.app](https://url-shortener-production-38b4.up.railway.app) 🔗

---

## 🛠 Tech Stack & Architecture

- **Core Language:** Go (Golang) 1.25 (Concurrency handling)
- **Architecture:** Hexagonal Architecture (Ports & Adapters)
- **Database:** PostgreSQL (Persistent storage)
- **Caching:** Redis (Cache-Aside pattern for high-speed retrieval)
- **DevOps:** Docker, Docker Compose (Multistage Builds)
- **Frontend:** Vanilla JS + CSS3
- **Cloud:** Railway (CI/CD)

## 🔌 API Usage

### Shorten URL
```bash
curl -X POST http://localhost:8080/api/shorten \
     -H "Content-Type: application/json" \
     -d '{"original_url": "[https://www.youtube.com/watch?v=dQw4w9WgXcQ](https://www.youtube.com/watch?v=dQw4w9WgXcQ)"}'
```
### Response
```
{
  "id": "aBcD12",
  "short_url": "http://localhost:8080/aBcD12"
}
```
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

### Day 6: Unit Testing & Quality Assurance
Introduced a robust testing strategy for the Core domain.
* **Unit Tests:** Implemented tests for the `Service` layer using `stretchr/testify`.
* **Mocking:** Created a `MockStore` to isolate business logic from infrastructure dependencies, enabling true unit testing without a running database.
* **Assertions:** Verified success scenarios and error handling logic.

### Day 7: Performance & Scalability (Redis)
Optimized the system for high concurrency using **Redis** as a caching layer.
* **Caching Strategy:** Implemented the **Cache-Aside** pattern.
    * *Read Path:* Checks Redis first (RAM); falls back to PostgreSQL (Disk) only on misses.
    * *Write Path:* Asynchronous cache updates to minimize user latency.
* **Architecture:** Injected `RedisClient` as a dependency into the Service, maintaining clean architecture principles.
* **Result:** Drastically reduced database load and improved response times for frequently accessed links.

### Day 8: DevOps & Dockerization
Transitioned from local execution to a containerized environment.
* **Multistage Build:** Created a `Dockerfile` that compiles the Go binary in a builder stage and deploys it into a lightweight **Alpine Linux** image (reducing image size significantly).
* **Orchestration:** Updated `docker-compose` to run the Application, PostgreSQL, and Redis in a unified network.
* **Networking:** Configured service-to-service communication using Docker DNS names (`db`, `cache`) instead of localhost.

### Day 9: Frontend & Static Assets
Developed a user-facing interface to make the tool accessible.

* **Modularization:** Separated logic (JS), styles (CSS), and structure (HTML) in a public/ directory.

* **Integration:** Configured Go's http.FileServer to serve static assets alongside the API.

### Day 10: Cloud Deployment (CI/CD)
Deployed the full stack application to the cloud using Railway.

* **CI/CD:** Configured automatic deployments from GitHub.

* **Production Config:** Managed environment variables and secrets securely in the cloud.

* **Public Access:** Generated a live SSL production URL.