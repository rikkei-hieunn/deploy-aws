package s3

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"start-ecs/configs"
)

type s3Storage struct {
	config     *configs.TickSystem
	svc        *s3.S3
	downloader *s3manager.Downloader
	uploader   *s3manager.Uploader
}

// NewS3Storage S3 storage constructor
func NewS3Storage(cfg *configs.TickSystem) (IS3Handler, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	sconf := &aws.Config{Region: aws.String(cfg.S3Region)}
	sess, err := session.NewSession(sconf)
	if err != nil {
		return nil, err
	}

	downloader := s3manager.NewDownloader(sess, func(d *s3manager.Downloader) {
		//Concurrency of 1 will download the parts sequentially.
		d.Concurrency = 1
	})

	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024 // minimum/default part size: 5MB
		u.Concurrency = 5            // default 5
		u.LeavePartsOnError = false
	})

	return &s3Storage{
		config:     cfg,
		svc:        s3.New(sess),
		downloader: downloader,
		uploader:   uploader,
	}, nil
}

// GetObject get object from S3
func (s *s3Storage) GetObject(path string) ([]byte, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.config.S3Bucket),
		Key:    aws.String(path),
	}

	size, err := s.getS3ObjectSize(path)
	if err != nil {
		return nil, err
	}

	bytesSlice := make([]byte, size)
	buff := aws.NewWriteAtBuffer(bytesSlice)
	_, err = s.downloader.Download(buff, input)
	if err != nil {
		return nil, err
	}

	return bytesSlice, err
}

// getS3ObjectSize get size of s3 file object
func (s *s3Storage) getS3ObjectSize(path string) (int64, error) {
	input := &s3.HeadObjectInput{
		Bucket: aws.String(s.config.S3Bucket),
		Key:    aws.String(path),
	}

	result, err := s.svc.HeadObject(input)
	if err != nil {
		return 0, err
	}

	return *result.ContentLength, nil
}

// PutObject put new object to S3
func (s *s3Storage) PutObject(path string, body bytes.Buffer) error {
	input := &s3manager.UploadInput{
		Body:   bytes.NewReader(body.Bytes()),
		Bucket: aws.String(s.config.S3Bucket),
		Key:    aws.String(path),
	}
	_, err := s.uploader.Upload(input)
	if err != nil {
		return err
	}

	return nil
}
