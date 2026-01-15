/*
Copyright 2025 The HAMi Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mars

import (
	"fmt"

	"github.com/Project-HAMi/HAMi/pkg/device"
)

const (
	MarsSDeviceAnno       = "mars-tech.com/node-gpu-devices"
	MarsAllocatedSDevices = "mars-tech.com/gpu-devices-allocated"
	MarsPredicateTime     = "mars-tech.com/predicate-time"

	MarsUseUUID   = "mars-tech.com/use-gpuuuid"
	MarsNoUseUUID = "mars-tech.com/nouse-gpuuuid"

	MarsSGPUQosPolicy     = "mars-tech.com/sgpu-qos-policy"
	MarsSGPUTopologyAware = "mars-tech.com/sgpu-topology-aware"
)

const (
	BestEffort = "best-effort"
	FixedShare = "fixed-share"
	BurstShare = "burst-share"
)

type MarsSDeviceInfo struct {
	UUID              string `json:"uuid"`
	BDF               string `json:"bdf,omitempty"`
	Model             string `json:"model,omitempty"`
	TotalDevCount     int32  `json:"totalDevCount,omitempty"`
	TotalCompute      int32  `json:"totalCompute,omitempty"`
	TotalVRam         int32  `json:"totalVRam,omitempty"`
	AvailableDevCount int32  `json:"availableDevCount,omitempty"`
	AvailableCompute  int32  `json:"availableCompute,omitempty"`
	AvailableVRam     int32  `json:"availableVRam,omitempty"`
	Numa              int32  `json:"numa,omitempty"`
	Healthy           bool   `json:"healthy,omitempty"`
	QosPolicy         string `json:"qosPolicy,omitempty"`
	LinkZone          int32  `json:"linkZone,omitempty"`
}
type NodeMarsSDeviceInfo []*MarsSDeviceInfo

type ContainerMarsSDevice struct {
	UUID    string `json:"uuid"`
	Compute int32  `json:"compute,omitempty"`
	VRam    int32  `json:"vRam,omitempty"`
}
type ContainerMarsSDevices []ContainerMarsSDevice
type PodMarsSDevice []ContainerMarsSDevices

func (ni NodeMarsSDeviceInfo) String() string {
	str := "\n"

	for _, i := range ni {
		str += fmt.Sprintf("MarsSDeviceInfo[%s]: TotalDevCount=%d, TotalCompute=%d, TotalVRam=%d, Numa=%d, Healthy=%t, QosPolicy=%s, LinkZone=%d\n",
			i.UUID, i.TotalDevCount, i.TotalCompute, i.TotalVRam, i.Numa, i.Healthy, i.QosPolicy, i.LinkZone)
	}

	return str
}

func (sdev *PodMarsSDevice) String() string {
	str := "\nPodMarsSDevice:\n"

	for ctrIdx, ctrDevices := range *sdev {
		str += fmt.Sprintf("  container[%d]:\n", ctrIdx)

		for _, device := range ctrDevices {
			str += fmt.Sprintf("    SDevice[%s]: Compute=%d, VRam=%d\n",
				device.UUID, device.Compute, device.VRam)
		}
	}

	return str
}

func convertMarsSDeviceToHAMIDevice(marsSDevices []*MarsSDeviceInfo) []*device.DeviceInfo {
	hamiDevices := make([]*device.DeviceInfo, len(marsSDevices))

	for idx, sdevice := range marsSDevices {
		hamiDevices[idx] = &device.DeviceInfo{
			ID:           sdevice.UUID,
			Index:        uint(idx),
			Count:        sdevice.TotalDevCount,
			Devmem:       sdevice.TotalVRam,
			Devcore:      sdevice.TotalCompute,
			Type:         MarsSGPUDevice,
			Numa:         int(sdevice.Numa),
			Mode:         "",
			MIGTemplate:  []device.Geometry{},
			Health:       sdevice.Healthy,
			DeviceVendor: MarsSGPUDevice,
			CustomInfo: map[string]any{
				"QosPolicy": sdevice.QosPolicy,
				"Model":     sdevice.Model,
				"LinkZone":  sdevice.LinkZone,
			},
		}
	}

	return hamiDevices
}

func convertHAMIPodDeviceToMarsPodDevice(hamiPodDevices device.PodSingleDevice) PodMarsSDevice {
	marsDevices := make(PodMarsSDevice, len(hamiPodDevices))

	for ctrIdx, ctrDevices := range hamiPodDevices {
		marsDevices[ctrIdx] = make(ContainerMarsSDevices, len(ctrDevices))
		for deviceIdx, device := range ctrDevices {
			marsDevices[ctrIdx][deviceIdx] = ContainerMarsSDevice{
				UUID:    device.UUID,
				VRam:    device.Usedmem,
				Compute: device.Usedcores,
			}
		}
	}

	return marsDevices
}
