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

type AreaCity struct {
	AreaCounty
	Counties []*AreaCounty `json:"counties"`
}

func (this *AreaCity) GetAreaByCode(code string) (Arear, bool) {
	if this.Code == code {
		return this, true
	}

	for _, county := range this.Counties {
		if arear, ok := county.GetAreaByCode(code); ok {
			return arear, true
		}
	}

	return nil, false
}

func (this *AreaCity) GetAreaByName(name string) (Arear, bool) {
	if this.Name == name {
		return this, true
	}

	for _, county := range this.Counties {
		if arear, ok := county.GetAreaByName(name); ok {
			return arear, true
		}
	}

	return nil, false
}

func (this *AreaCity) Children() []string {
	children := []string{}
	for _, county := range this.Counties {
        if county.Code != "" {
            children = append(children, county.Code)
        }
		children = append(children, county.Children()...)
	}

	return children
}

func (this *AreaCity) GetCode() string {
    return this.Code
}

func (this *AreaCity) GetName() string {
    return this.Name
}
