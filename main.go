package main

import (
	"io/ioutil"
	"path/filepath"

	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

const outdir = "generated/"

var log = logrus.New()

func main() {
	app := cli.NewApp()
	app.Name = "kubernetes-split-yaml"
	app.Usage = "Split the 'giant yaml file' into one file pr kubernetes resource"
	app.Action = func(c *cli.Context) error {

		handleFile(c.Args().Get(0), outdir)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Error running: %v", err)
	}
}

func handleFile(file string, outputDir string) {
	files := readAndSplitFile(file)

	os.Mkdir(outputDir, os.ModePerm)

	for _, fileContent := range files {

		// Start by removing all lines with templating in to create sane yaml
		cleanYaml := ""
		for _, line := range strings.Split(fileContent, "\n") {
			if !strings.Contains(line, "{{") {
				cleanYaml += line + "\n"
			}
		}

		var m KubernetesAPI
		err := yaml.Unmarshal([]byte(cleanYaml), &m)
		if err != nil {
			log.Fatalf("Could not unmarshal: %v \n---\n%v", err, fileContent)
		}

		if m.Kind == "" {
			log.Warn("yaml file with no kind - Ignoring")
		} else {

			name := m.Metadata.Name + "-" + getShortName(m.Kind) + ".yaml"
			filename := filepath.Join(outputDir, name)

			log.Infof("Creating file: %s", filename)

			err := ioutil.WriteFile(filename, []byte(fileContent), os.ModePerm)
			if err != nil {
				log.Fatalf("Failed creating file %s : %v", filename, err)
			}
		}
	}
}

func readAndSplitFile(file string) []string {

	fileContent, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed reading file %s : %v", file, err)
	}

	docs := strings.Split(string(fileContent), "\n---")
	// Trim whitespace in both ends of each yaml docs.
	// - Re-add a single newline last
	for i, doc := range docs {
		docs[i] = strings.TrimSpace(doc) + "\n"
	}
	return docs
}
