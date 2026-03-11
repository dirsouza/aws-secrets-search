package port

import (
	"context"

	"github.com/cliquefarma/aws-secrets-search/internal/core/domain"
)

// SecretReader é a porta de saída (driven) para leitura de secrets.
// Quem implementa essa interface é o adaptador de infraestrutura (ex: AWS).
type SecretReader interface {
	FetchAll(ctx context.Context) ([]domain.Secret, error)
}
