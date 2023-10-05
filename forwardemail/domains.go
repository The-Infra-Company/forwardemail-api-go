package forwardemail

import (
	"encoding/json"
	"fmt"
	"time"
)

type Domain struct {
	HasAdultContentProtection bool      `json:"has_adult_content_protection"`
	HasPhishingProtection     bool      `json:"has_phishing_protection"`
	HasExecutableProtection   bool      `json:"has_executable_protection"`
	HasVirusProtection        bool      `json:"has_virus_protection"`
	IsCatchallRegexDisabled   bool      `json:"is_catchall_regex_disabled"`
	Plan                      string    `json:"plan"`
	MaxRecipientsPerAlias     int       `json:"max_recipients_per_alias"`
	SmtpPort                  string    `json:"smtp_port"`
	Name                      string    `json:"name"`
	HasMxRecord               bool      `json:"has_mx_record"`
	HasTxtRecord              bool      `json:"has_txt_record"`
	HasRecipientVerification  bool      `json:"has_recipient_verification"`
	HasCustomVerification     bool      `json:"has_custom_verification"`
	VerificationRecord        string    `json:"verification_record"`
	Id                        string    `json:"id"`
	Object                    string    `json:"object"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
	Link                      string    `json:"link"`
}

func (c *Client) GetDomains() ([]Domain, error) {
	req, err := c.newRequest("GET", "/v1/domains")
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var items []Domain

	err = json.Unmarshal(res, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (c *Client) GetDomain(name string) (*Domain, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/v1/domains/%s", name))
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var item Domain

	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (c *Client) CreateDomain(name string) error {
	req, err := c.newRequest("POST", "/v1/domains")
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("domain", name)
	req.URL.RawQuery = q.Encode()

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteDomain(name string) error {
	req, err := c.newRequest("DELETE", fmt.Sprintf("/v1/domains/%s", name))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}