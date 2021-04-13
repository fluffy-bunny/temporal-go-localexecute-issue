package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	var result string
	err := workflow.ExecuteActivity(ctx, MyBeautifulActivities.ComposeGreeting, name).Get(ctx, &result)
	if err != nil {
		return "", err
	}
	localOptions := workflow.LocalActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithLocalActivityOptions(ctx, localOptions)
	err = workflow.ExecuteLocalActivity(ctx, MyBeautifulActivities.ComposeGreeting, name).Get(ctx, &result)
	return result, err
}
