package unit

import (
	"fmt"
	"io"

	"github.com/charmbracelet/log"
	"github.com/w-haibara/kotan/exec"
	"github.com/w-haibara/kotan/journal"
	"github.com/w-haibara/kotan/worker"
	"gopkg.in/ini.v1"
)

type Service struct {
	name          string
	status        status
	journalWriter io.Writer

	execStart string
	execStop  string
}

func LoadService(name string, path string) (*Service, error) {
	unit, err := ini.Load(path)
	if err != nil {
		log.Error("failed to load service file", "err", err, "path", path)
		return nil, err
	}

	section := unit.Section("Service")

	return &Service{
		name:          name,
		status:        statusLoaded,
		journalWriter: journal.NewWriter(name),

		execStart: section.Key("ExecStart").String(),
		execStop:  section.Key("ExecStop").String(),
	}, nil
}

func (s *Service) Name() string {
	return s.name
}

func (s *Service) Start() error {
	if s.status >= statusStarting {
		log.Warn("service is already running", "name", s.name)
		return nil
	}

	if err := SetStatus(s.name, statusStarting); err != nil {
		log.Error("failed to set status", "name", s.name, "err", err)
		return err
	}

	worker.Enqueue("start:"+s.name, func() {
		if err := exec.Exec(s.journalWriter, s.execStart); err != nil {
			log.Error("failed to start service", "name", s.name, "err", err)
			return
		}

		if err := SetStatus(s.name, statusRunning); err != nil {
			log.Error("failed to set status", "name", s.name, "err", err)
			return
		}
	})

	return nil
}

func (s *Service) Stop() error {
	if s.status >= statusStopping {
		log.Warn("service is already stopped", "name", s.name)
		return nil
	}

	if err := SetStatus(s.name, statusStopping); err != nil {
		log.Error("failed to set status", "name", s.name, "err", err)
		return err
	}

	worker.Enqueue("stop:"+s.name, func() {
		if err := exec.Exec(s.journalWriter, s.execStop); err != nil {
			log.Error("failed to stop service", "name", s.name, "err", err)
		}

		if err := SetStatus(s.name, statusStopped); err != nil {
			log.Error("failed to set status", "name", s.name, "err", err)
		}
	})

	return nil
}

func (s *Service) Status() string {
	status := s.status.String()

	return fmt.Sprintf("status: %s", status)
}

func (s *Service) SetStatus(status status) error {
	s.status = status

	log.Info("set status", "name", s.name, "state", status.String())

	return nil
}
