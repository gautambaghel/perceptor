/*
Copyright (C) 2018 Synopsys, Inc.

Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements. See the NOTICE file
distributed with this work for additional information
regarding copyright ownership. The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied. See the License for the
specific language governing permissions and limitations
under the License.
*/

package actions

import (
	"fmt"

	m "github.com/blackducksoftware/perceptor/pkg/core/model"
)

// ShouldScanLayer .....
type ShouldScanLayer struct {
	Layer   string
	Success chan bool
	Err     chan error
}

func NewShouldScanLayer(layer string) *ShouldScanLayer {
	return &ShouldScanLayer{Layer: layer, Success: make(chan bool), Err: make(chan error)}
}

// Apply .....
func (g *ShouldScanLayer) Apply(model *m.Model) {
	// ScanStatus:
	// unknown -> error
	// not scanned -> yes
	// complete -> no
	// running scan client -> no
	// running hub scan -> no
	layerInfo, ok := model.Layers[g.Layer]
	if !ok {
		g.Err <- fmt.Errorf("layer %s not found", g.Layer)
		return
	}
	switch layerInfo.ScanStatus {
	case m.ScanStatusUnknown:
		g.Err <- fmt.Errorf("layer %s status unknown", g.Layer)
	case m.ScanStatusNotScanned:
		g.Success <- true
	default:
		g.Success <- false
	}
}