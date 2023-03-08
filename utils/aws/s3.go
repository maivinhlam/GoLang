package external

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"strings"

	"terra-mapper.com/pkg/config"
	"terra-mapper.com/pkg/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetListFileFromBucket(mySession *session.Session, path string) (listItem []string, err error) {
	params := &s3.ListObjectsInput{
		Bucket: aws.String(config.Global.AwsS3Bucket),
		Prefix: aws.String(path),
	}

	resp, err := s3.New(mySession, aws.NewConfig()).ListObjects(params)
	if err != nil {
		return
	}

	for _, key := range resp.Contents {
		listItem = append(listItem, *key.Key)
	}

	return
}

func AwsPushDataToS3(ctx context.Context, mySession *session.Session, path, data, contentType string) error {
	utils.InfoLogger(ctx, "AwsPushDataToS3: "+path)
	_, err := s3.New(mySession).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(config.Global.AwsS3Bucket),
		Key:                  aws.String(path),
		Body:                 bytes.NewReader([]byte(data)),
		ContentType:          aws.String(contentType),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return err
}

func AwsMoveDataToAnotherFolder(ctx context.Context, mySession *session.Session, currentPath, targetPath string) error {
	bucket := config.Global.AwsS3Bucket

	svc := s3.New(mySession)

	var marker *string
	for {
		res, err := svc.ListObjects(
			&s3.ListObjectsInput{
				Bucket: aws.String(bucket),
				Marker: marker,
				Prefix: aws.String(currentPath),
			},
		)
		if err != nil {
			utils.WarnfLogger(ctx, "Failed to list objects: %s", err.Error())

			return nil
		}
		utils.InfoLogger(ctx, fmt.Sprint("AwsMoveDataToAnotherFolder: ", len(res.Contents), *res.IsTruncated, res.NextMarker))

		for _, obj := range res.Contents {
			srcKey := "/" + bucket + "/" + *obj.Key
			destKey := strings.Replace(*obj.Key, currentPath, "", 1)
			destKey = targetPath + destKey

			_, err = svc.CopyObject(
				&s3.CopyObjectInput{
					Bucket:     aws.String(bucket),
					CopySource: aws.String(url.QueryEscape(srcKey)),
					Key:        aws.String(destKey),
				},
			)
			fmt.Println("move "+srcKey, "to "+destKey)
			if err != nil {
				utils.WarnfLogger(ctx, "Failed to copy objects: %s", err.Error())

				continue
			}

			_, _ = svc.DeleteObject(
				&s3.DeleteObjectInput{
					Bucket: aws.String(bucket),
					Key:    obj.Key,
				},
			)
		}

		marker = res.NextMarker
		if marker == nil {
			break
		}
	}

	return nil
}

func AwsDeleteFolder(ctx context.Context, mySession *session.Session, path string) error {
	bucket := config.Global.AwsS3Bucket

	svc := s3.New(mySession)

	var marker *string
	for {
		res, err := svc.ListObjects(
			&s3.ListObjectsInput{
				Bucket: aws.String(bucket),
				Marker: marker,
				Prefix: aws.String(path),
			},
		)

		if err != nil {
			return err
		}

		fmt.Println(len(res.Contents), *res.IsTruncated, res.NextMarker)

		for _, obj := range res.Contents {
			_, err = svc.DeleteObject(
				&s3.DeleteObjectInput{
					Bucket: aws.String(bucket),
					Key:    obj.Key,
				},
			)

			if err != nil {
				utils.WarnfLogger(ctx, "Failed to delete objects "+*obj.Key+": %s", err.Error())
				continue
			}
		}

		marker = res.NextMarker
		if marker == nil {
			break
		}
	}

	return nil
}
