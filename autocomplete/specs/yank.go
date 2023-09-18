// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["yank"] = model.Subcommand{
		Name:        []string{"yank"},
		Description: `Yank terminal output to clipboard`,
		Options: []model.Option{{
			Name:        []string{"-i"},
			Description: `Ignore case differences between pattern and the input`,
		}, {
			Name:        []string{"-l"},
			Description: `Use the default delimiters except for space`,
		}, {
			Name:        []string{"-x"},
			Description: `Use alternate screen`,
		}, {
			Name:        []string{"-v"},
			Description: `Print the version`,
		}, {
			Name:        []string{"-d"},
			Description: `All input characters not present in delim will be recognized as fields`,
			Args: []model.Arg{{
				Name:        "delim",
				Description: `Custom delimiters`,
			}},
		}, {
			Name:        []string{"-g"},
			Description: `Use pattern to recognize fields`,
			Args: []model.Arg{{
				Name:        "pattern",
				Description: `Pattern to recognize fields`,
			}},
		}, {
			Name:        []string{"--"},
			Description: `Use a command as the yank command`,
			Args: []model.Arg{{
				Name:        "command",
				Description: `Command to use as the yank command`,
				IsVariadic:  true,
			}},
		}},
	}
}