// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// /unseal.go : Unseal the Vault (typically done after service restart)
// Original file timestamp: 2024.07.29 21:04:59

package hcpVaultLib

import (
	"encoding/json"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"log"
	"os"
)

type UnsealKey struct {
	Keys []string `json:"keys"`
}

func UnsealVault(client *vault.Client, jsonFilePath string, requiredKeys int) error {
	// Read the JSON file
	fileContent, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %w", err)
	}

	// Decrypt the fields if needed
	// You can add your decryption logic here

	// Parse the JSON content
	var unsealKeys UnsealKey
	if err := json.Unmarshal(fileContent, &unsealKeys); err != nil {
		return fmt.Errorf("failed to parse JSON file: %w", err)
	}

	// Unseal the Vault
	for i, key := range unsealKeys.Keys {
		if i >= requiredKeys {
			break
		}
		resp, err := client.Sys().Unseal(key)
		if err != nil {
			return fmt.Errorf("failed to unseal vault with key part %d: %w", i+1, err)
		}
		if resp.Sealed == false {
			log.Println("Vault is successfully unsealed")
			return nil
		}
	}

	return fmt.Errorf("failed to unseal the Vault with the provided key parts")
}
