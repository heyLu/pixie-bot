package main

import (
	"fmt"

	"os"
	"os/exec"

	irc "github.com/fluffle/goirc/client"
)

func main() {
	c := irc.SimpleClient("pixie-bot")

	c.HandleFunc("connected", func(conn *irc.Conn, line *irc.Line) {
		fmt.Println("connected, hi there.")
		conn.Join("#dotdotdot")
		conn.Privmsg("#dotdotdot", "may your evals be magnificient")
	})

	quit := make(chan bool)
	c.HandleFunc("disconnected", func(conn *irc.Conn, line *irc.Line) {
		fmt.Println("disconnected, bye.")
		quit <- true
	})

	c.HandleFunc("privmsg", func(conn *irc.Conn, line *irc.Line) {
		fmt.Println("privmsg in ", line.Target(), ": ", line.Text())
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
	cmd := exec.Command("/usr/bin/docker", "run", "-i", "--rm", "wunderseltsam/pixie", "-e", expr)
	return cmd.CombinedOutput()
}
