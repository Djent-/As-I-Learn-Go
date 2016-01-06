package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is important
	// create a false body to post using a fake form
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}

	// iocopy
	// copy the filehandler data to the fileWriter
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	// contentType variable holds the type of data posted
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	// post the form data to the server
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// convert the response body into a readable format
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)       // http status code returned
	fmt.Println(string(resp_body)) //response from the server
	return nil                     // no error
}

func main() {
	target_url := "http://localhost:9090/upload"
	filename := "./astaxie.pdf"
	postFile(filename, target_url)
}