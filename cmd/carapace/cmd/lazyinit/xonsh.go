package lazyinit

import (
	"fmt"
)

func Xonsh(completers []string) string {
	snippet := `from functools import lru_cache
from xonsh.completers.completer import add_one_completer
from xonsh.completers.tools import contextual_command_completer, sub_proc_get_output, RichCompletion
import json
import os

%v%v


@lru_cache
def get_supported_cmds():
    output, _ = sub_proc_get_output("carapace", "--list")
    if not output:
        return []
    return {line.split()[0] for line in output.decode().splitlines()}


@contextual_command_completer
def _carapace_lazy(context):
    """carapace lazy"""

    supported_cmds = get_supported_cmds()
    if context.command not in supported_cmds:
        return

    def fix_prefix(s):
        """quick fix for partially quoted prefix completion ('prefix',<TAB>)"""
        return s.translate(str.maketrans('', '', '\'"'))

    output, _ = sub_proc_get_output(
        'carapace', context.command, 'xonsh', *[a.value for a in context.args], fix_prefix(context.prefix)
    )
    if not output:
        return

    for c in json.loads(output):
        yield RichCompletion(
            c["Value"],
            display=c["Display"],
            description=c["Description"],
            prefix_len=len(context.raw_prefix),
            append_closing_quote=False,
            style=c["Style"],
        )
`
	complete := make([]string, len(completers))
	for index, completer := range completers {
		complete[index] = fmt.Sprintf(`'%v'`, completer)
	}
	snippet += `add_one_completer('carapace_lazy', _carapace_lazy, '>xompleter')`
	return fmt.Sprintf(snippet, pathSnippet("xonsh"), envSnippet("xonsh"))
}
