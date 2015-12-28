// Copyright 2015 mint.zhao.chiu@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package area

import "fmt"

var (
	defaultArea Arear
)

func init() {
	defaultArea = NewAreaByJson(areaJson)
}

func GetAreaNameByCode(code string) string {
	arear, f := defaultArea.GetAreaByCode(code)
	if !f {
		return ""
	}
	return arear.GetName()
}

func GetAreaCodeByName(name string) string {
	arear, f := defaultArea.GetAreaByName(name)
	if !f {
		return ""
	}

	return arear.GetCode()
}

func IsAreaByName(name string) bool {
	if GetAreaCodeByName(name) == "" {
		return false
	}

	return true
}

func IsAreaByCode(code string) bool {
	if GetAreaNameByCode(code) == "" {
		return false
	}

	return true
}

func Children(code string) []string {
	arear, f := defaultArea.GetAreaByCode(code)
	if !f {
		return []string{}
	}

	return arear.Children()
}

func AreaDetailByCode(code string) string {
	if code == "" {
		return "未知"
	}

	province := fmt.Sprintf("%s000", code[:3])
	pDetail := GetAreaNameByCode(province)
	if province == code {
		return pDetail
	}

	city := fmt.Sprintf("%s00", code[:4])
	cDetail := GetAreaNameByCode(city)
	if city == code {
		return fmt.Sprintf("%s%s", pDetail, cDetail)
	}

	return fmt.Sprintf("%s%s%s", pDetail, cDetail, GetAreaNameByCode(code))
}
