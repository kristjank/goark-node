package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

var BaseURL = ""

//ArkAPIResponseError struct to hold error response from api node
type ArkAPIResponseError struct {
	Success      bool   `json:"success,omitempty"`
	Message      string `json:"message,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
	Data         string `json:"data,omitempty"`
}

//Error interface function
func (e ArkAPIResponseError) Error() string {
	return fmt.Sprintf("ArkServiceApi: %v %v %v %v", e.Success, e.ErrorMessage, e.Data, e.Message)
}

//ArkClient sling rest pointer
type ArkClient struct {
	sling *sling.Sling
}

func init() {
	switchNetwork(MAINNET)
}

//NewArkClient creations with supported network
func NewArkClient(httpClient *http.Client) *ArkClient {
	return &ArkClient{
		sling: sling.New().Client(httpClient).Base(BaseURL).
			Add("nethash", EnvironmentParams.Network.Nethash).
			Add("version", EnvironmentParams.Network.ActivePeer.Version).
			Add("port", strconv.Itoa(EnvironmentParams.Network.ActivePeer.Port)).
			Add("Content-Type", "application/json"),
	}
}
