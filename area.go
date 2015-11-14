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

import (
	"bytes"
	"encoding/json"
)

type Arear interface {
	GetName(code string) (string, bool)
	GetCode(name string) (string, bool)
}

type Area struct {
	Countries []*AreaCountry `json:"countries"`
}

func (this *Area) GetName(code string) (string, bool) {
    for _, country := range this.Countries {
        if name, ok := country.GetName(code); ok {
            return name, true
        }
    }
    return "", false
}

func (this *Area) GetCode(name string) (string, bool) {
    for _, country := range this.Countries {
        if code, ok := country.GetCode(name); ok {
            return code, true
        }
    }
    return "", false
}

func NewAreaByJson(areaJson string) *Area {
	area := &Area{
		Countries: make([]*AreaCountry, 0),
	}
	if err := json.NewDecoder(bytes.NewReader([]byte(areaJson))).Decode(area); err != nil {
		return nil
	}

	return area
}

func NewArea(countries ...*AreaCountry) *Area {
	area := &Area{
		Countries: make([]*AreaCountry, 0),
	}

	for _, country := range countries {
		area.Countries = append(area.Countries, country)
	}

	return area
}

