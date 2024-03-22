package config

import (
	"os"
	"path/filepath"
)

// TODO: read from config file or env
var (
	UnitFileDir          = "./_units/"
	UnixDomainSocketPath = filepath.Join(os.TempDir(), "kotan", "kotan.sock")
)
