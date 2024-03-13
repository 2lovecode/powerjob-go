package base

import (
	"fmt"
	"strconv"
	"strings"
)

type Address struct {
	host string
	port int
}

func (addr *Address) ToFullAddress() string {
	return ToFullAddress(addr.host, addr.port)
}

func NewAddressFromIpv4(ipv4 string) *Address {
	split := strings.Split(ipv4, ":")
	host := ""
	port := 0
	if len(split) >= 2 {
		host = split[0]
		port, _ = strconv.Atoi(split[1])
	}

	return &Address{
		host: host,
		port: port,
	}
}

func ToFullAddress(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
