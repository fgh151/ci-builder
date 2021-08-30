package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func ComposeUp(path string) {

	createEnvFile(path)

	runCommand("docker-compose", "--project-directory", path, "down", "-v")

	var args []string

	args = append(args, "-f")
	args = append(args, path+"/docker-compose.yml")

	if _, err := os.Stat(path + "/docker-compose-development.yml"); err == nil {
		args = append(args, "-f")
		args = append(args, path+"/docker-compose-development.yml")
	}

	args = append(args, "--project-directory")
	args = append(args, path)
	args = append(args, "up")
	args = append(args, "-d")
	args = append(args, "--force-recreate")
	args = append(args, "--remove-orphans")

	runCommand("docker-compose", args...)
}

func runCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	checkErr(err)
}

func createEnvFile(path string) {
	fmt.Println("Создаем .env файл из примера")
	//Read all the contents of the  original file
	bytesRead, err := ioutil.ReadFile(path + "/.env.example")
	if err != nil {
		log.Fatal(err)
	}

	//Copy all the contents to the desitination file
	err = ioutil.WriteFile(path+"/.env", bytesRead, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
