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

import "flag"

type MarsConfig struct {
	// GPU
	ResourceCountName string `yaml:"resourceCountName"`

	// SGPU
	ResourceVCountName  string `yaml:"resourceVCountName"`
	ResourceVMemoryName string `yaml:"resourceVMemoryName"`
	ResourceVCoreName   string `yaml:"resourceVCoreName"`
	TopologyAware       bool   `yaml:"sgpuTopologyAware"`
}

func ParseConfig(fs *flag.FlagSet) {
	// GPU
	fs.StringVar(&MarsResourceCount, "mars-name", "mars-tech.com/gpu", "mars resource count")

	// SGPU
	fs.StringVar(&MarsResourceNameVCount, "mars-vcount", "mars-tech.com/sgpu", "mars vcount name")
	fs.StringVar(&MarsResourceNameVCore, "mars-vcore", "mars-tech.com/vcore", "mars vcore name")
	fs.StringVar(&MarsResourceNameVMemory, "mars-vmemory", "mars-tech.com/vmemory", "mars vmemory name")
	fs.BoolVar(&MarsTopologyAware, "sgpu-topology-aware", false, "sGPU topology aware enable")
}
