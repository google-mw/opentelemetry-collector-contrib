// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package translator

import (
	"go.opentelemetry.io/collector/consumer/pdata"
	semconventions "go.opentelemetry.io/collector/translator/conventions"
)

// ServiceData provides the shape for unmarshalling service version.
type ServiceData struct {
	Version         string `json:"version,omitempty"`
	CompilerVersion string `json:"compiler_version,omitempty"`
	Compiler        string `json:"compiler,omitempty"`
}

func makeService(resource pdata.Resource) *ServiceData {
	var (
		service *ServiceData
	)
	if resource.IsNil() {
		return service
	}
	verStr, ok := resource.Attributes().Get(semconventions.AttributeServiceVersion)
	if !ok {
		verStr, ok = resource.Attributes().Get(semconventions.AttributeContainerTag)
	}
	if ok {
		service = &ServiceData{
			Version: verStr.StringVal(),
		}
	}
	return service
}
