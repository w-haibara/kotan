package client

import (
	"net/rpc"

	"github.com/w-haibara/kotan/config"
	"github.com/w-haibara/kotan/daemon"

	"github.com/charmbracelet/log"
)

func ListUnit(req daemon.ListUnitReq) (*daemon.ListUnitResp, error) {
	client, err := dial()
	if err != nil {
		log.Error("failed to dial", "err", err)
		return nil, err
	}

	resp := daemon.ListUnitResp{}
	if err = client.Call("RPC.ListUnit", req, &resp); err != nil {
		log.Error("failed to call", "err", err)
		return nil, err
	}

	return &resp, nil
}

func StartUnit(req daemon.StartUnitReq) (*daemon.StartUnitResp, error) {
	client, err := dial()
	if err != nil {
		log.Error("failed to dial", "err", err)
		return nil, err
	}

	resp := daemon.StartUnitResp{}
	if err = client.Call("RPC.StartUnit", req, &resp); err != nil {
		log.Error("failed to call", "err", err)
		return nil, err
	}

	return &resp, nil
}

func StopUnit(req daemon.StopUnitReq) (*daemon.StopUnitResp, error) {
	client, err := dial()
	if err != nil {
		log.Error("failed to dial", "err", err)
		return nil, err
	}

	resp := daemon.StopUnitResp{}
	if err = client.Call("RPC.StopUnit", req, &resp); err != nil {
		log.Error("failed to call", "err", err)
		return nil, err
	}

	return &resp, nil
}

func dial() (*rpc.Client, error) {
	client, err := rpc.DialHTTP("unix", config.UnixDomainSocketPath)
	if err != nil {
		return nil, err
	}

	return client, nil
}
