package domain

// SearchResult agrupa os resultados de uma busca por um termo.
type SearchResult struct {
	Term    string
	Secrets []string
}

// NewSearchResult cria um resultado vazio para o termo.
func NewSearchResult(term string) *SearchResult {
	return &SearchResult{Term: term, Secrets: make([]string, 0)}
}

// Add registra o nome de uma secret encontrada.
func (r *SearchResult) Add(secretName string) {
	r.Secrets = append(r.Secrets, secretName)
}

// Count retorna a quantidade de secrets encontradas.
func (r *SearchResult) Count() int {
	return len(r.Secrets)
}

// HasMatches indica se houve correspondências.
func (r *SearchResult) HasMatches() bool {
	return len(r.Secrets) > 0
}
