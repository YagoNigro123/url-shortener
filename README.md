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