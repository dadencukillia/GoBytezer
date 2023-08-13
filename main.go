package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check needed arguments exist
	if len(os.Args) <= 1 {
		fmt.Println("Missing arguments: -o <file.name> -x <file.go>")
		return
	}

	// Results of arguments
	var from []string = []string{} // File path for data reading (for example: .png)
	var to []string = []string{}   // File path for data writing (.go)

	var mode string = "" // mode: change "from" variable or "to" variable

	Arguments := os.Args[1:] // arguments without file name
	for _, arg := range Arguments {
		if arg == "-o" {
			mode = "o" // setup mode for "from" variable changing
		} else if arg == "-x" {
			mode = "x" // setup mode for "to" variable changing
		} else {
			if mode == "o" {
				// write file path of file for reading
				from = append(from, arg)
			} else if mode == "x" {
				// write file path of file for writing
				to = append(to, arg)
			}
		}
	}

	// check if file reading path is not empty
	if len(from) == 0 {

		// check if file writing is not empty too
		if len(to) == 0 {
			fmt.Println("Missing arguments: -o <file.name> -x <file.go>")
		} else {
			fmt.Println("Missing argument: -o <file.name>")
		}

		// check if file writing path is not empty
	} else if len(to) == 0 {
		fmt.Println("Missing argument: -x <file.go>")

		// we can work with files
	} else {
		// path to read
		from_string := strings.ReplaceAll(strings.Join(from, " "), "\\", "/")

		// reading content
		rcontent, err := os.ReadFile(from_string)
		if err != nil {
			fmt.Println("Error of reading file: " + from_string)
			return
		}

		// name of variable
		filename_split := strings.Split(from_string, "/")
		filename := filename_split[len(filename_split)-1]
		filename = strings.ReplaceAll(strings.ReplaceAll(filename, ".", "_"), " ", "_") // Replacing invalid symbols of variable name

		// data for writing
		wcontent := collectData(filename, rcontent)

		// write data
		to_string := strings.ReplaceAll(strings.Join(to, " "), "\\", "/")

		if err = os.WriteFile(to_string, wcontent, 0777); err != nil {
			fmt.Println("Error of writing file: " + to_string)
			return
		}

		fmt.Println("Success!")
	}
}

// content for writing in the file (.go)
func collectData(variableName string, fileContent []byte) []byte {
	// Start file content
	content := []byte(fmt.Sprintf("package main\nvar %s []byte=[]byte{", variableName))

	// Generate middle file content
	content = append(content, bytes.ReplaceAll(bytes.TrimRight(bytes.TrimLeft([]byte(fmt.Sprint(fileContent)), "["), "]"), []byte(" "), []byte(","))...)

	// End file content
	content = append(content, byte("}"[0]))

	return content
}
