package root

import (
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/spf13/cobra"
)

type options struct {
	Cmd  *cobra.Command
	Args []string
}

var (
	createLong = templates.LongDesc(`
		Tekton pipelines automatically construct and mount a .git_credentials file to the $HOME dir.  For security reasons
		it's best that the service account used by the pipelines doesn't have permission to access kubernetes secrets and
		instead only access secrets available on the local filesystem.  This step parses the .git_credentials file and extracts
		the token which can then be used to set a local environment variable in the pipeline.
		
		GITHUB_TOKEN=$(step get git credential token)
`)

	createExample = templates.Examples(`
		# print the git token from a .git_credentials file to standard out
		step get git credential token
	`)
)

// NewCmdStep creates a command object for the "step get git credential token" command
func NewCmdStepParse() *cobra.Command {
	o := &options{}

	cmd := &cobra.Command{
		Use:     "step get git credential token",
		Short:   "Experimental Jenkins X lab step commands for getting a git token from a .git-credentials file",
		Long:    createLong,
		Example: createExample,
		Run: func(cmd *cobra.Command, args []string) {
			o.Args = args
			err := o.Run()
			helper.CheckErr(err)
		},
	}
	o.Cmd = cmd

	return cmd
}

// Run runs the command, if args are not nil they will be set on the command
func Run(args []string) error {
	cmd := NewCmdStepParse()
	if args != nil {
		args = args[1:]
		cmd.SetArgs(args)
	}
	return cmd.Execute()
}