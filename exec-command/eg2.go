package main
/* 写入标准输入 */
// golang中没有bzip2的压缩库，所以，可以通过把要压缩的文件传给命令去执行
import (
	"bytes"
	"compress/bzip2"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
)

// compress data using bzip2 without creating temporary files
func bzipCompress(d []byte) ([]byte, error) {
	var out bytes.Buffer
	cmd := exec.Command("bzip2", "-c", "-9")
	cmd.Stdin = bytes.NewBuffer(d)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func bzipDecompress(d []byte) ([]byte, error) {
	r := bzip2.NewReader(bytes.NewBuffer(d))
	var out bytes.Buffer
	_, err := io.Copy(&out, r)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func main() {
	path := "eg2.go"
	d, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("ioutil.ReadFile() failed with %s\n", err)
	}
	compressed, err := bzipCompress(d)
	if err != nil {
		log.Fatalf("bzipCompress() failed with %s\n", err)
	}
	fmt.Printf("Compressed %d bytes to %d, saving %d\n", len(d), len(compressed), len(d)-len(compressed))

	// verify compression was correct by decompressing and comparing
	// with the original
	uncompressed, err := bzipDecompress(compressed)
	if err != nil {
		log.Fatalf("bzipDecompress() failed with %s\n", err)
	}
	if !bytes.Equal(d, uncompressed) {
		log.Fatalf("decompressed data not the same as original data\n")
	}
}