package s3

import (
	"bytes"
	"context"
	"data-del/configs"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Client struct {
	config *configs.Server
	client *s3.Client
}
//NewS3Client create new s3 clients
func NewS3Client(ctx context.Context, cfg *configs.Server) (IS3Handler, error) {
	conf, err := config.LoadDefaultConfig(ctx, config.WithRegion(cfg.S3Region))
	if err != nil {
		return nil, err
	}

	return &s3Client{
		config: cfg,
		client: s3.NewFromConfig(conf),
	}, nil
}

// GetObject get content file aws s3
func (s *s3Client) GetObject(ctx context.Context, path string) (*bytes.Buffer, error) {
	rawObject, err := s.client.GetObject(
		ctx, &s3.GetObjectInput{
			Bucket: &s.config.S3Bucket,
			Key:    &path,
		})

	buf := new(bytes.Buffer)

	if err != nil {
		return buf, err
	}

	_, err = buf.ReadFrom(rawObject.Body)

	if err != nil {
		return buf, err
	}

	return buf, nil
}

// PutObject put new object to S3
func (s *s3Client) PutObject(ctx context.Context, path string, body bytes.Buffer) error {
	input := &s3.PutObjectInput{
		Body:   bytes.NewReader(body.Bytes()),
		Bucket: aws.String(s.config.S3Bucket),
		Key:    aws.String(path),
	}
	_, err := s.client.PutObject(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
//CheckObjectExists check object existed
func (s *s3Client) CheckObjectExists(ctx context.Context, path string) (bool, error) {
	input := &s3.ListObjectsInput{
		Bucket: &s.config.S3Bucket,
		Prefix: &path,
	}
	objects, err := s.client.ListObjects(ctx, input)
	if err != nil {
		return false, err
	}
	if len(objects.Contents) > 1 {
		return true, nil
	}

	return false, err
}