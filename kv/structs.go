// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : structs.go
// Original timestamp : 2024/08/01 18:52

package kv

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"strconv"
	"strings"
)

type KVManager struct {
	client *api.Client
}

func NewKVManager(client *api.Client) *KVManager {
	return &KVManager{client: client}
}

func (km *KVManager) getEngineVersion(path string) (int, error) {
	mounts, err := km.client.Sys().ListMounts()
	if err != nil {
		return 0, err
	}

	for mountPath, mount := range mounts {
		if strings.HasPrefix(path, mountPath) {
			if versionStr, ok := mount.Options["version"]; ok {
				version, err := strconv.Atoi(versionStr)
				if err != nil {
					return 0, fmt.Errorf("invalid version format for path: %s", path)
				}
				return version, nil
			}
			return 1, nil // Default to version 1 if no version is specified
		}
	}

	return 0, fmt.Errorf("KV engine not found for path: %s", path)
}
