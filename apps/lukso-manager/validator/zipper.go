package validator

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"lukso/apps/lukso-manager/shared"
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
	zipFile, err := os.Create(pathToFile)
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	writer := zip.NewWriter(zipFile)
	defer writer.Close()

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

		s := strings.Split(path, fmt.Sprintf("/%s/", folder))

		f, err := writer.Create(s[1])
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
