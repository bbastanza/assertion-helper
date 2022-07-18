package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// TODO if the more than one class is found give an option which namespace (find the namespace somehow)

type Config struct {
	ProjectRoot string `json:"project_root"`
}

func main() {
	Program()
}

func Program() {
	var className string

	fmt.Println("Enter Class Name")

	for {
		fmt.Scanf("%s", &className)

		if len(className) != 0 {
			fmt.Println("")
			break
		}
	}

	config, err := GetConfig()

	if err != nil {
		fmt.Println("Error getting config ", err)
		return
	}

	pattern := "public class " + className

	fmt.Println(config.ProjectRoot)

	rgCmd := exec.Command("/usr/bin/rg", pattern, config.ProjectRoot)

	result, err := rgCmd.CombinedOutput()

	res := string(result)

	if err != nil {
		fmt.Printf("Error running ripgrep")
		return
	}

	if !ContainsChar(string(res), ':') {
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

	// dont start building assertions until we find "public class *className"
	startProps := false
	for scanner.Scan() {

		// get the last scanned line
		t := scanner.Text()

		if IsSubstring(t, "public class "+className) {
			startProps = true
			continue
		}

		if !startProps {
			continue
		}

		// if we get here we are on the second class and can break
		if IsSubstring(t, "public class") {
			break
		}

		// TODO make this into its own function and have a list of all C# keywords
		// TODO maybe only for primitive types ie. string, int, char, decimal, float
		isProp := ContainsString(t, "public") &&
			!ContainsString(t, "class") && // this will rule out the class decloration
			!ContainsString(t, "enum") && // this will rule out enums
			!ArrContainsChar(strings.Split(t, " "), '(') // this will rule out methods

		if isProp {
			// split our the stdout line and remove any whitespace characters/strings
			arr := strings.Split(t, " ")
			arr = RemoveEmpty(arr)

			// we are looking for the idx + 2, so if we only have two continue
			if len(arr) <= 2 {
				continue
			}

			// get our index
			idx := IndexOf(arr, "public")

			// get our property name which is two over from public
			//      public sting PropertyName { get; set; }
			propName := arr[idx+2]

			assertion := "Assert.Equal(expected." + propName + ", result." + propName + ");"

			// assertion = ""
			fmt.Println(assertion)

			propCount++
		}
	} // End loop

	fmt.Printf("\nFound and created %d Assertions", propCount)
}

func IndexOf(arr []string, pattern string) int {
	for i, v := range arr {

		if v == pattern {
			return i
		}
	}

	return -1
}

func ContainsString(s string, pattern string) bool {

	arr := strings.Split(s, " ")

	for _, v := range arr {
		if v == pattern {
			return true
		}
	}

	return false
}

func ContainsChar(s string, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}

	return false
}

func ArrContainsChar(arr []string, r rune) bool {
	// loop through the array
	for _, v := range arr {

		// loop through each string in the array
		for _, x := range v {
			if x == r {
				return true
			}
		}
	}

	return false
}

// if the trimmed string is an empty string remove it from the output
func RemoveEmpty(arr []string) []string {
	var output []string

	for _, v := range arr {
		if len(strings.TrimSpace(v)) != 0 {
			output = append(output, v)
		}
	}

	return output
}

func IsSubstring(str string, sub string) bool {

Outer:
	for i, v := range str {

		// if the first letters don't match keep looping through partent
		if string(v) != string(sub[0]) {
			continue
		}

		// chop off any thing before the substring started
		adjustedStr := str[i:]

		// cannot have a substring longer than the parent
		if len(adjustedStr) < len(sub) {
			return false
		}

		// start the loop and compare
		for i, v := range adjustedStr {

			//if the substring is done we are done
			if i > len(sub)-1 {
				return true
			}

			// if two characters at the same don't match we don't have a substring
			if string(v) != string(sub[i]) {
				continue Outer
			}
		}

		return true
	}

	return false
}

func GetConfig() (Config, error) {
	configPath := "./config.json"

	config, err := ioutil.ReadFile(configPath)

	if err != nil {
		return Config{}, err
	}

	data := Config{}

	json.Unmarshal([]byte(config), &data)

	return data, nil
}
