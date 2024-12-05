package forwardemail

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Alias struct {
	Account                  Account   `json:"user"`
	Domain                   Domain    `json:"domain"`
	Name                     string    `json:"name"`
	Labels                   []string  `json:"labels"`
	Description              string    `json:"description"`
	IsEnabled                bool      `json:"is_enabled"`
	HasRecipientVerification bool      `json:"has_recipient_verification"`
	Recipients               []string  `json:"recipients"`
	Id                       string    `json:"id"`
	Object                   string    `json:"object"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

type AliasParameters struct {
	Recipients               *[]string
	Labels                   *[]string
	Description              string `json:"description"`
	HasRecipientVerification *bool
	IsEnabled                *bool
}

func (c *Client) GetAliases(domain string) ([]Alias, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/v1/domains/%s/aliases", domain))
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var items []Alias

	err = json.Unmarshal(res, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (c *Client) GetAlias(domain string, alias string) (*Alias, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/v1/domains/%s/aliases/%s", domain, alias))
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var item Alias

	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (c *Client) CreateAlias(domain string, alias string, parameters AliasParameters) (*Alias, error) {
	req, err := c.newRequest("POST", fmt.Sprintf("/v1/domains/%s/aliases", domain))
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("name", alias)
	if parameters.Description != "" {
		params.Add("description", parameters.Description)
	}

	for k, v := range map[string]*bool{
		"has_recipient_verification": parameters.HasRecipientVerification,
		"is_enabled":                 parameters.IsEnabled,
	} {
		if v != nil {
			params.Add(k, strconv.FormatBool(*v))
		}
	}

	for k, v := range map[string]*[]string{
		"recipients[]": parameters.Recipients,
		"labels[]":     parameters.Labels,
	} {
		if v != nil {
			for _, vv := range *v {
				params.Add(k, vv)
			}
		}
	}

	req.Body = io.NopCloser(strings.NewReader(params.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var item Alias

	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (c *Client) UpdateAlias(domain string, alias string, parameters AliasParameters) (*Alias, error) {
	req, err := c.newRequest("PUT", fmt.Sprintf("/v1/domains/%s/aliases/%s", domain, alias))
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("name", alias)
	if parameters.Description != "" {
		params.Add("description", parameters.Description)
	}

	for k, v := range map[string]*bool{
		"has_recipient_verification": parameters.HasRecipientVerification,
		"is_enabled":                 parameters.IsEnabled,
	} {
		if v != nil {
			params.Add(k, strconv.FormatBool(*v))
		}
	}

	for k, v := range map[string]*[]string{
		"recipients[]": parameters.Recipients,
		"labels[]":     parameters.Labels,
	} {
		if v != nil {
			for _, vv := range *v {
				params.Add(k, vv)
			}
		}
	}

	req.Body = io.NopCloser(strings.NewReader(params.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var item Alias

	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (c *Client) DeleteAlias(domain string, alias string) error {
	req, err := c.newRequest("DELETE", fmt.Sprintf("/v1/domains/%s/aliases/%s", domain, alias))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
