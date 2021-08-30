package geoRequest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	BASEURL  = "http://api.ipstack.com/"
	SBASEURL = "https://api.ipstack.com/"
)

type ApiWrapper struct {
	AccessKey string
	Https     bool
	Output    string
}

func (w ApiWrapper) RequestAndProcess(ipAddr []string, security int, language string, raw bool) error {
	var (
		baseUrl      string
		url          string
		jsonBody     baseResponse
		jsonBodyBulk []baseResponse
	)

	if w.Https {
		baseUrl = SBASEURL
	} else {
		baseUrl = BASEURL
	}

	if len(ipAddr) == 1 {
		url = fmt.Sprintf("%s%s?access_key=%s&security=%d&language=%s", baseUrl, string(ipAddr[0]), w.AccessKey, security, language)
		res, err := w.request(url)
		if err != nil {
			return err
		}
		err = json.Unmarshal(res, &jsonBody)
		if err != nil {
			return err
		}
		if !raw {
			w.printResults(jsonBody, url)
			return nil
		} else {
			err = w.printRaw(jsonBody)
			if err != nil {
				return err
			}
			return nil
		}
	} else {
		url = fmt.Sprintf("%s%s?access_key=%s&security=%d&language=%s", baseUrl, strings.Join(ipAddr, ","), w.AccessKey, security, language)
		res, err := w.request(url)
		if err != nil {
			return err
		}
		err = json.Unmarshal(res, &jsonBodyBulk)
		if err != nil {
			return err
		}
		if !raw {
			for _, i := range jsonBodyBulk {
				w.printResults(i, url)

			}
			return nil
		} else {
			err = w.printRaw(jsonBodyBulk)
			if err != nil {
				return err
			}
			return nil
		}
	}
}

func (w ApiWrapper) request(url string) ([]byte, error) {
	var errorResp responseError
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	_ = json.Unmarshal(body, &errorResp)
	if errorResp.Error.Code >= 100 {
		return nil, errors.New(fmt.Sprintf("code: %d, type: %s, info: %s, url: %s", errorResp.Error.Code, errorResp.Error.Type, errorResp.Error.Info, url))
	}
	return body, nil
}

func (w ApiWrapper) printResults(data baseResponse, url string) {
	fmt.Printf("[+] informations about %s\n\n", data.IP)
	fmt.Println("### Metadata")
	fmt.Printf("Type: %s\n\n", data.Type)
	fmt.Println("### Location data")
	fmt.Printf("Continent: %s\nCountry: %s\nRegion: %s\nCity: %s\nZip: %s\n\n", data.ContinentName, data.CountryName, data.RegionName, data.City, data.Zip)
	fmt.Println("API url: ", url)
}

func (w ApiWrapper) printRaw(data interface{}) error {
	p, err := json.Marshal(data)
	if err != nil {
		return err
	}
	fmt.Print(string(p))
	return nil
}
