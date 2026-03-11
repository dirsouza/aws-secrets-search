# AWS Secrets Search

Ferramenta CLI para buscar termos específicos nos valores das secrets do AWS Secrets Manager.

## O que faz

Busca por múltiplos termos dentro dos **valores** (não apenas nomes) das secrets armazenadas no AWS Secrets Manager, exibindo resultados coloridos e agrupados por termo de busca.

## Pré-requisitos

- Go 1.25.1 ou superior
- Credenciais AWS configuradas
- Acesso ao AWS Secrets Manager

## Instalação

```bash
# Clone o repositório
git clone <repo-url>
cd aws-secrets-search

# Instale as dependências
go mod download
```

## Configuração

Crie um arquivo `.env` na raiz do projeto com base no `.env.example`:

```env
AWS_ACCESS_KEY_ID=sua_access_key
AWS_SECRET_ACCESS_KEY=sua_secret_key
AWS_REGION=us-east-1
SEARCH_TERMS=termo1,termo2,termo3

# Opcional: apenas para credenciais temporárias
# AWS_SESSION_TOKEN=seu_session_token
```

**Variáveis obrigatórias:**

- `AWS_ACCESS_KEY_ID` - Chave de acesso AWS
- `AWS_SECRET_ACCESS_KEY` - Secret de acesso AWS
- `AWS_REGION` - Região AWS (ex: us-east-1)
- `SEARCH_TERMS` - Termos para buscar, separados por vírgula

**Variável opcional:**

- `AWS_SESSION_TOKEN` - Token de sessão (apenas necessário para credenciais temporárias, como ao assumir roles ou usar MFA)

## Uso

### Executar sem compilar

```bash
make run
```

### Compilar para o sistema atual

```bash
make build
./aws-secrets-search
```

### Compilar para múltiplas plataformas

```bash
# Compilar para todas as plataformas (Linux, Windows, macOS - 6 binários)
make build-all

# Ou compilar para plataforma específica com arquitetura personalizada:
make build-linux ARCH=arm64     # Linux ARM64
make build-windows ARCH=arm64   # Windows ARM64
make build-macos ARCH=arm64     # macOS Apple Silicon

# Os binários ficam em ./build/
# Arquiteturas suportadas: amd64 (padrão), arm64, 386, arm
```

### Limpar binários

```bash
make clean
```

### Ver todos os comandos

```bash
make help
```

## Exemplo de Saída

```
🔍 Searching for term: database

  ✓ Found in secret: production/api-config
  ✓ Found in secret: staging/backend-service
  ✓ Found in secret: development/app-settings

  📊 Total: 3 secret(s) found

──────────────────────────────────────────────────────────────────────────────────────────────────────────

🔍 Searching for term: redis-cache

  ⚠ No secrets found

═══════════════════════════════════════════════════════════════
  🎯 Search completed! Total matches: 3
═══════════════════════════════════════════════════════════════
```

## Arquitetura

Implementa **Arquitetura Hexagonal** (Ports & Adapters) seguindo princípios SOLID, Clean Code e Design Patterns.

### Estrutura

```
cmd/cli/                    # Entry point
internal/
  ├── core/                 # Núcleo da aplicação
  │   ├── domain/          # Entidades e regras de negócio
  │   ├── port/            # Interfaces (contratos)
  │   └── service/         # Casos de uso
  └── adapter/             # Adaptadores externos
      ├── driver/          # Adaptadores primários (CLI)
      └── driven/          # Adaptadores secundários (AWS)
```

### Camadas

- **Domain**: Entidades puras (Secret, SearchTerm, SearchResult)
- **Ports**: Interfaces que isolam o core (SearchService, Presenter, SecretReader)
- **Service**: Lógica de negócio (SecretSearcher)
- **Adapters**:
  - **Driver** (CLI): Interface de entrada que inicia a aplicação
  - **Driven** (AWS): Interface de saída para buscar secrets

### Princípios Aplicados

- **SOLID**: Separação de responsabilidades, inversão de dependência
- **DRY**: Eliminação de duplicação
- **KISS**: Simplicidade na implementação
- **YAGNI**: Apenas o necessário

## Dependências

- `github.com/aws/aws-sdk-go-v2` - SDK AWS
- `github.com/fatih/color` - Output colorido
- `github.com/joho/godotenv` - Carregamento de .env
- `golang.org/x/term` - Detecção de largura do terminal

## Tratamento de Erros

A aplicação fornece mensagens de erro claras e acionáveis:

- ❌ Arquivo `.env` não encontrado → instruções para criar
- ❌ Variáveis obrigatórias faltando → lista o que está faltando
- ❌ Falha na conexão AWS → verifica credenciais e região

## Licença

MIT License - Este software é open source e pode ser usado livremente.

**Você pode:**

- ✓ Fazer fork do projeto
- ✓ Modificar o código para uso pessoal

**Você deve:**

- Manter a atribuição de copyright original
- Incluir o aviso de licença em cópias

**Nota:** Modificações no repositório original requerem pull request e aprovação dos mantenedores.
