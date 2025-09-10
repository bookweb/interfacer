package commands

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func Execute() error {
	cmd := &cli.Command{
		Name:    "interfacer",
		Usage:   "generate interface and setter",
		Version: "dev",
		Flags: append(
			RootConfigFlags,
			RootFlags...,
		),
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Usage:   "generate interface and setter",
				Flags: append(
					GenerateFlags,
					AuthFlags...,
				),
				Action: generate,
			},
		},
		Action: generate,
	}

	// if err := cmd.Run(context.Background(), os.Args); err != nil {
	// 	log.Fatal(err)
	// }

	// version := config.VersionString(config.Version, config.Revision, config.BuildTime)

	// app := &cli.App{
	// 	Name:    "pstools",
	// 	Usage:   "service for managing server tools",
	// 	Version: version,
	// 	Flags: append(
	// 		RootConfigFlags,
	// 		RootFlags...,
	// 	),
	// 	Commands: []*cli.Command{
	// 		{
	// 			Name:    "serve",
	// 			Aliases: []string{"s"},
	// 			Usage:   "serves the LDAPManager API",
	// 			Flags: append(
	// 				ServeFlags,
	// 				AuthFlags...,
	// 			),
	// 			Action: serve,
	// 		},
	// 	},
	// 	Action: serve,
	// }

	return cmd.Run(context.Background(), os.Args)
}

var (
	// LdapHost configures the LDAP server host
	LdapHost = cli.StringFlag{
		Name:    "ldap-host",
		Value:   "localhost",
		Sources: cli.EnvVars("LDAP_HOST"),
		Usage:   "LDAP host",
	}
	// LdapPort configures the LDAP server port
	LdapPort = cli.IntFlag{
		Name:    "ldap-port",
		Value:   389,
		Sources: cli.EnvVars("LDAP_PORT"),
		Usage:   "LDAP port",
	}
	// LdapConfigFlags is a set of all LDAP CLI flags
	RootConfigFlags = []cli.Flag{
		&LdapHost,
		&LdapPort,
	}
)

var (
	// GroupsOu configures the LDAP group organizational unit
	GroupsOu = cli.StringFlag{
		Name:    "groups-ou",
		Value:   "groups",
		Sources: cli.EnvVars("GROUPS_OU"),
		Usage:   "group organizational unit",
	}
	// GroupsOu configures the LDAP group organizational unit
	UsersOu = cli.StringFlag{
		Name:    "users-ou",
		Value:   "users",
		Sources: cli.EnvVars("USERS_OU"),
		Usage:   "group organizational unit",
	}

	// LdapFlags is a collection of all LDAP CLI flags
	RootFlags = []cli.Flag{
		&GroupsOu,
		&UsersOu,
	}
)

var (
	// Key to sign the tokens with
	Key = cli.StringFlag{
		Name:    "key",
		Aliases: []string{"public-key", "signing-key"},
		Sources: cli.EnvVars("KEY"),
		Usage:   "private key to sign the tokens with",
	}
	// AuthFlags is a set of all CLI flags
	AuthFlags = []cli.Flag{
		&Key,
	}
)
