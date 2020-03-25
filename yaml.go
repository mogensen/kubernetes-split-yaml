package main

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

func getYamlInfo(yamlContent string) (string, string, error) {

	// Start by removing all lines with templating in to create sane yaml
	cleanYaml := ""
	for _, line := range strings.Split(yamlContent, "\n") {
		if !strings.Contains(line, "{{") {
			cleanYaml += line + "\n"
		}
	}

	var m KubernetesAPI
	err := yaml.Unmarshal([]byte(cleanYaml), &m)
	if err != nil {
		return "", "", fmt.Errorf("Could not unmarshal: %v \n---\n%v", err, yamlContent)
	}

	if m.Kind == "" {
		return "", "", fmt.Errorf("yaml file with kind missing")
	} else if m.Metadata.Name == "" {
		return "", "", fmt.Errorf("yaml file with name missing")
	}

	return m.Kind, m.Metadata.Name, nil
}
