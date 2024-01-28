package config

type Config struct {
	Todoist  Todoist
	SQLite   SQLite
	Server   Server
	Telegram Telegram
}

// Todoist contains configuration for communication with Todoist API.
type Todoist struct {
	DevBaseURL   string `koanf:"devBaseURL"`
	AuthzBaseURL string `koanf:"authzBaseURL"`
	App          App    `koanf:"app"`
}

// App contains details of Oauth2 application registered in Todoist.
type App struct {
	ClientID     string   `koanf:"clientID"`
	ClientSecret string   `koanf:"clientSecret"`
	RedirectURL  string   `koanf:"redirectURL"`
	Scopes       []string `koanf:"scopes"`
}

// SQLite contains configuration for communication with SQLite database.
type SQLite struct {
	Dbname       string `koanf:"dbname"`
	MaxIdleConns int    `koanf:"maxIdleConns"`
	MaxOpenConns int    `koanf:"maxOpenConns"`
}

// Server contains http server configuration.
type Server struct {
	Host     string `koanf:"host"`
	BasePath string `koanf:"basePath"`
	Port     string `koanf:"port"`
}

// Telegram contains configuration for communication with Telegram API.
type Telegram struct {
	Token string `koanf:"token"`
}
