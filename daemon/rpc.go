package daemon

type RPC struct{}

type ListUnitReq struct {
}
type ListUnitResp struct {
}

func (rpc *RPC) ListUnit(req *ListUnitReq, resp *ListUnitResp) error {
	return nil
}

type StartUnitReq struct {
}
type StartUnitResp struct {
}

func (rpc *RPC) StartUnit(req *StartUnitReq, resp *StartUnitResp) error {
	return nil
}

type StopUnitReq struct {
}
type StopUnitResp struct {
}

func (rpc *RPC) StopUnit(req *StopUnitReq, resp *StopUnitResp) error {
	return nil
}
