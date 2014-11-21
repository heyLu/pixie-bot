package main

import (
	"fmt"

	"os"
	"os/exec"
)

func main() {
	out, err := runPixie("(+ 3 4)")
	if err != nil {
		fmt.Println("Error: ", string(out))
		os.Exit(1)
	}
	fmt.Print(string(out))
}

func runPixie(expr string) ([]byte, error) {
	tmpl := "(print (eval (read-string \"%s\")))"
	expr = fmt.Sprintf(tmpl, expr)

	fmt.Printf("; running: %s\n", expr)
	cmd := exec.Command("/usr/bin/docker", "run", "-i", "--rm", "wunderseltsam/pixie", "-e", expr)
	return cmd.CombinedOutput()
}
