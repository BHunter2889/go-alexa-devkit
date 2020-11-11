package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BHunter2889/go-alexa-devkit/alexa"
	"log"
	"net/http"
	"context"
	"golang.org/x/net/context/ctxhttp"
)

type DeviceService struct {
	URL      string
	Id       string
	Token    string
	Endpoint string
}

// Client may be nil which will use the http.DefaultClient.
// Context Must be provided even if it will not be used.
// TODO - add non-context dependent service func.
func (s *DeviceService) GetDeviceLocation(ctx context.Context, client *http.Client) (*alexa.DeviceLocationResponse, error) {
	endp := fmt.Sprintf(s.Endpoint, s.Id)
	reqUrl := fmt.Sprintf("%s%s", s.URL, endp)
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)

	if err != nil {
		log.Print("Error creating new device location request")
		log.Print(err)
		log.Print(req)
		return nil, err
	}
	bearer := "Bearer " + s.Token
	req.Header.Add("Authorization", bearer)

	resp, err := ctxhttp.Do(ctx, , req)
	if err != nil || resp.StatusCode == 403 {
		if resp.StatusCode == 403 {
			err = errors.New(resp.Status)
		}
		log.Print("Error processing device location response")
		log.Print(err)
		log.Print(resp)
		return nil, err
	}
	defer resp.Body.Close()

	deviceLocationResponse := alexa.DeviceLocationResponse{}

	if err := json.NewDecoder(resp.Body).Decode(&deviceLocationResponse); err != nil {
		return nil, err
	}
	return &deviceLocationResponse, nil
}
