/*
Package summarizelog implements logics for handle log
*/
package summarizelog

import "context"

// ILogSummarizer interface start log summarize
type ILogSummarizer interface {
	Start(ctx context.Context) error
}

// IWorker defines worker interfaces
type IWorker interface {
	Start(ctx context.Context) []error
}
