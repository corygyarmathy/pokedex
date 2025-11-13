package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) FetchPokemon(name string) (respStruct *Pokemon, err error) {
	url := c.baseURL + "/pokemon/" + name + "/"

	data, err := c.doGET(url)
	if err != nil {
		return nil, fmt.Errorf("fetching pokemon %v: %w", name, err)
	}

	var payload Pokemon
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, fmt.Errorf("unmarshaling JSON: %w", err)
	}

	return &payload, nil
}
