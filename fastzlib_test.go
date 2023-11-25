package czlib

import (
	"bytes"
	"compress/zlib"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"testing"
)

var deflated []byte

func zip(b []byte) []byte {
	var out bytes.Buffer
	w := zlib.NewWriter(&out)
	w.Write(b)
	w.Close()
	return out.Bytes()
}

var raw []byte

func getRaw() []byte {
	if len(raw) != 0 {
		return raw
	}

	var (
		err error
	)

	payload := os.Getenv("PAYLOAD")
	if len(payload) == 0 {
		fmt.Println("provide PAYLOAD env var for path to test custom payload.")
		tmpFile, err := os.CreateTemp("", "payload")
		if err != nil {
			fmt.Printf("Error creating temp file: %s\n", err)
			panic(err)
		}

		buffer := make([]byte, 1021*17)
		rand.Read(buffer)
		tmpFile.Write(buffer)
		tmpFile.Close()
		os.Setenv("PAYLOAD", tmpFile.Name())
		payload = tmpFile.Name()
	}

	raw, err = os.ReadFile(payload)
	if err != nil {
		fmt.Printf("Error opening payload: %s\n", err)
	}

	if len(raw) == 0 {
		panic("payload is empty")
	}

	return raw
}

func init() {
	deflated = zip(getRaw())
	// fmt.Printf("%d byte test payload (%d orig)\n", len(deflated), len(raw))
}

// Generate an n-byte long []byte
func genData(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}

func TestAllZlib(t *testing.T) {
	type compressFunc func([]byte) ([]byte, error)
	funcs := []compressFunc{Compress, gzip, zzip}
	names := []string{"Compress", "gzip", "zzip"}
	for _, i := range []int{10, 128, 1000, 1024 * 10, 1024 * 100, 1024 * 1024, 1024 * 1024 * 7} {
		data, err := genData(i)
		if err != nil {
			t.Error(err)
			continue
		}
		for i, f := range funcs {
			comp, err := f(data)
			if err != nil {
				t.Fatalf("Compression failed on %v: %s", names[i], err)
			}
			decomp, err := Decompress(comp)
			if err != nil {
				t.Fatalf("Decompression failed on %v: %s", names[i], err)
			}
			if bytes.Compare(decomp, data) != 0 {
				t.Fatalf("deflate->inflate does not match original for %s", names[i])
			}
		}
	}
}

func TestEmpty(t *testing.T) {
	var empty []byte
	_, err := Compress(empty)
	if err != nil {
		t.Fatalf("unexpected error compressing empty slice")
	}
	_, err = Decompress(empty)
	if err == nil {
		t.Fatalf("unexpected success decompressing empty slice")
	}
}

func TestUnsafeZlib(t *testing.T) {
	for _, i := range []int{10, 128, 1000, 1024 * 10, 1024 * 100, 1024 * 1024, 1024 * 1024 * 7} {
		data, err := genData(i)
		if err != nil {
			t.Error(err)
			continue
		}
		comp, err := UnsafeCompress(data)
		if err != nil {
			t.Fatal(err)
		}
		decomp, err := UnsafeDecompress(comp)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(decomp, data) {
			t.Fatal("Compress -> Decompress on byte array failed to match original data.")
		}
		comp.Free()
		decomp.Free()
	}

	for _, i := range []int{10, 128, 1000, 1024 * 10, 1024 * 100, 1024 * 1024, 1024 * 1024 * 7} {
		data, err := genData(i)
		if err != nil {
			t.Error(err)
			continue
		}
		comp, err := UnsafeCompress2(data, GZipWbits)
		if err != nil {
			t.Fatal(err)
		}
		decomp, err := UnsafeDecompress(comp)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(decomp, data) {
			t.Fatal("Compress -> Decompress on byte array failed to match original data.")
		}
		comp.Free()
		decomp.Free()
	}
}

// Compression benchmarks
func BenchmarkCompressUnsafe(b *testing.B) {
	if raw == nil {
		b.Skip("You must provide PAYLOAD env var for benchmarking.")
	}
	b.SetBytes(int64(len(raw)))
	for i := 0; i < b.N; i++ {
		u, _ := UnsafeCompress(raw)
		u.Free()
	}
}

func BenchmarkCompress(b *testing.B) {
	if raw == nil {
		b.Skip("You must provide PAYLOAD env var for benchmarking.")
	}
	b.SetBytes(int64(len(raw)))
	for i := 0; i < b.N; i++ {
		Compress(raw)
	}
}

func BenchmarkCompressStream(b *testing.B) {
	if raw == nil {
		b.Skip("You must provide PAYLOAD env var for benchmarking.")
	}
	b.SetBytes(int64(len(raw)))
	for i := 0; i < b.N; i++ {
		gzip(raw)
	}
}

func BenchmarkCompressStdZlib(b *testing.B) {
	if raw == nil {
		b.Skip("You must provide PAYLOAD env var for benchmarking.")
	}
	b.SetBytes(int64(len(raw)))
	for i := 0; i < b.N; i++ {
		zzip(raw)
	}
}

// Decomression benchmarks

func BenchmarkDecompressUnsafe(b *testing.B) {
	if raw == nil {
		b.Skip("You must provide PAYLOAD env var for benchmarking.")
	}
	b.SetBytes(int64(len(raw)))
	for i := 0; i < b.N; i++ {
		u, _ := UnsafeDecompress(deflated)
		u.Free()
	}
}

func BenchmarkDecompress(b *testing.B) {
	if raw == nil {
		b.Skip("You must provide PAYLOAD env var for benchmarking.")
	}
	b.SetBytes(int64(len(raw)))
	for i := 0; i < b.N; i++ {
		Decompress(deflated)
	}
}

func BenchmarkDecompressStream(b *testing.B) {
	if raw == nil {
		b.Skip("You must provide PAYLOAD env var for benchmarking.")
	}
	b.SetBytes(int64(len(raw)))
	for i := 0; i < b.N; i++ {
		gunzip(deflated)
	}
}

func BenchmarkDecompressStdZlib(b *testing.B) {
	if raw == nil {
		b.Skip("You must provide PAYLOAD env var for benchmarking.")
	}
	b.SetBytes(int64(len(raw)))
	for i := 0; i < b.N; i++ {
		zunzip(deflated)
	}
}

// helpers

func gunzip(body []byte) ([]byte, error) {
	reader, err := NewReader(bytes.NewBuffer(body))
	if err != nil {
		return []byte{}, err
	}
	return io.ReadAll(reader)
}

// unzip a deflated []byte payload, returning the unzipped []byte and error
func zunzip(body []byte) ([]byte, error) {
	reader, err := zlib.NewReader(bytes.NewBuffer(body))
	if err != nil {
		return []byte{}, err
	}
	return io.ReadAll(reader)
}

func gzip(body []byte) ([]byte, error) {
	outb := make([]byte, 0, 16*1024)
	out := bytes.NewBuffer(outb)
	writer := NewWriter(out)
	n, err := writer.Write(body)
	if n != len(body) {
		return []byte{}, fmt.Errorf("compressed %d, expected %d", n, len(body))
	}
	if err != nil {
		return []byte{}, err
	}
	err = writer.Close()
	if err != nil {
		return []byte{}, err
	}
	return out.Bytes(), nil
}

func zzip(body []byte) ([]byte, error) {
	outb := make([]byte, 0, len(body))
	out := bytes.NewBuffer(outb)
	writer := zlib.NewWriter(out)
	n, err := writer.Write(body)
	if n != len(body) {
		return []byte{}, fmt.Errorf("compressed %d, expected %d", n, len(body))
	}
	if err != nil {
		return []byte{}, err
	}
	err = writer.Close()
	if err != nil {
		return []byte{}, err
	}
	return out.Bytes(), nil
}
