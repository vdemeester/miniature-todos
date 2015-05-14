package main

import (
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "miniatures-todo"
	app.Usage = "Run a TODOs webservice (back in memory)"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "bind, b",
			Value:  "127.0.0.1:9090",
			Usage:  "host string to bind on",
			EnvVar: "MIMIATURE_TODOS_HOST",
		},
	}

	app.Action = func(c *cli.Context) {
		host := "127.0.0.1:9090"
		if c.IsSet("bind") {
			host = c.String("bind")
		}
		router := NewRouter()
		log.Printf("Serving HTTP on %s", host)
		log.Fatal(http.ListenAndServe(host, router))
	}

	app.Run(os.Args)
}
