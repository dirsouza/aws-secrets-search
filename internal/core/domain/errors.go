package domain

import "fmt"

// ErrMissingConfig indica que uma configuração obrigatória está ausente.
type ErrMissingConfig struct {
	Field string
}

func (e *ErrMissingConfig) Error() string {
	return fmt.Sprintf("missing required configuration: %s", e.Field)
}

// ErrAWSConnection indica falha na conexão com a AWS.
type ErrAWSConnection struct {
	Cause error
}

func (e *ErrAWSConnection) Error() string {
	return fmt.Sprintf("AWS connection failed: %v", e.Cause)
}

func (e *ErrAWSConnection) Unwrap() error {
	return e.Cause
}
