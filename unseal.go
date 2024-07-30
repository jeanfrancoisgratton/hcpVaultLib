// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// /unseal.go : Unseal the Vault (typically done after service restart)
// Original file timestamp: 2024.07.29 21:04:59

package hcpVaultLib

import (
	"encoding/json"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"os"
)

type VaultOperatorInfo struct {
	Keys []string `json:"keys"`
}

func UnsealVault(client *vault.Client, jsonFilePath string, requiredKeys int) *cerr.CustomError {
	// Read the JSON file
	fileContent, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return &cerr.CustomError{Title: "Failed to read JSON file", Message: err.Error()}
	}

	// Decrypt the fields if needed
	// You can add your decryption logic here

	// Parse the JSON content
	var unsealKeys VaultOperatorInfo
	if err := json.Unmarshal(fileContent, &unsealKeys); err != nil {
		return &cerr.CustomError{Title: "Failed to parse JSON file", Message: err.Error()}
	}

	// Unseal the Vault
	for i, key := range unsealKeys.Keys {
		key := hf.DecodeString(key, "")
		if i >= requiredKeys {
			break
		}
		resp, err := client.Sys().Unseal(key)
		// the return statement is commented for now: do we really need to return if a key part is wrong ? We might need a boolean to deal with this
		// something like var analSecurity bool, where if true we uncomment the lines below.
		// otherwise: we keep the loop going, it'll fail if all key parts have been processed, anyway
		if err != nil {
			//return &cerr.CustomError{Fatality: cerr.Continuable, Title: fmt.Sprintf("Failed to unseal vault with key part %d", i+1), Message: err.Error()}
			continue // <-- for now.. Remove this line if we use that boolean var I've mentioned, and uncomment the above
		}
		if resp.Sealed == false {
			fmt.Printf("%s\n", hf.Green("Vault is unsealed"))
			return nil
		}
	}

	return &cerr.CustomError{Title: "failed to unseal the Vault with the provided key parts"}
}
