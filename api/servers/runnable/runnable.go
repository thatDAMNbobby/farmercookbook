package runnable

import "context"

type Runnable interface {
	Start(ctx context.Context)
	Stop(ctx context.Context)
}
