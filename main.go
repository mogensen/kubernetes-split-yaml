package main

import (
	"io/ioutil"
	"path/filepath"

	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
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

		kind, metadataName, err := getYamlInfo(fileContent)
		if err != nil {
			log.Warnf("Ignoring %v", err)
		}

		name := metadataName + "-" + getShortName(kind) + ".yaml"
		filename := filepath.Join(outputDir, name)

		log.Infof("Creating file: %s", filename)

		err = ioutil.WriteFile(filename, []byte(fileContent), os.ModePerm)
		if err != nil {
			log.Fatalf("Failed creating file %s : %v", filename, err)
		}
	}
}

func readAndSplitFile(file string) []string {

	fileContent, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed reading file %s : %v", file, err)
	}

	docs := strings.Split(string(fileContent), "\n---")

	res := []string{}
	// Trim whitespace in both ends of each yaml docs.
	// - Re-add a single newline last
	for _, doc := range docs {
		content := strings.TrimSpace(doc)
		// Ignore empty docs
		if content != "" {
			res = append(res, content+"\n")
		}
	}
	return res
}
