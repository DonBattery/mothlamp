package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/spf13/viper"
)

func welcome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("MothLamp\n"))
}

// custom404 handles all unhandlet request with a common Wrong way 404 message.
func custom404() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("Wrong way 404 🐸\n"))
	})
}

func uploadFile(res http.ResponseWriter, req *http.Request) {
	// Parse request form
	if err := req.ParseMultipartForm(32 << 20); err != nil {
		multiWriter(res, "Error parsing request form %v", err)
	}
	// get the file ment to be uploaded from the request form
	file, fileHeader, err := req.FormFile("uploadfile")
	if err != nil {
		multiWriter(res, "Error opening Form file %v", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("Error closing incoming file: %v\n", closeErr)
		}
	}()
	// set local file's path
	localFilePath := path.Join(viper.GetString("MOTHLAMP_DIR"), fileHeader.Filename)
	// check if already exists, return if overwrite is not set
	if _, err := os.Stat(localFilePath); !os.IsNotExist(err) && req.FormValue("overwrite") != "true" {
		multiWriter(res, "%v already exists. use -F overwrite=true\n", fileHeader.Filename)
		return
	}
	// prepare local file for copy
	localFile, err := os.OpenFile(localFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		multiWriter(res, "Error opening local file: %v\n", err)
	}
	defer func() {
		if closeErr := localFile.Close(); closeErr != nil {
			fmt.Printf("Error closing local file: %v", closeErr)
		}
	}()
	// actually try to copy bytes from form file to local file
	bytes, err := io.Copy(localFile, file)
	if err != nil {
		multiWriter(res, "Error writing to local file %v", err)
	} else {
		multiWriter(res, "%v uploaded. Received %v bytes.\n", fileHeader.Filename, bytes)
	}

}

func multiWriter(writer io.Writer, format string, args ...interface{}) {
	bytes, err := fmt.Fprintf(writer, format, args...)
	if err != nil {
		fmt.Printf("Error writing response : %v\n", err)
	} else {
		fmt.Printf("%v bytes successfully writen to response\n", bytes)
		fmt.Printf(format, args...)
		fmt.Println()
	}
}
