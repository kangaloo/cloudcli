# build

go build cloudcli.go

# run a command

./cloudcli -h
NAME:
   cloudcli - Aliyun API command line tool

USAGE:
   cloudcli \[global options\] command \[command options] \[arguments...]

VERSION:
   1.0.0

AUTHOR:
   Li Xiangyang <lixy4@belink.com>

COMMANDS:
     oss      aliyun OSS API tool
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -c file                           specify the configuration file
   -d                                debug
   --endpoint endpoint, -e endpoint  Aliyun API endpoint
   --ak accessKey                    Aliyun API accessKey
   --aks accessKeySecret             Aliyun API accessKeySecret
   --help, -h                        show help
   --version, -v                     print the version

./cloudcli oss -h
NAME:
   cloudcli oss - aliyun OSS API tool

USAGE:
   cloudcli oss command \[command options] \[arguments...]

COMMANDS:
     upload, ul            upload files to a oss bucket
     download, dl          download objects from oss
     list, ls              list all objects in a bucket
     list_bucket, lsbk     list all objects in a bucket
     create, ct            create bucket
     delete_bucket, delbk  delete a bucket
     delete, del           delete object

OPTIONS:
   --help, -h  show help
