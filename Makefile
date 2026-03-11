.PHONY: run build build-all build-linux build-windows build-macos clean help

# Variáveis
BINARY_NAME=aws-secrets-search
MAIN_PATH=./cmd/cli
BUILD_DIR=./build

# Ajuda
help:
	@echo "\nComandos disponíveis:\n"
	@echo "  make help         - Mostra esta mensagem"
	@echo "  make run          - Executa a aplicação sem compilar"
	@echo "  make build        - Compila o binário para o SO atual"
	@echo "  make build-all    - Compila para Linux, Windows e macOS"
	@echo "  make build-linux  - Compila apenas para Linux"
	@echo "  make build-windows- Compila apenas para Windows"
	@echo "  make build-macos  - Compila apenas para macOS"
	@echo "  make clean        - Remove arquivos compilados"

# Executar sem compilar
run:
	@echo "🚀 Executando aplicação..."
	@go run $(MAIN_PATH)

# Compilar o binário para o SO atual
build: clean
	@echo "🔨 Compilando para o sistema atual..."
	@go build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "✓ Compilado com sucesso: $(BINARY_NAME)"

# Compilar para todas as plataformas
build-all: clean build-linux build-windows build-macos
	@echo "\n✓ Compilação completa para todas as plataformas!"
	@echo "📦 Binários disponíveis em: $(BUILD_DIR)/"

# Linux
build-linux:
	@echo "🐧 Compilando para Linux..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	@echo "✓ $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64"

# Windows
build-windows:
	@echo "🪟  Compilando para Windows..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)
	@GOOS=windows GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-arm64.exe $(MAIN_PATH)
	@echo "✓ $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe (Intel/AMD)"
	@echo "✓ $(BUILD_DIR)/$(BINARY_NAME)-windows-arm64.exe (ARM)"

# macOS
build-macos:
	@echo "🍎 Compilando para macOS..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-macos-amd64 $(MAIN_PATH)
	@GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-macos-arm64 $(MAIN_PATH)
	@echo "✓ $(BUILD_DIR)/$(BINARY_NAME)-macos-amd64 (Intel)"
	@echo "✓ $(BUILD_DIR)/$(BINARY_NAME)-macos-arm64 (Apple Silicon)"

# Limpar arquivos compilados
clean:
	@echo "🧹 Limpando..."
	@rm -f $(BINARY_NAME)
	@rm -rf $(BUILD_DIR)
	@echo "✓ Arquivos removidos"
