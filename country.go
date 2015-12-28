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

type AreaCountry struct {
	AreaCounty
	Provinces []*AreaProvince `json:"provinces"`
}

func (this *AreaCountry) GetAreaByCode(code string) (Arear, bool) {
	if this.Code == code {
		return this, true
	}

	for _, province := range this.Provinces {
		if arear, ok := province.GetAreaByCode(code); ok {
			return arear, true
		}
	}

	return nil, false
}

func (this *AreaCountry) GetAreaByName(name string) (Arear, bool) {
	if this.Name == name {
		return this, true
	}

	for _, province := range this.Provinces {
		if arear, ok := province.GetAreaByName(name); ok {
			return arear, true
		}
	}

	return nil, false
}

func (this *AreaCountry) Children() []string {
	children := []string{}
	for _, province := range this.Provinces {
		if province.Code != "" {
			children = append(children, province.Code)
		}
		children = append(children, province.Children()...)
	}

	return children
}

func (this *AreaCountry) GetCode() string {
	return this.Code
}

func (this *AreaCountry) GetName() string {
	return this.Name
}
