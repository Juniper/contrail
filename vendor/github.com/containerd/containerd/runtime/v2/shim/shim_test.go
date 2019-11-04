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

package shim

import (
	"context"
	"testing"
)

func TestShimOptWithValue(t *testing.T) {
	ctx := context.TODO()
	ctx = context.WithValue(ctx, OptsKey{}, Opts{Debug: true})

	o := ctx.Value(OptsKey{})
	if o == nil {
		t.Fatal("opts nil")
	}
	op, ok := o.(Opts)
	if !ok {
		t.Fatal("opts not of type Opts")
	}
	if !op.Debug {
		t.Fatal("opts.Debug should be true")
	}
}
