package uno

// UnoApi is an Uno Application for creating APIs
type Uno struct {
	Server *Server
	DB     *DB
	Config *Config
	Logger Logger
}
