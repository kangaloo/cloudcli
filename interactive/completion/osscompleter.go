package completion

import (
	"github.com/c-bata/go-prompt"
	"github.com/kangaloo/cloudcli/interactive/cloud"
)

// todo 自动完成提示中排除已经输入过的参数

func ossListCompleter(args []string) []prompt.Suggest {

	// todo 可用共用函数代替
	if len(args) > 3 {
		switch args[len(args)-2] {
		case "-b":
			// bucket completer
			return prompt.FilterHasPrefix(cloud.BucketCompleter(), args[len(args)-1], true)
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

func ossDelCompleter(args []string) []prompt.Suggest {

	if len(args) > 3 {
		switch args[len(args)-2] {
		case "-b":
			return nil
		case "-o":
			return nil
		case "-n":
			return nil
		case "--prefix":
			return nil
		}
	}

	return prompt.FilterHasPrefix(
		[]prompt.Suggest{
			{Text: "-b", Description: "bucket name"},
			{Text: "-o", Description: "delete one object by name"},
			{Text: "-n", Description: "delete a specified number of objects (default: 10)"},
			{Text: "--prefix", Description: "delete objects by prefix"},
			{Text: "--all", Description: "delete all objects in a bucket"},
			{Text: "--del", Description: "delete the list objects in a bucket"},
		},
		args[len(args)-1],
		true,
	)
}

func ossCreateCompleter(args []string) []prompt.Suggest {

	if len(args) > 3 {
		return nil
	}

	return prompt.FilterHasPrefix(
		[]prompt.Suggest{},
		args[len(args)-1],
		true,
	)
}

func ossDownloadCompleter(args []string, d prompt.Document) []prompt.Suggest {

	if len(args) > 3 {
		switch args[len(args)-2] {
		case "-b", "-o", "-p":
			return nil
		case "-f":
			return fileSystemCompleter(d)
		}
	}

	return prompt.FilterHasPrefix(
		[]prompt.Suggest{
			{Text: "-b", Description: "bucket name"},
			{Text: "-o", Description: "object name"},
			{Text: "-f", Description: "download object as file"},
			{Text: "-p", Description: "prat size number (default: 102400)"},
			{Text: "--overwrite", Description: "overwrite a exist file when download"},
		},
		args[len(args)-1],
		true,
	)
}
