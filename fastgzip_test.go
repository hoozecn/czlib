package czlib

import (
	"bytes"
	zlib "compress/gzip"
	"crypto/rand"
	"fmt"
	"io"
	"testing"
)

var gzipped []byte

func init() {
	var err error
	gzipped, err = zzip2(getRaw())
	if err != nil {
		panic(err)
	}
}

// Compression BenchmarkGzips
func BenchmarkGzipCompressUnsafe(b *testing.B) {
	if raw == nil {
		b.Skip("You must provide PAYLOAD env var for BenchmarkGziping.")
	}
	b.SetBytes(int64(len(raw)))
	for i := 0; i < b.N; i++ {
		u, _ := UnsafeCompress2(raw, GZipWbits)
		u.Free()
	}
}

func BenchmarkGzipCompress(b *testing.B) {
	if raw == nil {
		b.Skip("You must provide PAYLOAD env var for BenchmarkGziping.")
	}
	b.SetBytes(int64(len(raw)))
	for i := 0; i < b.N; i++ {
		Compress2(raw, GZipWbits)
	}
}

func BenchmarkGzipCompressStream(b *testing.B) {
	if raw == nil {
		b.Skip("You must provide PAYLOAD env var for BenchmarkGziping.")
	}
	b.SetBytes(int64(len(raw)))
	for i := 0; i < b.N; i++ {
		gzip2(raw)
	}
}

func BenchmarkGzipCompressStdZlib(b *testing.B) {
	if raw == nil {
		b.Skip("You must provide PAYLOAD env var for BenchmarkGziping.")
	}
	b.SetBytes(int64(len(raw)))
	for i := 0; i < b.N; i++ {
		zzip2(raw)
	}
}

func TestDecompressGzip(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	w := zlib.NewWriter(buffer)
	origin := make([]byte, 1024*1024)
	rand.Read(origin)
	w.Write(origin)
	w.Close()

	decomp, err := Decompress(buffer.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(decomp, origin) {
		t.Fatal("Decompress on byte array failed to match original data.")
	}

	unsafeDecomp, err := UnsafeDecompress(buffer.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	defer unsafeDecomp.Free()

	if !bytes.Equal(unsafeDecomp, origin) {
		t.Fatal("Decompress on byte array failed to match original data.")
	}
}

// Decomression BenchmarkGzips

func BenchmarkGzipDecompressUnsafe(b *testing.B) {
	if raw == nil {
		b.Skip("You must provide PAYLOAD env var for BenchmarkGziping.")
	}
	b.SetBytes(int64(len(raw)))
	for i := 0; i < b.N; i++ {
		u, _ := UnsafeDecompress(gzipped)
		u.Free()
	}
}

func BenchmarkGzipDecompress(b *testing.B) {
	if raw == nil {
		b.Skip("You must provide PAYLOAD env var for BenchmarkGziping.")
	}
	b.SetBytes(int64(len(raw)))
	for i := 0; i < b.N; i++ {
		Decompress(gzipped)
	}
}

func BenchmarkGzipDecompressStream(b *testing.B) {
	if raw == nil {
		b.Skip("You must provide PAYLOAD env var for BenchmarkGziping.")
	}
	b.SetBytes(int64(len(raw)))
	for i := 0; i < b.N; i++ {
		gunzip2(gzipped)
	}
}

func BenchmarkGzipDecompressStdZlib(b *testing.B) {
	if raw == nil {
		b.Skip("You must provide PAYLOAD env var for BenchmarkGziping.")
	}
	b.SetBytes(int64(len(raw)))
	for i := 0; i < b.N; i++ {
		_, err := zunzip2(gzipped)
		if err != nil {
			b.Fatalf("zunzip2 failed: %v", err)
		}
	}
}

// helpers

func gunzip2(body []byte) ([]byte, error) {
	reader, err := NewReader(bytes.NewBuffer(body))
	if err != nil {
		return []byte{}, err
	}
	return io.ReadAll(reader)
}

// unzip a gzipped []byte payload, returning the unzipped []byte and error
func zunzip2(body []byte) ([]byte, error) {
	reader, err := zlib.NewReader(bytes.NewBuffer(body))
	if err != nil {
		return []byte{}, err
	}
	return io.ReadAll(reader)
}

func gzip2(body []byte) ([]byte, error) {
	outb := make([]byte, 0, 16*1024)
	out := bytes.NewBuffer(outb)
	writer := NewWriterLevelWbits(out, DefaultCompression, GZipWbits)
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

func zzip2(body []byte) ([]byte, error) {
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
