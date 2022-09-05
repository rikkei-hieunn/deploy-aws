// Package filebus to define filebus instance
package filebus

import (
	"fmt"
	"insert-calendar-infos/configs"
	"insert-calendar-infos/model"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type filebusHandler struct {
	cfg    *configs.TickFileBus
	client http.Client
}

// NewFilebusHandler init filebus handler
func NewFilebusHandler(config *configs.TickFileBus) IFilebusHandler {
	return &filebusHandler{
		cfg: config,
		client: http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// Download download file use filebus API
func (f *filebusHandler) Download(path, file string) ([]byte, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf(model.ServerURL+"%s:%d"+f.cfg.URLDownload,
		f.cfg.Hostname, f.cfg.Port), nil)
	if err != nil {
		return nil, fmt.Errorf("download file from Filebus fail, error: " + err.Error())
	}

	query := request.URL.Query()
	query.Add("id", f.cfg.Username)
	query.Add("pw", f.cfg.Password)
	query.Add("file", file)
	query.Add("path", path)
	request.URL.RawQuery = query.Encode()

	if response, err := f.client.Do(request); err != nil {
		return nil, fmt.Errorf("download file from Filebus fail, error: " + err.Error())
	} else if response.StatusCode == 404 {
		return nil, fmt.Errorf("download file from Filebus fail, response code: 404")
	} else if response.StatusCode == 200 || response.StatusCode == 300 {
		responseData, errReadBody := ioutil.ReadAll(response.Body)
		if errReadBody != nil {
			return nil, fmt.Errorf("download file from Filebus fail, error: " + errReadBody.Error())
		}
		_ = response.Body.Close()

		return responseData, nil
	} else {
		return nil, fmt.Errorf("download file from Filebus fail, response code: " + strconv.Itoa(response.StatusCode))
	}
}
