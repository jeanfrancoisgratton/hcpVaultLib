// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : init.go
// Original timestamp : 2024/08/01 13:56

// This is where we create the framework around the policies:
// - the structure
// - the constructor

package policies

import (
	"github.com/hashicorp/vault/api"
)

type PolicyManager struct {
	client *api.Client
}

func NewPolicyManager(client *api.Client) *PolicyManager {
	return &PolicyManager{client: client}
}
