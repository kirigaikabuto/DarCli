package main

import (
	"fmt"
	yt "github.com/kkdai/youtube"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	directory string
)

func downloadFileFromUrl(url string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	currentDir = currentDir + "/videos"
	if directory != "" {
		currentDir = directory
	}
	y := yt.NewYoutube(false, true)
	if err := y.DecodeURL(url); err != nil {
		log.Fatal(err)
	}
	fileName := fmt.Sprintf("%s_%s.mp4", y.Title, y.Author)
	if err := y.StartDownload(currentDir, fileName, "high", 0); err != nil {
		log.Fatal(err)
	}
	return nil
}

func main() {
	cli.VersionFlag = cli.BoolFlag{
		Name:  "print-version, V",
		Usage: "print only the version",
	}
	app := cli.NewApp()
	app.Name = "Template"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "upload_directory, u_dir",
			Value:       "",
			Usage:       "directory where upload files",
			Destination: &directory,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "download",
			Aliases: []string{"d"},
			Usage:   "download something by url",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "url, u"},
				cli.StringFlag{Name: "file, f"},
			},
			Action: func(c *cli.Context) error {
				url := c.String("url")
				file := c.String("file")
				if url != "" {
					err := downloadFileFromUrl(url)
					if err != nil {
						return err
					}
				} else if file != "" {
					data, err := ioutil.ReadFile(file)
					if err != nil {
						log.Fatal(err)
					}
					urls := strings.Split(string(data), "\n")
					for i, v := range urls {
						fmt.Println(i, v)
						v = strings.TrimSpace(v)
						err = downloadFileFromUrl(v)
						if err != nil {
							return err
						}
					}
				}
				return nil

			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
