package ziggeo

// Application application object
type Application struct {
	application *Ziggeo
}

// NewApplication new application constructor
func NewApplication(application *Ziggeo) *Application {
	app := new(Application)
	app.application = application
	return app
}

// Get Read application
func (a *Application) Get() ([]byte, error) {
	return a.application.Connect().Get("/v1/application", emptyData)
}

// Update update application
func (a *Application) Update(data map[string]string) ([]byte, error) {
	return a.application.Connect().Post("/v1/application", data, "")

}

// GetStatus get status
func (a *Application) GetStatus(data map[string]string) ([]byte, error) {
	return a.application.APIConnect().Get("/server/v1/application/stats", data)
}
