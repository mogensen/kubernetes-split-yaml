package main

import (
	"fmt"
	"testing"
)

func Test_getShortName(t *testing.T) {
	tests := []struct {
		name string
		have string
		want string
	}{
		{have: "ServiceAccount", want: "sa"},
		{have: "horizontalpodautoscaler", want: "hpa"},
		{have: "Pod", want: "pod"},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%s -> %s", tt.have, tt.want)
		t.Run(name, func(t *testing.T) {
			if got := getShortName(tt.have, nil); got != tt.want {
				t.Errorf("getShortName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getShortName_CustomMap(t *testing.T) {
	tests := []struct {
		name string
		have string
		want string
	}{
		{have: "ServiceAccount", want: "sa"},
		{have: "horizontalpodautoscaler", want: "hpa"},
		{have: "Pod", want: "pod"},
		{have: "Deployment", want: "dep"},
	}

	customShortNameMap, err := loadCustomShortnameMapFile("./testdata/custom_sn_map.json")
	if err != nil {
		t.Fatalf("Could not load custom shortname map file: %v", err)
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%s -> %s", tt.have, tt.want)
		t.Run(name, func(t *testing.T) {
			if got := getShortName(tt.have, customShortNameMap); got != tt.want {
				t.Errorf("getShortName() = %v, want %v", got, tt.want)
			}
		})
	}
}
