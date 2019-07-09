package dnserrors

import (
	"fmt"
	"net"
	"unicode/utf8"
)

//LookupDNSHost try to lookup in DNS "fqdn" and prints result
func LookupDNSHost(fqdn string) (addrs []string, err error){
	if utf8.RuneCountInString(fqdn) <= 0 {
		return nil, nil
	}
	fmt.Println("Looking up host...", fqdn)
	addr, err := net.LookupHost(fqdn)
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("Request timedout")
		} else if err.Temporary() {
			fmt.Println("Temporary error")
		} else {
			fmt.Println("generic error ", err)
		}
		return nil, err
	}
	return addr, nil
}
