// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// auth/auth.go :
// Original file timestamp: 2024.07.30 22:15:37

package auth

import (
	"github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
)

// GetAuthMethodMountPaths retrieves the mount paths for all enabled auth methods in Vault.
func GetAuthMethodMountPaths(client *api.Client) (map[string]string, *cerr.CustomError) {
	authMethods := make(map[string]string)

	// List enabled auth methods
	authList, err := client.Sys().ListAuth()
	if err != nil {
		return nil, &cerr.CustomError{Title: "failed to list auth methods", Message: err.Error()}
	}

	// Populate the map with the auth type and corresponding mount path
	for path, auth := range authList {
		authMethods[auth.Type] = path
	}

	return authMethods, nil
}
