#!/bin/bash

# Obtener el directorio donde está el script, independientemente de desde dónde se ejecute
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ENV_FILE="$SCRIPT_DIR"
TEMP_ENV="$SCRIPT_DIR/.env.tmp"
EXISTS_ENV=false
ENV="$SCRIPT_DIR/.env"

# Función para mostrar el uso del script
show_usage() {
    echo "Uso: $0 [--with-sql]"
    echo "  Sin parámetros: Solo levanta los contenedores"
    echo "  --with-sql: Borra contenedores existentes, volúmenes y carga el SQL"
}

# Función para verificar que el .env existe y tiene las variables necesarias
check_env() {

    if [ -f "$ENV_FILE/.env" ]; then
        cp "$ENV_FILE/.env" "$TEMP_ENV"
        EXISTS_ENV=true
    fi


    cd $ENV_FILE
    cd ..
    ENV_FILE=$(pwd)/.env

    echo $ENV_FILE

    if [ ! -f "$ENV_FILE" ]; then
        echo "❌ Error: Archivo $ENV_FILE no encontrado"
        exit 1
    fi

    # Verificar variables requeridas
    required_vars=("DB_PASS" "DB_DATABASE" "DB_USER")
    missing_vars=0

    for var in "${required_vars[@]}"; do
        if ! grep -q "^${var}=" "$ENV_FILE"; then
            echo "❌ Error: Variable $var no encontrada en .env"
            missing_vars=1
        fi
    done

    if [ $missing_vars -eq 1 ]; then
        exit 1
    fi

    cp "$ENV_FILE" "$ENV"
    sleep 3
    
}

# Función para limpiar todo y reiniciar
clean_and_restart() {
    echo "🗑️  Deteniendo contenedores existentes..."
    # Cambiar al directorio del script antes de ejecutar docker compose
    cd "$SCRIPT_DIR"
    docker compose --env-file "$ENV_FILE" down

    echo "🗑️  Forzando eliminación de contenedores relacionados..."
    docker rm -f mysql_container 2>/dev/null || true

    echo "🗑️  Eliminando volumen mysql_data..."
    docker volume rm backend_mysql_data 2>/dev/null || true

    # echo " Probando mapeo de variables..."
    # docker compose --env-file "$ENV_FILE" config

    echo "🚀 Iniciando contenedores limpios..."
    docker compose up -d

   

   

    if [ $? -ne 0 ]; then
        echo "❌ Error al iniciar los contenedores"
        exit 1
    fi

    echo "⏳ Esperando que MySQL esté listo..."
    sleep 10
}

# Función para solo levantar los contenedores
start_containers() {
    # Cambiar al directorio del script antes de ejecutar docker compose
    cd "$SCRIPT_DIR"

    # echo " Probando mapeo de variables..."
    # docker compose --env-file "$ENV_FILE" config

    echo "🚀 Iniciando contenedores..."
    
    docker compose up -d


    if [ $? -ne 0 ]; then
        echo "❌ Error al iniciar los contenedores"
        exit 1
    fi
}

# Verificar el .env antes de empezar
check_env

# Verificar argumentos
case "$1" in
    "--with-sql")
        clean_and_restart
        ;;
    "")
        start_containers
        ;;
    "--help"|"-h")
        show_usage
        exit 0
        ;;
    *)
        echo "❌ Parámetro no válido: $1"
        show_usage
        exit 1
        ;;
esac

# Verificar estado final
echo "🔍 Verificando estado de los contenedores..."
if docker compose ps | grep -q "mysql_container.*Up"; then
    echo "✅ Contenedores iniciados correctamente"
    if [ "$EXISTS_ENV" = true ]; then
        cp "$TEMP_ENV" "$ENV"
        rm "$TEMP_ENV"
    fi
    if [ "$EXISTS_ENV" = false ]; then
        rm "$ENV"
    fi
    exit 0
else
    echo "❌ Error: Los contenedores no están ejecutándose correctamente"
    exit 1
fi