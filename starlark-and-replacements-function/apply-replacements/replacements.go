package main

import (
	"fmt"
	sdk "github.com/GoogleContainerTools/kpt-functions-catalog/thirdparty/kyaml/fnsdk"
	"sigs.k8s.io/kustomize/api/filters/replacement"
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type Replacements struct {
	Replacements []types.Replacement `json:"replacements,omitempty" yaml:"replacements,omitempty"`
}

// Config initializes Replacements from a functionConfig sdk.KubeObject
func (si *Replacements) Config(functionConfig *sdk.KubeObject) error {
	si.Replacements = []types.Replacement{}
	if err := functionConfig.As(si); err != nil {
		return fmt.Errorf("unable to convert functionConfig to %s:\n%w",
			"replacements", err)
	}
	return nil
}

// Transform set image out of place and returns a new []*sdk.KubeObject
func (si *Replacements) Transform(items []*sdk.KubeObject) ([]*sdk.KubeObject, error) {
	var transformedItems []*sdk.KubeObject
	var nodes []*yaml.RNode
	for _, obj := range items {
		objRN := obj.ToRNode()
		nodes = append(nodes, objRN)
	}
	transformedNodes, err := replacement.Filter{
		Replacements:  si.Replacements,
	}.Filter(nodes)
	if err != nil {
		return nil, err
	}
	for _, n := range transformedNodes {
		transformedItems = append(transformedItems, sdk.NewFromRNode(n))
	}
	return transformedItems, nil
}
