.PHONY: run build build-all build-linux build-windows build-macos clean help

# Variáveis
BINARY_NAME=aws-secrets-search
MAIN_PATH=./cmd/cli
BUILD_DIR=./build

# Arquitetura (pode ser sobrescrita: make build-linux ARCH=arm64)
ARCH?=amd64

# Ajuda
help:
	@echo "\nComandos disponíveis:\n"
	@echo "  make help         - Mostra esta mensagem"
	@echo "  make run          - Executa a aplicação sem compilar"
	@echo "  make build        - Compila o binário para o SO atual"
	@echo "  make build-all    - Compila para Linux, Windows e macOS (todas as arquiteturas)"
	@echo "  make build-linux  - Compila para Linux (padrão: amd64)"
	@echo "  make build-windows- Compila para Windows (padrão: amd64)"
	@echo "  make build-macos  - Compila para macOS (padrão: amd64)"
	@echo "  make clean        - Remove arquivos compilados"
	@echo "\nPersonalizar arquitetura:"
	@echo "  make build-linux ARCH=arm64"
	@echo "  make build-windows ARCH=arm64"
	@echo "  make build-macos ARCH=arm64"
	@echo "\nArquiteturas suportadas: amd64, arm64, 386, arm"

# Executar sem compilar
run:
	@echo "🚀 Executando aplicação..."
	@go run $(MAIN_PATH)

# Compilar o binário para o SO atual
build: clean
	@echo "🔨 Compilando para o sistema atual..."
	@go build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "✓ Compilado com sucesso: $(BINARY_NAME)"

# Compilar para todas as plataformas (com todas as arquiteturas comuns)
build-all: clean
	@echo "🔨 Compilando para todas as plataformas..."
	@mkdir -p $(BUILD_DIR)
	@echo "🐧 Linux..."
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	@GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(MAIN_PATH)
	@echo "🪟  Windows..."
	@GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)
	@GOOS=windows GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-arm64.exe $(MAIN_PATH)
	@echo "🍎 macOS..."
	@GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-macos-amd64 $(MAIN_PATH)
	@GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-macos-arm64 $(MAIN_PATH)
	@echo "\n✓ Compilação completa! 6 binários gerados:"
	@echo "  • Linux: amd64, arm64"
	@echo "  • Windows: amd64, arm64"
	@echo "  • macOS: amd64 (Intel), arm64 (Apple Silicon)"
	@echo "📦 Binários disponíveis em: $(BUILD_DIR)/"

# Linux
build-linux:
	@echo "🐧 Compilando para Linux ($(ARCH))..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=$(ARCH) go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-$(ARCH) $(MAIN_PATH)
	@echo "✓ $(BUILD_DIR)/$(BINARY_NAME)-linux-$(ARCH)"

# Windows
build-windows:
	@echo "🪟  Compilando para Windows ($(ARCH))..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=windows GOARCH=$(ARCH) go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-$(ARCH).exe $(MAIN_PATH)
	@echo "✓ $(BUILD_DIR)/$(BINARY_NAME)-windows-$(ARCH).exe"

# macOS
build-macos:
	@echo "🍎 Compilando para macOS ($(ARCH))..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=darwin GOARCH=$(ARCH) go build -o $(BUILD_DIR)/$(BINARY_NAME)-macos-$(ARCH) $(MAIN_PATH)
	@echo "✓ $(BUILD_DIR)/$(BINARY_NAME)-macos-$(ARCH)"

# Limpar arquivos compilados
clean:
	@echo "🧹 Limpando..."
	@rm -f $(BINARY_NAME)
	@rm -rf $(BUILD_DIR)
	@echo "✓ Arquivos removidos"
