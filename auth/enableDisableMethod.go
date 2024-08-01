// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// auth/enableDisableMethod.go :
// Original file timestamp: 2024.07.31 06:03:36

package auth

import (
	"github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
)

func (am *AuthManager) EnableAuthMethod(path, method string) *cerr.CustomError {
	options := &api.EnableAuthOptions{Type: method}
	if err := am.client.Sys().EnableAuthWithOptions(path, options); err != nil {
		return &cerr.CustomError{Title: "Unable to enable the Auth method", Message: err.Error()}
	}
	return nil
}

func (am *AuthManager) DisableAuthMethod(path string) *cerr.CustomError {
	if err := am.client.Sys().DisableAuth(path); err != nil {
		return &cerr.CustomError{Title: "Unable to disable the Auth method", Message: err.Error()}
	}
	return nil
}

func (am *AuthManager) ListAuthMethods() (map[string]*api.AuthMount, *cerr.CustomError) {
	m, err := am.client.Sys().ListAuth()
	if err != nil {
		return nil, &cerr.CustomError{Title: "Unable to list Auth methods", Message: err.Error()}
	}
	return m, nil
}
