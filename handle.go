package main

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/joint-online-judge/go-horse/pkg/task"
	"github.com/joint-online-judge/go-tiger/pkg/json"
	"github.com/sirupsen/logrus"
)

func SubmitRecordTask(ctx context.Context, t *asynq.Task) error {
	var p task.SubmitRecordPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	logrus.Printf("SubmitRecordTask payload: %+v", p)
	return nil
}
