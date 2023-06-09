package completion

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// taken from https://github.com/urfave/cli/blob/master/autocomplete/zsh_autocomplete
var zsh_script = `
#compdef temporal
_cli_zsh_autocomplete() {
  local -a opts
  local cur
  cur=${words[-1]}
  if [[ "$cur" == "-"* ]]; then
    opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} ${cur} --generate-bash-completion)}")
  else
    opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} --generate-bash-completion)}")
  fi
  if [[ "${opts[1]}" != "" ]]; then
    _describe 'values' opts
  else
    _files
  fi
  return
}
compdef _cli_zsh_autocomplete temporal
`

var bash_script = `
#! /bin/bash
_cli_bash_autocomplete() {
  if [[ "${COMP_WORDS[0]}" != "source" ]]; then
    local cur opts base
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    if [[ "$cur" == "-"* ]]; then
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} ${cur} --generate-bash-completion )
    else
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --generate-bash-completion )
    fi
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
  fi
}
complete -o bashdefault -o default -o nospace -F _cli_bash_autocomplete temporal
`

func NewCompletionCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:      "bash",
			Usage:     "bash completion output",
			UsageText: "source <(temporal completion bash)",
			Action: func(c *cli.Context) error {
				fmt.Fprintln(os.Stdout, bash_script)
				return nil
			},
		},
		{
			Name:      "zsh",
			Usage:     "zsh completion output",
			UsageText: "source <(temporal completion zsh)",
			Action: func(c *cli.Context) error {
				fmt.Fprintln(os.Stdout, zsh_script)
				return nil
			},
		},
	}
}
