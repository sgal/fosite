//go:build tools
// +build tools

package fosite

import (
	_ "github.com/ecordell/optgen"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/mattn/goveralls"

	_ "github.com/ory/go-acc"
)
