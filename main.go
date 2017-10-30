package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/hpcloud/tail"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "rtail"
	app.Usage = "serve a file, or stdin to http"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "f,follow",
			Usage: "follow a file",
		},
		cli.BoolFlag{
			Name:  "V,verbose",
			Usage: "enable verbose output",
		},
		cli.StringFlag{
			Name:  "addr",
			Value: "localhost:8000",
			Usage: "listen addr ie localhost:8000",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.Bool("V") {
			log.SetLevel(log.DebugLevel)
			log.Debug("Setting Verbose")
		}
		addr := c.String("addr")
		lines := make(chan string)
		if c.Bool("f") {
			fmt.Println(1, addr)
			filename := c.Args().First()
			log.WithField("filename", filename).Debug("Opening file for reading")
			t, err := tail.TailFile(filename, tail.Config{Follow: true})
			if err != nil {
				log.WithError(err).Errorf("Tail failed: %v", err)
				return err
			}
			go func() {
				for line := range t.Lines {
					lines <- line.Text
				}
			}()
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Split(bufio.ScanLines)
			go func() {
				for scanner.Scan() {
					lines <- scanner.Text()
				}
			}()
		}
		h := func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "text/plain")
			for line := range lines {
				if _, err := io.Copy(w, strings.NewReader(line+"\n")); err != nil {
					log.WithError(err).Error("IO Copy failed")
					return
				}
				w.(http.Flusher).Flush()
			}

		}
		return http.ListenAndServe(addr, http.HandlerFunc(h))
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
