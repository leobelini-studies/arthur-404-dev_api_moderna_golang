 #!/bin/bash

# Nome do seu programa Go
APP_NAME="go_opportunities"

# Diretório de saída para os binários compilados
OUTPUT_DIR="builds"

# Verificar se o diretório de saída existe, se não, criá-lo
mkdir -p "$OUTPUT_DIR"

# Compilar para Windows 64-bit
echo "Compilando para Windows 64-bit..."
GOOS=windows GOARCH=amd64 go build -o "$OUTPUT_DIR/$APP_NAME-windows-amd64.exe"

# Compilar para Windows ARM
echo "Compilando para Windows ARM..."
GOOS=windows GOARCH=arm go build -o "$OUTPUT_DIR/$APP_NAME-windows-arm.exe"

# Compilar para Linux 64-bit
echo "Compilando para Linux 64-bit..."
GOOS=linux GOARCH=amd64 go build -o "$OUTPUT_DIR/$APP_NAME-linux-amd64"

# Compilar para Linux ARM
echo "Compilando para Linux ARM..."
GOOS=linux GOARCH=arm go build -o "$OUTPUT_DIR/$APP_NAME-linux-arm"

# Compilar para macOS 64-bit
echo "Compilando para macOS 64-bit..."
GOOS=darwin GOARCH=amd64 go build -o "$OUTPUT_DIR/$APP_NAME-darwin-amd64"

# Compilar para macOS ARM
echo "Compilando para macOS ARM..."
GOOS=darwin GOARCH=arm64 go build -o "$OUTPUT_DIR/$APP_NAME-darwin-arm64"

echo "Compilação concluída. Binários estão em $OUTPUT_DIR."
