package location4ip

//
// GetIpLocationByMixed 获取IP位置信息
//
func GetIpLocationByMixed(ip string) (*IpLocation, error) {
	location, err := GetIpLocationByIp2Location(ip)
	if err != nil {
		return nil, err
	}

	l1, err := GetIpLocationByIp2Region(ip)
	if l1 != nil && err == nil {
		if l1.Country != "" {
			location.Country = l1.Country
		}
		if l1.Region != "" {
			location.Region = l1.Region
		}
		if l1.City != "" {
			location.City = l1.City
		}
	}
	return location, nil
}