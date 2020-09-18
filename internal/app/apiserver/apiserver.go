package apiserver

// APIServer struct
type APIServer struct {
	config *Config
}

// Initialize APIServer
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

//Start APIServer
func (s *APIServer) Start() error {
	return nil
}
