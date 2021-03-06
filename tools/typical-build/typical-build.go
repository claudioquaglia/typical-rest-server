package main

import (
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
	"github.com/typical-go/typical-go/pkg/typrls"
	"github.com/typical-go/typical-rest-server/pkg/dockerrx"
	"github.com/typical-go/typical-rest-server/pkg/typdocker"
	"github.com/typical-go/typical-rest-server/pkg/typrest"
	pg "github.com/typical-go/typical-rest-server/tools/pg-tool/pkg/util"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "typical-rest-server",
	ProjectVersion: "0.8.36",
	ProjectLayouts: []string{"internal", "pkg"},

	Cmds: []typgo.Cmd{

		// test
		&typgo.TestCmd{
			Action: &typgo.StdTest{},
		},

		&typannot.AnnotateCmd{
			Annotators: []typannot.Annotator{
				&typapp.CtorAnnotation{},
				&typapp.DtorAnnotation{},
				&typrest.AppCfgAnnotation{DotEnv: true},
			},
		},

		// compile
		&typgo.CompileCmd{
			Action: &typgo.StdCompile{},
		},

		// run
		&typgo.RunCmd{
			Before: typgo.BuildSysRuns{"annotate", "compile"},
			Action: &typgo.StdRun{},
		},

		// clean
		&typgo.CleanCmd{
			Action: &typgo.StdClean{},
		},

		// mock
		&typmock.MockCmd{},

		// docker
		&typdocker.DockerCmd{
			Composers: []typdocker.Composer{
				&dockerrx.PostgresWithEnv{
					Name:        "pg01",
					UserEnv:     "PG_USER",
					PasswordEnv: "PG_PASSWORD",
					PortEnv:     "PG_PORT",
				},
				&dockerrx.RedisWithEnv{
					Name:        "redis01",
					PasswordEnv: "REDIS_PASSWORD",
					PortEnv:     "REDIS_PORT",
				},
			},
		},

		// pg
		&pg.PSQLCmd{
			Name:         "pg",
			HostEnv:      "PG_HOST",
			PortEnv:      "PG_PORT",
			UserEnv:      "PG_USER",
			PasswordEnv:  "PG_PASSWORD",
			DBNameEnv:    "PG_DBNAME",
			MigrationSrc: "databases/pg/migration",
			SeedSrc:      "databases/pg/seed",
		},

		// release
		&typrls.ReleaseCmd{
			Before:     typgo.BuildSysRuns{"test", "compile"},
			Validation: typrls.DefaultValidation,
			Summary:    typrls.DefaultSummary,
			Releaser:   &typrls.Github{Owner: "typical-go", Repo: "typical-rest-server"},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
