package main
import (
	"os"

	"AntarJemput-Be-C/app"
	"AntarJemput-Be-C/cli"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())
}
