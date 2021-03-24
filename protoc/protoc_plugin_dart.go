package protoc

type ProtocPluginDart struct {
	Out     string
	Opts    string
	Targets []string
}

func (p ProtocPluginDart) MakeArgs() []string {
	args := []string{"--dart_out=" + p.Opts + ":" + p.OutDir()}

	return append(args, p.Targets...)
}

func (p ProtocPluginDart) OutDir() string {
	return p.Out
}
