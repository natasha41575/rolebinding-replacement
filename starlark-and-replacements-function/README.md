# replace the rolebinding group with starlark to generate the string, and 
# a go function leveraging kustomize's replacements filter to propagate the value

## Usage

### Build the image locally
`cd apply-replacements`

`docker build -t gcr.io/kpt-fn/apply-replacements:unstable .`

### Replace the rolebinding group
`kpt fn render --image-pull-policy=ifNotPresent`

You will see the group name change in rolebinding.yaml to "project-order-service-app-admin@example.com".
You will see the value of data.group in value-store.yaml to also change to "project-order-service-app-admin@example.com".

