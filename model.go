package main

type KubernetesAPI struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name   string `yaml:"name"`
		Labels struct {
			Source string `yaml:"source"`
		} `yaml:"labels"`
	} `yaml:"metadata"`
}
