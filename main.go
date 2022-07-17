package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// if the more than one class is found give an option which namespace (find the namespace somehow)
// loop through class file starting at (if line contains class XClass start creating assertions)
// print out all the assertions

func main() {
	var className string

	fmt.Println("Enter Class Name")

	for {
		fmt.Scanf("%s", &className)

		if len(className) != 0 {
			break
		}
	}

	fmt.Println("")

	pattern := "class " + className

	rgCmd := exec.Command("/usr/bin/rg", pattern, "/home/stanzu10/Repos/repos2/Bridge/")

	result, err := rgCmd.CombinedOutput()

	res := string(result)

	if err != nil {
		fmt.Printf("Error running ripgrep")
	}

	filePath := string(res)[:strings.Index(res, ":")]

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Error opening filepath")
	}

	scanner := bufio.NewScanner(file)

	propCount := 0
	for scanner.Scan() {
		t := scanner.Text()
		if containsString(t, "public") && !containsString(t, "class") {
			arr := strings.Split(t, " ")
			arr = removeEmpty(arr)

			if len(arr) > 2 {
				idx := indexOf(arr, "public")

				if idx+2 > len(arr) {
					continue
				}

				propName := arr[idx+2]

				assertion := "Assert.Equal(one." + propName + ", two." + propName + ");"

				fmt.Println(assertion)

				propCount++
			}
		}
	}

	fmt.Printf("\nFound and created %d Assertions", propCount)
}

func indexOf(arr []string, pattern string) int {
	for i, v := range arr {

		if v == pattern {
			return i
		}
	}

	return -1
}
func containsString(s string, pattern string) bool {

	arr := strings.Split(s, " ")

	for _, v := range arr {
		if v == pattern {
			return true
		}
	}
	return false
}

func removeEmpty(arr []string) []string {
	var output []string

	for _, v := range arr {
		if len(strings.TrimSpace(v)) != 0 {
			output = append(output, v)
		}
	}

	return output
}
