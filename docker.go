package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func up(path string) {

	createEnvFile(path)

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

	cmd := exec.Command("docker-compose", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	checkerr(err)
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
