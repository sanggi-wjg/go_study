package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Requester struct {
	URL           string
	ContainerName string
}

type OBSInfo struct {
	URL   string
	TOKEN string
}

func (r *Requester) request() {
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
	fmt.Println(string(body))
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

func main() {
	requester := Requester{
		URL:           "http://api.wms.pickby.us:8000/cloud/obs/%s/token",
		ContainerName: "KR01",
	}
	requester.request()

}
