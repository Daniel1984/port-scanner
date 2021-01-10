package cfg

import "github.com/port-scanner/pkg/logger"

// Cfg will hold program wide dependencies for easy DI
type Cfg struct {
	*logger.Logger
}

// New returns pointer to Cfg
func New() *Cfg {
	log := logger.New()

	return &Cfg{log}
}

func (c *Cfg) GetAPIPort() string {
	return ":8080"
}
