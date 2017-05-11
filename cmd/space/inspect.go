package space

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewInspectCmd creates a new space inspect command.
func NewInspectCmd(alauda client.APIClient) *cobra.Command {
	inspectCmd := &cobra.Command{
		Use:   "inspect NAME",
		Short: "Inspect a space",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("space inspect expects NAME")
			}
			return doInspect(alauda, args[0])
		},
	}

	return inspectCmd
}

func doInspect(alauda client.APIClient, name string) error {
	fmt.Println("[alauda] Inspecting", name)

	util.InitializeClient(alauda)

	result, err := alauda.InspectSpace(name)
	if err != nil {
		return err
	}

	err = printSpace(result)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}

func printSpace(space *client.Space) error {
	marshalled, err := json.MarshalIndent(space, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(marshalled))

	return nil
}
