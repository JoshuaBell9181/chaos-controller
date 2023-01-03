// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.

/*

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

package v1beta1_test

import (
	"github.com/DataDog/chaos-controller/api/v1beta1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sort"
)

var _ = Describe("TargetInjections", func() {
	Describe("GetTargetNames", func() {
		var targetInjections v1beta1.TargetInjections

		Context("with three targets", func() {
			BeforeEach(func() {
				targetInjections = v1beta1.TargetInjections{"target-1": {}, "target-2": {}, "target-3": {}}
			})

			It("should return the list of targets name", func() {
				targetNames := targetInjections.GetTargetNames()
				sort.Strings(targetNames)
				Expect(targetNames).Should(BeEquivalentTo([]string{"target-1", "target-2", "target-3"}))
			})
		})

		Context("without targets", func() {
			BeforeEach(func() {
				targetInjections = v1beta1.TargetInjections{}
			})

			It("should return the list of targets name", func() {
				Expect(targetInjections.GetTargetNames()).Should(BeEquivalentTo([]string{}))
			})
		})
	})
})

var _ = Describe("Check if a target exist into DisruptionStatus targets list", func() {
	var disruptionStatus *v1beta1.DisruptionStatus

	BeforeEach(func() {
		disruptionStatus = &v1beta1.DisruptionStatus{
			TargetInjections: v1beta1.TargetInjections{"test-1": {}},
		}

	})

	AfterEach(func() {
		disruptionStatus = nil
	})

	Context("with an empty target", func() {
		It("should return false", func() {
			Expect(disruptionStatus.HasTarget("")).Should(BeFalse())
		})
	})

	Context("with an existing target", func() {
		It("should return true", func() {
			Expect(disruptionStatus.HasTarget("test-1")).Should(BeTrue())
		})
	})

	Context("with an non existing target", func() {
		It("should return false", func() {
			Expect(disruptionStatus.HasTarget("test-2")).Should(BeFalse())
		})
	})
})