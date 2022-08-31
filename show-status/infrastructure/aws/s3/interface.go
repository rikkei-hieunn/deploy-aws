/*
Package s3 implements logics about S3.
*/
package s3

// IS3Handler interface S3 handler
type IS3Handler interface {
	GetObject(path string) ([]byte, error)
	PutObject(path string, body []byte) error
}
