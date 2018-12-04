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
		Usage: "prat size `number`",
		Value: 1024 * 100,
	},
	cli.BoolFlag{
		Name:  "o",
		Usage: "specify the `object` name when upload a single file",
	},
	cli.BoolFlag{
		Name:  "r",
		Usage: "upload directories recursively",
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
	Flags:     ossListBkFlags,
	Action:    oss.ListBucket,
}

var ossListBkFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "b",
		Usage: "`bucket` name",
	},
	cli.BoolFlag{
		Name:  "s",
		Usage: "show size",
	},
	cli.IntFlag{
		Name:  "n",
		Usage: "number of objects",
	},
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
		Usage: "do not show the header",
		//	注意 -s 参数对表头的影响
	},
	cli.BoolTFlag{
		Name:  "s",
		Usage: "show object size",
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
	cli.StringFlag{
		Name:  "suffix",
		Usage: "list objects by `suffix`",
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
	cli.StringFlag{
		Name:  "prefix",
		Usage: "delete objects by `prefix`",
	},
	cli.StringFlag{
		Name:  "suffix",
		Usage: "delete objects by `suffix`",
	},
	cli.BoolFlag{
		Name:  "all",
		Usage: "delete all objects in a bucket",
	},
}
