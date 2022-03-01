# replace the role binding group with starlark

## Usage

### Fetch the package
`kpt pkg get REPO_URI[.git]/PKG_PATH[@VERSION] replacements`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree starkark`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Replace the rolebinding group
`kpt fn eval --fn-config=configure-rolebinding-group.yaml --image=gcr.io/kpt-fn/starlark:v0.3.0`
You will see the group name change in rolebinding.yaml to "myproject-myns-appadmin@example.com".

