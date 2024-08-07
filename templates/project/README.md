# GoPostulate

GoPostulate is a full-stack web application framework inspired by NextJS, leveraging the power of Go and complementary technologies. It is designed to facilitate the rapid development of performant, Server-Side Rendered (SSR) Single Page Applications (SPA) with dynamic capabilities.

## Core Technologies

- Go: The primary language for backend development
- Gorm: Object-Relational Mapping (ORM) for database interactions
- Gin: High-performance HTTP web framework
- Templ: Typed templating language for Go
- Typed-HTMx-Go: Type-safe Go bindings for HTMX
- TypeScript: For additional frontend logic
- ESBuild: Fast JavaScript bundler and minifier

## Key Features

- File-based routing system similar to NextJS
- Integrated CLI for project scaffolding and management
- Built-in development server with hot reloading
- Streamlined production build process
- Type-safe templating with Templ
- Efficient asset management and bundling with ESBuild
- Database integration with Gorm
- HTMX integration for dynamic frontend updates without full page reloads

## Project Structure

```
gopostulate/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── handlers/
│   │   └── .gitkeep
│   ├── models/
│   │   └── .gitkeep
├── web/
│   ├── components/
│   │   ├── Header/
│   │   │   ├── header.templ
│   │   │   ├── header.ts
│   │   │   ├── header.css
│   │   └── .gitkeep
│   ├── pages/
│   │   ├── index.templ
│   │   ├── index.ts
│   │   ├── index.css
│   │   └── .gitkeep
│   ├── static/
│   │   ├── css/
│   │   │   └── main.css
│   │   ├── js/
│   │   │   └── main.js
│   │   └── .gitkeep
│   ├── templates/
│   │   └── .gitkeep
├── migrations/
│   └── .gitkeep
├── app/
│   ├── controllers/
│   │   ├── index.go
│   │   └── .gitkeep
│   ├── routes/
│   │   └── router.go
│   └── models/
│       └── .gitkeep
├── tests/
│   └── .gitkeep
├── go.mod
└── README.md
```
