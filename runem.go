package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func run(s string) {
	fmt.Printf("Package : %s\n", s)
	os.Chdir(s)
	out, err := exec.Command("go", "test").Output()
	if err != nil {
		log.Fatal("No test files...nothing to see here!")
	}
	fmt.Printf("%s", out)
	os.Chdir("..")
}

func main() {
	cmd := exec.Command("bash", "-c", "ls -l | grep '^d' | awk '{print $9}' > gotests")
	cmd.Run()

	file, err := os.Open("gotests")
	if err != nil {
		fmt.Println("Failed opening file.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		run(scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
