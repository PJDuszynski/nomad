package command

import (
	"fmt"
	"strings"

	"github.com/hashicorp/nomad/api"
	"github.com/mitchellh/cli"
	"github.com/posener/complete"
)

// Ensure LoginCommand satisfies the cli.Command interface.
var _ cli.Command = &LoginCommand{}

// LoginCommand implements cli.Command.
type LoginCommand struct {
	Meta

	authMethodType string
	authMethodName string

	template string
	json     bool
}

// Help satisfies the cli.Command Help function.
func (l *LoginCommand) Help() string {
	helpText := `
Usage: nomad login [options]

  Info is used to read the services registered to a single service name.

  When ACLs are enabled, this command requires a token with the 'read-job'
  capability for the service namespace.

General Options:

  ` + generalOptionsUsage(usageOptsNoNamespace) + `

Login Options:

  -json
    Output the ACL token in JSON format.

  -t
    Format and display the ACL token using a Go template.
`
	return strings.TrimSpace(helpText)
}

// Synopsis satisfies the cli.Command Synopsis function.
func (l *LoginCommand) Synopsis() string {
	return "Display an individual Nomad service registration"
}

func (l *LoginCommand) AutocompleteFlags() complete.Flags {
	return mergeAutocompleteFlags(l.Meta.AutocompleteFlags(FlagSetClient),
		complete.Flags{
			"-json": complete.PredictNothing,
			"-t":    complete.PredictAnything,
		})
}

// Name returns the name of this command.
func (l *LoginCommand) Name() string { return "login" }

// Run satisfies the cli.Command Run function.
func (l *LoginCommand) Run(args []string) int {

	flags := l.Meta.FlagSet(l.Name(), FlagSetClient)
	flags.Usage = func() { l.Ui.Output(l.Help()) }
	flags.BoolVar(&l.json, "json", false, "")
	flags.StringVar(&l.template, "t", "", "")
	if err := flags.Parse(args); err != nil {
		return 1
	}
	args = flags.Args()

	if len(args) != 0 {
		l.Ui.Error("This command takes no arguments")
		l.Ui.Error(commandErrorText(l))
		return 1
	}

	// Validate other flags to ensure they are correctly set.
	// Switch on the auth type and call the sub-function; only OIDC now.
	// Format the ACL token.

	_, err := l.Meta.Client()
	if err != nil {
		l.Ui.Error(fmt.Sprintf("Error initializing client: %s", err))
		return 1
	}

	if l.json || l.template != "" {
		out, err := Format(l.json, l.template, "INSERT_TOKEN")
		if err != nil {
			l.Ui.Error(err.Error())
			return 1
		}
		l.Ui.Output(out)
		return 0
	}
	return 0
}

func (l *LoginCommand) loginOIDC() (*api.ACLToken, error) {

	// Check auth name; lookup default if needed.
	// Start the callback server.
	// Perform the auth-url API request.
	// Listen to the callback server for a response or error; use a timeout.
	// Perform the auth-complete API request.

	return nil, nil
}
