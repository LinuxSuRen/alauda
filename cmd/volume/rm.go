package volume

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewRmCmd creates a new volume rm command.
func NewRmCmd(alauda client.APIClient) *cobra.Command {
	rmCmd := &cobra.Command{
		Use:   "rm NAME",
		Short: "Remove a volume",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("volume rm expects NAME")
			}
			return doRm(alauda, args[0])
		},
	}

	return rmCmd
}

func doRm(alauda client.APIClient, name string) error {
	fmt.Println("[alauda] Removing", name)

	util.InitializeClient(alauda)

	id, err := GetVolumeID(alauda, name)
	if err != nil {
		return err
	}

	err = alauda.RemoveVolume(id)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
