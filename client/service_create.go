package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// CreateServiceData defines the request body for the CreateService API.
type CreateServiceData struct {
	Version            string                `json:"version"`
	Name               string                `json:"service_name"`
	Cluster            string                `json:"region_name"`
	Space              string                `json:"space_name"`
	ImageName          string                `json:"image_name"`
	ImageTag           string                `json:"image_tag"`
	Command            string                `json:"run_command"`
	Entrypoint         string                `json:"entrypoint"`
	TargetState        string                `json:"target_state"`
	TargetInstances    int                   `json:"target_num_instances"`
	InstanceSize       string                `json:"instance_size"`
	CustomInstanceSize ServiceInstanceSize   `json:"custom_instance_size"`
	ScalingMode        string                `json:"scaling_mode"`
	Ports              []int                 `json:"ports"`
	NetworkMode        string                `json:"network_mode"`
	Env                map[string]string     `json:"instance_envvars"`
	LoadBalancers      []ServiceLoadBalancer `json:"load_balancers"`
	Volumes            []ServiceVolume       `json:"volumes"`
	Configs            []ServiceConfig       `json:"mount_points"`
}

// CreateService creates and deploys a new service.
func (client *Client) CreateService(data *CreateServiceData) error {
	url := client.buildURL("services", "")

	request, err := client.buildCreateServiceRequest(data)
	if err != nil {
		return err
	}

	response, err := request.Post(url)
	if err != nil {
		return err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) buildCreateServiceRequest(data *CreateServiceData) (*rest.Request, error) {
	request := rest.NewRequest(client.Token())

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request.SetBody(body)

	return request, nil
}
