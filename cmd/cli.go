package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ansiz/doubanSearchCLI/pkg/crawler"
	"github.com/urfave/cli"
)

var (
	verbose     bool
	bookCrawler *crawler.Crawler
)

func main() {
	app := cli.NewApp()
	app.Name = "Douban searcher"
	app.Usage = "Search data from Douban"
	app.Version = "0.1.0"
	app.Before = appInit
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "verbose",
			Destination: &verbose,
			Usage:       "run in verbose mode",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "search items by specified keyword, output the list",
			Subcommands: []cli.Command{
				{
					Name:  "book",
					Usage: "list books",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:        "verbose, v",
							Usage:       "run in verbose mode",
							Destination: &bookCrawler.Verbose,
						},
						cli.BoolFlag{
							Name:        "local, l",
							Usage:       "run in local mode",
							Destination: &bookCrawler.LocalMode,
						},
						cli.StringFlag{
							Name:        "file, f",
							Usage:       "specify local data file",
							Destination: &bookCrawler.LocalFile,
						},
						cli.IntFlag{
							Name:        "interval, i",
							Usage:       "the request interval, unit: second",
							Destination: &bookCrawler.Interval,
						},
						cli.IntFlag{
							Name:        "start",
							Usage:       "start at",
							Destination: &bookCrawler.Start,
						},
						cli.IntFlag{
							Name:        "page, p",
							Usage:       "total page",
							Destination: &bookCrawler.Page,
							Value:       1,
						},
						cli.StringFlag{
							Name:        "keyword, k",
							Usage:       "the keyword",
							Destination: &bookCrawler.Keyword,
						},
					},
					Action: func(c *cli.Context) error {
						err := bookCrawler.SearchList()
						if err != nil {
							return cli.NewExitError(err, 1)
						}
						return nil
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

func init() {
	var err error
	bookCrawler, err = crawler.New()
	if err != nil {
		log.Fatal(fmt.Sprintf("init crawler failed: %v", err))
		os.Exit(2)
	}
}

func appInit(ctx *cli.Context) error {
	return nil
}
