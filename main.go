package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Lockfile struct {
	Packages []string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing path argument")
	}
	yarn := os.Args[1]
	content, err := ioutil.ReadFile(yarn)
	contentString := string(content)
	if err != nil {
		log.Fatal(err)
	}

	listOfLines := strings.Split(contentString, "\n")[4:]
	//lockfile := Lockfile{Packages: []string{}}
	amountOfLines := len(listOfLines)
	for i := 0; i < amountOfLines; {
		line := listOfLines[i]
		if len(line) > 2 && line[:2] != "  " && string(line[len(line)-1]) == ":" {
			fmt.Println("package found", line)
		} else if line != "" {
			items := strings.Split(line[2:], " ")
			if items[0] == "dependencies:" {
				fmt.Println("\nDependencies found")
				count := getDependencies(listOfLines[i+1:])
				i += count
			} else {
				fmt.Println("key ", items[0])
				fmt.Println("value ", items[1])
			}
		} else {
			fmt.Println("Package ended.\n")
		}
		i += 1
	}
}

func getDependencies(lines[]string) int {
	count := 0
	for _, line := range lines {
		if string(line) == "" {
			break
		} else {
			items := strings.Split(string(line)[4:], " ")
			fmt.Println("Dependency name:", items[0])
			fmt.Println("Dependency version:", items[1])
			count++
		}
	}
	return count
}
