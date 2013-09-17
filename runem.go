package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func run(s string) {
	fmt.Printf("\033[1m[ Package : %s ]\033[0m\n", s)
	os.Chdir(s)
	out, err := exec.Command("go", "test").Output()
	if err != nil {
		log.Fatal("No test files...nothing to see here!")
	}
	fmt.Printf("%s", out)
	os.Chdir("..")
}

func main() {
	cmd := exec.Command("bash", "-c", "ls -l | grep '^d' | awk '{print $9}'")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		run(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
}
