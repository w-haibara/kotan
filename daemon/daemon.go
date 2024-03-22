package daemon

import (
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/w-haibara/kotan/config"
	"github.com/w-haibara/kotan/unit"
)

func init() {
	if err := unit.LoadAll(); err != nil {
		log.Error("failed to load units", "err", err)
		panic(err.Error())
	}
}

func Run() {
	log.Info("daemon is starting")

	go catchSignals()

	registerRpcs()

	l, err := net.Listen("unix", sockPath())
	if err != nil {
		log.Error("failed to listen", "err", err)
		panic(err.Error())
	}
	defer l.Close()
	log.Info("listening", "path", l.Addr().String())

	rpc.HandleHTTP()

	server := &http.Server{
		ReadHeaderTimeout: 3 * time.Second,
	}
	if err := server.Serve(l); err != nil {
		log.Error("failed to serve", "err", err)
		panic(err.Error())
	}
}

func catchSignals() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	log.Info("daemon is gracefully stopping")
	// TODO: gracefully stop

	<-sigs
	log.Info("daemon is force stopping")
	os.Exit(0)
}

func registerRpcs() {
	if err := rpc.Register(new(RPC)); err != nil {
		log.Error("failed to register rpc", "err", err)
		panic(err.Error())
	}
}

func sockPath() string {
	path := config.UnixDomainSocketPath
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(path), 0750); err != nil {
			log.Error("failed to create directory", "err", err, "path", path)
			panic(err.Error())
		}
	} else if err != nil {
		log.Error("failed to stat", "err", err, "path", path)
		panic(err.Error())
	}

	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		log.Error("failed to remove", "err", err, "path", path)
		panic(err.Error())
	}

	return path
}
