package base

type URL struct {
	serverType ServerType
	address    *Address
	location   *HandlerLocation
}
