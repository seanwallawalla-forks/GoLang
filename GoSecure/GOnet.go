/*
show all of the network traffic for my computer.
save the output to a txt file in the current directory.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
)

// geo location info
type GeoIP struct {
	Ip          string `json:"query"`
	Status      string `json:"status"`
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	City        string `json:"city"`
	Zipcode     string `json:"zip"`
	Currency    string `json:"currency"`
	ISP         string `json:"isp"`
	ORG         string `json:"org"`
	VPN         bool   `json:"proxy"`
}

var (
	ipaddrs   []string
	output    []GeoIP
	arraySize int
	err       error
	geo       GeoIP
	resp      *http.Response
	body      []byte
)

// print information about the ip address
func get_info(address string) GeoIP {
	resp, err = http.Get("http://ip-api.com/json/" + address + "?fields=message,status,country,countryCode,city,zip,currency,isp,org,proxy,query")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &geo)
	if err != nil {
		fmt.Println(err)
	}
	return geo
}

func get_ip(wg *sync.WaitGroup) {
	addrs, er := net.InterfaceAddrs()
	if er != nil {
		fmt.Println(er)
	}
	arraySize = len(addrs)
	ipaddrs = make([]string, 0, arraySize)

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipaddrs = append(ipaddrs, ipnet.IP.String())
				// get the info for each ip address and save it to a map
				geo = get_info(ipnet.IP.String())
				output = append(output, geo)
			}
		}
	}
	wg.Done()
}

func main() {
	fmt.Println("Net Info:")
	wg := new(sync.WaitGroup)
	wg.Add(1)
	get_ip(wg)
	wg.Wait()
	fmt.Println(output)

}
