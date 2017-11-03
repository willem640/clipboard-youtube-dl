package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/BrianAllred/goydl"
	"github.com/tjgq/clipboard"
)

func main() {
	f()
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from bad URL", r)
		}
	}()
	c := make(chan string, 1)
	youtubeDl := goydl.NewYoutubeDl()
	// go io.Copy(os.Stdout, youtubeDl.Stdout)
	// go io.Copy(os.Stderr, youtubeDl.Stderr)
	for {
		clipboard.Notify(c)
		got := <-c
		println("got: ", got)
		_, err := url.Parse(got)
		if err != nil {
			fmt.Println("URL recieved not valid: ", err)
			continue
		}
		cmd, err := youtubeDl.Download(got)
		if err != nil {
			log.Print("that wasn't a good URL, ", err)
		} else {
			println("starting download")
			cmd.Wait()
			println("finished download")
		}
	}
}
