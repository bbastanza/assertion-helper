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
			fmt.Println("")
			break
		}
	}

	pattern := "class " + className

	rgCmd := exec.Command("/usr/bin/rg", pattern, "/home/stanzu10/Repos/repos2/Bridge/")

	result, err := rgCmd.CombinedOutput()

	res := string(result)

	if err != nil {
		fmt.Printf("Error running ripgrep")
		return
	}

	if !containsChar(string(res), ':') {
		fmt.Printf("No file class found")
		return
	}

	filePath := string(res)[:strings.Index(res, ":")]

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Error opening filepath")
		return
	}

	// create a new scanner with our file
	scanner := bufio.NewScanner(file)

	propCount := 0

	for scanner.Scan() {
		// get the last scanned line
		t := scanner.Text()

		isProp := containsString(t, "public") &&
			!containsString(t, "class") && // this will rule out the class decloration
			!arrContainsChar(strings.Split(t, " "), '(') // this will rule out methods

		if isProp {
			// split our the stdout line and remove any whitespace characters/strings
			arr := strings.Split(t, " ")
			arr = removeEmpty(arr)

			// we are looking for the idx + 2, so if we only have two continue
			if len(arr) <= 2 {
				continue
			}

			// get our index
			idx := indexOf(arr, "public")

			// get our property name which is two over from public
			// public sting PropertyName { get; set; }
			propName := arr[idx+2]

			assertion := "Assert.Equal(one." + propName + ", two." + propName + ");"

			fmt.Println(assertion)

			propCount++
		}
	} // End loop

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

func containsChar(s string, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}

	return false
}

func arrContainsChar(arr []string, r rune) bool {
	for _, v := range arr {
		for _, x := range v {
			if x == r {
				return true
			}
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
