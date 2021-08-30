package main

import (
	"flag"
	"github.com/iiestIT/geoRequest"
	"log"
	"net"
)

var (
	accessKey string
	https bool
	ip string
	hostname string
	language string
	security int
	raw bool
)

func init() {
	flag.StringVar(&accessKey, "key", "", "Specify your ipstack API key. For more information visit the official ipstack website: https://ipstack.com/.")
	flag.BoolVar(&https, "https", false, "If you purchased the basic plan or higher, you should set this option true. Default is false.")
	flag.StringVar(&ip, "ip", "", "If you want to get information about a single ip address, use this argument.")
	flag.StringVar(&hostname, "hostname", "", "If you want to get information about a single hostname, use this argument.")
	flag.StringVar(&language, "lan", "en", "If you want to modify the response language of the api, use this argument.")
	flag.IntVar(&security, "sec", 0, "If you purchased the professional plus plan or higher, you can set this argument to 1 to enable the security module.")
	flag.BoolVar(&raw, "raw", false, "If you want to get the raw api output, set this argument to true")

	flag.Parse()
}

func main() {
	api := geoRequest.ApiWrapper{
		AccessKey: accessKey,
		Https: https,
		Output: "json",
	}

	if len(ip) > 0 {
		err := api.RequestAndProcess([]string{ip}, security, language, raw)
		if err != nil {
			log.Fatalln(err)
		}
		return
	} else {
		if len(hostname) > 0{
			addr, err := net.LookupHost(hostname)
			if err != nil {
				log.Fatalln(err)
				return
			}
			err = api.RequestAndProcess([]string{addr[0]}, security, language, raw)
			if err != nil {
				log.Fatalln(err)
			}
			return
		}
		log.Fatalln("You need to set 'ip' or 'hostname' argument!")
		return
	}
}
