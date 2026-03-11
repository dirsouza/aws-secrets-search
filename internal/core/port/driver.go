package port

import (
	"context"

	"github.com/cliquefarma/aws-secrets-search/internal/core/domain"
)

// SearchService é a porta de entrada (driver) que expõe os casos de uso.
// Quem consome essa interface é o adaptador primário (ex: CLI).
type SearchService interface {
	Search(ctx context.Context, rawTerms string) ([]domain.SearchResult, error)
}

// Presenter é a porta de saída (driven) para apresentação de resultados.
// Quem implementa essa interface é o adaptador de apresentação (ex: CLI colorido).
type Presenter interface {
	RenderMatch(secretName string)
	RenderTermStart(term string)
	RenderTermSummary(result *domain.SearchResult)
	RenderSeparator()
	RenderFinalSummary(totalMatches int)
	RenderWarning(message string)
	RenderError(message string, hints []string)
}
