package background

import (
	"github.com/2lovecode/powerjob-go/remote/framework/transporter"
	"github.com/2lovecode/powerjob-go/worker/background/discovery"
)

type OmsLogHandler struct {
	workerAddress          string
	transporter            transporter.Transporter
	serverDiscoveryService discovery.ServerDiscoveryService
}

func NewOmsLogHandler(workerAddress string, transporter transporter.Transporter, serverDiscoveryService discovery.ServerDiscoveryService) *OmsLogHandler {
	return &OmsLogHandler{
		workerAddress:          workerAddress,
		transporter:            transporter,
		serverDiscoveryService: serverDiscoveryService,
	}
}
