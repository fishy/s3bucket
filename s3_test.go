package s3bucket_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/fishy/s3bucket"
)

func Example() {
	ctx := context.Background()
	region := "us-east-1"
	bkt := s3bucket.Open("test-bucket", region)

	obj := "test/object"
	content := `Lorem ipsum dolor sit amet,
consectetur adipiscing elit,
sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.

Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.

Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.

Excepteur sint occaecat cupidatat non proident,
sunt in culpa qui officia deserunt mollit anim id est laborum.`

	_, err := bkt.Read(ctx, obj)
	fmt.Println(bkt.IsNotExist(err)) // true

	if err := bkt.Write(ctx, obj, strings.NewReader(content)); err != nil {
		log.Fatal(err)
	}
	reader, err := bkt.Read(ctx, obj)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf)) // content
	if err := bkt.Delete(ctx, obj); err != nil {
		log.Fatal(err)
	}
}

func TestEmpty(t *testing.T) {
	// Empty test to silence the no test warning of go test
}
