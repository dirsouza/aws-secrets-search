package cli

import (
	"context"
	"log"
	"os"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	awsadapter "github.com/cliquefarma/aws-secrets-search/internal/adapter/driven/aws"
	"github.com/cliquefarma/aws-secrets-search/internal/core/domain"
	"github.com/cliquefarma/aws-secrets-search/internal/core/service"
	"github.com/joho/godotenv"
)

// App é o adaptador driver que orquestra a execução via CLI.
// Faz a composição das dependências (Composition Root) e executa o serviço.
type App struct {
	presenter *ColorPresenter
	config    *appConfig
}

type appConfig struct {
	accessKeyID     string
	secretAccessKey string
	sessionToken    string
	region          string
	searchTerms     string
}

// NewApp cria e configura a aplicação CLI.
func NewApp() (*App, error) {
	log.SetFlags(0)

	presenter := NewColorPresenter()

	cfg, err := loadConfig(presenter)
	if err != nil {
		return nil, err
	}

	return &App{presenter: presenter, config: cfg}, nil
}

// Run executa a busca de secrets.
func (a *App) Run(ctx context.Context) error {
	// Constrói o client AWS (adaptador driven)
	awsCfg, err := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithRegion(a.config.region),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			a.config.accessKeyID,
			a.config.secretAccessKey,
			a.config.sessionToken,
		)),
	)
	if err != nil {
		a.presenter.RenderError("Failed to load AWS configuration", []string{err.Error()})
		return &domain.ErrAWSConnection{Cause: err}
	}

	client := secretsmanager.NewFromConfig(awsCfg)

	// Monta o hexágono: porta driven → serviço ← porta driver
	reader := awsadapter.NewSecretManagerReader(client)
	searcher := service.NewSecretSearcher(reader, a.presenter)

	// Executa via porta de entrada
	_, err = searcher.Search(ctx, a.config.searchTerms)
	return err
}

// loadConfig carrega e valida configurações.
func loadConfig(presenter *ColorPresenter) (*appConfig, error) {
	_ = godotenv.Load()

	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		presenter.RenderWarning(".env file not found")
	}

	cfg := &appConfig{
		accessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		secretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		sessionToken:    os.Getenv("AWS_SESSION_TOKEN"),
		region:          envOrDefault("AWS_DEFAULT_REGION", "us-east-1"),
		searchTerms:     os.Getenv("SEARCH_TERMS"),
	}

	if err := validateConfig(cfg, presenter); err != nil {
		return nil, err
	}

	return cfg, nil
}

// validateConfig valida campos obrigatórios e exibe mensagens claras.
func validateConfig(cfg *appConfig, presenter *ColorPresenter) error {
	if cfg.accessKeyID == "" || cfg.secretAccessKey == "" || cfg.sessionToken == "" {
		err := &domain.ErrMissingConfig{Field: "AWS credentials"}
		presenter.RenderError("Missing required AWS credentials", []string{
			"The following environment variables must be set:",
			"• AWS_ACCESS_KEY_ID",
			"• AWS_SECRET_ACCESS_KEY",
			"• AWS_SESSION_TOKEN",
			"💡 Tip: Create a .env file based on .env.example",
		})
		return err
	}

	if cfg.searchTerms == "" {
		err := &domain.ErrMissingConfig{Field: "SEARCH_TERMS"}
		presenter.RenderError("Missing search terms", []string{
			"The SEARCH_TERMS environment variable must be set",
			"💡 Example: SEARCH_TERMS=\"redis;rabbitmq;postgres\"",
		})
		return err
	}

	return nil
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
