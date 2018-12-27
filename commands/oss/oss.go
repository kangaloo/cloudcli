package oss

import (
	"github.com/kangaloo/cloudcli/cloud/oss"
	"github.com/urfave/cli"
)

var Oss = &cli.Command{
	Name:        "oss",
	Usage:       "aliyun OSS API tool",
	Subcommands: ossSubCmds,
}

var ossSubCmds = cli.Commands{
	*ossUpload,
	*ossDownload,
	*ossList,
	*ossListBucket,
	*ossCreate,
	*ossDelBucket,
	*ossDel,
}

var ossUpload = &cli.Command{
	Name:      "upload",
	ShortName: "ul",
	Usage:     "upload files to a oss bucket",
	Flags:     ossUploadFlags,
	Action:    oss.Upload,
}

var ossUploadFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "b",
		Usage: "`bucket` name",
	},
	cli.IntFlag{
		Name:  "p",
		Usage: "part size `number`",
		Value: 1024 * 100,
	},
	cli.StringFlag{
		Name: "o",
		Usage: "specify the `object` name when upload a single file; " +
			"if not specified, same as the file name",
	},
	cli.StringFlag{
		Name:  "f",
		Usage: "upload the specified `file`",
	},
	cli.StringFlag{
		Name:  "prefix",
		Usage: "objects `prefix`",
	},
	cli.StringFlag{
		Name:  "r",
		Usage: "upload directories and files in specified `directory`",
	},
	cli.BoolFlag{
		Name:  "R",
		Usage: "upload directories and files in current work directory",
	},
	cli.BoolFlag{
		Name:  "overwrite",
		Usage: "overwrite a file already in a bucket",
	},
}

var ossListBucket = &cli.Command{
	Name:      "list_bucket",
	ShortName: "lsbk",
	Usage:     "list all objects in a bucket",
	Action:    oss.ListBucket,
}

var ossList = &cli.Command{
	Name:      "list",
	ShortName: "ls",
	Usage:     "list all objects in a bucket",
	Flags:     ossListFlags,
	Action:    oss.ListObjects,
}

var ossListFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "b",
		Usage: "`bucket` name",
	},
	cli.BoolFlag{
		Name:  "q",
		Usage: "do not show the header and object size",
	},
	cli.IntFlag{
		Name:  "n",
		Usage: "`number` of listed objects",
		Value: 10,
	},
	cli.StringFlag{
		Name:  "prefix",
		Usage: "list objects by `prefix`",
	},
	cli.BoolFlag{
		Name:  "all",
		Usage: "list all objects in a bucket",
	},
	cli.StringFlag{
		Name:  "marker",
		Usage: "oss `marker`",
	},
}

var ossDownload = &cli.Command{
	Name:      "download",
	ShortName: "dl",
	Usage:     "download objects from oss",
	Flags:     ossDownloadFlags,
	Action:    oss.Download,
}

var ossDownloadFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "b",
		Usage: "`bucket` name",
	},
	cli.StringFlag{
		Name:  "o",
		Usage: "`object` name",
	},
	cli.StringFlag{
		Name:  "f",
		Usage: "download object as `file`",
	},
	cli.IntFlag{
		Name:  "p",
		Usage: "prat size `number`",
		Value: 1024 * 100,
	},
	cli.BoolFlag{
		Name:  "overwrite",
		Usage: "overwrite a exist file when download",
	},
}

var ossCreate = &cli.Command{
	Name:      "create",
	ShortName: "ct",
	Usage:     "create bucket",
	Flags:     ossCreateFlags,
	Action:    oss.CreateBucket,
}

var ossCreateFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "b",
		Usage: "`bucket` name",
	},
}

var ossDelBucket = &cli.Command{
	Name:      "delete_bucket",
	ShortName: "delbk",
	Usage:     "delete a bucket",
	Flags:     ossDelBkFlags,
	Action:    oss.DelBucket,
}

var ossDelBkFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "b",
		Usage: "`bucket` name",
	},
}

var ossDel = &cli.Command{
	Name:      "delete",
	ShortName: "del",
	Usage:     "delete object",
	Flags:     ossDelFlags,
	Action:    oss.DelObject,
}

var ossDelFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "b",
		Usage: "`bucket` name",
	},
	cli.StringFlag{
		Name:  "o",
		Usage: "delete one `object` by name",
	},
	cli.IntFlag{
		Name:  "n",
		Usage: "delete a specified `number` of objects",
		Value: 10,
	},
	cli.StringFlag{
		Name:  "prefix",
		Usage: "delete objects by `prefix`",
	},
	cli.BoolFlag{
		Name:  "all",
		Usage: "delete all objects in a bucket",
	},
	cli.BoolFlag{
		Name:  "del",
		Usage: "delete the list objects in a bucket",
	},
}
