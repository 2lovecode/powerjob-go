package engine

import "github.com/2lovecode/powerjob-go/remote/framework/transporter"

type Output struct {
	transporter transporter.Transporter
}

func (o *Output) GetTransporter() transporter.Transporter {
	return o.transporter
}
