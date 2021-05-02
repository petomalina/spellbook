package protoc

import (
	"context"
	"github.com/magefile/mage/sh"
	"os"
)

type PluginDart struct {
	Out      string
	Opts     string
	Targets  []string
	Includes []string
}

func (p PluginDart) MakeArgs() []string {
	args := []string{"--dart_out=" + p.Opts + ":" + p.OutDir()}

	return append(args, p.Targets...)
}

func (p PluginDart) OutDir() string {
	return p.Out
}

func (p PluginDart) Build(ctx context.Context) error {
	_ = os.MkdirAll(p.Out, os.ModePerm)

	args := append([]string{}, p.Includes...)
	args = append(args, p.MakeArgs()...)

	return sh.RunV("protoc", args...)
}
