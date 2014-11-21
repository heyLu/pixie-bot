package main

import (
	"flag"
	"fmt"
	"strings"

	"os"
	"os/exec"

	irc "github.com/fluffle/goirc/client"
)

var name = flag.String("name", "pixie-bot", "the name of the bot")
var channel = flag.String("channel", "#pixie-lang", "the channel to connect to")

func main() {
	flag.Parse()

	c := irc.SimpleClient(*name)

	c.HandleFunc("connected", func(conn *irc.Conn, line *irc.Line) {
		fmt.Println("connected, hi there.")
		conn.Join(*channel)
	})

	quit := make(chan bool)
	c.HandleFunc("disconnected", func(conn *irc.Conn, line *irc.Line) {
		fmt.Println("disconnected, bye.")
		quit <- true
	})

	c.HandleFunc("privmsg", func(conn *irc.Conn, line *irc.Line) {
		fmt.Println("privmsg in ", line.Target(), ": ", line.Text())
		text := strings.TrimSpace(line.Text())
		if strings.HasPrefix(text, ",") {
			out, err := runPixie(text[1:])
			if err != nil {
				conn.Privmsg(line.Target(), fmt.Sprint("error: ", string(out)))
			} else {
				conn.Privmsg(line.Target(), fmt.Sprint("=> ", string(out)))
			}
		}
	})

	if err := c.ConnectTo("irc.freenode.net"); err != nil {
		fmt.Printf("Error connecting: %s\n", err)
		os.Exit(1)
	}

	<-quit
}

func runPixie(expr string) ([]byte, error) {
	tmpl := "(print (eval (read-string \"%s\")))"
	expr = fmt.Sprintf(tmpl, expr)

	fmt.Printf("; running: %s\n", expr)
	cmd := exec.Command("/usr/bin/docker", "run", "-i", "--rm", "-u", "nobody", "wunderseltsam/pixie", "-e", expr)
	return cmd.CombinedOutput()
}
