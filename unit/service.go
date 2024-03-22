package unit

import (
	"io"

	"github.com/charmbracelet/log"
	"github.com/w-haibara/kotan/exec"
	"github.com/w-haibara/kotan/journal"
	"gopkg.in/ini.v1"
)

type Service struct {
	name          string
	loaded        bool
	runnning      bool
	journalWriter io.Writer

	execStart string
	execStop  string
}

func NewService(name string, path string) (*Service, error) {
	unit, err := ini.Load(path)
	if err != nil {
		log.Error("failed to load service file", "err", err, "path", path)
		return nil, err
	}

	section := unit.Section("Service")

	return &Service{
		name:          name,
		loaded:        true,
		runnning:      false,
		journalWriter: journal.NewWriter(name),

		execStart: section.Key("ExecStart").String(),
		execStop:  section.Key("ExecStop").String(),
	}, nil
}

func (s *Service) Name() string {
	return s.name
}

func (s *Service) Start() error {
	if s.runnning {
		log.Warn("service is already running", "name", s.name)
		return nil
	}

	if err := exec.Exec(s.journalWriter, s.execStart); err != nil {
		log.Error("failed to start service", "name", s.name, "err", err)
		return err
	}

	s.runnning = true

	return nil
}

func (s *Service) Stop() error {
	if !s.runnning {
		log.Warn("service is already stopped", "name", s.name)
		return nil
	}

	if err := exec.Exec(s.journalWriter, s.execStop); err != nil {
		log.Error("failed to stop service", "name", s.name, "err", err)
		return err
	}

	s.runnning = false

	return nil
}

func (s *Service) Status() error {
	log.Warn("status is not implemented yet")
	return nil
}
