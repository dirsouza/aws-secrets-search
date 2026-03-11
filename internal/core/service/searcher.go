package service

import (
	"context"
	"strings"

	"github.com/cliquefarma/aws-secrets-search/internal/core/domain"
	"github.com/cliquefarma/aws-secrets-search/internal/core/port"
)

// SecretSearcher implementa a porta SearchService.
// Contém a lógica de negócio pura, sem conhecer detalhes de infraestrutura.
type SecretSearcher struct {
	reader    port.SecretReader
	presenter port.Presenter
}

// NewSecretSearcher cria o serviço injetando as portas de saída.
func NewSecretSearcher(reader port.SecretReader, presenter port.Presenter) *SecretSearcher {
	return &SecretSearcher{reader: reader, presenter: presenter}
}

// Search executa a busca agrupada por termo e retorna os resultados.
func (s *SecretSearcher) Search(ctx context.Context, rawTerms string) ([]domain.SearchResult, error) {
	terms, err := parseTerms(rawTerms)
	if err != nil {
		return nil, err
	}

	secrets, err := s.reader.FetchAll(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]domain.SearchResult, 0, len(terms))

	for i, term := range terms {
		s.presenter.RenderTermStart(term.Raw)
		result := s.matchTerm(term, secrets)
		s.presenter.RenderTermSummary(&result)
		results = append(results, result)

		if i < len(terms)-1 {
			s.presenter.RenderSeparator()
		}
	}

	totalMatches := countMatches(results)
	s.presenter.RenderFinalSummary(totalMatches)

	return results, nil
}

// matchTerm busca um termo em todos os secrets.
func (s *SecretSearcher) matchTerm(term domain.SearchTerm, secrets []domain.Secret) domain.SearchResult {
	result := *domain.NewSearchResult(term.Raw)

	for _, secret := range secrets {
		if term.MatchesIn(secret.Value) {
			result.Add(secret.Name)
			s.presenter.RenderMatch(secret.Name)
		}
	}

	return result
}

// parseTerms converte a string de entrada em uma lista de SearchTerm.
func parseTerms(raw string) ([]domain.SearchTerm, error) {
	parts := strings.Split(raw, ";")
	terms := make([]domain.SearchTerm, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		term, err := domain.NewSearchTerm(part)
		if err != nil {
			return nil, err
		}
		terms = append(terms, term)
	}

	return terms, nil
}

// countMatches soma o total de matches de todos os resultados.
func countMatches(results []domain.SearchResult) int {
	total := 0
	for _, r := range results {
		total += r.Count()
	}
	return total
}
