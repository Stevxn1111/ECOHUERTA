# ECOHUERTA
cat << 'EOF' > README.md
# ECOHUERTA

Aplicación web para la gestión y consulta de flora por departamentos de Colombia, desarrollada con Go, MariaDB y frontend web.

## Requisitos
- **Linux:** Docker Engine y Docker Compose v2.
- **Windows:** Docker Desktop (con WSL 2 habilitado).

## Cómo ejecutar el proyecto

1. Clonar o descomprimir el proyecto y abrir una terminal en la carpeta raíz (`ECOHUERTA-main`).
2. Crear el archivo de variables de entorno a partir de la plantilla:
   - **Linux / macOS:**
     ```bash
     cp .env.example .env
     ```
   - **Windows (PowerShell):**
     ```powershell
     Copy-Item .env.example .env
     ```
3. Construir y levantar los contenedores:
   ```bash
   docker compose up -d --build

4. Abrir en el navegador:http://localhost:8080/Login/login.html