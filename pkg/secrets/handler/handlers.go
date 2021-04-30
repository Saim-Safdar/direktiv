package handler

import (
	"fmt"
	"strings"
)

type SecretType string

const (
	DBSecretType SecretType = "string"
)

var (
	secretHandlers = map[SecretType]SecretsHandlerInstantiator{
		DBSecretType: setupDB,
	}
)

type SecretsHandlerInstantiator func() (SecretsHandler, error)

type SecretsHandler interface {
	AddSecret(namespace, name string, secret []byte) error
	RemoveSecret(namespace, name string) error
	RemoveSecrets(namespace string) error
	GetSecret(namespace, name string) ([]byte, error)
	GetSecrets(namespace string) ([]string, error)
}

func (x *SecretType) Setup() (SecretsHandler, error) {
	return secretHandlers[*x]()
}

func ParseType(s string) (SecretType, error) {
	original := s

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	f := SecretType(s)
	if _, ok := secretHandlers[f]; !ok {
		return "db", fmt.Errorf("unrecognized secrets backend type '%s'", original)
	}

	return f, nil
}

func ListRegisteredTypes() []SecretType {
	stList := make([]SecretType, 0)
	for k := range secretHandlers {
		stList = append(stList, k)
	}

	return stList
}

// String returns a string representation of the SecretType.
func (x SecretType) String() string {
	return string(x)
}

func RegisterNewSecretType(newType SecretType, handlerFunc SecretsHandlerInstantiator) error {
	if _, exists := secretHandlers[newType]; exists {
		return fmt.Errorf("refusing to register secret backend type '%s': already registered", newType)
	}

	secretHandlers[newType] = handlerFunc
	return nil
}
