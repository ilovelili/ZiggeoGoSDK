package ziggeo

import "strings"

// Ziggeo main object
type Ziggeo struct {
	Token          string
	PrivateKey     string
	EncryptionKey  string
	configObj      *Config
	connectObj     *Connect
	apiConnectObj  *Connect
	videosObj      *Video
	applicationObj *Application
}

// NewZiggeo constructor
func NewZiggeo(token, privateKey, encryptionKey string) *Ziggeo {
	ziggeo := new(Ziggeo)
	ziggeo.Token = token
	ziggeo.PrivateKey = privateKey
	ziggeo.EncryptionKey = encryptionKey

	ziggeo.configObj = NewConfig()
	server_api_url := ziggeo.Config().server_api_url()
	for k, v := range ziggeo.Config().Regions {
		if strings.HasPrefix(ziggeo.Token, k) {
			server_api_url = v
			break
		}
	}
	ziggeo.connectObj = NewConnect(ziggeo, server_api_url)

	api_url := ziggeo.Config().api_url()
	for k, v := range ziggeo.Config().APIRegions {
		if strings.HasPrefix(ziggeo.Token, k) {
			api_url = v
			break
		}
	}
	ziggeo.apiConnectObj = NewConnect(ziggeo, api_url)
	return ziggeo
}

func (z *Ziggeo) Config() *Config {
	return z.configObj
}

func (z *Ziggeo) Connect() *Connect {
	return z.connectObj
}

func (z *Ziggeo) APIConnect() *Connect {
	return z.apiConnectObj
}

func (z *Ziggeo) Videos() *Video {
	if z.videosObj == nil {
		z.videosObj = NewVideo(z)
	}
	return z.videosObj
}

func (z *Ziggeo) Application() *Application {
	if z.applicationObj == nil {
		z.applicationObj = NewApplication(z)
	}
	return z.applicationObj
}
