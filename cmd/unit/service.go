package unit

import (
	"io"

	"github.com/w-haibara/kotan/exec"
	"github.com/w-haibara/kotan/journal"

	"github.com/charmbracelet/log"
	"gopkg.in/ini.v1"
)

type Service struct {
	Name          string
	Type          string
	ExecStart     string
	ExecStop      string
	journalWriter io.Writer
}

func NewService(name string, path string) (*Service, error) {
	unit, err := ini.Load(path)
	if err != nil {
		log.Error("failed to load service file", "err", err, "path", path)
		return nil, err
	}

	section := unit.Section("Service")

	return &Service{
		Name:          name,
		ExecStart:     section.Key("ExecStart").String(),
		ExecStop:      section.Key("ExecStop").String(),
		journalWriter: journal.NewWriter(name),
	}, nil
}

func (s *Service) Start() error {
	if err := exec.Exec(s.journalWriter, s.ExecStart); err != nil {
		log.Error("failed to start service", "name", s.Name, "err", err)
		return err
	}

	return nil
}

func (s *Service) Stop() error {
	if err := exec.Exec(s.journalWriter, s.ExecStop); err != nil {
		log.Error("failed to stop service", "name", s.Name, "err", err)
		return err
	}

	return nil
}

func (s *Service) Status() error {
	log.Warn("status is not implemented yet")
	return nil
}
