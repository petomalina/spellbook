package terraform

import (
	"context"
	"github.com/magefile/mage/sh"
	"strings"
)

type Plugin struct {
	// Name is the user-defined name for filtering purposes
	// [Optional]
	Name string

	// Dir is the directory that terraform commands should be executed on
	// [Optional], defaults to current dir
	Dir string

	// AutoApprove automatically approves applies
	// [Optional], defaults to false
	AutoApprove bool

	// VarFile is a variables file to be used during plan & apply
	VarFile string

	// Target is a -target flag passed into the plan & apply commands
	Target string

	// ImpersonateGoogleSA defines which service account should be impersonated for the action
	// [Optional]
	ImpersonateGoogleSA string
}

func (p Plugin) Init(ctx context.Context) error {
	return p.executeCommon(ctx, "init")
}

func (p Plugin) Plan(ctx context.Context) error {
	return p.executeCommon(ctx, "plan")
}

func (p Plugin) Apply(ctx context.Context) error {
	return p.executeCommon(ctx, "apply")
}

func (p Plugin) Destroy(ctx context.Context) error {
	return p.executeCommon(ctx, "destroy")
}

func (p Plugin) Import(ctx context.Context, resource, id string) error {
	return p.executeCommon(ctx, "import", resource, id)
}

// executeCommon executes suite of common commands like plan, apply, destroy, import, etc.
// that have common flag and env interface
func (p Plugin) executeCommon(ctx context.Context, subcommand string, args ...string) error {
	globalArgs := p.makeGlobalFlags()
	subcommandArgs := append(p.makeSubcommandFlags(subcommand), args...)

	envMap, err := p.makeEnvMap()
	if err != nil {
		return err
	}

	return sh.RunWithV(envMap, "terraform", append(globalArgs, subcommandArgs...)...)
}

func (p Plugin) makeGlobalFlags() []string {
	var args []string

	if p.Dir != "" {
		args = append(args, "-chdir="+p.Dir)
	}

	return args
}

func (p Plugin) makeSubcommandFlags(subcommand string) []string {
	args := []string{subcommand, "-input=false"}

	if subcommand == "apply" && p.AutoApprove {
		args = append(args, "-auto-approve")
	}

	if (subcommand == "apply" || subcommand == "plan") && p.VarFile != "" {
		args = append(args, "-var-file="+p.VarFile)
	}

	if (subcommand == "apply" || subcommand == "plan") && p.Target != "" {
		args = append(args, "-target="+p.Target)
	}

	return args
}

func (p Plugin) makeEnvMap() (map[string]string, error) {
	env := map[string]string{}

	if p.ImpersonateGoogleSA != "" {
		cmdArgs := []string{
			"--impersonate-service-account",
			p.ImpersonateGoogleSA,
			"--quiet",
			"--verbosity=error",
			"auth",
			"print-access-token",
		}

		output, err := sh.Output("gcloud", cmdArgs...)
		if err != nil {
			return nil, err
		}

		env["GOOGLE_OAUTH_ACCESS_TOKEN"] = output
	}

	return env, nil
}

func FilterPluginsByNames(plugs []Plugin, names string) []Plugin {
	if names == "" {
		return plugs
	}

	var res []Plugin

	filter := strings.Split(names, ",")
	for _, plug := range plugs {
		for _, f := range filter {
			if plug.Name == f {
				res = append(res, plug)
				break
			}
		}
	}

	return res
}
