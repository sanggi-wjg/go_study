package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Requester struct {
	URL           string
	ContainerName string
}

type OBSInfo struct {
	OBS_URL   string
	OBS_TOKEN string
}

type UploadResult struct {
	Filepath   string
	Success    bool
	Status     string
	StatusCode int
}

func (r *Requester) request() []byte {
	url := fmt.Sprintf(r.URL, r.ContainerName)
	resp, err := http.Get(url)
	if err = checkStatusCode(resp, err); err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func (obs *OBSInfo) uploadToOBS(filePaths []string) []UploadResult {
	results := make([]UploadResult, len(filePaths))
	channel := make(chan UploadResult)

	for _, path := range filePaths {
		go tryUpload(obs, path, channel)
	}

	for i := 0; i < len(filePaths); i++ {
		results[i] = <-channel
	}

	return results
}

func tryUpload(obs *OBSInfo, path string, c chan UploadResult) {
	file, err := os.Open(path)
	if err != nil {
		c <- UploadResult{Success: false, Filepath: path, Status: "file open failed"}
	}
	defer file.Close()

	client := &http.Client{}
	request, err := http.NewRequest("PUT", obs.OBS_URL, file)
	request.Header.Add("X-Auth-Token", obs.OBS_TOKEN)
	if err != nil {
		c <- UploadResult{Success: false, Filepath: path, Status: "request setter failed"}
	}

	resp, err := client.Do(request)
	defer resp.Body.Close()
	if err != nil {
		c <- UploadResult{Success: false, Filepath: path, Status: "request operation failed"}
	}
	if resp.StatusCode != 201 {
		c <- UploadResult{Success: false, Filepath: path, Status: resp.Status, StatusCode: resp.StatusCode}
	}
	c <- UploadResult{Success: true, Filepath: path, Status: resp.Status, StatusCode: resp.StatusCode}
}

func checkStatusCode(resp *http.Response, err error) error {
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("[%d] %s", resp.StatusCode, resp.Status))
	}
	return nil
}

func parseJson(respBody []byte) OBSInfo {
	var obsInfo OBSInfo
	err := json.Unmarshal(respBody, &obsInfo)
	if err != nil {
		panic(err)
	}
	return obsInfo
}

func isModTimeAfter(info fs.FileInfo) bool {
	day, _ := time.ParseDuration(TargetModTime)
	targetDay := time.Now().Add(-day)

	if info.ModTime().Before(targetDay) {
		return true
	}
	return false
}

func walkFiles() []string {
	var paths []string

	err := filepath.Walk(LogRoot, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() == false && isModTimeAfter(info) {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return paths
}

const (
	LogRoot string = "go_study/logs/"
	//TargetModTime string = "168h"
	TargetModTime string = "1h"
)

func main() {
	requester := Requester{
		URL:           "http://api.wms.pickby.us:8000/cloud/obs/%s/token",
		ContainerName: "KR01",
	}
	respBody := requester.request()
	obsInfo := parseJson(respBody)

	filePaths := walkFiles()
	results := obsInfo.uploadToOBS(filePaths)

	for _, res := range results {
		fmt.Println(res)
	}
}
