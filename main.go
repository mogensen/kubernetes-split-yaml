package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"gopkg.in/urfave/cli.v1"
	"gopkg.in/yaml.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const outdir = "generated/"

func main() {
	app := cli.NewApp()
	app.Name = "kubernetes-split-yaml"
	app.Usage = "Split the 'giant yaml file' into one file pr kubernetes resource"
	app.Action = func(c *cli.Context) error {

		files := readAndSplitFile(c.Args().Get(0))

		os.Mkdir(outdir, os.ModePerm)

		for key, value := range files {

			var m KubernetesAPI

			err := yaml.Unmarshal([]byte(value), &m)
			check(err)

			if m.Kind == "" {
				fmt.Println("yaml file with no kind")
			} else {

				fmt.Println("File: ", key, m.Kind)

				filename := outdir + m.Metadata.Name + "-" + m.Kind + ".yaml"
				fmt.Println(filename)

				f, err := os.Create(filename)
				check(err)
				defer f.Close()

				_, err = f.WriteString(value)
				check(err)
			}
		}

		return nil
	}

	err := app.Run(os.Args)
	check(err)
}

func readAndSplitFile(logfile string) map[int]string {

	f, err := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)
	check(err)

	defer f.Close()

	rd := bufio.NewReader(f)
	c := make(map[int]string)
	count := 0

	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Panicf("read file line error: %v", err)
		}
		prettyline := strings.TrimSpace(line)
		if prettyline == "---" {
			count++
		} else {
			c[count] += line
		}
	}
	return c
}
