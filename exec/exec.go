package exec

import (
	"io"
	"os/exec"

	"github.com/charmbracelet/log"
)

func Exec(w io.Writer, script string) error {
	log.Info("execute command", "script", script)
	cmd := exec.Command("bash", "-c", script)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Error("failed to get stdout pipe", "err", err)
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Error("failed to get stderr pipe", "err", err)
		return err
	}

	go func() {
		if _, err := io.Copy(w, stdout); err != nil {
			log.Error("failed to copy stdout", "err", err)
		}
	}()
	go func() {
		if _, err := io.Copy(w, stderr); err != nil {
			log.Error("failed to copy stderr", "err", err)
		}
	}()

	if err := cmd.Start(); err != nil {
		log.Error("failed to start command", "err", err)
		return err
	}

	if err := cmd.Wait(); err != nil {
		log.Error("failed to wait command", "err", err)
		return err
	}

	return nil
}
