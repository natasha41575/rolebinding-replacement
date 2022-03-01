# replace the role binding group with starlark

## Usage

### Replace the rolebinding group
`kpt fn eval --fn-config=configure-rolebinding-group.yaml --image=gcr.io/kpt-fn/starlark:v0.3.0`

You will see the group name change in rolebinding.yaml to "myproject-myns-appadmin@example.com".

