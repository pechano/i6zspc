package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestPath(t *testing.T) {
	filename := loadi6z()
	fmt.Println("The chosen file is: " + filename)
	folder := filepath.Dir(filename)
	folder = filepath.Join(folder, "temp")
	fmt.Println("The temp folder is: " + folder)

	defer os.RemoveAll(folder)
}

func TestExtract(t *testing.T) {
	inputPath := "/home/johanneslb/Desktop/insectosec.i6z"
	tempfolder := extractFiles(inputPath)
	fmt.Println(tempfolder)
}
