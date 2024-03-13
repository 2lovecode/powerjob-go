package transporter

import "github.com/2lovecode/powerjob-go/remote/framework/base"

type Transporter interface {
	GetProtocol() Protocol
	Tell(url base.URL, request interface{})
	Ask(url base.URL, request interface{})
}
