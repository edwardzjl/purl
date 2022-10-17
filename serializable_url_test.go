package url

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalJson(t *testing.T) {
	pu, _ := Parse("http://foo.bar.com")
	actual, err := json.Marshal(pu)
	require.NoError(t, err)
	expected, _ := json.Marshal("http://foo.bar.com")
	assert.Equal(t, expected, actual)
}

func TestUnmarshalJson(t *testing.T) {
	data := "\"http://foo.bar.com\""
	var pu URL
	err := json.Unmarshal([]byte(data), &pu)
	require.NoError(t, err)
	expected, _ := Parse("http://foo.bar.com")
	assert.Equal(t, expected, &pu)
}

func TestValue(t *testing.T) {
	u, _ := url.Parse("http://foo.bar.com")
	pu := URL{*u}
	actual, err := pu.Value()
	require.NoError(t, err)
	assert.Equal(t, "http://foo.bar.com", actual)
}

func TestScan(t *testing.T) {
	pu := URL{}
	err := pu.Scan("http://foo.bar.com")
	require.NoError(t, err)
	assert.Equal(t, "http://foo.bar.com", pu.String())
}
