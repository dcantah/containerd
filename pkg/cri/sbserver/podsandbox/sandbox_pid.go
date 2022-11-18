/*
   Copyright The containerd Authors.

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

package podsandbox

import (
	"context"
	"fmt"

	"github.com/containerd/containerd/api/services/sandbox/v1"
)

func (c *Controller) PID(ctx context.Context, sandboxID string) (*sandbox.ControllerPIDResponse, error) {
	s, err := c.sandboxStore.Get(sandboxID)
	if err != nil {
		return nil, fmt.Errorf("an error occurred when try to find sandbox %q: %w",
			sandboxID, err)
	}
	return &sandbox.ControllerPIDResponse{Pid: s.Status.Get().Pid}, nil
}
