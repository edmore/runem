package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func runTest(s string) {
	fmt.Printf("\033[1m[ %s ]\033[0m\n", s)
	os.Chdir(s)
	out, _ := exec.Command("go", "test").Output()
	fmt.Printf("%s", out)
	if s != "." {
		os.Chdir("..")
	}
}

func main() {
	cmd_string := `ls -al | \
                       awk '$1 ~ /^d/ && $NF ~ /(^[.]$|^[a-zA-Z]+$)/ \
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
