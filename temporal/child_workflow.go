package temporal

import (
	"fmt"
	"go.temporal.io/sdk/workflow"
	"time"
)

// @@@SNIPSTART samples-go-child-workflow-example-child-workflow-definition
// SampleChildWorkflow is a Workflow Definition
func SampleChildWorkflow(ctx workflow.Context, name string, i int) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	logger := workflow.GetLogger(ctx)
	greeting := "Hello " + name + "!+num:%v"
	greeting = fmt.Sprintf(greeting, i)
	logger.Info("Child workflow execution: " + greeting)
	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	workflow.ExecuteActivity(ctx, WriteMessage, greeting).Get(ctx, &result)
	return greeting, nil
}
func SampleChildWorkflowB(ctx workflow.Context, name string, i int) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	logger := workflow.GetLogger(ctx)
	greeting := "Hello " + name + "!+num:%v"
	greeting = fmt.Sprintf(greeting, i)
	logger.Info("Child workflow execution: " + greeting)
	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	workflow.ExecuteActivity(ctx, ReadMessage, greeting).Get(ctx, &result)
	return greeting, nil
}
