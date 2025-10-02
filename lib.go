package gowordstat

import (
	"errors"
	"fmt"
)

const (
	apiURLFormat       = "https://api.wordstat.yandex.net%s"
	endpointGetRegions = "/v1/getRegionsTree"
)

type Client interface {
	GetRegions() ([]Region, error)
}

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
type Region map[string]any

func (c *client) GetRegions() ([]Region, error) {
	return request[[]Region](fmt.Sprintf(apiURLFormat, endpointGetRegions), nil, c.oauthToken)
}
