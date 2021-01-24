// +build !linux android

package device

import (
	"github.com/octetsecurity/wireguard-go/conn"
	"github.com/octetsecurity/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
