package model

type AppInfo struct {
	appId int
}

func (a *AppInfo) GetAppId() int {
	return a.appId
}
