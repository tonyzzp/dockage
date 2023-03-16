package cmds

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v3"
)

func cmdStartAction(ctx *cli.Context) error {
	log.Println("---------")
	log.Println(ctx.Command.Name)
	log.Println(ctx.Args())
	log.Println("dir", ctx.String("dir"))
	log.Println("port", ctx.Int("port"))
	var server = gin.New()
	server.Any("/", func(ctx *gin.Context) {
		ctx.String(200, "hello gin")
	})
	return server.Run(":" + strconv.Itoa(ctx.Int("port")))
}

var CmdStart = &cli.Command{
	Name:        "start",
	Action:      cmdStartAction,
	Description: "启动服务端",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "dir",
			Value: "",
			Usage: "工作目录",
		},
		&cli.IntFlag{
			Name:  "port",
			Value: 8000,
			Usage: "http端口",
		},
	},
}
