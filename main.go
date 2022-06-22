// create a directory wuth the argument specified
// execute go mod github.com/deanvlue/[name of the app]
// create a main.go file
// add a simple hello world

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	arguments := os.Args

	if len(arguments) <= 1 {
		log.Fatal("gonew: Generates a new Go Lang Project\n Usage: gonew APP_NAME ")
	}

	project_name := arguments[1]

	// Obtain arguments
	path_strings := []string{"./", project_name}
	path_folder := strings.Join(path_strings, "")
	fmt.Printf("Creating %s\n", path_folder)
	// Validating if the folder exists
	err := os.Mkdir(path_folder, 0755)
	if err != nil {
		log.Println(err)
	}

	// Create the main.go file
	path_main_file := path_folder + "/main.go"

	fl, err := os.Create(path_main_file)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Creando main.go")
	defer fl.Close()

	go_program := []string{"package main\n", "import \"fmt\" \n", "func main(){", "\tfmt.Println(\"Hola Charly!\")\n", "}"}

	for _, line := range go_program {
		fl.WriteString(line + "\n")
	}

	// Create the go.mod file
	change_directory := fmt.Sprintf("./%s", project_name)
	fmt.Println(change_directory)
	mod_command := fmt.Sprintf("github.com/deanvlue/%s", project_name)
	out := exec.Command("go", "mod", "init", mod_command)
	out.Dir = change_directory
	err = out.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

	log.Println("Initializing git")
	git := exec.Command("git", "init")
	git.Dir = change_directory
	err = git.Run()
	if err != nil {
		log.Fatal(err)
	}
}
