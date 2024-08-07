# Getting Started with GoPostulate

## Creating a New Project

To create a new GoPostulate project, use the CLI:

```bash
gopostulate new --name my-project
cd my-project
```

This will create a new directory with the following structure:

```
my-project/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── handlers/
│   │   └── handlers.go
│   └── models/
│       └── models.go
├── web/
│   ├── css/
│   │   └── main.css
│   ├── js/
│   │   └── main.js
│   └── templates/
│       └── index.html
├── migrations/
├── app/
│   ├── controllers/
│   └── models/
├── tests/
├── go.mod
└── .gitignore
```

## Project Structure

- `cmd/app/`: Contains the main application entry point.
- `internal/`: Houses internal application code.
- `web/`: Stores frontend assets and templates.
- `migrations/`: For database migration scripts.
- `app/`: Application-specific code (models, controllers, etc.).
- `tests/`: For your application tests.

## Next Steps

1. Customize your `go.mod` file with your project's module name.
2. Start building your application by adding routes in `internal/handlers/handlers.go`.
3. Create models in `app/models/` and controllers in `app/controllers/`.
4. Add your frontend code in the `web/` directory.
5. Run your application with `go run cmd/app/main.go`.

For more detailed information on using GoPostulate, refer to our [documentation](https://github.com/yourusername/gopostulate/docs).
