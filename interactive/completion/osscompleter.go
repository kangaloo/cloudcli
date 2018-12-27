package completion

import "github.com/c-bata/go-prompt"

func ossListCompleter(args []string) []prompt.Suggest {

	if len(args) > 3 {

		switch args[len(args)-2] {
		case "-b":
			// bucket completer
			return nil
		case "--prefix":
			// prefix completer
			return nil
		}
	}

	return prompt.FilterHasPrefix(
		[]prompt.Suggest{
			{Text: "-b", Description: "bucket name"},
			{Text: "-q", Description: "do not show the header and object size"},
			{Text: "-n", Description: "number of listed objects (default: 10)"},
			{Text: "--prefix", Description: "list objects by prefix"},
			{Text: "--all", Description: "list all objects in a bucket"},
		},
		args[len(args)-1],
		true,
	)
}
