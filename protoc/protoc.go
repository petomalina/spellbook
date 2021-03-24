package protoc

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

var (
	Plugins  []ProtocPlugin
	Includes []string
)

type ProtocPlugin interface {
	MakeArgs() []string
	OutDir() string
}

func Build() {
	// includes all in the current folder
	includes := append([]string{"-I."}, Includes...)

	for _, plug := range Plugins {
		_ = os.MkdirAll(plug.OutDir(), os.ModePerm)

		args := append([]string{}, includes...)
		args = append(args, plug.MakeArgs()...)

		fmt.Println(args)

		sh.RunV("protoc", args...)
	}
}
