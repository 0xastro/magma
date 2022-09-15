/*
Copyright 2022 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sas_test

import (
	"testing"

	"magma/dp/cloud/go/services/dp/active_mode_controller/message_generator/sas"
	"magma/dp/cloud/go/services/dp/active_mode_controller/protos/active_mode"
)

func TestDeregistrationRequestGenerator(t *testing.T) {
	cbsd := &active_mode.Cbsd{CbsdId: "some_id"}
	g := &sas.DeregistrationRequestGenerator{}
	actual := g.GenerateRequests(cbsd)
	expected := []*request{{
		requestType: "deregistrationRequest",
		data: `{
	"cbsdId": "some_id"
}`,
	}}
	assertRequestsEqual(t, expected, actual)
}