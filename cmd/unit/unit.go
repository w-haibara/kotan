package unit

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"slices"

	"github.com/charmbracelet/log"
	"github.com/w-haibara/kotan/config"
)

type Unit interface {
	Start() error
	Stop() error
	Status() error
}

const (
	unitTypeService = ".service"
	unitTypeTimer   = ".timer"
	unitTypeTarget  = ".target"
)

var (
	unitTypes = []string{
		unitTypeService,
		unitTypeTimer,
		unitTypeTarget,
	}
)

func List() ([]string, error) {
	files, err := os.ReadDir(config.UnitFileDir)
	if err != nil {
		log.Error("failed to read unit files", "err", err)
		return nil, err
	}

	var names []string
	for _, file := range files {
		name := file.Name()
		if slices.Contains(unitTypes, filepath.Ext(name)) {
			names = append(names, name)
		}
	}

	return names, nil
}

func Load(name string) (Unit, error) {
	ext := filepath.Ext(name)
	if !slices.Contains(unitTypes, ext) {
		return nil, fmt.Errorf("unknown unit type: %s", ext)
	}

	path := path.Join(config.UnitFileDir, name)

	switch ext {
	case unitTypeService:
		service, err := NewService(name, path)
		if err != nil {
			log.Error("failed to load service file", "err", err)
			return nil, err
		}
		return service, nil
	case unitTypeTimer:
		log.Warn("timer unit is not implemented yet")
		return nil, nil
	case unitTypeTarget:
		log.Warn("target unit is not implemented yet")
		return nil, nil
	default:
		return nil, fmt.Errorf("unknown unit type: %s", ext)
	}
}
