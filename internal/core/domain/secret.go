package domain

import "regexp"

// Secret representa uma secret com nome e valor.
type Secret struct {
	Name  string
	Value string
}

// SearchTerm encapsula um termo de busca com regex compilada.
type SearchTerm struct {
	Raw     string
	pattern *regexp.Regexp
}

// NewSearchTerm cria um SearchTerm com regex case-insensitive.
func NewSearchTerm(raw string) (SearchTerm, error) {
	pattern, err := regexp.Compile(`(?i)` + regexp.QuoteMeta(raw) + `\S*`)
	if err != nil {
		return SearchTerm{}, err
	}
	return SearchTerm{Raw: raw, pattern: pattern}, nil
}

// MatchesIn verifica se o termo e encontrado no conteudo.
func (st SearchTerm) MatchesIn(content string) bool {
	return st.pattern.MatchString(content)
}
