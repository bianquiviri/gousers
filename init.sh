#!/bin/bash

# Colores para la salida
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}Iniciando la aplicación gousers...${NC}"

# 1. Verificar/Crear la red bianquiviri_network
if [ ! "$(docker network ls | grep bianquiviri_network)" ]; then
  echo -e "Creando red ${GREEN}bianquiviri_network${NC}..."
  docker network create bianquiviri_network
else
  echo -e "La red ${GREEN}bianquiviri_network${NC} ya existe."
fi

# 2. Levantar los contenedores
echo -e "Levantando el contenedor de la API con Docker Compose..."
docker compose up -d --build

# 3. Mostrar estado y URL de Swagger
echo -e "\n${GREEN}¡Aplicación inicializada con éxito!${NC}"
echo -e "La API está corriendo en: ${BLUE}http://localhost:8084${NC}"
echo -e "Documentación Swagger: ${BLUE}http://localhost:8084/swagger/index.html${NC}"
echo -e "\nContenedores activos:"
docker ps --filter name=gousers_api
