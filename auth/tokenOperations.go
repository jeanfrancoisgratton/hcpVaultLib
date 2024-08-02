// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : tokenOperations.go
// Original timestamp : 2024/08/02 10:31

package auth

import (
	"github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
)

// CreateToken creates a new token with the given parameters
func (am *AuthManager) CreateToken(policies []string, ttl string) (*api.Secret, *cerr.CustomError) {
	tokenData := map[string]interface{}{
		"policies": policies,
		"ttl":      ttl,
	}
	secret, err := am.client.Logical().Write("auth/token/create", tokenData)
	if err != nil {
		return nil, &cerr.CustomError{Title: "Unable to create auth token", Message: err.Error()}
	}

	return secret, nil
}

// DeleteToken deletes a token by its accessor
func (am *AuthManager) DeleteToken(accessor string) error {
	path := "auth/token/revoke-accessor/" + accessor
	_, err := am.client.Logical().Write(path, nil)
	return err
}
