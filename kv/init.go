// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : init.go
// Original timestamp : 2024/08/01 18:52

package kv

import (
	"github.com/hashicorp/vault/api"
)

type KVManager struct {
	client *api.Client
}

func NewKVManager(client *api.Client) *KVManager {
	return &KVManager{client: client}
}
