package utils

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"time"
)

func FileUpload(file multipart.File) (string, error) {

	now := time.Now() // current local time

	//store file temporarily
	tempFile, err := ioutil.TempFile("files/"+now.Local().String(), "*.yaml")

	if err != nil {
		return "", errors.New("failed to store file temporarily")
	}

	//read file
	readfile, err := ioutil.ReadAll(file)

	if err != nil {
		return "", errors.New("failes to read temp file")
	}

	//store finally
	tempFile.Write(readfile)

	return tempFile.Name(), nil

}
