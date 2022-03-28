# YAML Query (`yq`), a Go Bonzai Branch

![WIP](https://img.shields.io/badge/status-wip-red)
![Go Version](https://img.shields.io/github/go-mod/go-version/rwxrob/yq)
[![GoDoc](https://godoc.org/github.com/rwxrob/yq?status.svg)](https://godoc.org/github.com/rwxrob/yq)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

This is the `yq` Go expression evaluator calling the same `yqlib` code
but wrapped in Bonzai (instead of Cobra) making it more portable and
composable into one's own Bonzai command trees along with other
commands. A simplified, high-level function calling into the `yqlib`
with reasonable defaults is also provided in the [`pkg`](pkg) allowing
other non-Bonzai applications to easily duplicate the same
functionality.

## Install

This command can be installed as a standalone program or composed into
a Bonzai command tree.

Standalone

```
go install github.com/rwxrob/yq/yq@latest
```

Composed

```go
package z

import (
	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/yq"
)

var Cmd = &bonzai.Cmd{
	Name:     `cmds`,
	Commands: []*bonzai.Cmd{help.Cmd, yq.Cmd},
}
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C yq yq
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.
