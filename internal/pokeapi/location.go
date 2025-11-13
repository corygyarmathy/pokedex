package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) FetchLocationAreas(fetchURL *string) (respStruct *locationAreaListResp, err error) {
	url := c.baseURL + "/location-area"
	if fetchURL != nil {
		url = *fetchURL
	}

	data, err := c.doGET(url)
	if err != nil {
		return nil, fmt.Errorf("fetching locations: %w", err)
	}

	var payload locationAreaListResp
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, fmt.Errorf("unmarshaling JSON: %w", err)
	}

	return &payload, nil
}

func (c *Client) FetchLocationArea(area string) (respStruct *locationAreaResp, err error) {
	url := c.baseURL + "/location-area/" + area + "/"

	data, err := c.doGET(url)
	if err != nil {
		return nil, fmt.Errorf("fetching location %v: %w", area, err)
	}

	var payload locationAreaResp
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, fmt.Errorf("unmarshaling JSON: %w", err)
	}

	return &payload, nil
}
