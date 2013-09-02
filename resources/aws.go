//
// Copyright (c) 2013 Jake Brukhman/East River Labs. All rights reserved.
//
package resources

import (
	"bytes"
	"fmt"
	"github.com/kr/s3/s3util"
	"io"
	"os"
)

// environment keys for credentials, bucket
const (
	SecretKey    = "S3_SECRET_ACCESS_KEY"
	AccessKey    = "S3_ACCESS_KEY_ID"
	ObfBucketKey = "S3_BUCKET_NAME"
)

const (
	// format string for data files, where
	// the first placeholder is bucket name,
	// followed by resourceId
	ObfUrlFmt = "https://%s.s3.amazonaws.com/recordings/%s"
)

// The name of the Octopus data bucket
// in s3, which varies depending on staging/prod.
var ObfBucket string

func init() {
	ObfBucket = os.Getenv(ObfBucketKey)
	if ObfBucket == "" {
		panic("cannot read bucket")
	}

	var (
		sk = os.Getenv(SecretKey)
		ak = os.Getenv(AccessKey)
	)

	if sk == "" || ak == "" {
		panic("could not read AWS credentials from environment")
	}

	s3util.DefaultConfig.SecretKey = sk
	s3util.DefaultConfig.AccessKey = ak
}

// resourceUrl will return a fully-formed resourceId url
func resourceUrl(bucket, resourceId string) string {
	return fmt.Sprintf(ObfUrlFmt, bucket, resourceId)
}

func GetUrl(url string) (resource io.Reader, err error) {
	r, err := s3util.Open(url, nil)
	if err != nil {
		return
	}

	// now, we have the file;
	// let's read it into memory
	buf := new(bytes.Buffer)
	if _, err = io.Copy(buf, r); err != nil {
		return
	}

	// close the remote resource
	r.Close()
	return buf, nil
}

// Resource will read the S3 resource into memory and
// return the reader.
func Resource(resourceId string) (resource io.Reader, err error) {
	url := resourceUrl(ObfBucket, resourceId)
	return GetUrl(url)
}
