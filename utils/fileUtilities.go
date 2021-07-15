package utils

import (
	"io"
	"mime/multipart"
	"os"
)

func FileUpload(file multipart.File, handler *multipart.FileHeader) (string, error) {

	defer file.Close() //close the file when we finish

	f, err := os.OpenFile("files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return "", err
	}

	defer f.Close()

	io.Copy(f, file)

	//here we save our file to our path
	return handler.Filename, nil

}
