/*
Copyright 2022 The Katalyst Authors.

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

package region

import (
	"github.com/kubewharf/katalyst-core/pkg/agent/sysadvisor/metacache"
	"github.com/kubewharf/katalyst-core/pkg/agent/sysadvisor/types"
	"github.com/kubewharf/katalyst-core/pkg/metrics"
)

type QoSRegionDedicatedNuma struct {
	*QoSRegionBase
}

// NewQoSRegionDedicatedNuma returns a share qos region instance
func NewQoSRegionDedicatedNuma(name string, ownerPoolName string, regionType QoSRegionType,
	regionPolicy types.CPUProvisionPolicyName, metaCache *metacache.MetaCache, emitter metrics.MetricEmitter) QoSRegion {
	r := &QoSRegionDedicatedNuma{
		QoSRegionBase: NewQoSRegionBase(name, ownerPoolName, regionType, regionPolicy, metaCache, emitter),
	}
	return r
}

func (r *QoSRegionDedicatedNuma) TryUpdateControlKnob() {
}

func (r *QoSRegionDedicatedNuma) GetControlKnobUpdated() (types.ControlKnob, error) {
	return nil, nil
}

func (r *QoSRegionDedicatedNuma) GetHeadroom() (int, error) {
	return 0, nil
}
