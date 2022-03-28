package yq

import (
	"log"

	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/inc/help"
)

var Cmd = &bonzai.Cmd{

	Name:      `yq`,
	Summary:   `query YAML and JSON files`,
	Usage:     `(h|help|<expression>) [<file>...]`,
	Version:   `v0.0.1`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2.0`,
	Commands:  []*bonzai.Cmd{help.Cmd},

	Description: `
		The <yq> command allows YAML (and JSON, since all JSON is YAML)
		files to be queried using a simple syntax that is nearly identical
		to the popular stedolan/jq tool (written in C). In fact, <yq> uses
		the same <yqlib> that the <yq> tool does (just without the Cobra).

		The first argument is the <expression> and almost always begins with
		a dot (.). See the <https://github.com/stedolan/jq> project for
		exact syntax (until more documentation is updated here to contain
		the same).

		The remaining arguments are the names of one or more files. If no
		file argument is passed standard input is assumed (per UNIX filter
		philosophy). Note that the special dash (-) filename is not
		supported even though it was in the original <yq> tool.`,

	Call: func(_ *bonzai.Cmd, args ...string) error {
		log.Printf("would run yq stuff")
		return nil
	},
}
