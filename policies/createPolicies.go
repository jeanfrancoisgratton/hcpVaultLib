// hcpVaultLib
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : createPolicies.go
// Original timestamp : 2024/08/01 14:08

package policies

import "fmt"

func (pm *Policy) CreateReadOnlyPolicy(path string) error {
	policy := fmt.Sprintf(ReadOnlyPolicy, struct{ Path string }{Path: path})
	return pm.CreatePolicy("read-only", policy)
}

func (pm *Policy) CreateReadWritePolicy(path string) error {
	policy := fmt.Sprintf(ReadWritePolicy, struct{ Path string }{Path: path})
	return pm.CreatePolicy("read-write", policy)
}

func (pm *Policy) CreateDenyAllPolicy(path string) error {
	policy := fmt.Sprintf(DenyAllPolicy, struct{ Path string }{Path: path})
	return pm.CreatePolicy("deny-all", policy)
}
