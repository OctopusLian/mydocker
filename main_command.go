package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"mydocker/container"
)

var runCommand = cli.Command{
	Name: "run",
	Usage: `Create a container with namespace and cgroups limit
			mydocker run -ti [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 { // 判断参数是否包含command
			return fmt.Errorf("Missing container command")
		}
		cmd := context.Args().Get(0) // 获取用户指定的command
		tty := context.Bool("ti")
		Run(tty, cmd) // 启动容器
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		logrus.Infof("init come on")
		cmd := context.Args().Get(0) // 获取传递过来的 command 参数
		logrus.Infof("command %s", cmd)
		err := container.RunContainerInitProcess(cmd, nil) // 执行容器初始化操作
		return err
	},
}
