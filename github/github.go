package github

import (
	"io"
	"net/http"
)

// FetchGoCodeFromGist busca o código Go de um Gist do GitHub pelo ID.
func FetchGoCodeFromGist(gistID string) (string, error) {
	url := "https://gist.githubusercontent.com/" + gistID + "/raw"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	code, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(code), nil
}
