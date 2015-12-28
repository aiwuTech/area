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
	GetAreaByCode(code string) (Arear, bool)
	GetAreaByName(name string) (Arear, bool)
	Children() []string
	GetCode() string
	GetName() string
}

type Area struct {
	Countries []*AreaCountry `json:"countries"`
}

func (this *Area) GetAreaByCode(code string) (Arear, bool) {
	for _, country := range this.Countries {
		if arear, ok := country.GetAreaByCode(code); ok {
			return arear, true
		}
	}
	return nil, false
}

func (this *Area) GetAreaByName(name string) (Arear, bool) {
	for _, country := range this.Countries {
		if arear, ok := country.GetAreaByName(name); ok {
			return arear, true
		}
	}
	return nil, false
}

func (this *Area) Children() []string {
	children := []string{}
	for _, country := range this.Countries {
		if country.Code != "" {
			children = append(children, country.Code)
		}
		children = append(children, country.Children()...)
	}

	return children
}

func (this *Area) GetCode() string {
	return ""
}

func (this *Area) GetName() string {
	return ""
}

func (this *Area) Parent() Arear {
    return nil
}

func (this *Area) String() string {
    return ""
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
