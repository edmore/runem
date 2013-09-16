package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

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
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
