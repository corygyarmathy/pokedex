package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchLocationAreas(pageURL *string) (respStruct *locationAreaListResp, err error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, ok := c.cache.Get(url)
	if !ok || data == nil { // URL not in cache or no cached data, retrieve via HTTP GET
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("creating request: %w", err)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("making request: %w", err)
		}
		defer func() {
			if cerr := resp.Body.Close(); cerr != nil && err == nil {
				err = fmt.Errorf("closing response body: %w", cerr)
			}
		}()
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("reading data: %w", err)
		}

		c.cache.Add(url, data)
	}

	var payload locationAreaListResp
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, fmt.Errorf("unmarshaling JSON: %w", err)
	}

	return &payload, nil
}
