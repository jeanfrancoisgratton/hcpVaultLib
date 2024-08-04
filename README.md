# hcpVaultLib
___

A Hashicorp Vault client package

This lib (package) was first written for my own use and was not even published on Github at first.
As I intend on expanding the package, now is a good time to publish and document the package.


## What's in there:

Mainly functions to:
- seal/unseal a Vault
- login with a token, userpass or approle methods
- read/write a secret from a given secret engine
- create/delete/edit/assign policies
- create/delete users
- enable/disable auth methods