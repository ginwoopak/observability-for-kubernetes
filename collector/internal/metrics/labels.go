// Copyright 2015 Google Inc. All Rights Reserved.
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

// Copyright 2018 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package metrics

// Definition of labels supported in Set.

var (
	LabelMetricSetType = LabelDescriptor{
		Key:         "type",
		Description: "Type of the metrics set (container, pod, namespace, node, cluster)",
	}
	MetricSetTypeSystemContainer = "sys_container"
	MetricSetTypePodContainer    = "pod_container"
	MetricSetTypePod             = "pod"
	MetricSetTypeNamespace       = "ns"
	MetricSetTypeNode            = "node"
	MetricSetTypeCluster         = "cluster"

	LabelCluster = LabelDescriptor{
		Key:         "cluster",
		Description: "The name of the kubernetes cluster",
	}
	LabelPodId = LabelDescriptor{
		Key:         "pod_id",
		Description: "The unique ID of the pod",
	}
	LabelPodName = LabelDescriptor{
		Key:         "pod_name",
		Description: "The name of the pod",
	}
	LabelNamespaceName = LabelDescriptor{
		Key:         "namespace_name",
		Description: "The name of the namespace",
	}
	LabelWorkloadName = LabelDescriptor{
		Key:         "workload_name",
		Description: "Workload name, derived from top level Deployment or DaemonSet",
	}
	LabelWorkloadKind = LabelDescriptor{
		Key:         "workload_kind",
		Description: "Workload Kind, derived from top level Deployment or DaemonSet",
	}
	LabelDesired = LabelDescriptor{
		Key:         "desired",
		Description: "The desired number for the Pod",
	}
	LabelAvailable = LabelDescriptor{
		Key:         "available",
		Description: "The available number for the Pod",
	}
	LabelReason = LabelDescriptor{
		Key:         "reason",
		Description: "The failure reason for the Pod",
	}
	LabelMessage = LabelDescriptor{
		Key:         "message",
		Description: "The failure message for the Pod",
	}
	LabelPodNamespaceUID = LabelDescriptor{
		Key:         "namespace_id",
		Description: "The UID of namespace of the pod",
	}
	LabelContainerName = LabelDescriptor{
		Key:         "container_name",
		Description: "User-provided name of the container or full container name for system containers",
	}
	LabelLabels = LabelDescriptor{
		Key:         "labels",
		Description: "Comma-separated list of user-provided labels",
	}
	LabelNodename = LabelDescriptor{
		Key:         "nodename",
		Description: "nodename where the container ran",
	}
	LabelNodeRole = LabelDescriptor{
		Key:         "node_role",
		Description: "Node role worker or control-plane",
	}
	LabelHostname = LabelDescriptor{
		Key:         "hostname",
		Description: "Hostname where the container ran",
	}
	LabelResourceID = LabelDescriptor{
		Key:         "resource_id",
		Description: "Identifier(s) specific to a metric",
	}
	LabelHostID = LabelDescriptor{
		Key:         "host_id",
		Description: "Identifier specific to a host. Set by cloud provider or user",
	}
	LabelContainerBaseImage = LabelDescriptor{
		Key:         "container_base_image",
		Description: "User-defined image name that is run inside the container",
	}
	// The label is populated only for GCM
	LabelCustomMetricName = LabelDescriptor{
		Key:         "custom_metric_name",
		Description: "User-defined name of the exported custom metric",
	}
	LabelGCEResourceID = LabelDescriptor{
		Key:         "resource_id",
		Description: "Resource id for nodes specific for GCE.",
	}
	LabelGCEResourceType = LabelDescriptor{
		Key:         "resource_type",
		Description: "Resource types for nodes specific for GCE.",
	}
	LabelNodeSchedulable = LabelDescriptor{
		Key:         "schedulable",
		Description: "Node schedulable status.",
	}
	LabelVolumeName = LabelDescriptor{
		Key:         "volume_name",
		Description: "The name of the volume.",
	}
	LabelPVCName = LabelDescriptor{
		Key:         "pvc_name",
		Description: "The name of the persistent volume claim.",
	}
	LabelAcceleratorMake = LabelDescriptor{
		Key:         "make",
		Description: "Make of the accelerator (nvidia, amd, google etc.)",
	}
	LabelAcceleratorModel = LabelDescriptor{
		Key:         "model",
		Description: "Model of the accelerator (tesla-p100, tesla-k80 etc.)",
	}
	LabelAcceleratorID = LabelDescriptor{
		Key:         "accelerator_id",
		Description: "ID of the accelerator",
	}
)

type LabelDescriptor struct {
	// Key to use for the label.
	Key string `json:"key,omitempty"`

	// Description of the label.
	Description string `json:"description,omitempty"`
}

var commonLabels = []LabelDescriptor{
	LabelNodename,
	LabelHostname,
	LabelHostID,
}

var containerLabels = []LabelDescriptor{
	LabelContainerName,
	LabelContainerBaseImage,
}

var podLabels = []LabelDescriptor{
	LabelPodName,
	LabelPodId,
	LabelPodNamespaceUID,
	LabelLabels,
}

var metricLabels = []LabelDescriptor{
	LabelResourceID,
}

var customMetricLabels = []LabelDescriptor{
	LabelCustomMetricName,
}

var acceleratorLabels = []LabelDescriptor{
	LabelAcceleratorMake,
	LabelAcceleratorModel,
	LabelAcceleratorID,
}

// Labels exported to GCM. The number of labels that can be exported to GCM is limited by 10.
var gcmLabels = []LabelDescriptor{
	LabelMetricSetType,
	LabelPodName,
	LabelNamespaceName,
	LabelHostname,
	LabelHostID,
	LabelContainerName,
	LabelContainerBaseImage,
	LabelCustomMetricName,
	LabelResourceID,
}

var gcmNodeAutoscalingLabels = []LabelDescriptor{
	LabelGCEResourceID,
	LabelGCEResourceType,
	LabelHostname,
}

func CommonLabels() []LabelDescriptor {
	result := make([]LabelDescriptor, len(commonLabels))
	copy(result, commonLabels)
	return result
}

func ContainerLabels() []LabelDescriptor {
	result := make([]LabelDescriptor, len(containerLabels))
	copy(result, containerLabels)
	return result
}

func PodLabels() []LabelDescriptor {
	result := make([]LabelDescriptor, len(podLabels))
	copy(result, podLabels)
	return result
}

func MetricLabels() []LabelDescriptor {
	result := make([]LabelDescriptor, len(metricLabels)+len(customMetricLabels))
	copy(result, metricLabels)
	copy(result, customMetricLabels)
	return result
}

func SupportedLabels() []LabelDescriptor {
	result := CommonLabels()
	result = append(result, PodLabels()...)
	return append(result, MetricLabels()...)
}

func GcmLabels() map[string]LabelDescriptor {
	result := make(map[string]LabelDescriptor, len(gcmLabels))
	for _, l := range gcmLabels {
		result[l.Key] = l
	}
	return result
}
func GcmNodeAutoscalingLabels() map[string]LabelDescriptor {
	result := make(map[string]LabelDescriptor, len(gcmNodeAutoscalingLabels))
	for _, l := range gcmNodeAutoscalingLabels {
		result[l.Key] = l
	}
	return result
}
