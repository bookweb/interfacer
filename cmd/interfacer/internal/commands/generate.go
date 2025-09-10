package commands

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/bookweb/interfacer/internal/engines"
	"github.com/spf13/afero"
	"github.com/urfave/cli/v3"
)

var (
	// Type configures the static file root dir
	Type = cli.StringFlag{
		Name:    "type",
		Value:   "",
		Sources: cli.EnvVars("TYPE"),
		Usage:   "type name; must be set",
	}
	// Lock configures the static file root dir
	Lock = cli.StringFlag{
		Name:    "lock",
		Value:   "",
		Sources: cli.EnvVars("LOCK"),
		Usage:   "lock name",
	}
	// Receiver configures the static file root dir
	Receiver = cli.StringFlag{
		Name:    "receiver",
		Value:   "",
		Sources: cli.EnvVars("RECEIVER"),
		Usage:   "receiver name; default first letter of type name",
	}
	// Output configures the static file root dir
	Output = cli.StringFlag{
		Name:    "output",
		Value:   "",
		Sources: cli.EnvVars("OUTPUT"),
		Usage:   "output file name; default <type_name>.gen.go",
	}
	// GenerateFlags is the set of all service CLI flags
	GenerateFlags = []cli.Flag{
		&Type,
		&Lock,
		&Receiver,
		&Output,
	}
)

func generate(ctx context.Context, cmd *cli.Command) error {
	// typeName := flags.String("type", "", "type name; must be set")
	typeName := cmd.String("type")
	// lockName := flags.String("lock", "", "lock name")
	lockName := cmd.String("lock")
	// receiver := flags.String("receiver", "", "receiver name; default first letter of type name")
	receiver := cmd.String("receiver")
	// output := flags.String("output", "", "output file name; default <type_name>.gen.go")
	output := cmd.String("output")

	var dir string
	if cliArgs := cmd.Args(); cliArgs.Len() > 0 {
		dir = cliArgs.Get(0)
	} else {
		// Default: process whole package in current directory.
		dir = "."
	}

	if !isDir(dir) {
		return errors.New("Specified argument is not a directory.")
	}

	src, err := engines.Parse(dir)
	if err != nil {
		return err
	}

	var options = []engines.Option{
		engines.Type(typeName),
		engines.Output(output),
		engines.Receiver(receiver),
		engines.Lock(lockName),
	}

	err = engines.Generate(afero.NewOsFs(), src, options...)
	if err != nil {
		return err
	}

	return nil
}

func isDir(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}
