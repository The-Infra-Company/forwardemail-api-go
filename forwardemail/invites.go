package forwardemail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"time"
)

type Invite struct {
	Email     string    `json:"email"`
	Group     string    `json:"group"`
	ID        string    `json:"id"`
	Object    string    `json:"object"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Client) CreateDomainInvite(domain, email, group string) (*Invite, error) {
	req, err := c.newRequest("POST", fmt.Sprintf("/v1/domains/%s/invites", domain))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	params := url.Values{}
	params.Add("email", email)
	params.Add("group", group)

	req.Body = io.NopCloser(bytes.NewBufferString(params.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	var item Invite
	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &item, nil
}
