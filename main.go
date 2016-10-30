package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli"
	"errors"
)

//configuration params
var (
	gdKey            string
	gdSecret         string
	gdDomain         string
	gdRecord         string
	gdRecordType     string
	gdRecordTTL      int
	gdDaemon         bool
	gdUpdateInterval int
)

func main() {
	app := cli.NewApp()
	app.Name = "nulldaddy-ddns"
	app.Version = "0.0.1"

	app.Usage = "Poor man's DDNS!"
	app.UsageText = "nulldaddy-ddns [global options] (mode)"
	app.Author = "Ioannis Polyzos <i.polyzos@null-box.org>"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "key, k",
			Usage:       "The developer's 'KEY'.",
			EnvVar:      "GODADDY_KEY",
			Destination: &gdKey,
		},
		cli.StringFlag{
			Name:        "secret, s",
			Usage:       "The developer's 'SECRET'.",
			EnvVar:      "GODADDY_SERCRET",
			Destination: &gdSecret,
		},
		cli.StringFlag{
			Name:        "domain, dn",
			Usage:       "The 'DOMAIN' you would like to update.",
			EnvVar:      "GODADDY_DOMAIN",
			Destination: &gdDomain,
		},
		cli.StringFlag{
			Name:        "record, r",
			Usage:       "The domain 'RECORD' you would like to update.",
			EnvVar:      "GODADDY_RECORD",
			Destination: &gdRecord,
		},
		cli.StringFlag{
			Name:        "record-type, t",
			Value:       "A",
			Usage:       "The 'TYPE' of domain record you would like to update.",
			EnvVar:      "GODADDY_RECORD_TYPE",
			Destination: &gdRecordType,
		},
		cli.IntFlag{
			Name:        "record-ttl, tt",
			Value:       600,
			Usage:       "The 'TTL' of domain record you would like to update.",
			EnvVar:      "GODADDY_RECORD_TTL",
			Destination: &gdRecordTTL,
		},
		cli.BoolFlag{
			Name:        "daemon, d",
			Usage:       "Run godaddy-dns as a'DAEMON'.",
			EnvVar:      "NULLDADDY_DAEMON",
			Destination: &gdDaemon,
		},
		cli.IntFlag{
			Name:        "interval, i",
			Value:       1800,
			Usage:       "The 'INTERVAL' between update in sec.",
			EnvVar:      "NULLDADDY_INTERVAL",
			Destination: &gdUpdateInterval,
		},
	}

	app.Action = func(c *cli.Context) error {

		if gdDaemon {
			log.Println("- Daemon mode enabled.")
		}

		//update domain
		updateDomainRecord(gdKey, gdSecret, gdDomain, gdRecordType, gdRecord)

		//if daemon mode enabled
		if gdDaemon {
			// keep to update the domain periodically
			interval := time.Duration(gdUpdateInterval)
			for range time.Tick(time.Second * interval) {
				updateDomainRecord(gdKey, gdSecret, gdDomain, gdRecordType, gdRecord)
			}
		}
		return nil
	}

	app.Run(os.Args)
}

// update domain record with the new IP address
func updateDomainRecord(gdKey string, gdSecret string, gdDomain string, gdRecordType string, gdRecord string) {
	//get current external IP address
	ip, err := discoverExternalIp()
	if err != nil {
		panic(err.Error())
	}

	//update record
	err = updateRecord(ip)
	if err != nil {
		panic(err.Error())
	}

	log.Printf("- record updated.")

}

//discover current external IP address
func discoverExternalIp() (ip string, err error) {
	ret := &IPAddr{}

	// use ipify.org API to retrieve IP addr
	url := "https://api.ipify.org?format=json"

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	response, _ := netClient.Get(url)

	// read body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	//unmarshal response to return object
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}

	return ret.IP, nil
}

func updateRecord(ip string) error {
	//initiate request data
	ru := &RecordUpdate{
		Data: ip,
		TTL:  gdRecordTTL,
	}

	//marshall record update data
	recordUpdate, err := json.Marshal(ru)
	if err != nil {
		return err
	}

	//create custom transport for httpclient
	transport := &http.Transport{
		///ignore ssl verify
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	//instatiate http client
	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	//format godaddy rest api url
	url := fmt.Sprintf("https://api.godaddy.com/v1/domains/%v/records/%v/%v", gdDomain, gdRecordType, gdRecord)

	// create request
	request, err := http.NewRequest("PUT", url, bytes.NewReader(recordUpdate))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("sso-key %v:%v", gdKey, gdSecret))

	response, err := netClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}

	return nil
}
