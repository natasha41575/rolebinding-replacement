# replace the role binding group with a go function leveraging kustomize's replacements filter

## Usage

### Build the image locally
`cd apply-replacements`

`docker build -t gcr.io/kpt-fun/apply-replacements:unstable .`

### Replace the rolebinding group
`kpt fn eval --fn-config=configure-rolebinding-group.yaml --image-pull-policy=never --image=gcr.io/kpt-fun/apply-replacements:unstable`

You will see the group name change in rolebinding.yaml to "myproject-myns-appadmin@example.com".

