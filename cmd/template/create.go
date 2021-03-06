package template

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type createOptions struct {
	filePath string
	space    string
}

func newCreateCmd(alauda client.APIClient) *cobra.Command {
	var opts createOptions

	createCmd := &cobra.Command{
		Use:   "create NAME",
		Short: "Creates an app template",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("template create expects NAME")
			}
			return doCreate(alauda, args[0], &opts)
		},
	}

	createCmd.Flags().StringVarP(&opts.filePath, "file", "f", "./alauda-compose.yml", "Compose yaml file template")
	createCmd.Flags().StringVarP(&opts.space, "space", "s", "", "Space to create the template in")

	return createCmd
}

func doCreate(alauda client.APIClient, name string, opts *createOptions) error {
	fmt.Printf("[alauda] Creating %s using %s\n", name, opts.filePath)

	util.InitializeClient(alauda)

	space, err := util.ConfigSpace(opts.space)
	if err != nil {
		return err
	}

	data := client.CreateAppTemplateData{
		Name:        name,
		Description: "",
		Space:       space,
	}

	absPath, err := filepath.Abs(opts.filePath)
	if err != nil {
		return err
	}

	err = alauda.CreateAppTemplate(&data, absPath)
	if err != nil {
		return err
	}

	return nil
}
