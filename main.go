package main

import (
	"errors"
	"fmt"
	"github.com/kkdai/youtube"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	path string
	url  string
	file string
)

func DownloadVideo(videoUrl string, wg *sync.WaitGroup) {
	defer wg.Done()
	yt := youtube.NewYoutube(false, true)
	err := yt.DecodeURL(videoUrl)
	if err != nil {
		log.Fatal(err)
	}
	fileName := fmt.Sprintf("%s_%s.mp4", yt.Title, yt.Author)
	err = yt.StartDownload(path, fileName, "high", 0)
	if err != nil {
		log.Fatal(err)
	}
}

func DownloadCommand(c *cli.Context) error {
	if url != "" {
		//err := DownloadVideo(url)
		//if err != nil {
		//	return err
		//}
	} else if file != "" {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		text := string(data)
		urls := strings.Split(text, "\r\n")
		var wg sync.WaitGroup
		for _, v := range urls {
			wg.Add(1)
			go DownloadVideo(v, &wg)
		}
		wg.Wait()
	} else {
		return errors.New("we need flags")
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "Dar app"
	app.Usage = "App for downloading videos from youtube"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path,p",
			Value:       "./videos",
			Usage:       "it is used for path",
			Destination: &path,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "download",
			Aliases: []string{"down"},
			Usage:   "it is command for downloading videos",
			Action:  DownloadCommand,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "url,u",
					Usage:       "it is used for url",
					Destination: &url,
				},
				cli.StringFlag{
					Name:        "file,f",
					Usage:       "file contained urls",
					Destination: &file,
				},
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		a := c.Args().Get(0)
		fmt.Println("First Argument", a)
		b := c.Args().Get(1)
		fmt.Println("Second Argument", b)
		//path := c.String("path")
		fmt.Println("Path flag", path)
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
