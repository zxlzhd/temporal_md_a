package t_start

import (
	"context"
	"github.com/zxlzhd/temporal_md_a/temporal"
	"go.temporal.io/sdk/client"
	"log"
)

// @@@SNIPSTART samples-go-child-workflow-example-execution-starter
func StartRun() error {
	// The client is a heavyweight object that should be created only once per process.
	c, err := client.Dial(client.Options{
		HostPort: temporal.TemporalPath,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	err = startWriteFlow(c)
	if err != nil {
		log.Fatalln("Unable to create startwriteflow", err)
	}
	err = startReadFlow(c)
	if err != nil {
		log.Fatalln("Unable to create startreadflow", err)
	}
	return err
}
func startWriteFlow(c client.Client) error {
	workflowID := "parent-workflow_" + "write_1"
	workflowOptions := client.StartWorkflowOptions{
		ID:           workflowID,
		TaskQueue:    temporal.TaskQueueName,
		CronSchedule: temporal.CronConfig,
	}

	workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, temporal.ParentWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
		return err
	}
	log.Println("Started workflow",
		"WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())
	return nil
}
func startReadFlow(c client.Client) error {
	workflowIDB := "parent-b-workflow_" + "read_1" //uuid.New()
	workflowOptionsB := client.StartWorkflowOptions{
		ID:           workflowIDB,
		TaskQueue:    temporal.TaskQueueNameB,
		CronSchedule: temporal.CronConfigB,
	}

	workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptionsB, temporal.ParentWorkflowB)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
		return err
	}
	log.Println("Started workflow",
		"WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())
	return nil
}

// @@@SNIPEND
