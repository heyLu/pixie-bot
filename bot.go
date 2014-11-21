package main

import (
	"fmt"

	"os"
	"os/exec"
)

func main() {
	pixieStr := "(+ 3 4)"
	pixieTmpl := "(print (eval (read-string \"%s\")))"
	pixieExpr := fmt.Sprintf(pixieTmpl, pixieStr)

	fmt.Printf("; running: %s\n", pixieExpr)
	cmd := exec.Command("/usr/bin/docker", "run", "-i", "--rm", "wunderseltsam/pixie", "-e", pixieExpr)
	out, err := cmd.CombinedOutput()
	if err !=  nil {
		fmt.Println("Error: ", err, string(out))
		os.Exit(1)
	}
	fmt.Print(string(out))
}
