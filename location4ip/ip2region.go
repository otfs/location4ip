package location4ip

import (
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
	"location4ip/config"
)

//
// GetIpLocationByIp2Region 获取IP位置信息（基于ip2region）
//
func GetIpLocationByIp2Region(ip string) (*IpLocation, error) {
	db, err := ip2region.New(config.Settings.Ip2RegionDbFile)
	if err != nil {
		return nil, err
	}

	record, err := db.BinarySearch(ip)
	if err != nil {
		return nil, err
	}

	location := new(IpLocation)
	location.Ip = ip
	location.Country = record.Country
	location.Region = record.Province
	location.City = record.City
	location.EmptyInvalidValues()
	return location, nil
}
