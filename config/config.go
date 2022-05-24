package config

type Config struct {
	UPLOAD_DESTINATION   string   // upload destination
	ALLOWED_CONTENT_TYPE []string // allowed content type
}

var _config *Config = nil

func GetConfig() *Config {
	if _config == nil {
		_config = &Config{
			UPLOAD_DESTINATION:   "public/upload",
			ALLOWED_CONTENT_TYPE: []string{"image/jpeg", "image/png", "image/gif", "image/bmp"},
		}
	}
	return _config
}
