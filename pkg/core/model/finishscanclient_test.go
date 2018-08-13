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

package model

import (
	"fmt"
	"testing"
)

// TestScanClientFails .....
func TestScanClientFails(t *testing.T) {
	model := NewModel()
	image := *NewImage("abc", "4.0", DockerImageSha("23bcf2dae3"))
	model.addImage(image, 0)
	model.setImageScanStatus(image.Sha, ScanStatusInQueue)
	model.setImageScanStatus(image.Sha, ScanStatusRunningScanClient)
	model.finishRunningScanClient(&image, fmt.Errorf("oops, unable to run scan client"))

	if model.Images[image.Sha].ScanStatus != ScanStatusInQueue {
		t.Logf("expected ScanStatus of InQueue, got %s", model.Images[image.Sha].ScanStatus)
		t.Fail()
	}

	nextImage := model.getNextImageFromScanQueue()
	if image != *nextImage {
		t.Logf("expected nextImage of %v, got %v", image, nextImage)
		t.Fail()
	}
}
