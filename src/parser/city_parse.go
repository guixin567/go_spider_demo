package parser

import (
	"regexp"
	"zhengxin2/src/model"
)

func CityParse(bytes []byte) (result model.ParseResult) {
	//regCityParse := `<tbody><tr><th><a href="http://album.zhenai.com/u/(.+?)" target="_blank">(.+?)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>(.+?)</td> <td><span class="grayL">居住地：</span>(.+?)</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>\d{2,3}</td> <td><span class="grayL">学.+?历：</span>(.+?)</td> <!----></tr> <tr><td width="180"><span class="grayL">婚况：</span>(.+?)</td> <td width="180"><span class="grayL">身.+?高：</span>(.+?)</td></tr></tbody>`
	regCityParse := `<tbody><tr><th><a href="http://album.zhenai.com/u/(.+?)" target="_blank">(.+?)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>(.+?)</td> <td><span class="grayL">居住地：</span>(.+?)</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>(\d{2,3})</td> <td><span class="grayL">学.+?历：</span>(.+?)</td> </tr> <tr><td width="180"><span class="grayL">婚况：</span>(.+?)</td> <td width="180"><span class="grayL">身.+?高：</span>(.+?)</td></tr></tbody>`
	compile := regexp.MustCompile(regCityParse)
	values := compile.FindAllStringSubmatch(string(bytes), -1)

	for _, value := range values {
		result.Items = append(result.Items, model.Person{
			ID:      value[1],
			Name:    value[2],
			Age:     value[5],
			Marital: value[7],
			Height:  value[8],
		})
	}
	return result
}
