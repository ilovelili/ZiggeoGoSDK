package ziggeo

var emptyData = map[string]string{}

// Video video object
type Video struct {
	application *Ziggeo
}

// NewVideo new video constructor
func NewVideo(application *Ziggeo) *Video {
	video := new(Video)
	video.application = application
	return video
}

// Index Query an array of videos (will return at most 50 videos by default). Newest videos come first.
func (v *Video) Index(data map[string]string) ([]byte, error) {
	return v.application.Connect().Get("/v1/videos/", data)
}

// Count videos count
func (v *Video) Count(data map[string]string) ([]byte, error) {
	return v.application.Connect().Get("/v1/videos/count", data)
}

// Get Read an existing video
func (v *Video) Get(token_or_key string) ([]byte, error) {
	return v.application.Connect().Get("/v1/videos/"+token_or_key+"", emptyData)
}

// GetBulk bulk get
func (v *Video) GetBulk(data map[string]string) ([]byte, error) {
	return v.application.Connect().Get("/v1/videos/get_bulk", data)
}

// GetStatus get video status
func (v *Video) GetStatus(token_or_key string) ([]byte, error) {
	return v.application.Connect().Get("/v1/videos/"+token_or_key+"/stats", emptyData)
}

// StatusBulk bulk get status
func (v *Video) StatusBulk(data map[string]string) ([]byte, error) {
	return v.application.Connect().Get("/v1/videos/stats_bulk", data)
}

// DownloadVideo download video
func (v *Video) DownloadVideo(token_or_key string) ([]byte, error) {
	return v.application.Connect().Get("/v1/videos/"+token_or_key+"/video", emptyData)
}

// DownloadImage download image
func (v *Video) DownloadImage(token_or_key string) ([]byte, error) {
	return v.application.Connect().Get("/v1/videos/"+token_or_key+"/image", emptyData)
}

// Update update video
func (v *Video) Update(token_or_key string, data map[string]string) ([]byte, error) {
	return v.application.Connect().Post("/v1/videos/"+token_or_key+"", data, "")
}

// UpdateBulk bulk update video
func (v *Video) UpdateBulk(token_or_key string, data map[string]string) ([]byte, error) {
	return v.application.Connect().Post("/v1/videos/update_bulk", data, "")
}

// Delete delete an existing video
func (v *Video) Delete(token_or_key string) ([]byte, error) {
	return v.application.Connect().Delete("/v1/videos/"+token_or_key+"", emptyData)
}

// Create create a video
func (v *Video) Create(data map[string]string, file string) ([]byte, error) {
	return v.application.Connect().Post("/v1/videos/", data, file)
}
