// Copyright 2018 SEQSENSE, Inc.
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

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	decrypt = flag.Bool("decrypt", false, "decrypt using kms")
)

func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Fprintf(os.Stderr, "usage: s3cat [options] bucket s3path\noptions:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	bucket := flag.Arg(0)
	path := flag.Arg(1)

	sess := session.Must(session.NewSession())

	var out string

	obj := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	}

	dl := s3manager.NewDownloader(sess)
	buf := &aws.WriteAtBuffer{}
	_, err := dl.Download(buf, obj)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if *decrypt {
		svc := kms.New(sess)

		res, err := svc.Decrypt(&kms.DecryptInput{CiphertextBlob: buf.Bytes()})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		out = string(res.Plaintext)
	} else {
		out = string(buf.Bytes())
	}

	fmt.Print(out)
	os.Exit(0)
}
