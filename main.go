package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//var bucket string
//var region string
//var expireTime time.Duration

func generatePresignedUrl(svc *s3.S3, bucket string, key string, expireTime time.Duration) (string, error) {
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	return req.Presign(expireTime)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: s3pre KEY_URL(format: s3://bucket/key) -r=REGION -e=EXPIRE_TIME")
		os.Exit(1)
	}

	url := os.Args[1]

	var region string
	var expire string
	flagSet := flag.NewFlagSet("", flag.ExitOnError)
	flagSet.StringVar(&region, "r", "ap-northeast-2", "Region")
	flagSet.StringVar(&expire, "e", "15m", "Expire time in minutes (default 15m)")
	flagSet.Parse(os.Args[2:])

	var err error
	log.Println("region: " + region)
	log.Println("expireTime: " + expire)
	expireTime, err := time.ParseDuration(expire)
	if err != nil {
		log.Println("Wrong value input for Expire Time")
		os.Exit(1)
	}

	part1 := strings.Split(url, "s3://")
	part2 := strings.Split(part1[1], "/")
	bucket := part2[0]
	key := strings.Join(part2[1:len(part2)], "/")
	log.Println("Bucket:" + bucket)
	log.Println("Key: " + key)

	awsConfig := &aws.Config{
		Region: aws.String(region),
	}

	sess := session.Must(session.NewSession(awsConfig))
	svc := s3.New(sess)

	urlStr, err := generatePresignedUrl(svc, bucket, key, expireTime)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(urlStr)
}
