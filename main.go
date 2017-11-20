package main

import (
	"log"

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
		// path := fmt.Sprintf("~/Desktop/%s.mp4", info.Title)
		// println("saving to ", path)
		// youtubeDl.Options.Output.Value = path
		cmd, err := youtubeDl.Download(got)
		youtubeDl.Options.Output.Value = info.Title

		println("starting download")
		cmd.Wait()
		println("finished download")
	}
}
