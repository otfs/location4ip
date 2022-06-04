package main

import (
	"github.com/ip2location/ip2location-go/v9"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

//
// IpLocationService Ip位置服务
//
type IpLocationService interface {

	// GetIpLocation 获取Ip位置信息
	GetIpLocation(ip string) (*IpLocation, error)
}

type DefaultLocationService struct {
}

func (service DefaultLocationService) GetIpLocation(ip string) (*IpLocation, error) {
	locationService := new(Ip2LocationLocationService)
	return locationService.GetIpLocation(ip)
}

//
// Ip2LocationLocationService 基于ip2location的位置服务
//
type Ip2LocationLocationService struct {
}

func (service Ip2LocationLocationService) GetIpLocation(ip string) (*IpLocation, error) {
	db, err := ip2location.OpenDB("db/ip2location.bin")
	if err != nil {
		return nil, err
	}

	record, err := db.Get_all(ip)
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

//
// Ip2RegionLocationService 基于ip2region的位置服务
//
type Ip2RegionLocationService struct {
}

func (service Ip2RegionLocationService) GetIpLocation(ip string) (*IpLocation, error) {
	db, err := ip2region.New("db/ip2region.db")
	if err != nil {
		return nil, err
	}

	record, err := db.BinarySearch(ip)
	if err != nil {
		return nil, err
	}

	location := new(IpLocation)
	location.Country = record.Country
	location.Region = record.Province
	location.City = record.City
	location.EmptyInvalidValues()
	return location, nil
}

//
// IpLocation 位置信息
//
type IpLocation struct {
	Ip          string  `json:"ip"`          // IP地址
	CountryCode string  `json:"countryCode"` // 国家编码
	Country     string  `json:"country"`     // 国家
	Region      string  `json:"region"`      // 地区（省）
	City        string  `json:"city"`        // 城市
	Zipcode     string  `json:"zipcode"`     // 邮政编码
	Longitude   float32 `json:"longitude"`   // 纬度
	Latitude    float32 `json:"latitude"`    // 纬度
}

func (location *IpLocation) EmptyInvalidValues() {
	invalids := map[string]byte{
		"-": 1,
		"0": 1,
		"This parameter is unavailable for selected data file. Please upgrade the data file.": 1,
	}

	if _, ok := invalids[location.Country]; ok {
		location.Country = ""
	}
	if _, ok := invalids[location.Region]; ok {
		location.Region = ""
	}
	if _, ok := invalids[location.City]; ok {
		location.City = ""
	}
	if _, ok := invalids[location.Zipcode]; ok {
		location.Zipcode = ""
	}
	if _, ok := invalids[location.CountryCode]; ok {
		location.CountryCode = ""
	}
}
