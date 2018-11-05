package root

type MongoConfig struct {
	URL       string `json:"url"`
	DBName    string `json:"dbName"`
	UserTable string `json:"userTable"`
}

type ServerConfig struct {
	Port string `json:"port"`
}

type AuthConfig struct {
	Secret string `json:"secret"`
}

type Config struct {
	Mongo  *MongoConfig  `json:"mongo"`
	Server *ServerConfig `json:"server"`
	Auth   *AuthConfig   `json:"auth"`
}
