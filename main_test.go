package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"text/template"

	"github.com/google/go-cmp/cmp"
)

func Test_outFile(t *testing.T) {
	tests := []struct {
		tname       string
		outdir      string
		template    string
		nameRe      string
		namespaceRe string
		fileRe      string
		kind        string
		name        string
		namespace   string
		fileExp     string
	}{
		{
			tname:       "tpl_flat",
			outdir:      "gen",
			template:    TemplateFlat,
			nameRe:      NameRe,
			namespaceRe: NamespaceRe,
			fileRe:      FileRe,
			kind:        "Pod",
			namespace:   "foo",
			name:        "bar",
			fileExp:     filepath.Join("gen", "bar-pod.yaml"),
		},
		{
			tname:       "tpl_ns",
			outdir:      "gen",
			template:    TemplateNS,
			nameRe:      NameRe,
			namespaceRe: NamespaceRe,
			fileRe:      FileRe,
			kind:        "Pod",
			namespace:   "foo",
			name:        "bar",
			fileExp:     filepath.Join("gen", "foo", "bar.Pod.yaml"),
		},
		{
			tname:       "ns/name.kind-otherRe",
			outdir:      "gen",
			template:    TemplateNS,
			nameRe:      NameRe,
			namespaceRe: NamespaceRe,
			fileRe:      "foo.*bar",
			kind:        "Pod",
			namespace:   "foo",
			name:        "bar",
			fileExp:     filepath.Join("gen", "foo", "bar.Pod.yaml"),
		},
		{
			tname:       "no-ns/name.kind-otherRe",
			outdir:      "gen",
			template:    TemplateNS,
			nameRe:      NameRe,
			namespaceRe: NamespaceRe,
			fileRe:      FileRe,
			kind:        "ClusterRole",
			namespace:   "",
			name:        "bar",
			fileExp:     filepath.Join("gen", "_no_ns_", "bar.ClusterRole.yaml"),
		},
		{
			tname:       "nonmatching-ns",
			outdir:      "gen",
			template:    TemplateNS,
			nameRe:      NameRe,
			namespaceRe: "qqq",
			fileRe:      FileRe,
			kind:        "Pod",
			namespace:   "foo",
			name:        "bar",
			fileExp:     "",
		},
		{
			tname:       "no-filename-match",
			outdir:      "gen",
			template:    TemplateNS,
			nameRe:      NameRe,
			namespaceRe: NamespaceRe,
			fileRe:      "foo.*BAR",
			kind:        "Pod",
			namespace:   "foo",
			name:        "bar",
			fileExp:     "",
		},
		{
			tname:       "no-name-match",
			outdir:      "gen",
			template:    TemplateNS,
			nameRe:      "qqq",
			namespaceRe: NamespaceRe,
			fileRe:      FileRe,
			kind:        "Pod",
			namespace:   "foo",
			name:        "bar",
			fileExp:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.tname, func(t *testing.T) {
			m := &KubernetesAPI{
				Kind: tt.kind,
			}
			tpl, err := template.New("test").Parse(tt.template)
			if err != nil {
				t.Errorf("template: %s", tt.template)
			}
			m.Metadata.Name = tt.name
			m.Metadata.Namespace = tt.namespace

			filters := &Filters{
				name:      tt.nameRe,
				namespace: tt.namespaceRe,
				filename:  tt.fileRe,
			}
			got, err := outFile(tt.outdir, tpl, filters, m, nil)
			if got != tt.fileExp {
				t.Errorf("outFile() got = '%v', want '%v'", got, tt.fileExp)
			}
		})
	}
}

func Test_handleFile(t *testing.T) {
	// determine input files
	match, err := filepath.Glob("testdata/*.yaml")
	if err != nil {
		t.Fatal(err)
	}

	for _, in := range match {
		out := in + ".golden"
		runTest(t, in, out)
	}
}

func runTest(t *testing.T, in, out string) {
	t.Run(in, func(t *testing.T) {
		f := filepath.Base(in)

		outDir, err := ioutil.TempDir(os.TempDir(), f+"-")
		if err != nil {
			log.Fatal(err)
		}
		defer os.RemoveAll(outDir)

		handleFile(in, outDir, TemplateFlat, "", &Filters{})

		wantFiles, err := filepath.Glob(filepath.Join(out, "*"))
		if err != nil {
			t.Errorf("could not find test files: %v", err)
		}
		gotFiles, err := filepath.Glob(filepath.Join(outDir, "*"))
		if err != nil {
			t.Errorf("could not find result files: %v", err)
		}

		if len(gotFiles) != len(wantFiles) {
			t.Errorf("handleFile() = %v, want %v", len(gotFiles), len(wantFiles))
		}
		t.Logf("wantFiles: %s", wantFiles)

		for _, wantFile := range wantFiles {
			fileName := filepath.Base(wantFile)
			want, _ := ioutil.ReadFile(wantFile)
			generatedFileName := filepath.Join(outDir, filepath.Base(wantFile))

			t.Logf("Validating content of file: \n - Generated: %s\n - GoldenTest: %s", generatedFileName, wantFile)

			got, _ := ioutil.ReadFile(generatedFileName)
			if !reflect.DeepEqual(got, want) {
				t.Errorf(fileName + "\n" + cmp.Diff(string(got), string(want)))
			}
		}
	})
}
