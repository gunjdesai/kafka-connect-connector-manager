package network

import (
	"github.com/gunjdesai/kafka-connect-connector-manager/constants"
	"github.com/gunjdesai/kafka-connect-connector-manager/globals"
	"io"
	"log"
	"net/http"
)

func buildHttpRequest(requestType string, url string, body io.Reader) (*http.Client, *http.Request) {

	client := &http.Client{}
	request, err := http.NewRequest(requestType, url, body)

	if err != nil {
		log.Fatal(err)
	}

	switch globals.Config.AuthType {

	case constants.BASIC_AUTH:
		request.SetBasicAuth(globals.Config.BasicAuth.Username, globals.Config.BasicAuth.Password)
	}

	return client, request

}

func MakeHttpRequest(requestType string, url string, body io.Reader) (*http.Response, error) {

	client, request := buildHttpRequest(requestType, url, body)
	request.Header.Add("Content-Type", "application/json")
	return client.Do(request)
}
