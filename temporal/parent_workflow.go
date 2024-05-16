package temporal

import (
	"fmt"
	"go.temporal.io/sdk/workflow"
)

// @@@SNIPSTART samples-go-child-workflow-example-parent-workflow-definition
// SampleParentWorkflow is a Workflow Definition
// This Workflow Definition demonstrates how to start a Child Workflow Execution from a Parent Workflow Execution.
// Each Child Workflow Execution starts a new Run.
// The Parent Workflow Execution is notified only after the completion of last Run of the Child Workflow Execution.
func ParentWorkflow(ctx workflow.Context) (string, error) {
	logger := workflow.GetLogger(ctx)

	var result string
	for i := 0; i < 10; i++ {
		cwo := workflow.ChildWorkflowOptions{
			WorkflowID: fmt.Sprintf("ABC-SIMPLE-CHILD-WORKFLOW-ID-%v", i),
		}
		ctxChild := workflow.WithChildOptions(ctx, cwo)
		err := workflow.ExecuteChildWorkflow(ctxChild, SampleChildWorkflow, "World", i).Get(ctx, &result)
		if err != nil {
			logger.Error("Parent execution received child execution failure.", "Error", err)
			return "", err
		}

		logger.Info("Parent execution completed.", "Result", result)
	}
	return result, nil
}
func ParentWorkflowB(ctx workflow.Context) (string, error) {
	logger := workflow.GetLogger(ctx)

	var result string
	for i := 0; i < 10; i++ {
		cwo := workflow.ChildWorkflowOptions{
			WorkflowID: fmt.Sprintf("ABC-B-SIMPLE-CHILD-WORKFLOW-ID-%v", i),
		}
		ctxChild := workflow.WithChildOptions(ctx, cwo)
		err := workflow.ExecuteChildWorkflow(ctxChild, SampleChildWorkflowB, "WorldB", i).Get(ctx, &result)
		if err != nil {
			logger.Error("Parent execution received child execution failure.", "Error", err)
			return "", err
		}

		logger.Info("Parent execution completed.", "Result", result)
	}
	return result, nil
}

// @@@SNIPEND
