package typicmd

import (
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typictx"
	"github.com/urfave/cli"
)

// StandardCommands return standard commands for typical task tool
func StandardCommands(ctx *typictx.Context) []*typictx.Command {
	return []*typictx.Command{
		{
			Name:       "build",
			ShortName:  "b",
			Usage:      "Build the binary",
			ActionFunc: BuildBinary,
		},
		{
			Name:      "run",
			ShortName: "r",
			Usage:     "Run the binary",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "no-build",
					Usage: "Run the binary without build",
				},
			},
			ActionFunc: RunBinary,
		},
		{
			Name:       "test",
			ShortName:  "t",
			Usage:      "Run the testing",
			ActionFunc: RunTest,
		},
		{
			Name:  "release",
			Usage: "Release the distribution",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "no-test",
					Usage: "Release without run automated test",
				},
				cli.BoolFlag{
					Name:  "no-readme",
					Usage: "Release without generate readme",
				},
			},
			ActionFunc: ReleaseDistribution,
		},
		{
			Name:  "mock",
			Usage: "Generate mock class",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "no-delete",
					Usage: "Generate mock class with delete previous generation",
				},
			},
			ActionFunc: GenerateMock,
		},
		{
			Name:  "readme",
			Usage: "Generate readme document",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "no-commit",
					Usage: "Generate readme without auto commit",
				},
			},
			ActionFunc: GenerateReadme,
		},
		{
			Name:       "clean",
			Usage:      "Clean project from generated file during build time",
			ActionFunc: CleanProject,
		},
		{
			Name:  "docker",
			Usage: "Docker utitly",
			SubCommands: []*typictx.Command{
				{
					Name:       "compose",
					Usage:      "Generate docker-compose.yaml",
					ActionFunc: GenerateDockerCompose,
				},
				{
					Name:  "up",
					Usage: "Create and start containers",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "no-gen",
							Usage: "Create and start containers without generate docker-compose.yaml",
						},
					},
					ActionFunc: DockerUp,
				},
				{
					Name:       "down",
					Usage:      "Stop and remove containers, networks, images, and volumes",
					ActionFunc: DockerDown,
				},
			},
		},
	}
}