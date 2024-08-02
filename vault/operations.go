// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : operations.go
// Original timestamp : 2024/08/01 19:41

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/vault/api"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"os"
)

func (vm *VaultManager) UnsealVault(client *api.Client, jsonFilePath string, requiredKeys int) *cerr.CustomError {
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

func (vm *VaultManager) BackupVault(path string) *cerr.CustomError {
	resp, err := vm.client.Logical().Write("sys/storage/raft/snapshot", nil)
	if err != nil {
		return &cerr.CustomError{Title: "Error writing snapshot", Message: err.Error()<}
	}

	snapshot, ok := resp.Data["snapshot"].(string)
	if !ok {
		return &cerr.CustomError{Title: fmt.Sprintf("Unexpected response format: %v", resp.Data),
			Message: err.Error()}
	}

	if err := os.WriteFile(path, []byte(snapshot), 0644); err != nil {
		return &cerr.CustomError{Title: fmt.Sprintf("Error writing %s", path), Message: err.Error()}
	}
	return nil
}

func (vm *VaultManager) RestoreVault(path string) *cerr.CustomError {
	data, err := os.ReadFile(path)
	if err != nil {
		return &cerr.CustomError{Title: fmt.Sprintf("Error reading file %s", path), Message: err.Error()}
	}

	req := map[string]interface{}{
		"snapshot": string(data),
	}

	if _, err = vm.client.Logical().Write("sys/storage/raft/snapshot/restore", req); err != nil {
		return &cerr.CustomError{Title: "Restoring", Message: err.Error()}
	}
	return nil
}
