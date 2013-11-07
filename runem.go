// Runem - tool for running all your tests from a project root directory.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
)

const (
	Reset  = "\x1b[0m"
	Bright = "\x1b[1m"
	FgRed  = "\x1b[31m"
)

func runTest(s string) {
	fmt.Printf(Bright+"[ %s ]"+Reset+"\n", s)
	os.Chdir(s)

	out, _ := exec.Command("go", "test").Output()
	r, _ := regexp.Compile("FAIL")
	if r.Match(out) {
		fmt.Printf(FgRed+"%s"+Reset, out)
	} else {
		fmt.Printf("%s", out)
	}

	if s != "." {
		os.Chdir("..")
	}
}

func main() {
	cmd_string := `ls -al | \
                       awk '$1 ~ /^d/ && $NF ~ /(^[.]$|^[a-zA-Z0-9]+$)/ \
                       {print $NF}'`
	cmd := exec.Command("bash", "-c", cmd_string)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	cmd.Start()
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		runTest(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
}
