package discovery

import "github.com/2lovecode/powerjob-go/common/model"

type ServerDiscoveryService interface {
	AssertApp() *model.AppInfo
	GetCurrentServerAddress() string
	TimingCheck(timingPool interface{})
}
