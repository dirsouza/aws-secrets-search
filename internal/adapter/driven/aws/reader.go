package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/cliquefarma/aws-secrets-search/internal/core/domain"
)

// SecretManagerReader é o adaptador driven que implementa port.SecretReader
// usando o AWS Secrets Manager como fonte de dados.
type SecretManagerReader struct {
	client *secretsmanager.Client
}

// NewSecretManagerReader cria o adaptador AWS.
func NewSecretManagerReader(client *secretsmanager.Client) *SecretManagerReader {
	return &SecretManagerReader{client: client}
}

// FetchAll busca todos os secrets paginados e retorna com seus valores.
func (r *SecretManagerReader) FetchAll(ctx context.Context) ([]domain.Secret, error) {
	secrets := make([]domain.Secret, 0)
	var nextToken *string

	for {
		response, err := r.client.ListSecrets(ctx, &secretsmanager.ListSecretsInput{
			NextToken: nextToken,
		})
		if err != nil {
			return nil, err
		}

		for _, entry := range response.SecretList {
			secret, err := r.fetchValue(ctx, entry.Name)
			if err != nil {
				continue
			}
			secrets = append(secrets, secret)
		}

		if response.NextToken == nil {
			break
		}
		nextToken = response.NextToken
	}

	return secrets, nil
}

// fetchValue obtém o valor de um secret específico.
func (r *SecretManagerReader) fetchValue(ctx context.Context, name *string) (domain.Secret, error) {
	response, err := r.client.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: name,
	})
	if err != nil {
		return domain.Secret{}, err
	}

	return domain.Secret{
		Name:  *name,
		Value: *response.SecretString,
	}, nil
}
