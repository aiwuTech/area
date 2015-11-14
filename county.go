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

type AreaCounty struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func (this *AreaCounty) GetName(code string) (string, bool) {
    if this.Code == code {
        return this.Name, true
    }

    return "", false
}

func (this *AreaCounty) GetCode(name string) (string, bool) {
    if this.Name == name {
        return this.Code, true
    }

    return "", false
}

func (this *AreaCounty) String() string {
	return fmt.Sprintf("[%s]%s", this.Code, this.Name)
}

func NewAreaCounty(name string, code string) *AreaCounty {
	return &AreaCounty{
		Name: name,
		Code: code,
	}
}