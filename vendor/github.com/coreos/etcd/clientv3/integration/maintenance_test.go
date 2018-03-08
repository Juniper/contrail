// Copyright 2017 The etcd Authors
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

package integration

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/integration"
	"github.com/coreos/etcd/lease"
	"github.com/coreos/etcd/mvcc"
	"github.com/coreos/etcd/mvcc/backend"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestMaintenanceHashKV(t *testing.T) {
	defer testutil.AfterTest(t)

	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)

	for i := 0; i < 3; i++ {
		if _, err := clus.RandClient().Put(context.Background(), "foo", "bar"); err != nil {
			t.Fatal(err)
		}
	}

	var hv uint32
	for i := 0; i < 3; i++ {
		cli := clus.Client(i)
		// ensure writes are replicated
		if _, err := cli.Get(context.TODO(), "foo"); err != nil {
			t.Fatal(err)
		}
		hresp, err := cli.HashKV(context.Background(), clus.Members[i].GRPCAddr(), 0)
		if err != nil {
			t.Fatal(err)
		}
		if hv == 0 {
			hv = hresp.Hash
			continue
		}
		if hv != hresp.Hash {
			t.Fatalf("#%d: hash expected %d, got %d", i, hv, hresp.Hash)
		}
	}
}

func TestMaintenanceMoveLeader(t *testing.T) {
	defer testutil.AfterTest(t)

	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)

	oldLeadIdx := clus.WaitLeader(t)
	targetIdx := (oldLeadIdx + 1) % 3
	target := uint64(clus.Members[targetIdx].ID())

	cli := clus.Client(targetIdx)
	_, err := cli.MoveLeader(context.Background(), target)
	if err != rpctypes.ErrNotLeader {
		t.Fatalf("error expected %v, got %v", rpctypes.ErrNotLeader, err)
	}

	cli = clus.Client(oldLeadIdx)
	_, err = cli.MoveLeader(context.Background(), target)
	if err != nil {
		t.Fatal(err)
	}

	leadIdx := clus.WaitLeader(t)
	lead := uint64(clus.Members[leadIdx].ID())
	if target != lead {
		t.Fatalf("new leader expected %d, got %d", target, lead)
	}
}

// TestMaintenanceSnapshotError ensures that context cancel/timeout
// before snapshot reading returns corresponding context errors.
func TestMaintenanceSnapshotError(t *testing.T) {
	defer testutil.AfterTest(t)

	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)

	// reading snapshot with canceled context should error out
	ctx, cancel := context.WithCancel(context.Background())
	rc1, err := clus.RandClient().Snapshot(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer rc1.Close()

	cancel()
	_, err = io.Copy(ioutil.Discard, rc1)
	if err != context.Canceled {
		t.Errorf("expected %v, got %v", context.Canceled, err)
	}

	// reading snapshot with deadline exceeded should error out
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rc2, err := clus.RandClient().Snapshot(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer rc2.Close()

	time.Sleep(2 * time.Second)

	_, err = io.Copy(ioutil.Discard, rc2)
	if err != nil && err != context.DeadlineExceeded {
		t.Errorf("expected %v, got %v", context.DeadlineExceeded, err)
	}
}

// TestMaintenanceSnapshotErrorInflight ensures that inflight context cancel/timeout
// fails snapshot reading with corresponding context errors.
func TestMaintenanceSnapshotErrorInflight(t *testing.T) {
	defer testutil.AfterTest(t)

	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)

	// take about 1-second to read snapshot
	clus.Members[0].Stop(t)
	dpath := filepath.Join(clus.Members[0].DataDir, "member", "snap", "db")
	b := backend.NewDefaultBackend(dpath)
	s := mvcc.NewStore(b, &lease.FakeLessor{}, nil)
	rev := 100000
	for i := 2; i <= rev; i++ {
		s.Put([]byte(fmt.Sprintf("%10d", i)), bytes.Repeat([]byte("a"), 1024), lease.NoLease)
	}
	s.Close()
	b.Close()
	clus.Members[0].Restart(t)

	// reading snapshot with canceled context should error out
	ctx, cancel := context.WithCancel(context.Background())
	rc1, err := clus.RandClient().Snapshot(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer rc1.Close()

	donec := make(chan struct{})
	go func() {
		time.Sleep(300 * time.Millisecond)
		cancel()
		close(donec)
	}()
	_, err = io.Copy(ioutil.Discard, rc1)
	if err != nil && err != context.Canceled {
		t.Errorf("expected %v, got %v", context.Canceled, err)
	}
	<-donec

	// reading snapshot with deadline exceeded should error out
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rc2, err := clus.RandClient().Snapshot(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer rc2.Close()

	// 300ms left and expect timeout while snapshot reading is in progress
	time.Sleep(700 * time.Millisecond)
	_, err = io.Copy(ioutil.Discard, rc2)
	if err != nil && err != context.DeadlineExceeded {
		t.Errorf("expected %v, got %v", context.DeadlineExceeded, err)
	}
}
