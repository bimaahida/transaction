package config

type Database struct {
	Server   string
	Database string
	User     string
	Password string
}

type Server struct {
	Address string
}

type Context struct {
	Timeout int8
}

type Configuration struct {
	Environment string
	Database    Database
	Server      Server
	Context     Context
}
