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
	Name() string
	Start() error
	Stop() error
	Status() error
}

const (
	unitTypeService = ".service"
	unitTypeTimer   = ".timer"
	unitTypeTarget  = ".target"
)

var unitTypes = []string{
	unitTypeService,
	unitTypeTimer,
	unitTypeTarget,
}

var unitMap = make(map[string]Unit)

func LoadAll() error {
	files, err := os.ReadDir(config.UnitFileDir)
	if err != nil {
		log.Error("failed to read unit files", "err", err)
		return err
	}

	for _, file := range files {
		name := file.Name()
		if slices.Contains(unitTypes, filepath.Ext(name)) {
			if err := Load(file.Name()); err != nil {
				log.Error("failed to load unit", "name", name, "err", err)
				return err
			}
		}
	}

	return nil
}

func Load(name string) error {
	ext := filepath.Ext(name)
	if !slices.Contains(unitTypes, ext) {
		return fmt.Errorf("unknown unit type: %s", ext)
	}

	path := path.Join(config.UnitFileDir, name)

	switch ext {
	case unitTypeService:
		service, err := NewService(name, path)
		if err != nil {
			log.Error("failed to load service file", "err", err)
			return err
		}

		unitMap[name] = service

		return nil
	case unitTypeTimer:
		log.Warn("timer unit is not implemented yet")
		return nil
	case unitTypeTarget:
		log.Warn("target unit is not implemented yet")
		return nil
	default:
		return fmt.Errorf("unknown unit type: %s", ext)
	}
}

type UnitInfo struct {
	Name string
}

func List() map[string]UnitInfo {
	units := make(map[string]UnitInfo)
	for name := range unitMap {
		units[name] = UnitInfo{
			Name: name,
		}
	}

	return units
}

func Find(name string) (Unit, error) {
	unit, ok := unitMap[name]
	if !ok {
		return nil, fmt.Errorf("unit not found: %s", name)
	}

	return unit, nil
}
