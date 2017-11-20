package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/TableMountain/goydl"
	"github.com/tjgq/clipboard"
)

func main() {
	f()
}

func f() {
	c := make(chan string, 1)
	youtubeDl := goydl.NewYoutubeDl()

	// go io.Copy(os.Stdout, youtubeDl.Stdout)
	// go io.Copy(os.Stderr, youtubeDl.Stderr)
	for {
		clipboard.Notify(c)
		got := <-c
		youtubeDl.VideoURL = got
		info, err := youtubeDl.GetInfo()
		if err != nil {
			log.Print(err)
			continue
		}
		println("got: ", got, "\nas: ", info.Title)
		_, err = url.Parse(got)
		if err != nil {
			fmt.Println("URL recieved not valid: ", err)
			continue
		}
		path := fmt.Sprintf("~/Desktop/%s",info.Title)
		youtubeDl.Options.Output.Value = path
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from bad URL", r)
				}
			}()
			cmd, err := youtubeDl.Download(got)
			youtubeDl.Options.Output.Value = info.Title
			if err != nil {
				log.Print("that wasn't a good URL, ", err)
			} else {
				println("starting download")
				cmd.Wait()
				println("finished download")
			}
		}()
	}
}
