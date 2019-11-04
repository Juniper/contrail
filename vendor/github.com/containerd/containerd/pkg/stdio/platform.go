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

package stdio

import (
	"context"
	"sync"

	"github.com/containerd/console"
)

// Platform handles platform-specific behavior that may differs across
// platform implementations
type Platform interface {
	CopyConsole(ctx context.Context, console console.Console, stdin, stdout, stderr string,
		wg *sync.WaitGroup) (console.Console, error)
	ShutdownConsole(ctx context.Context, console console.Console) error
	Close() error
}
