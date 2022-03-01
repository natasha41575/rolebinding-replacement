# replace the role binding group with kustomize and replacements

## Usage

### Replace the rolebinding group
`kustomize build .`

You will receive an out of place yaml stream, where the rolebinding's group name
has been changed to "myproject-myns-appadmin@example.com".
