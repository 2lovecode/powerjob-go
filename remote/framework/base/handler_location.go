package base

import "fmt"

type HandlerLocation struct {
	rootPath   string
	methodPath string
}

func (hl *HandlerLocation) ToPath() string {
	return fmt.Sprintf("/%s/%s", hl.rootPath, hl.methodPath)
}
