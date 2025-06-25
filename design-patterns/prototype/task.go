package main

type Clonnable interface {
	Clone() Clonnable
}

type EnvironmentVariable struct {
	Key   string
	Value string
}

type Container struct {
	Name    string
	Image   string
	EnvVars []EnvironmentVariable
}

type DeploymentTemplate struct {
	AppName    string
	Replicas   int
	Containers []Container
	Metadata   map[string]string
}

func (d *DeploymentTemplate) Clone() *DeploymentTemplate {
	clone := *d
	return &clone
}
