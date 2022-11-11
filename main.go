package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Package struct {
	Name         string
	Version      string
	Checksum     string
	Source       string
	Dependencies []SubPackage
}
type SubPackage struct {
	Name    		string
	Version 		string
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
	amountOfLines := len(listOfLines)
	lockfile := make([]Package, 0)
	var pkg Package
	for i := 0; i < amountOfLines; {
		line := listOfLines[i]
		if len(line) > 2 && line[:2] != "  " && string(line[len(line)-1]) == ":" {
			fmt.Println("package found", line)
			pkg.Name = removeVersionFromName(line)
		} else if line != "" {
			items := strings.Split(line[2:], " ")
			if items[0] == "dependencies:" {
				fmt.Println("\n Dependencies found")
				count, subPkgs := getDependencies(listOfLines[i+1:])
				fmt.Println(subPkgs)
				pkg.Dependencies = subPkgs
				i += count
			} else {
				fmt.Println("key ", items[0])
				fmt.Println("value ", items[1])
				if items[0] == "integrity" {
					pkg.Checksum = items[1]
				}
				if items[0] == "version" {
					pkg.Version = items[1]
				}
				if items[0] == "resolved" {
					pkg.Source = items[1]
				}
			}
		} else {
			fmt.Println("Package ended.\n")
			lockfile = append(lockfile, pkg)
			pkg = Package{}
		}
		i += 1
	}
	fmt.Println("---------------------------------------------------------")
	fmt.Println(lockfile)
	fmt.Println("---------------------------------------------------------")
}
func getDependencies(lines []string) (int, []SubPackage) {
	count := 0
	subPkgs := make([]SubPackage, 0)
	var subPkg SubPackage
	for _, line := range lines {
		if string(line) == "" {
			break
		} else {
			items := strings.Split(string(line)[4:], " ")
			fmt.Println("dependency name", items[0])
			fmt.Println("dependency version", items[1])
			count++
			subPkg.Name = items[0]
			subPkg.Version = items[1]
			subPkgs = append(subPkgs, subPkg)
			subPkg = SubPackage{}
		}
	}
	return count, subPkgs
}
func removeVersionFromName(line string) string {
	return strings.Split(line, "@")[0]
}
