# Go REST API - gousers

Este es un proyecto de API REST desarrollado en Go utilizando el framework Gin y GORM para la persistencia de datos. El proyecto incluye documentación con Swagger y está completamente contenedorizado con Docker.

## Requisitos

- Docker y Docker Compose
- Go 1.24 (opcional, para desarrollo local sin Docker)

## Estructura del Proyecto

- `cmd/api/main.go`: Punto de entrada de la aplicación.
- `internal/`: Lógica interna (modelos, handlers, base de datos).
- `docs/`: Documentación Swagger autogenerada.
- `database/`: Scripts SQL para la inicialización de la base de datos.

## Ejecución con Docker

El proyecto está configurado para ejecutarse en el puerto **8084** dentro de la red **bianquiviri_network**. Se asume que el contenedor de la base de datos ya está activo en esta red.

1. Asegúrate de que la red externa exista y el contenedor de BD esté corriendo:
   ```bash
   docker network create bianquiviri_network
   ```

2. Levanta el contenedor de la API:
   ```bash
   docker compose up -d
   ```

Esto iniciará:
- La API en un contenedor llamado `gousers_api`.

## Documentación API (Swagger)

Una vez que el proyecto esté en ejecución, puedes acceder a la documentación interactiva en:
[http://localhost:8084/swagger/index.html](http://localhost:8084/swagger/index.html)

## Endpoints Principales

- `GET /api/v1/users`: Obtener todos los usuarios.
- `POST /api/v1/users`: Crear un nuevo usuario.
- `GET /api/v1/users/:id`: Obtener un usuario por ID.

## Desarrollo Local

Si deseas ejecutarlo localmente sin Docker:
1. Configura tu `.env` con las credenciales de tu base de datos local.
2. Ejecuta:
   ```bash
   go run cmd/api/main.go
   ```

## Tests

Para ejecutar los tests unitarios localmente:
```bash
go test ./...
```

Para ejecutar los tests **dentro del contenedor** de la API:
```bash
docker exec gousers_api go test ./...
```
