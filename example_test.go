package clapp_test

import (
	"encoding/json"
	"fmt"

	"github.com/pirmd/clapp"
)

type config struct {
	Species string
}

func Example() {
	cfg := &config{
		Species: "Tamarin", //Species defaults to Tamarin if not altered by a command flag
	}

	cmd := &clapp.Command{
		Name:        "monkey",
		Usage:       "A command-line monkey simulator.",
		Description: "monkey is a command-line companion that provides you with endless untertaining interactions.",

		Config: &clapp.Config{
			Unmarshaller: json.Unmarshal,
			Files:        clapp.DefaultConfigFiles("config.yaml"),
			Var:          cfg,
		},

		ShowHelp:    clapp.ShowUsage,
		ShowVersion: clapp.ShowVersion,
	}

	cmd.Flags = clapp.Flags{
		{
			Name:  "species",
			Usage: "Monkey's species to simulate.",
			Var:   &cfg.Species,
		},
	}

	var fruit string
	cmd.SubCommands.Add(&clapp.Command{
		Name:  "eat",
		Usage: "Give something to eat to your monkey. Be aware that lelectronic monkey eats only bananas.",
		Args: clapp.Args{
			{
				Name:  "fruit",
				Usage: "Fruit to eat.",
				Var:   &fruit,
			},
		},
		Execute: func() error {
			if fruit != "banana" {
				return fmt.Errorf("%s only eats bananas", cfg.Species)
			}

			fmt.Printf("%s loves eating %s", cfg.Species, fruit)
			return nil
		},
	})

	cmd.MustRun([]string{"eat", "banana"}) //usually cmd.MustRun(os.Args[1:])
	//Output: Tamarin loves eating banana
}
