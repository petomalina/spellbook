package protoc

import (
	"context"
	"github.com/magefile/mage/sh"
	"os"
	"path/filepath"
	"strings"
)

type PluginGoGRPC struct {
	Out       string
	GoOpts    string
	GRPCOpts  string
	Targets   []string
	StripPath string
	Includes  []string
}

func (p PluginGoGRPC) MakeArgs() []string {
	args := []string{"--go-grpc_out=" + p.OutDir(), "--go_out=" + p.OutDir()}

	if p.GRPCOpts != "" {
		args = append(args, "--go-grpc_opt="+p.GRPCOpts)
	}

	if p.GoOpts != "" {
		args = append(args, "--go_opt="+p.GoOpts)
	}

	return append(args, p.Targets...)
}

func (p PluginGoGRPC) OutDir() string {
	return filepath.Join(p.Out, filepath.Dir(strings.TrimPrefix(p.Targets[0], p.StripPath)))
}

func (p PluginGoGRPC) Build(ctx context.Context) error {
	_ = os.MkdirAll(p.Out, os.ModePerm)

	args := append([]string{}, p.Includes...)
	args = append(args, p.MakeArgs()...)

	return sh.RunV("protoc", args...)
}
