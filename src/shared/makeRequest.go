package shared

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func MakeRequest(endpoint string, obj interface{}) error {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)

	if err != nil {
		return err
	}

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(body, &obj)
}
