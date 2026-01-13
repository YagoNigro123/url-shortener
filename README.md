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