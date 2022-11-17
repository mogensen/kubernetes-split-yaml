package main

import (
	"encoding/json"
	"os"
	"strings"
)

type shortNameMap map[string]string

func getShortName(kind string, shortNameMap shortNameMap) string {
	if shortNameMap == nil {
		shortNameMap = getDefaultShortNameMap()
	}

	kindLower := strings.ToLower(kind)

	shortname, exists := shortNameMap[kindLower]
	if !exists {
		return kindLower
	}

	return shortname
}

func getDefaultShortNameMap() shortNameMap {
	return shortNameMap{
		"service":                  "svc",
		"serviceaccount":           "sa",
		"rolebinding":              "rb",
		"clusterrolebinding":       "crb",
		"clusterrole":              "cr",
		"horizontalpodautoscaler":  "hpa",
		"poddisruptionbudget":      "pdb",
		"customresourcedefinition": "crd",
		"configmap":                "cm",
	}
}

func loadCustomShortnameMapFile(path string) (shortNameMap, error) {
	customShortnameMapFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer customShortnameMapFile.Close()

	decoder := json.NewDecoder(customShortnameMapFile)

	shortnameMap := getDefaultShortNameMap()

	if err := decoder.Decode(&shortnameMap); err != nil {
		return nil, err
	}

	return shortnameMap, nil
}
