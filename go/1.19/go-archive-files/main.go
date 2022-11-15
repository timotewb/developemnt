package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ncruces/zenity"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func duplicateFile(f string, i int) (string, error) {
	/*
		purpose: give me a diplicate file and i will let you know what to name it
		i = iteration, number of duplicates, sued for suffix
		f = file name, path and extension - cannot have suffix
		x = extension including leading period
		n = file name excluding path and extension
		nx = fiel name including extension, excluding path
	*/
	x := filepath.Ext(f)
	nx := filepath.Base(f)
	n := strings.ReplaceAll(filepath.Base(f), x, "")
	p := strings.ReplaceAll(f, nx, "")
	i_str := fmt.Sprintf("%04d", i)

	newF := filepath.Join(p, n+"_d"+i_str+x)

	if _, err := os.Stat(newF); err == nil {
		newF, err := duplicateFile(f, i+1)
		if err != nil {
			return "", fmt.Errorf("%s", err)
		} else {
			return newF, nil
		}
	} else if errors.Is(err, os.ErrNotExist) {
		return newF, nil
	} else {
		return "", fmt.Errorf("file does or doesnot exist ")
	}
}

func MoveFile(sourcePath, destPath string) error {

	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("writing to output file failed: %s", err)
	}
	err = outputFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Stat(sourcePath)
	if err != nil {
		fmt.Println(err)
	}
	os.Chtimes(destPath, file.ModTime(), file.ModTime())
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("failed removing original file: %s", err)
	}
	return nil
}

func main() {
	inDir, err := zenity.SelectFile(
		zenity.Filename(""),
		zenity.Directory(),
		zenity.DisallowEmpty(),
		zenity.Title("Select input directory"),
	)
	if err != nil {
		zenity.Error(err.Error(),
			zenity.Title("Error"),
			zenity.ErrorIcon,
		)
		log.Fatal(err)
	}
	outDir, err := zenity.SelectFile(
		zenity.Filename(""),
		zenity.Directory(),
		zenity.DisallowEmpty(),
		zenity.Title("Select output directory"),
	)
	if err != nil {
		zenity.Error(err.Error(),
			zenity.Title("Error"),
			zenity.ErrorIcon,
		)
		log.Fatal(err)
	}
	err = zenity.Question("Input: \n- "+inDir+"\nOutput:\n -"+outDir+"\n\nAre you sure you want to proceed?",
		zenity.Title("Question"),
		zenity.QuestionIcon)
	if err != nil {
		zenity.Error(err.Error(),
			zenity.Title("Error"),
			zenity.ErrorIcon,
		)
		log.Fatal(err)
	}

	//--- start the progress bar
	dlg, err := zenity.Progress(
		zenity.Title("Archiving Files"),
		zenity.Pulsate())
	if err != nil {
		zenity.Error(err.Error(),
			zenity.Title("Error"),
			zenity.ErrorIcon,
		)
		log.Fatal(err)
	}
	defer dlg.Close()

	//--- get all files into list
	files, err := FilePathWalkDir(inDir)
	if err != nil {
		panic(err)
	}
	if err != nil {
		zenity.Error(err.Error(),
			zenity.Title("Error"),
			zenity.ErrorIcon,
		)
		log.Fatal(err)
	}

	//--- move files files
	for _, f := range files {
		fn := filepath.Base(f)
		var ext string
		//--- if file as no extension
		if strings.ToLower(filepath.Ext(f)) == "" {
			ext = "no_ext"
		} else {
			ext = strings.ToLower(filepath.Ext(f)[1:])
		}

		//--- check if file is hidden or tmp
		ignoreExt := []string{"ds_store"}
		if fn[:2] != "._" && !stringInSlice(ext, ignoreExt) {
			fmt.Println("Archiving:", f)
			if err != nil {
				fmt.Println(err)
			}
			file, err := os.Stat(f)
			if err != nil {
				fmt.Println(err)
			}
			year, month, day := file.ModTime().Date()
			month1 := fmt.Sprintf("%02d", month)
			day1 := fmt.Sprintf("%02d", day)
			newpath := filepath.Join(outDir, ext, strconv.Itoa(year), month1, day1)
			newF := filepath.Join(newpath, fn)
			duppath := filepath.Join(outDir, "_duplicates", ext, strconv.Itoa(year), month1, day1)

			//--- check if file already exists - TODO
			if _, err := os.Stat(newF); err == nil {
				// this is what we want to put it... but what if it already exists?
				dupF := filepath.Join(duppath, fn)
				// we check and give it a siffix
				dupF, err := duplicateFile(dupF, 1)
				if err != nil {
					fmt.Println(err)
				} else {
					//--- create folder
					os.MkdirAll(duppath, os.ModePerm)
					//---move file
					MoveFile(f, dupF)
				}

			} else if errors.Is(err, os.ErrNotExist) {
				//--- create folder
				os.MkdirAll(newpath, os.ModePerm)
				//---move file
				MoveFile(f, newF)

			} else {
				fmt.Println("Error: file does or doesnot exist ", newF)

			}
		}
	}

	zenity.Info("Archiving complete.",
		zenity.Title("Information"),
		zenity.InfoIcon)

}

// https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
// env GOOS=target-OS GOARCH=target-architecture go build package-import-path
// env GOOS=windows GOARCH=amd64 go build github.com/timotewb/go-archive-files
