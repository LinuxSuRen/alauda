package service

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

func newStartCmd(alauda client.APIClient) *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "start NAME",
		Short: "Start a service",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("service start expects NAME")
			}
			return doStart(alauda, args[0])
		},
	}

	return startCmd
}

func doStart(alauda client.APIClient, name string) error {
	fmt.Println("[alauda] Starting", name)

	util.InitializeClient(alauda)

	appName, serviceName, err := parseName(name)
	if err != nil {
		return err
	}

	params := client.ServiceParams{
		App: "",
	}

	if appName != "" {
		params.App = appName
	}

	err = alauda.StartService(serviceName, &params)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
