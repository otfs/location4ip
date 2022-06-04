package location4ip

import (
	"github.com/ip2location/ip2location-go/v9"
	"location4ip/config"
	"sync"
)

var ip2locationDb *ip2location.DB
var ip2locationDbLock sync.Mutex

//
// GetIpLocationByIp2Location 获取IP位置信息（基于ip2location）
//
func GetIpLocationByIp2Location(ip string) (*IpLocation, error) {
	if ip2locationDb == nil {
		if err := initIp2Location(); err != nil {
			return nil, err
		}
	}

	record, err := ip2locationDb.Get_all(ip)
	if err != nil {
		return nil, err
	}

	location := new(IpLocation)
	location.Ip = ip
	location.CountryCode = record.Country_short
	location.Country = record.Country_long
	location.Region = record.Region
	location.City = record.City
	location.Zipcode = record.Zipcode
	location.Longitude = record.Longitude
	location.Latitude = record.Latitude
	location.EmptyInvalidValues()
	return location, nil
}

func initIp2Location() error {
	ip2locationDbLock.Lock()
	defer ip2locationDbLock.Unlock()

	if ip2locationDb != nil {
		return nil
	}
	dbpath := config.Settings.Ip2LocationDbFile
	db, err := ip2location.OpenDB(dbpath)
	if err != nil {
		return err
	}
	ip2locationDb = db
	return nil
}
