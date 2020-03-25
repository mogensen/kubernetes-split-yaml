# kubernetes-split-yaml

[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fmogensen%2Fkubernetes-split-yaml%2Fbadge%3Fref%3Dmaster&style=flat)](https://actions-badge.atrox.dev/mogensen/kubernetes-split-yaml/goto?ref=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mogensen/kubernetes-split-yaml)](https://goreportcard.com/report/github.com/mogensen/kubernetes-split-yaml)
[![codecov](https://codecov.io/gh/mogensen/kubernetes-split-yaml/branch/master/graph/badge.svg)](https://codecov.io/gh/mogensen/kubernetes-split-yaml)

Split the 'giant yaml file' into one file pr kubernetes resource

## Installation

If you have golang installed you can use `go get`.

```bash
$ go get -v github.com/mogensen/kubernetes-split-yaml
```
This will download the source and install the binary `kubernetes-split-yaml`

## Usage

```
$ kubernetes-split-yaml giant-k8s-file.yaml
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)