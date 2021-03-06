package service

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

func newRestartCmd(alauda client.APIClient) *cobra.Command {
	restartCmd := &cobra.Command{
		Use:   "restart NAME",
		Short: "Restart a service",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("service restart expects NAME")
			}
			return doRestart(alauda, args[0])
		},
	}

	return restartCmd
}

func doRestart(alauda client.APIClient, name string) error {
	fmt.Println("[alauda] Restarting", name)

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

	err = alauda.RestartService(serviceName, &params)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
