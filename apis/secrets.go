package apis

import (
	"encoding/json"
	"io"
	"net/http"
)

// Get secret from a local secret store

func GetSecret(url string, key string) (string, error) {
	//getResponse, err := http.Get(DAPR_HOST + ":" + DAPR_HTTP_PORT + "/v1.0/secrets/" + DAPR_SECRET_STORE + "/" + SECRET_NAME)

	getResponse, err := http.Get(url + "/" + key)

	if err != nil {
		return "", nil
	}
	result, err := io.ReadAll(getResponse.Body)
	if err != nil {
		return "", nil
	}

	var mapData map[string]string
	if err = json.Unmarshal(result, &mapData); err != nil {
		return "", nil
	} else {

	}

	return mapData["dsn"], nil
}
