# GoPostulate

> Note: This project is currently in the very early stages of development and is not yet ready for use of any kind.

GoPostulate is a full-stack web application framework inspired by NextJS, leveraging the power of Go and complementary technologies. It is designed to facilitate the rapid development of performant, Server-Side Rendered (SSR) Single Page Applications (SPA) with dynamic capabilities.

## Core Technologies

-   **Go**: The primary language for backend development
-   **Gorm**: Object-Relational Mapping (ORM) for database interactions
-   **Gin**: High-performance HTTP web framework
-   **Templ**: Typed templating language for Go
-   **Typed-HTMx-Go**: Type-safe Go bindings for HTMX
-   **TypeScript**: For additional frontend logic
-   **ESBuild**: Fast JavaScript bundler and minifier

## Key Features

-   File-based routing system similar to NextJS
-   Integrated CLI for project scaffolding and management
-   Built-in development server with hot reloading
-   Streamlined production build process
-   Type-safe templating with Templ
-   Efficient asset management and bundling with ESBuild
-   Database integration with Gorm
-   HTMX integration for dynamic frontend updates without full page reloads

## Quick Start

To create a new GoPostulate project:

```bash
go run cmd/gopostulate/main.go new --name my-project
cd my-project
go run cmd/app/main.go
```

## Project Structure

GoPostulate provides a standardized project structure to help you organize your code effectively:

```
my-project/
├── cmd/
├── internal/
├── web/
├── migrations/
├── app/
├── tests/
├── go.mod
└── .gitignore
```

For more details on the project structure and how to use GoPostulate, please refer to our [Getting Started Guide](docs/getting-started.md).

## Documentation

For more detailed information, please refer to the documentation in the `docs/` directory:

-   [Getting Started](docs/getting-started.md)
-   [Routing](docs/routing.md)
-   [Configuration](docs/configuration.md)

## Contributing

We welcome contributions! Please see our [Contributing Guide](docs/contributing.md) for more details.

## License

GoPostulate is released under the MIT License.
