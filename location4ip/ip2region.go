package location4ip

import (
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"io/ioutil"
	"location4ip/config"
	"strings"
	"sync"
)

var ip2regionDb []byte
var ip2regionDbLock sync.Mutex

//
// GetIpLocationByIp2Region 获取IP位置信息（基于ip2region）
//
func GetIpLocationByIp2Region(ip string) (*IpLocation, error) {
	if ip2locationDb == nil {
		if err := initIp2Region(); err != nil {
			return nil, err
		}
	}

	searcher, err := xdb.NewWithBuffer(ip2regionDb)
	if err != nil {
		return nil, err
	}
	defer searcher.Close()

	record, err := searcher.SearchByStr(ip)
	if err != nil {
		return nil, err
	}
	items := strings.Split(record, "|")

	location := new(IpLocation)
	location.Ip = ip
	location.Country = items[0]
	location.CountryCode = items[1]
	location.Region = items[2]
	location.City = items[3]
	location.EmptyInvalidValues()
	return location, nil
}

func initIp2Region() error {
	ip2regionDbLock.Lock()
	defer ip2regionDbLock.Unlock()

	if ip2regionDb != nil {
		return nil
	}
	var err error
	ip2regionDb, err = ioutil.ReadFile(config.Settings.Ip2RegionDbFile)
	if err != nil {
		return err
	}
	return nil
}
