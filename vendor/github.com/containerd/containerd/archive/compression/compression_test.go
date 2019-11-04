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

package compression

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	// Force initPigz to be called, so tests start with the same initial state
	gzipDecompress(context.Background(), strings.NewReader(""))
	os.Exit(m.Run())
}

// generateData generates data that composed of 2 random parts
// and single zero-filled part within them.
// Typically, the compression ratio would be about 67%.
func generateData(t *testing.T, size int) []byte {
	part0 := size / 3             // random
	part2 := size / 3             // random
	part1 := size - part0 - part2 // zero-filled
	part0Data := make([]byte, part0)
	if _, err := rand.Read(part0Data); err != nil {
		t.Fatal(err)
	}
	part1Data := make([]byte, part1)
	part2Data := make([]byte, part2)
	if _, err := rand.Read(part2Data); err != nil {
		t.Fatal(err)
	}
	return append(part0Data, append(part1Data, part2Data...)...)
}

func testCompressDecompress(t *testing.T, size int, compression Compression) DecompressReadCloser {
	orig := generateData(t, size)
	var b bytes.Buffer
	compressor, err := CompressStream(&b, compression)
	if err != nil {
		t.Fatal(err)
	}
	if n, err := compressor.Write(orig); err != nil || n != size {
		t.Fatal(err)
	}
	compressor.Close()
	compressed := b.Bytes()
	t.Logf("compressed %d bytes to %d bytes (%.2f%%)",
		len(orig), len(compressed), 100.0*float32(len(compressed))/float32(len(orig)))
	if compared := bytes.Compare(orig, compressed); (compression == Uncompressed && compared != 0) ||
		(compression != Uncompressed && compared == 0) {
		t.Fatal("strange compressed data")
	}

	decompressor, err := DecompressStream(bytes.NewReader(compressed))
	if err != nil {
		t.Fatal(err)
	}
	decompressed, err := ioutil.ReadAll(decompressor)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(orig, decompressed) {
		t.Fatal("strange decompressed data")
	}

	return decompressor
}

func TestCompressDecompressGzip(t *testing.T) {
	oldUnpigzPath := unpigzPath
	unpigzPath = ""
	defer func() { unpigzPath = oldUnpigzPath }()

	decompressor := testCompressDecompress(t, 1024*1024, Gzip)
	wrapper := decompressor.(*readCloserWrapper)
	_, ok := wrapper.Reader.(*gzip.Reader)
	if !ok {
		t.Fatalf("unexpected compressor type: %T", wrapper.Reader)
	}
}

func TestCompressDecompressPigz(t *testing.T) {
	if _, err := exec.LookPath("unpigz"); err != nil {
		t.Skip("pigz not installed")
	}

	decompressor := testCompressDecompress(t, 1024*1024, Gzip)
	wrapper := decompressor.(*readCloserWrapper)
	_, ok := wrapper.Reader.(*io.PipeReader)
	if !ok {
		t.Fatalf("unexpected compressor type: %T", wrapper.Reader)
	}
}

func TestCompressDecompressUncompressed(t *testing.T) {
	testCompressDecompress(t, 1024*1024, Uncompressed)
}

func TestDetectPigz(t *testing.T) {
	// Create fake PATH with unpigz executable, make sure detectPigz can find it
	tempPath, err := ioutil.TempDir("", "containerd_temp_")
	if err != nil {
		t.Fatal(err)
	}

	filename := "unpigz"
	if runtime.GOOS == "windows" {
		filename = "unpigz.exe"
	}

	fullPath := filepath.Join(tempPath, filename)

	if err := ioutil.WriteFile(fullPath, []byte(""), 0111); err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(tempPath)

	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", tempPath)
	defer os.Setenv("PATH", oldPath)

	if pigzPath := detectPigz(); pigzPath == "" {
		t.Fatal("failed to detect pigz path")
	} else if pigzPath != fullPath {
		t.Fatalf("wrong pigz found: %s != %s", pigzPath, fullPath)
	}

	os.Setenv(disablePigzEnv, "1")
	defer os.Unsetenv(disablePigzEnv)

	if pigzPath := detectPigz(); pigzPath != "" {
		t.Fatalf("disable via %s doesn't work", disablePigzEnv)
	}
}

func TestCmdStream(t *testing.T) {
	out, err := cmdStream(exec.Command("sh", "-c", "echo hello; exit 0"), nil)
	if err != nil {
		t.Fatal(err)
	}

	buf, err := ioutil.ReadAll(out)
	if err != nil {
		t.Fatalf("failed to read from stdout: %s", err)
	}

	if string(buf) != "hello\n" {
		t.Fatalf("unexpected command output ('%s' != '%s')", string(buf), "hello\n")
	}
}

func TestCmdStreamBad(t *testing.T) {
	out, err := cmdStream(exec.Command("sh", "-c", "echo hello; echo >&2 bad result; exit 1"), nil)
	if err != nil {
		t.Fatalf("failed to start command: %v", err)
	}

	if buf, err := ioutil.ReadAll(out); err == nil {
		t.Fatal("command should have failed")
	} else if err.Error() != "exit status 1: bad result\n" {
		t.Fatalf("wrong error: %s", err.Error())
	} else if string(buf) != "hello\n" {
		t.Fatalf("wrong output: %s", string(buf))
	}
}
