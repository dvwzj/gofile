package params

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
)

type UploadFileParams struct {
	FolderId   *string
	FileName   *string
	FileReader io.Reader
	Server     *string
}

type UploadFile func(*UploadFileParams) error

func WithFile(file *os.File) UploadFile {
	return func(params *UploadFileParams) error {
		if file == nil {
			return errors.New("file is nil")
		}
		params.FileReader = file
		segments := strings.Split(file.Name(), string(os.PathSeparator))
		fileName := segments[len(segments)-1]
		params.FileName = &fileName
		return nil
	}
}

func WithReader(fileReader io.Reader, fileName string) UploadFile {
	return func(params *UploadFileParams) error {
		if fileReader == nil {
			return errors.New("fileReader is nil")
		}
		params.FileReader = fileReader
		params.FileName = &fileName
		return nil
	}
}

func WithBytes(data []byte, fileName string) UploadFile {
	return func(params *UploadFileParams) error {
		if data == nil {
			return errors.New("data is nil")
		}
		params.FileReader = bytes.NewReader(data)
		params.FileName = &fileName
		return nil
	}
}

func WithPath(filePath string) UploadFile {
	return func(params *UploadFileParams) error {
		if filePath == "" {
			return errors.New("filePath is empty")
		}
		if strings.HasPrefix(filePath, "http") || strings.HasPrefix(filePath, "https") {
			client := resty.New()
			resp, err := client.R().SetDoNotParseResponse(false).Get(filePath)
			if err != nil {
				return err
			}
			buf := new(bytes.Buffer)
			_, err = buf.WriteString(resp.String())
			if err != nil {
				return err
			}
			params.FileReader = buf
			fileName := resp.Header().Get("Content-Disposition")
			if fileName != "" {
				fileName = strings.Trim(strings.Split(fileName, "filename=")[1], "\"")
			}
			if fileName == "" {
				segments := strings.Split(strings.Split(filePath, "?")[0], "/")
				validSegments := []string{}
				for _, segment := range segments {
					if segment != "" {
						validSegments = append(validSegments, segment)
					}
				}
				fileName = validSegments[len(validSegments)-1]
			}
			if !strings.Contains(fileName, ".") {
				contentType := resp.Header().Get("Content-Type")
				if contentType != "" {
					fileName = fileName + "." + strings.Split(contentType, "/")[1]
				} else {
					detectedType := http.DetectContentType(buf.Bytes())
					if detectedType != "application/octet-stream" {
						fileName = fileName + "." + strings.Split(detectedType, "/")[1]
					}
				}
			}
			if fileName == "" {
				fileName = "unknown"
			}
			params.FileName = &fileName
		} else {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()
			buf := new(bytes.Buffer)
			_, err = buf.ReadFrom(file)
			if err != nil {
				return err
			}
			params.FileReader = buf
			segments := strings.Split(filePath, string(os.PathSeparator))
			fileName := segments[len(segments)-1]
			params.FileName = &fileName
		}
		return nil
	}
}

type UploadFileOption func(*UploadFileParams) error

func WithFolderId(folderId string) UploadFileOption {
	return func(params *UploadFileParams) error {
		params.FolderId = &folderId
		return nil
	}
}

func WithServerName(serverName string) UploadFileOption {
	return func(params *UploadFileParams) error {
		params.Server = &serverName
		return nil
	}
}

func WithFileName(fileName string) UploadFileOption {
	return func(params *UploadFileParams) error {
		params.FileName = &fileName
		return nil
	}
}
