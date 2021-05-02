package protoc

import (
	"context"
	"github.com/magefile/mage/sh"
	"os"
	"path/filepath"
	"strings"
)

type PluginGo struct {
	Out       string
	Opts      string
	Targets   []string
	StripPath string
	Includes  []string
}

func (p PluginGo) MakeArgs() []string {
	args := []string{"--go_out=" + p.OutDir()}

	if p.Opts != "" {
		args = append(args, "--go_opt="+p.Opts)
	}

	return append(args, p.Targets...)
}

func (p PluginGo) OutDir() string {
	return filepath.Join(p.Out, filepath.Dir(strings.TrimPrefix(p.Targets[0], p.StripPath)))
}

func (p PluginGo) Build(ctx context.Context) error {
	_ = os.MkdirAll(p.OutDir(), os.ModePerm)

	args := append([]string{}, p.Includes...)
	args = append(args, p.MakeArgs()...)

	return sh.RunV("protoc", args...)
}
