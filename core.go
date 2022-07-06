package core

import (
	g "github.com/davemostarda/test-project-go/generators"
)

func Generate(val string) (string, error) {
	res := g.TestMethod(val)

	return res, nil
}
