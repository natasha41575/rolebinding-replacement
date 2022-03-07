We assume that the user has already used `kube-gen`
to generate value-store.yaml, a ConfigMap that will be used to store the value to propagate.

From there, we have two functions:

1. Starlark to generate the rolebinding group
2. Replacements to propagate the value to the rolebinding. 

## Usage

### Build the image locally
`cd apply-replacements`

`docker build -t gcr.io/kpt-fn/apply-replacements:unstable .`

### Replace the rolebinding group
`kpt fn render --image-pull-policy=ifNotPresent`

You will see the group name change in rolebinding.yaml to "project-order-service-app-admin@example.com".
You will see the value of data.group in value-store.yaml to also change to "project-order-service-app-admin@example.com".

