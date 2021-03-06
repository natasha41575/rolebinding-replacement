package main

import (
	"os"

	sdk "github.com/GoogleContainerTools/kpt-functions-catalog/thirdparty/kyaml/fnsdk"
)

func main() {
	if err := sdk.AsMain(sdk.ResourceListProcessorFunc(applyReplacements)); err != nil {
		os.Exit(1)
	}
}

func applyReplacements(rl *sdk.ResourceList) error {
	si := Replacements{}
	if err := si.Config(rl.FunctionConfig); err != nil {
		return err
	}
	transformedItems, err := si.Transform(rl.Items)
	if err != nil {
		return err
	}
	rl.Items = transformedItems
	return nil
}
