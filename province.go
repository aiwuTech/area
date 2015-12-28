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

type AreaProvince struct {
	AreaCounty
	Cities []*AreaCity `json:"cities"`
}

func (this *AreaProvince) GetAreaByCode(code string) (Arear, bool) {
	if this.Code == code {
		return this, true
	}

	for _, city := range this.Cities {
		if arear, ok := city.GetAreaByCode(code); ok {
			return arear, true
		}
	}

	return nil, false
}

func (this *AreaProvince) GetAreaByName(name string) (Arear, bool) {
	if this.Name == name {
		return this, true
	}

	for _, city := range this.Cities {
		if arear, ok := city.GetAreaByName(name); ok {
			return arear, true
		}
	}

	return nil, false
}

func (this *AreaProvince) Children() []string {
	children := []string{}
	for _, city := range this.Cities {
		if city.Code != "" {
			children = append(children, city.Code)
		}
		children = append(children, city.Children()...)
	}

	return children
}

func (this *AreaProvince) GetCode() string {
    return this.Code
}

func (this *AreaProvince) GetName() string {
    return this.Name
}

