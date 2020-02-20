package main

// KubernetesAPI is a minimal struct for unmarshaling kubernetes configs into
type KubernetesAPI struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
}
