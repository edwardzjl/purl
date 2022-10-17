package url

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"net/url"
)

// extends url.URL to support driver.Scan and driver.Value
// see https://husobee.github.io/golang/database/2015/06/12/scanner-valuer.html
type URL struct {
	url.URL
}

func Parse(rawURL string) (*URL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	return &URL{*u}, nil
}

func (pu URL) MarshalJSON() ([]byte, error) {
	return json.Marshal(pu.String())
}

func (pu *URL) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	u, err := Parse(value)
	if err != nil {
		return err
	}
	*pu = *u
	return nil
}

func (mUrl URL) Value() (driver.Value, error) {
	return mUrl.String(), nil
}

func (mUrl *URL) Scan(value interface{}) error {
	if value == nil {
		*mUrl = URL{url.URL{}}
		return nil
	}
	bv, err := driver.String.ConvertValue(value)
	if err != nil {
		return err
	}
	if v, ok := bv.(string); ok {
		_url, err := url.Parse(v)
		if err != nil {
			return err
		}
		*mUrl = URL{*_url}
		return nil
	}
	return errors.New("failed to scan URL")
}
