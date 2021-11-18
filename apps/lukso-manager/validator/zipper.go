package validator

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"lukso/shared"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func zipFolder(network string, folder string) (filePath string) {
	baseFolder := shared.NetworkDir + network

	now := time.Now()
	sec := now.Unix()

	pathToFile := baseFolder + "/" + folder + "-" + fmt.Sprint(sec) + ".zip"
	file, err := os.Create(pathToFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.
		s := strings.Split(path, "/"+folder+"/")

		f, err := w.Create(s[1])
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	targpath := baseFolder + "/" + folder
	basepath := path
	relpath, _ := filepath.Rel(basepath, targpath)

	err = filepath.Walk(relpath, walker)
	if err != nil {
		panic(err)
	}

	return pathToFile
}
