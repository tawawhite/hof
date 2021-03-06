package design

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

Outdir :: "./"

_LibImport :: [
	schema.Import & {Path: CLI.Package + "/lib"},
]

CLI :: schema.Cli & {
	Name:    "hof"
	Package: "github.com/hofstadter-io/hof"

	Usage: "hof"
	Short: "Polyglot Code Gereration Framework"
	Long:  Short

	Releases: CliReleases

	OmitRun: true

  Pflags: CliPflags

  // EnablePProf: true


	Imports: [
		{Path: "github.com/hofstadter-io/mvs/lib"},
	]

	PersistentPrerun: true
	PersistentPrerunBody: """
    lib.InitLangs()
  """

	Commands: [
    // meta
    AuthCommand,
    ConfigCommand,

    // hof
    NewCommand,
    GenCommand,
    StudiosCommand,

    // extern
    ModCommand,
    RunCommand,
    CueCommand,

	]
}

