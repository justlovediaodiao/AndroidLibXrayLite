package libv2ray

import (
	"encoding/json"

	"github.com/justlovediaodiao/https-proxy/client"
)

func StartVpoint(c *client.Config) error {
	return client.Start(c)
}

func CloseVpoint() error {
	return client.Close()
}

func LoadJSONConfig(content string) (*client.Config, error) {
	var c client.Config
	err := json.Unmarshal([]byte(content), &c)
	return &c, err
}
