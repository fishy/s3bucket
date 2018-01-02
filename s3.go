// Package s3bucket provides an implementation of Bucket for AWS S3.
package s3bucket

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/fishy/fsdb/bucket"
)

// Make sure *S3Bucket satisifies bucket.Bucket interface.
var _ bucket.Bucket = (*S3Bucket)(nil)

// S3Bucket is an implementation of bucket with S3.
type S3Bucket struct {
	svc        *s3.S3
	bucket     *string
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

// Open opens an s3 bucket.
//
// The bucket must already exist, otherwise all operations will fail.
func Open(bucket string, region string) *S3Bucket {
	svc := s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	})))
	return &S3Bucket{
		svc:        svc,
		bucket:     aws.String(bucket),
		uploader:   s3manager.NewUploaderWithClient(svc),
		downloader: s3manager.NewDownloaderWithClient(svc),
	}
}

func (bkt *S3Bucket) Read(ctx context.Context, name string) (
	io.ReadCloser,
	error,
) {
	buf := new(aws.WriteAtBuffer)
	_, err := bkt.downloader.DownloadWithContext(
		ctx,
		buf,
		&s3.GetObjectInput{
			Bucket: bkt.bucket,
			Key:    aws.String(name),
		},
	)
	if err != nil {
		return nil, err
	}
	return ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}

func (bkt *S3Bucket) Write(
	ctx context.Context,
	name string,
	data io.Reader,
) error {
	_, err := bkt.uploader.UploadWithContext(
		ctx,
		&s3manager.UploadInput{
			Bucket: bkt.bucket,
			Key:    aws.String(name),
			Body:   data,
		},
	)
	return err
}

// Delete deletes an object from the bucket.
//
// Please note that S3 will NOT return error when deleting a non-exist object.
func (bkt *S3Bucket) Delete(ctx context.Context, name string) error {
	_, err := bkt.svc.DeleteObjectWithContext(
		ctx,
		&s3.DeleteObjectInput{
			Bucket: bkt.bucket,
			Key:    aws.String(name),
		},
	)
	return err
}

// IsNotExist checks whether err is an awserr.Error with code ErrCodeNoSuchKey.
func (bkt *S3Bucket) IsNotExist(err error) bool {
	if aerr, ok := err.(awserr.Error); ok {
		return aerr.Code() == s3.ErrCodeNoSuchKey
	}
	return false
}
