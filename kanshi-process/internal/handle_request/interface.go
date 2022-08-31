// Package handlerequest for define service to handle request
package handlerequest

import "context"

// IWorker worker process send request to each port
type IWorker interface {
	Process(ctx context.Context, host string, port int)
}

// IRequestHandler management request check status toiawase and furiwake server
type IRequestHandler interface {
	StartSendRequest(ctx context.Context) error
}
