package discov

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/resolver"
	"net"
	"strings"
)

func parseTarget(t resolver.Target) (scheme, authority, endpoint string, err error) {
	scheme, authority, endpoint = t.Scheme, t.Authority, t.Endpoint
	if scheme != "discov" {
		err = fmt.Errorf("the scheme %q matched the discov builder incorrectly", scheme)
		return
	}
	if authority != "etcd" && authority != "headless" {
		err = fmt.Errorf("invalid authority, must be %q or %q", "etcd", "headless")
		return
	}
	if endpoint == "" {
		err = errors.New("endpoint is empty")
		return
	}
	return
}

func parseEndpoint(endpoint string) (dnsName, port string, err error) {
	s := strings.Split(endpoint, ":")
	if len(s) != 2 {
		err = errors.New("invalid endpoint, must be in form of `dnsName:port`")
		return
	}
	dnsName, port = s[0], s[1]
	return
}

func getEtcdKeyPrefix(srv string) (keyPrefix string) {
	keyPrefix = fmt.Sprintf("/srv/%s", srv)
	return
}

func formatIP(addr string) (addrIP string, ok bool) {
	ip := net.ParseIP(addr)
	if ip == nil {
		return "", false
	}
	if ip.To4() != nil {
		return addr, true
	}
	return "[" + addr + "]", true
}
