// +build !ignore_autogenerated

/*
INTEL CONFIDENTIAL
Copyright (c) 2018 Intel Corporation

The source code contained or described herein and all documents related to
the source code ("Material") are owned by Intel Corporation or its suppliers
or licensors. Title to the Material remains with Intel Corporation or its
suppliers and licensors. The Material contains trade secrets and proprietary
and confidential information of Intel or its suppliers and licensors. The
Material is protected by worldwide copyright and trade secret laws and treaty
provisions. No part of the Material may be used, copied, reproduced, modified,
published, uploaded, posted, transmitted, distributed, or disclosed in any way
without Intel's prior express written permission.
No license under any patent, copyright, trade secret or other intellectual
property right is granted to or conferred upon you by disclosure or delivery
of the Materials, either expressly, by implication, inducement, estoppel or
otherwise. Any license under such intellectual property rights must be express
and approved by Intel in writing.
*/
// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	aggregator "github.com/nervanasystems/carbon/applications/test-exp-service/pkg/apis/aggregator"
	common "github.com/nervanasystems/carbon/applications/test-exp-service/pkg/apis/aggregator/common"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_Run_To_aggregator_Run,
		Convert_aggregator_Run_To_v1_Run,
		Convert_v1_RunList_To_aggregator_RunList,
		Convert_aggregator_RunList_To_v1_RunList,
		Convert_v1_RunSpec_To_aggregator_RunSpec,
		Convert_aggregator_RunSpec_To_v1_RunSpec,
		Convert_v1_RunStatus_To_aggregator_RunStatus,
		Convert_aggregator_RunStatus_To_v1_RunStatus,
		Convert_v1_RunStatusStrategy_To_aggregator_RunStatusStrategy,
		Convert_aggregator_RunStatusStrategy_To_v1_RunStatusStrategy,
		Convert_v1_RunStrategy_To_aggregator_RunStrategy,
		Convert_aggregator_RunStrategy_To_v1_RunStrategy,
	)
}

func autoConvert_v1_Run_To_aggregator_Run(in *Run, out *aggregator.Run, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_RunSpec_To_aggregator_RunSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_RunStatus_To_aggregator_RunStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Run_To_aggregator_Run is an autogenerated conversion function.
func Convert_v1_Run_To_aggregator_Run(in *Run, out *aggregator.Run, s conversion.Scope) error {
	return autoConvert_v1_Run_To_aggregator_Run(in, out, s)
}

func autoConvert_aggregator_Run_To_v1_Run(in *aggregator.Run, out *Run, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_aggregator_RunSpec_To_v1_RunSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_aggregator_RunStatus_To_v1_RunStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_aggregator_Run_To_v1_Run is an autogenerated conversion function.
func Convert_aggregator_Run_To_v1_Run(in *aggregator.Run, out *Run, s conversion.Scope) error {
	return autoConvert_aggregator_Run_To_v1_Run(in, out, s)
}

func autoConvert_v1_RunList_To_aggregator_RunList(in *RunList, out *aggregator.RunList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]aggregator.Run)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_RunList_To_aggregator_RunList is an autogenerated conversion function.
func Convert_v1_RunList_To_aggregator_RunList(in *RunList, out *aggregator.RunList, s conversion.Scope) error {
	return autoConvert_v1_RunList_To_aggregator_RunList(in, out, s)
}

func autoConvert_aggregator_RunList_To_v1_RunList(in *aggregator.RunList, out *RunList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Run)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_aggregator_RunList_To_v1_RunList is an autogenerated conversion function.
func Convert_aggregator_RunList_To_v1_RunList(in *aggregator.RunList, out *RunList, s conversion.Scope) error {
	return autoConvert_aggregator_RunList_To_v1_RunList(in, out, s)
}

func autoConvert_v1_RunSpec_To_aggregator_RunSpec(in *RunSpec, out *aggregator.RunSpec, s conversion.Scope) error {
	out.ExperimentName = in.ExperimentName
	out.PodSelector = in.PodSelector
	out.PodCount = in.PodCount
	out.Parameters = *(*[]string)(unsafe.Pointer(&in.Parameters))
	out.Metrics = *(*map[string]string)(unsafe.Pointer(&in.Metrics))
	out.State = common.RunState(in.State)
	return nil
}

// Convert_v1_RunSpec_To_aggregator_RunSpec is an autogenerated conversion function.
func Convert_v1_RunSpec_To_aggregator_RunSpec(in *RunSpec, out *aggregator.RunSpec, s conversion.Scope) error {
	return autoConvert_v1_RunSpec_To_aggregator_RunSpec(in, out, s)
}

func autoConvert_aggregator_RunSpec_To_v1_RunSpec(in *aggregator.RunSpec, out *RunSpec, s conversion.Scope) error {
	out.ExperimentName = in.ExperimentName
	out.PodSelector = in.PodSelector
	out.PodCount = in.PodCount
	out.Parameters = *(*[]string)(unsafe.Pointer(&in.Parameters))
	out.Metrics = *(*map[string]string)(unsafe.Pointer(&in.Metrics))
	out.State = common.RunState(in.State)
	return nil
}

// Convert_aggregator_RunSpec_To_v1_RunSpec is an autogenerated conversion function.
func Convert_aggregator_RunSpec_To_v1_RunSpec(in *aggregator.RunSpec, out *RunSpec, s conversion.Scope) error {
	return autoConvert_aggregator_RunSpec_To_v1_RunSpec(in, out, s)
}

func autoConvert_v1_RunStatus_To_aggregator_RunStatus(in *RunStatus, out *aggregator.RunStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1_RunStatus_To_aggregator_RunStatus is an autogenerated conversion function.
func Convert_v1_RunStatus_To_aggregator_RunStatus(in *RunStatus, out *aggregator.RunStatus, s conversion.Scope) error {
	return autoConvert_v1_RunStatus_To_aggregator_RunStatus(in, out, s)
}

func autoConvert_aggregator_RunStatus_To_v1_RunStatus(in *aggregator.RunStatus, out *RunStatus, s conversion.Scope) error {
	return nil
}

// Convert_aggregator_RunStatus_To_v1_RunStatus is an autogenerated conversion function.
func Convert_aggregator_RunStatus_To_v1_RunStatus(in *aggregator.RunStatus, out *RunStatus, s conversion.Scope) error {
	return autoConvert_aggregator_RunStatus_To_v1_RunStatus(in, out, s)
}

func autoConvert_v1_RunStatusStrategy_To_aggregator_RunStatusStrategy(in *RunStatusStrategy, out *aggregator.RunStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1_RunStatusStrategy_To_aggregator_RunStatusStrategy is an autogenerated conversion function.
func Convert_v1_RunStatusStrategy_To_aggregator_RunStatusStrategy(in *RunStatusStrategy, out *aggregator.RunStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1_RunStatusStrategy_To_aggregator_RunStatusStrategy(in, out, s)
}

func autoConvert_aggregator_RunStatusStrategy_To_v1_RunStatusStrategy(in *aggregator.RunStatusStrategy, out *RunStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_aggregator_RunStatusStrategy_To_v1_RunStatusStrategy is an autogenerated conversion function.
func Convert_aggregator_RunStatusStrategy_To_v1_RunStatusStrategy(in *aggregator.RunStatusStrategy, out *RunStatusStrategy, s conversion.Scope) error {
	return autoConvert_aggregator_RunStatusStrategy_To_v1_RunStatusStrategy(in, out, s)
}

func autoConvert_v1_RunStrategy_To_aggregator_RunStrategy(in *RunStrategy, out *aggregator.RunStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1_RunStrategy_To_aggregator_RunStrategy is an autogenerated conversion function.
func Convert_v1_RunStrategy_To_aggregator_RunStrategy(in *RunStrategy, out *aggregator.RunStrategy, s conversion.Scope) error {
	return autoConvert_v1_RunStrategy_To_aggregator_RunStrategy(in, out, s)
}

func autoConvert_aggregator_RunStrategy_To_v1_RunStrategy(in *aggregator.RunStrategy, out *RunStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_aggregator_RunStrategy_To_v1_RunStrategy is an autogenerated conversion function.
func Convert_aggregator_RunStrategy_To_v1_RunStrategy(in *aggregator.RunStrategy, out *RunStrategy, s conversion.Scope) error {
	return autoConvert_aggregator_RunStrategy_To_v1_RunStrategy(in, out, s)
}
