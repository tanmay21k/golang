// Package main is the entry point for the roadmapResources application.
// It demonstrates various Go language features including constants, iota, and documentation.
package main

import (
	constantpkg "github.com/tanmay21k/golang/roadmapResources/constantPkg"
	iotapkg "github.com/tanmay21k/golang/roadmapResources/iota"
)

func main() {
	constantpkg.ConstantPrinter()
	iotapkg.IotaConstant()
	iotapkg.IotaOutsideConst()
}