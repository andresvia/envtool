// This package makes use of github.com/andresvia/envlib
package main

import (
	"github.com/andresvia/envlib"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {

	app := cli.NewApp()

	app.Commands = []cli.Command{
		cli.Command{
			Name: "persist",
			Flags: []cli.Flag{
				cli.StringSliceFlag{
					Name:  "select",
					Value: &cli.StringSlice{},
				},
				cli.StringSliceFlag{
					Name:  "edit",
					Value: &cli.StringSlice{},
				},
			},
			Action: func(ctx *cli.Context) error {
				environ := []string{}
				for _, sel := range ctx.StringSlice("select") {
					environ = append(environ, sel+"="+os.Getenv(sel))
				}
				for _, edit := range ctx.StringSlice("edit") {
					envlib.Init(environ)
					envlib.Select(envlib.NewNameset(ctx.StringSlice("select")...))
					envlib.MergeWith(edit)
					envlib.Persist(edit)
				}
				return nil
			},
		},
	}

	app.Run(os.Args)

}
