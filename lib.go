package gowordstat

import (
	"errors"
	"fmt"
)

const (
	apiURLFormat       = "https://api.wordstat.yandex.net%s"
	endpointGetRegions = "/v1/getRegionsTree"
)

type Client interface{}

type client struct {
	oauthToken string
}

func NewClient(oauthToken string) (Client, error) {
	if oauthToken == "" {
		return nil, errors.New("auth token not set")
	}

	return &client{
		oauthToken: oauthToken,
	}, nil
}

// TODO
type Regions struct{}

func (c *client) GetRegions() (Regions, error) {
	return request[Regions](fmt.Sprintf(apiURLFormat, endpointGetRegions), nil, c.oauthToken)
}
