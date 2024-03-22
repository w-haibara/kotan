package daemon

import (
	"github.com/charmbracelet/log"
	"github.com/w-haibara/kotan/unit"
)

type RPC struct{}

type ListUnitReq struct{}
type ListUnitResp struct {
	UnitInfo map[string]unit.UnitInfo
}

func (rpc *RPC) ListUnit(req *ListUnitReq, resp *ListUnitResp) error {
	log.Info("Call ListUnit", "req", req)

	*resp = ListUnitResp{
		UnitInfo: unit.List(),
	}

	log.Info("Call ListUnit", "resp", resp)

	return nil
}

type StartUnitReq struct {
	Name string
}
type StartUnitResp struct {
}

func (rpc *RPC) StartUnit(req *StartUnitReq, resp *StartUnitResp) error {
	log.Info("Call StartUnit", "req", req)

	unit, err := unit.Find(req.Name)
	if err != nil {
		log.Error("failed to find unit", "err", err)
		return err
	}

	if err := unit.Start(); err != nil {
		log.Error("failed to start unit", "err", err)
		return err
	}

	log.Info("Call StartUnit", "resp", resp)

	return nil
}

type StopUnitReq struct {
	Name string
}
type StopUnitResp struct {
}

func (rpc *RPC) StopUnit(req *StopUnitReq, resp *StopUnitResp) error {
	log.Info("Call StopUnit", "req", req)

	unit, err := unit.Find(req.Name)
	if err != nil {
		log.Error("failed to find unit", "err", err)
		return err
	}

	if err := unit.Stop(); err != nil {
		log.Error("failed to stop unit", "err", err)
		return err
	}

	log.Info("Call StopUnit", "resp", resp)

	return nil
}
