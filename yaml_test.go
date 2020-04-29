package main

import (
	"strings"
	"testing"
)

func Test_getYamlInfo(t *testing.T) {
	type args struct {
		yamlContent string
	}
	tests := []struct {
		name          string
		args          args
		wantKind      string
		wantName      string
		wantNamespace string
		wantErr       bool
		wantErrMsg    string
	}{
		{
			name: "Error on empty file",
			args: args{
				yamlContent: "",
			},
			wantKind:   "",
			wantName:   "",
			wantErr:    true,
			wantErrMsg: "yaml file with kind missing",
		},
		{
			name: "Error on non-yaml file",
			args: args{
				yamlContent: "Non-yaml file....?",
			},
			wantErr:    true,
			wantErrMsg: "Could not unmarshal",
		},
		{
			name: "Error on missing kind",
			args: args{
				yamlContent: `
apiVersion: v1
metadata:
  name: simple
`,
			},
			wantErr:    true,
			wantErrMsg: "yaml file with kind missing",
		},
		{
			name: "Error on missing name",
			args: args{
				yamlContent: `kind: Deployment`,
			},
			wantErr:    true,
			wantErrMsg: "yaml file with name missing",
		},
		{
			name: "Can handle simple file",
			args: args{
				yamlContent: `
apiVersion: v1
kind: Pod
metadata:
  name: simple
  namespace: foo
`,
			},
			wantName:      "simple",
			wantNamespace: "foo",
			wantKind:      "Pod",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getYamlInfo(tt.args.yamlContent)
			if (err != nil) != tt.wantErr {
				t.Errorf("getYamlInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && !strings.HasPrefix(err.Error(), tt.wantErrMsg) {
				t.Errorf("getYamlInfo() error = %v, wantErr %v", err.Error(), tt.wantErrMsg)
				return
			}
			if tt.wantErr {
				return
			}

			if got.Kind != tt.wantKind {
				t.Errorf("getYamlInfo() got = %v, want %v", got.Kind, tt.wantKind)
			}
			if got.Metadata.Name != tt.wantName {
				t.Errorf("getYamlInfo() got = %v, want %v", got.Metadata.Name, tt.wantName)
			}
		})
	}
}
