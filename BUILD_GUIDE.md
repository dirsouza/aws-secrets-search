# Guia de Execução dos Builds

Este documento explica como executar os binários compilados em cada sistema operacional.

## 1. Compilar para plataformas específicas

### Compilar para todas as plataformas (6 binários)

```bash
make build-all
```

Gera automaticamente:

- Linux: amd64, arm64
- Windows: amd64, arm64
- macOS: amd64 (Intel), arm64 (Apple Silicon)

### Compilar para plataforma específica

**Com arquitetura padrão (amd64):**

```bash
make build-linux
make build-windows
make build-macos
```

**Especificando a arquitetura:**

```bash
# Linux ARM64 (Raspberry Pi, servidores ARM)
make build-linux ARCH=arm64

# Windows ARM64 (Surface Pro X)
make build-windows ARCH=arm64

# macOS Apple Silicon (M1/M2/M3)
make build-macos ARCH=arm64

# Arquiteturas de 32-bit
make build-linux ARCH=386
make build-windows ARCH=386
```

**Arquiteturas suportadas:**

- `amd64` - 64-bit Intel/AMD (padrão)
- `arm64` - 64-bit ARM (Apple Silicon, ARM servers)
- `386` - 32-bit Intel/AMD
- `arm` - 32-bit ARM

## 2. Executar em cada sistema

### 🐧 Linux (amd64)

```bash
# Dê permissão de execução (apenas na primeira vez)
chmod +x ./build/aws-secrets-search-linux-amd64

# Execute
./build/aws-secrets-search-linux-amd64
```

### 🪟 Windows (amd64 - Intel/AMD)

**No PowerShell ou CMD:**

```powershell
.\build\aws-secrets-search-windows-amd64.exe
```

**No Git Bash:**

```bash
./build/aws-secrets-search-windows-amd64.exe
```

### 🪟 Windows (arm64 - ARM)

**No PowerShell ou CMD:**

```powershell
.\build\aws-secrets-search-windows-arm64.exe
```

### 🍎 macOS (Intel - amd64)

```bash
# Dê permissão de execução (apenas na primeira vez)
chmod +x ./build/aws-secrets-search-macos-amd64

# Execute
./build/aws-secrets-search-macos-amd64
```

### 🍎 macOS (Apple Silicon - arm64)

```bash
# Dê permissão de execução (apenas na primeira vez)
chmod +x ./build/aws-secrets-search-macos-arm64

# Execute
./build/aws-secrets-search-macos-arm64
```

## 3. Configuração necessária

**Importante:** Antes de executar em qualquer sistema, você precisa:

1. Criar o arquivo `.env` no mesmo diretório do executável
2. Configurar as variáveis AWS

**Exemplo de `.env`:**

```env
AWS_ACCESS_KEY_ID=sua_access_key
AWS_SECRET_ACCESS_KEY=sua_secret_key
AWS_REGION=us-east-1
SEARCH_TERMS=termo1,termo2,termo3
```

## 4. Distribuição

Para distribuir os binários:

1. Compile: `make build-all`
2. Compartilhe o binário correspondente ao SO do usuário
3. Inclua instruções para criar o `.env`
4. ⚠️ **NUNCA** distribua seu arquivo `.env` com credenciais

## 5. Troubleshooting

### Linux/macOS: "Permission denied"

```bash
chmod +x ./build/aws-secrets-search-*
```

### macOS: "não pode ser aberto porque o desenvolvedor não pode ser verificado"

```bash
# Remova a quarentena do macOS
xattr -d com.apple.quarantine ./build/aws-secrets-search-macos-*
```

### Windows: "O Windows protegeu seu PC"

- Clique em "Mais informações" > "Executar assim mesmo"

### Qualquer sistema: "arquivo .env não encontrado"

- Crie o arquivo `.env` no mesmo diretório do executável
- Configure as variáveis AWS necessárias
