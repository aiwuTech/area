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

func (this *AreaProvince) GetName(code string) (string, bool) {
    if this.Code == code {
        return this.Name, true
    }

    for _, city := range this.Cities {
        if name, ok := city.GetName(code); ok {
            return name, true
        }
    }

    return "", false
}

func (this *AreaProvince) GetCode(name string) (string, bool) {
    if this.Name == name {
        return this.Code, true
    }

    for _, city := range this.Cities {
        if code, ok := city.GetCode(name); ok {
            return code, true
        }
    }

    return "", false
}