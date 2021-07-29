package utils

import (
	"context"
	"time"
)

type RetryContext struct {
	RetryCount int
	Errors     []error
	Success    bool
}

type Retry struct {
	MaxRetries    int
	RetryDelaySec int
	RetryFunc     func() error
}

func NewRetry(maxRetries int, retryDelaySec int, retryFunc func() error) *Retry {
	return &Retry{
		MaxRetries:    maxRetries,
		RetryDelaySec: retryDelaySec,
		RetryFunc:     retryFunc,
	}
}

func (r *Retry) ExecWithRetries(ctx context.Context) (*RetryContext, error) {
	retryCtx := RetryContext{}
	success := false
	for !success {
		retryCtx.RetryCount++
		if err := r.RetryFunc(); err != nil {
			retryCtx.Errors = append(retryCtx.Errors, err)

			if retryCtx.RetryCount == r.MaxRetries {
				return &retryCtx, err
			}

			w := make(chan struct{}, 1)
			go func() {
				time.Sleep(time.Duration(r.RetryDelaySec) * time.Second)
				w <- struct{}{}
			}()

			select {
			case <-ctx.Done():
				return &retryCtx, context.Canceled
			case <-w:
			}
		} else {
			success = true
		}
	}

	retryCtx.Success = true
	return &retryCtx, nil
}
