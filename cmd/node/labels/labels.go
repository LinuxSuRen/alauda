package labels

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type labelsOptions struct {
	cluster string
}

// NewLabelsCmd creates a new node labels command.
func NewLabelsCmd(alauda client.APIClient) *cobra.Command {
	var opts labelsOptions

	labelsCmd := &cobra.Command{
		Use:   "labels IP",
		Short: "Get labels of a node",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("node labels expects IP")
			}
			return doLabels(alauda, args[0], &opts)
		},
	}

	labelsCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")

	labelsCmd.AddCommand(
		newSetCmd(alauda),
	)

	return labelsCmd
}

func doLabels(alauda client.APIClient, name string, opts *labelsOptions) error {
	fmt.Println("[alauda] Getting labels for", name)

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	result, err := alauda.InspectNode(name, cluster)
	if err != nil {
		return err
	}

	printLabelsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printLabelsResult(result *client.Node) {
	header := buildLabelsTableHeader()
	content := buildLabelsTableContent(result)

	util.PrintTable(header, content)
}

func buildLabelsTableHeader() []string {
	return []string{"KEY", "VALUE", "EDITABLE"}
}

func buildLabelsTableContent(node *client.Node) [][]string {
	var content [][]string

	for _, label := range node.Labels {
		content = append(content, []string{label.Key, label.Value, strconv.FormatBool(label.Editable)})
	}

	return content
}
