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

func ossUploadCompleter(args []string, d prompt.Document) []prompt.Suggest {

	if len(args) > 3 {
		switch args[len(args)-2] {
		case "-b":
			// bucket completer
			return nil
		case "-o", "--prefix":
			return nil
		case "-f":
			return fileSystemCompleter(d)
		case "-r":
			return pathCompleter(d)
		}
	}

	return prompt.FilterHasPrefix(
		[]prompt.Suggest{
			{Text: "-b", Description: "bucket name"},
			{Text: "-o", Description: "specify the object name when upload a single file; default same as the file name"},
			{Text: "-f", Description: "upload the specified file"},
			{Text: "-r", Description: "upload directories and files in specified directory"},
			{Text: "-R", Description: "upload directories and files in current work directory"},
			{Text: "--prefix", Description: "objects prefix"},
			{Text: "--overwrite", Description: "overwrite file already in a bucket"},
		},
		args[len(args)-1],
		true,
	)
}
