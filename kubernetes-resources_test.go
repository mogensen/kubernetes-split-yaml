package main

import (
	"fmt"
	"testing"
)

func Test_getShortName(t *testing.T) {
	type args struct {
		kind string
	}
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
			if got := getShortName(tt.have); got != tt.want {
				t.Errorf("getShortName() = %v, want %v", got, tt.want)
			}
		})
	}
}
