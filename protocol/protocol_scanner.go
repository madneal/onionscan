package protocol

import (
	"github.com/neal1991/onionscan/config"
	"github.com/neal1991/onionscan/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
