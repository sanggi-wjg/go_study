package server

type Settings struct {
	Debug         bool
	Port          string
	SiteTitle     string
	SiteDeveloper string
	FavIconURL    string
	FavIconPath   string
	StaticURL     string
	StaticRoot    string
	MediaURL      string
	MediaRoot     string
	TemplateRoot  string
}

func NewSettings() *Settings {
	settings := Settings{
		Debug:         true,
		Port:          ":8080",
		SiteTitle:     "AMK_DEMO",
		SiteDeveloper: "쫑꾸",
		FavIconURL:    "/favicon.ico",
		FavIconPath:   "./static/favicon.ico",
		StaticURL:     "/static",
		StaticRoot:    "./static",
		MediaURL:      "/media",
		MediaRoot:     "./media",
		TemplateRoot:  "templates/**/*",
	}

	// 이후 Linux에서 경로 등의 이슈가 있다면 이런식으로 하자
	// var os = runtime.GOOS
	// if os == "windows" {}
	return &settings
}
