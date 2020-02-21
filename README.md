# CLAPP
[![GoDoc](https://godoc.org/github.com/pirmd/clapp?status.svg)](https://godoc.org/github.com/pirmd/clapp)&nbsp; 
[![Go Report Card](https://goreportcard.com/badge/github.com/pirmd/clapp)](https://goreportcard.com/report/github.com/pirmd/clapp)&nbsp;

`clapp` provides simple means to build a command-line application, featuring
several levels of sub-commands, flags and args parsing, help and manpage
generation, config file parser and build version integration with git.

`clapp` is probably not as features complete as other command-line application
builders existing in the wild but tries to stay simple.

## EXAMPLE

```go
package main

import (
    "fmt"
    "encoding/json"

    "github.com/pirmd/clapp"
)

type config struct {
    Species   string
}

func main() {
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
```

A slightly more complete example with documentation generation can be found in
[gostore](https://github.com/pirmd/gostore).

## INSTALLATION
Everything should work fine using go standard commands (`build`, `get`,
`install`...).

## USAGE
Running `go doc github.com/pirmd/clapp` should give you helpful guidelines on
available features.

## CONTRIBUTION
If you feel like to contribute, just follow github guidelines on
[forking](https://help.github.com/articles/fork-a-repo/) then [send a pull
request](https://help.github.com/articles/creating-a-pull-request/)

[modeline]: # ( vim: set fenc=utf-8 spell spl=en: )
