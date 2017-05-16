package lb

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewUnbindCmd creates a new lb unbind command.
func NewUnbindCmd(alauda client.APIClient) *cobra.Command {
	var opts bindOptions

	unbindCmd := &cobra.Command{
		Use:   "unbind NAME",
		Short: "Remove the bindings between service endpoints and the load balancer",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("lb unbind expects NAME")
			}
			return doUnbind(alauda, args[0], &opts)
		},
	}

	unbindCmd.Flags().StringSliceVarP(&opts.publish, "publish", "p", []string{}, "Published endpoints to remove the bindings from the load balancer (servicename:listenerPort:containerPort")

	return unbindCmd
}

func doUnbind(alauda client.APIClient, name string, opts *bindOptions) error {
	fmt.Println("[alauda] Unbinding listeners from", name)

	util.InitializeClient(alauda)

	data, err := parseUnbindPublish(opts.publish)
	if err != nil {
		return err
	}

	err = alauda.BindLoadBalancer(name, data)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}

func parseUnbindPublish(publish []string) (*client.BindLoadBalancerData, error) {
	var listeners = []client.BindListenerData{}

	for _, desc := range publish {
		serviceName, listenerPort, containerPort, protocol, err := parsePublish(desc)
		if err != nil {
			return nil, err
		}

		if serviceName == "" {
			return nil, errors.New("no service name specified for listener")
		}

		if listenerPort == 0 {
			return nil, errors.New("no listener port specified for listener")
		}

		if protocol != "" {
			return nil, errors.New("invalid publish descriptor, expected serviceName:listenerPort:containerPort")
		}

		listener := client.BindListenerData{
			ServiceName:   serviceName,
			ListenerPort:  listenerPort,
			ContainerPort: containerPort,
			Protocol:      protocol,
		}
		listeners = append(listeners, listener)
	}

	data := client.BindLoadBalancerData{
		Action:    "unbind",
		Listeners: listeners,
	}

	return &data, nil
}
