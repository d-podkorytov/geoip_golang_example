package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/oschwald/geoip2-golang"
)

func main() {
	db, err := geoip2.Open("GeoIP2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}

  defer db.Close()
	
  if len(os.Args[1])<=1 {fmt.Printf("It needs IP\nRuning example:\n %v 1.1.1.1\n",os.Args[0])
                         return}
                               
	ip := net.ParseIP(os.Args[1])

// If you are using strings that may be invalid, check that ip is not nil

  if ip==nil {fmt.Printf("Can not parse IP '%v'\n",os.Args[1])
                               return}

	record, err := db.City(ip)
	
  if err != nil { log.Fatal(err)}
  
	fmt.Printf("Record : %#v\n", record)
	
  if len(record.Subdivisions) > 0 { fmt.Printf("English name: %v\n", record.Subdivisions[0].Names["en"]) }

	fmt.Printf("\n\nISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
}
