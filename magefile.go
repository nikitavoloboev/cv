// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Build PDF file.
func Build() {
	sh.RunV("xelatex", "cv.tex")
}
