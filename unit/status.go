package unit

type status int

const (
	statusUnknown status = iota
	statusLoaded
	statusStarting
	statusRunning
	statusStopping
	statusStopped
)

var statusMap = map[string]status{
	"unknown":  statusUnknown,
	"loaded":   statusLoaded,
	"starting": statusStarting,
	"running":  statusRunning,
	"stopping": statusStopping,
	"stopped":  statusStopped,
}

func (s status) String() string {
	for k, v := range statusMap {
		if v == s {
			return k
		}
	}

	return "unknown"
}
