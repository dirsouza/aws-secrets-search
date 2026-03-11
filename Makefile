.PHONY: run build clean help

# Variáveis
BINARY_NAME=aws-secrets-search
MAIN_PATH=./cmd/cli

# Ajuda
help:
	@echo "\nComandos disponíveis:\n"
	@echo "  make help   - Mostra esta mensagem"
	@echo "  make run    - Executa a aplicação sem compilar"
	@echo "  make build  - Compila o binário"
	@echo "  make clean  - Remove arquivos compilados"

# Executar sem compilar
run:
	@echo "🚀 Executando aplicação..."
	@go run $(MAIN_PATH)

# Compilar o binário
build:
	@echo "🔨 Compilando..."
	@go build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "✓ Compilado com sucesso: $(BINARY_NAME)"

# Limpar arquivos compilados
clean:
	@echo "🧹 Limpando..."
	@rm -f $(BINARY_NAME)
	@echo "✓ Arquivos removidos"
