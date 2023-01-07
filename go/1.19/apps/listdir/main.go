package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// helper
	var sendHelp bool = false
	for _, h := range os.Args {
		if strings.ToLower(h) == "help" {
			sendHelp = true
		}
	}
	if sendHelp {
		fmt.Print(`
----------------------------------------
listdir help
----------------------------------------

API to return an array of files in a given directory.

Call this api using the below format in the body:
{
	'Name':'listdir',
	'Args':[
		'<directory, string, required>',
		'<search recursively, bool, optional, default false, valid responses: 	
			'1',
			't',
			'T',
			'TRUE',
			'true',
			'True',
			'0',
			'f',
			'F',
			'FALSE',
			'false',
			'False'>
	]
}

Example call to list all filss in directory /tmp/test not recursively:
{
	'Name':'listdir',
	'Args':[
		'/tmp/test'
	]
}

Example format of data returned:
['<file1>','<file2>'...]

----------------------------------------
`)
		return
	}

	var boolValue bool
	var err error
	if len(os.Args) > 2 {
		boolValue, err = strconv.ParseBool(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}

	if (len(os.Args) == 2) && (!boolValue) {
		getfilesnonrecursive()
		return
	} else if len(os.Args) == 3 {
		if boolValue {
			fmt.Println("Functionality under development.")
		} else {
			getfilesnonrecursive()
			return
		}
	} else if len(os.Args) > 3 {
		fmt.Println("Functionality under development.")
	} else {
		fmt.Println("Unexpected number of Agrs provided. Please call this api with 'help' in Args for help.")
	}
}

func getfilesnonrecursive() {
	var f []string
	files, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			f = append(f, "'"+file.Name()+"'")
		}
		// fmt.Println(file.Name(), file.IsDir())
	}
	fmt.Print("[" + strings.Join(f, ", ") + "]")
}
