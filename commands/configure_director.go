package commands

import (
	"fmt"

	"github.com/pivotal-cf/jhanda/commands"
	"github.com/pivotal-cf/jhanda/flags"
)

type ConfigureDirector struct {
	service directorService
	Options struct {
		NetworkAssignment string `short:"n" long:"network-assignment" description:"assigns networks and AZs"`
		DirectorConfiguration string `short:"d" long:"director-configuration" description:"properties for director configuration"`
	}
}

//go:generate counterfeiter -o ./fakes/director_service.go --fake-name DirectorService . directorService
type directorService interface {
	NetworkAndAZ(jsonBody string) error
	Properties(jsonBody string) error
}

func NewConfigureDirector(service directorService) ConfigureDirector {
	return ConfigureDirector{service: service}
}

func (c ConfigureDirector) Execute(args []string) error {
	_, err := flags.Parse(&c.Options, args)
	if err != nil {
		return fmt.Errorf("could not parse configure-director flags: %s", err)
	}
	if c.Options.NetworkAssignment != "" {
		err = c.service.NetworkAndAZ(c.Options.NetworkAssignment)
		if err != nil {
			return fmt.Errorf("network and AZs couldn't be applied: %s", err)
		}
	}
	if c.Options.DirectorConfiguration != "" {
		err = c.service.Properties(c.Options.DirectorConfiguration)
		if err != nil {
			return fmt.Errorf("properties couldn't be applied: %s", err)
		}
	}
	return nil
}

func (c ConfigureDirector) Usage() commands.Usage {
	return commands.Usage{
		Description:      "This authenticated command configures the director.",
		ShortDescription: "configures the director",
		Flags:            c.Options,
	}
}
