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

package api

import (
	"time"

	"github.com/blackducksoftware/perceptor/pkg/hub"
)

// Model .....
type Model struct {
	Pods               map[string]*Pod
	Images             map[string]*ModelImageInfo
	ImageScanQueue     []string
	ImageHubCheckQueue []string
	HubVersion         string
	Config             *ModelConfig
	Timings            *ModelTimings
}

// ModelConfig .....
type ModelConfig struct {
	HubHost string
	HubUser string
	//	HubPasswordEnvVar   string
	HubPort             int
	Port                int
	LogLevel            string
	ConcurrentScanLimit int
}

// ModelTime ...
type ModelTime struct {
	duration     time.Duration
	Minutes      int
	Seconds      int
	Milliseconds int
}

// NewModelTime consumes a time.Duration and calculates the minutes, seconds,
// and milliseconds
func NewModelTime(duration time.Duration) *ModelTime {
	return &ModelTime{
		duration:     duration,
		Minutes:      int(duration / time.Minute),
		Seconds:      int(duration / time.Second),
		Milliseconds: int(duration / time.Millisecond),
	}
}

// ModelTimings ...
type ModelTimings struct {
	HubClientTimeout               ModelTime
	CheckHubForCompletedScansPause ModelTime
	CheckHubThrottle               ModelTime
	CheckForStalledScansPause      ModelTime
	StalledScanClientTimeout       ModelTime
	RefreshImagePause              ModelTime
	EnqueueImagesForRefreshPause   ModelTime
	RefreshThresholdDuration       ModelTime
	ModelMetricsPause              ModelTime
	HubReloginPause                ModelTime
}

// ModelImageInfo .....
type ModelImageInfo struct {
	ScanStatus             string
	TimeOfLastStatusChange string
	ScanResults            *hub.ImageScan
	ImageSha               string
	ImageNames             []string
}
