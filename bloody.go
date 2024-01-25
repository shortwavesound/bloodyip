package bloodyip

import (
	"encoding/json"
	"io"
	"net/http"
)

var CheckBelarus bool = false

func IsMyIPBloody() (bool, error) {
	res, err := http.Get("https://ipinfo.io/")
	if err != nil {
		return false, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	ipInfo := make(map[string]interface{})
	err = json.Unmarshal(data, &ipInfo)
	if err != nil {
		return false, err
	}

	if ipInfo["country"].(string) == "RU" || CheckBelarus && ipInfo["country"].(string) == "BY" {
		return true, nil
	}

	return false, nil
}
