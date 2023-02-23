package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var version string
var buildstamp string

func main() {
	app := &cli.App{
		Name:  "个人命令工具箱",
		Usage: "Happy life use this tool",
		Commands: []*cli.Command{
			{
				Name:  "env",
				Usage: "环境相关",
				Subcommands: []*cli.Command{
					{
						Name:        "alias",
						Aliases:     []string{"a"},
						Usage:       "检查本地环境",
						Description: "运行本地环境检查,检查本地go版本与各种基础配置",
						Action: func(c *cli.Context) error {
							return AddAlias(c.Args().Get(0))
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
