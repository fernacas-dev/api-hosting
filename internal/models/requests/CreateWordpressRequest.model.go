package requests

type CreateWordpressServiceRequest struct {
	ContainerName  string `json:"containerName"`
	ContainerImage string `json:"containerImage"`
	VolumeName     string `json:"volumeName"`
	NetworkName    string `json:"networkName"`
}
