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

func startCompare(s string, sub string) bool {

	substringLength := len(sub)

	if len(s) < substringLength {
		return false
	}

	for i, v := range s {
		if i > len(sub)-1 {
			break
		}

		if string(v) != string(sub[i]) {
			return false
		}
	}

	return true
}

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

	pattern := "public class " + className

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

	// TODO we need to start at the line that conains the className
	// and then end at the next className

	startProps := false
	for scanner.Scan() {
		// get the last scanned line
		t := scanner.Text()

		if isSubstring(t, "public class "+className) {
			startProps = true
			continue
		}

		if !startProps {
			continue
		}

		// if we get here we are on the second class and can break
		if isSubstring(t, "public class") {
			break
		}

		// TODO make this into its own function and have a list of all C# keywords
		// TODO maybe only for primitive types ie. string, int, char, decimal, float
		isProp := containsString(t, "public") &&
			!containsString(t, "class") && // this will rule out the class decloration
			!containsString(t, "enum") && // this will rule out enums
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

			// assertion = ""
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

// TODO this isn't 100 percent correct because it is checking for the first letter
// so if it appears before the substring starts we have a problem
// for this use case it's probably okay though because we are looking for "public class"
func isSubstring(s string, sub string) bool {

	for i, v := range s {

		if string(v) == string(sub[0]) {
			return startCompare(s[i:], sub)
		}
	}

	return false
}
