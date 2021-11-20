package parser

import (
	"regexp"
	"zhengxin2/src/model"
)

// CityListParse /正则将字节解析成ParseResult
func CityListParse(bytes []byte) (result model.ParseResult) {

	regCityList := `<a href="(http://www.zhenai.com/zhenghun/.+?)" data-v-1573aa7c>(.+?)</a>`
	compile := regexp.MustCompile(regCityList)
	values := compile.FindAllStringSubmatch(string(bytes), -1)

	for _, value := range values {
		result.Requests = append(result.Requests, model.Request{Url: value[1], ParseFunction: CityParse})
		result.Items = append(result.Items, value[2])
	}
	return result
}
