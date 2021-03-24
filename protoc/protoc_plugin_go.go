package protoc

import (
	"path/filepath"
	"strings"
)

type ProtocPluginGo struct {
	Out       string
	Opts      string
	Targets   []string
	StripPath string
}

func (p ProtocPluginGo) MakeArgs() []string {
	args := []string{"--go_out=" + p.OutDir()}

	if p.Opts != "" {
		args = append(args, "--go_opt="+p.Opts)
	}

	return append(args, p.Targets...)
}

func (p ProtocPluginGo) OutDir() string {
	return filepath.Join(p.Out, filepath.Dir(strings.TrimPrefix(p.Targets[0], p.StripPath)))
}

type ProtocPluginGoGRPC struct {
	Out       string
	Opts      string
	Targets   []string
	StripPath string
}

func (p ProtocPluginGoGRPC) MakeArgs() []string {
	args := []string{"--go-grpc_out=" + p.OutDir()}

	if p.Opts != "" {
		args = append(args, "--go-grpc_opt="+p.Opts)
	}

	return append(args, p.Targets...)
}

func (p ProtocPluginGoGRPC) OutDir() string {
	return filepath.Join(p.Out, filepath.Dir(strings.TrimPrefix(p.Targets[0], p.StripPath)))
}
