// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : structs.go
// Original timestamp : 2024/08/01 13:56

// This is where we create the framework around the policies:
// - the structure
// - the constructor

package policies

import "github.com/hashicorp/vault/api"

const (
	ReadOnlyPolicy = `
path "{{.Path}}" {
  capabilities = ["read", "list"]
}`

	ReadWritePolicy = `
path "{{.Path}}" {
  capabilities = ["create", "read", "update", "delete", "list"]
}`

	DenyAllPolicy = `
path "{{.Path}}" {
  capabilities = []
}`
)

type Policy struct {
	client *api.Client
}

func NewPolicyManager(client *api.Client) *Policy {
	return &Policy{client: client}
}
