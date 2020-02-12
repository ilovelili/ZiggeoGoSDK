package ziggeo

// Config ziggeo config
type Config struct {
	Regions    map[string]string
	APIRegions map[string]string
}

func (z *Config) server_api_url() string {
	return "https://srvapi.ziggeo.com"
}

func (z *Config) api_url() string {
	return "https://api-us-east-1.ziggeo.com"
}

// NewConfig constructor
func NewConfig() *Config {
	config := new(Config)
	config.Regions = map[string]string{"r1": "https://srvapi-eu-west-1.ziggeo.com"}
	config.APIRegions = map[string]string{"r1": "https://api-eu-west-1.ziggeo.com"}
	return config
}
