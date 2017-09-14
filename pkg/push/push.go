package push

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"vvoid.pw/archivebuffer"
)

const (
	errorCreatingTarball     = "PushCurrentDirFiles: Error creating tarball: %v"
	errorCreatingGzip        = "PushCurrentDirFiles: Error creating gzip: %v"
	errorWritingToBuffer     = "PushCurrentDirFiles: Error writing to buffer: %v"
	errorWritingToBody       = "PushCurrentDirFiles: Error writing to body: %v"
	errorPushingFiles        = "PushCurrentDirFiles: Error pushing files: %v"
	errorParsingResponseBody = "PushCurrentDirFiles: Error parsing response body: %v"
)

type PushResponse struct {
	URL   string `json:"push-url"`
	Token string `json:"token"`
	Error string `json:"error"`
}

func PushCurrentDirFiles(pimbaServerURL, bucketName, token string) ([]string, error) {
	currentDir, _ := os.Getwd()
	tarball, err := archivebuffer.NewTarballBuffer(currentDir)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(errorCreatingTarball, err))
	}
	tarGz, err := archivebuffer.NewGzipBuffer(tarball)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(errorCreatingGzip, err))
	}

	pushURL := "http://" + pimbaServerURL + "/v1/push"
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	filename := "/pimba.tar.gz"
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(errorWritingToBuffer, err))
	}

	_, err = io.Copy(fileWriter, tarGz)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(errorWritingToBody, err))
	}

	if bucketName != "" {
		err = bodyWriter.WriteField("name", bucketName)
		if err != nil {
			return nil, errors.New(fmt.Sprintf(errorWritingToBody, err))
		}
	}

	err = bodyWriter.WriteField("token", token)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(errorWritingToBody, err))
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(pushURL, contentType, bodyBuf)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(errorPushingFiles, err))
	}
	defer resp.Body.Close()

	pushResp := PushResponse{}
	err = parseResponseBody(resp.Body, &pushResp)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(errorParsingResponseBody, err))
	}

	if pushResp.Error != "" {
		return nil, errors.New(pushResp.Error)
	}

	return []string{pushResp.URL, pushResp.Token}, nil
}

func parseResponseBody(body io.Reader, pushResp *PushResponse) error {
	reader, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	sl := strings.Split(string(reader), "]")
	jsonContent := []byte(sl[len(sl)-1])
	err = json.Unmarshal(jsonContent, pushResp)
	if err != nil {
		return err
	}

	return nil
}
