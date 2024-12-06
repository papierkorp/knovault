# Knovault

Knovault is a flexible, plugin-based knowledge management system written in Go. It features a customizable theme system, plugin architecture, and markdown support.

## Features

- 🔌 Plugin System - Extensible architecture for adding new functionality
- 🎨 Theme Support - Customizable appearance with swappable themes
- 📝 Markdown Support - Native markdown parsing and rendering
- 🔄 Hot Reload - Development environment with automatic reloading
- 🚀 HTMX Integration - Modern, minimal JavaScript approach
- 📱 Responsive Design - Mobile-friendly interface

## Prerequisites

- Docker (for development)
- Go 1.22 or higher (for local development)
- Make

## Quick Start

1. **Clone the repository**

   ```bash
   git clone [repository-url]
   cd knovault
   ```

2. **Development with Docker (Recommended)**

   ```bash
   make docker-dev-build
   make docker-dev-run
   ```

3. **Local Development**
   ```bash
   make dev
   ```

The application will be available at `http://localhost:1323`

## Project Structure

```
.
├── cmd/                    # Application entry point
├── internal/
│   ├── assetManager/      # Plugin and theme management
│   ├── globals/           # Global variables and state
│   ├── server/            # HTTP server and routing
│   └── types/             # Type definitions and interfaces
├── data/                  # Content storage
├── static/                # Static assets
└── docker/                # Docker configuration
```

## Development

See [Developer Quick Start Guide](docs/dev-quickstart.md) for detailed development instructions.

For creating new plugins or themes, refer to:

- [Asset Manager Documentation](internal/assetManager/README.md)

## Docker Development Environment

The project includes a development-focused Docker environment that provides:

- Live code reloading
- Template generation
- Dependency management
- Development tools

See [Development Environment Documentation](docs/docker-dev.md) for details.

## Built-in Plugins

- **CustomCSS**: Custom styling support
- **FileManager**: File system operations
- **MarkdownParser**: Markdown processing
- **ThemeChanger**: Theme switching functionality

## Built-in Themes

- **defaultTheme**: Default application appearance
- **dark**: Dark mode theme

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to your branch
5. Create a Pull Request

## License

[License Type] - See LICENSE file for details
