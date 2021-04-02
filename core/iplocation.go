package core

import (
	"fmt"
	"github.com/ip2location/ip2location-go"
)

func getLocation(ip string) {
	ip2location.Open("./IP2LOCATION-LITE-DB5.BIN")

	country := ip2location.Get_country_long(ip)
    region := ip2location.Get_region(ip)
    city := ip2location.Get_city(ip)
    latitude := ip2location.Get_latitude(ip)
    longitude := ip2location.Get_longitude(ip)

	ip2location.Close()
}