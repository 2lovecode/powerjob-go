package engine

import "github.com/2lovecode/powerjob-go/remote/framework/base"

type Config struct {
	serverType base.ServerType
	typ        string
	address    *base.Address
	actorList  []interface{}
}
