package location4ip

import "location4ip/config"

// GetIpLocationFunc 获取Ip位置信息
type GetIpLocationFunc func(ip string) (*IpLocation, error)

var ipLocationHandles = map[string]GetIpLocationFunc{
	ProviderIp2Region:   GetIpLocationByIp2Region,
	ProviderIp2Location: GetIpLocationByIp2Location,
}

// GetIpLocation 获取IP位置信息
func GetIpLocation(ip string) (*IpLocation, error) {
	provider := config.Settings.Provider
	ipLocationHandle := ipLocationHandles[provider]
	return ipLocationHandle(ip)
}

const (
	ProviderIp2Location = "ip2location"
	ProviderIp2Region   = "ip2region"
)

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
