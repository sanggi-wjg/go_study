package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
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

func (obs *OBSInfo) uploadToOBS([]string) error {
	// https://stackoverflow.com/questions/27656898/how-to-upload-file-using-golang-code
	request, err := http.NewRequest("PUT", obs.OBS_URL, strings.NewReader(""))
	request.Header.Add("X-Auth-Token", obs.OBS_TOKEN)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
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

	if info.ModTime().After(targetDay) {
		return true
	}
	return false
}

func walkFiles() []string {
	var files []string
	err := filepath.Walk(LogRoot, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() == false {
			if isModTimeAfter(info) {
				files = append(files, path)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

const (
	LogRoot       string = "go_study/logs/"
	TargetModTime string = "168h"
)

func main() {
	requester := Requester{
		URL:           "http://api.wms.pickby.us:8000/cloud/obs/%s/token",
		ContainerName: "KR01",
	}
	respBody := requester.request()
	obsInfo := parseJson(respBody)
	fmt.Println(obsInfo)

	files := walkFiles()
	fmt.Println(files)

	obsInfo.uploadToOBS(files)
}
