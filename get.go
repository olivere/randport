package randport

import (
	"net"
)

// Get returns a random TCP port. It asks the OS to return a random
// address by swifly listening on ":0". The listener will then be
// closed immediately and Get will return the assigned port.
//
// Warning: Get will panic if listen fails.
func Get() int {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	return lis.Addr().(*net.TCPAddr).Port
}
