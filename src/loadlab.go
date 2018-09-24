package loadlab

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)


type LoadLab struct {
	token string
}

func request(resource string, token string) (io.Reader, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://api.loadlab.co/v1/%v/", resource),
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Token %v", token))
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		println(resp.Status)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	return bytes.NewBuffer(respBytes), nil
}

func plans(loadlab LoadLab) (io.Reader, error) {
	plans, err := request("plans", loadlab.token)
	if err != nil {
		panic(err)
	}
	return plans, nil
}

func jobs(loadlab LoadLab) (io.Reader, error) {
	plans, err := request("jobs", loadlab.token)
	if err != nil {
		panic(err)
	}
	return plans, nil
}

func sites(loadlab LoadLab) (io.Reader, error) {
	plans, err := request("sites", loadlab.token)
	if err != nil {
		panic(err)
	}
	return plans, nil
}
