package challenge

import (
	_ "github.com/docker/distribution/registry/client/auth/challenge"
	_ "unsafe"
)

// borrowed from https://github.com/distribution/distribution/blob/v2.7.1/registry/client/auth/challenge/authchallenge.go
//go:linkname parseValueAndParams github.com/docker/distribution/registry/client/auth/challenge.parseValueAndParams
func parseValueAndParams(header string) (value string, params map[string]string)
