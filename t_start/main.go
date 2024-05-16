package t_start

import (
  "context"
  "github.com/zxlzhd/temporal_md_a/temporal"
  "log"
  "sync"

  "github.com/pborman/uuid"
  "go.temporal.io/sdk/client"
)

// @@@SNIPSTART samples-go-child-workflow-example-execution-starter
func MainS() {
  // The client is a heavyweight object that should be created only once per process.
  c, err := client.Dial(client.Options{
    HostPort: temporal.TemporalPath,
  })
  if err != nil {
    log.Fatalln("Unable to create client", err)
  }
  defer c.Close()
  var wg sync.WaitGroup
  wg.Add(2)
  go func() {
    defer wg.Done()
    // This Workflow ID can be a user supplied business logic identifier.
    workflowID := "parent-workflow_" + uuid.New()
    workflowOptions := client.StartWorkflowOptions{
      ID:           workflowID,
      TaskQueue:    temporal.TaskQueueName,
      CronSchedule: temporal.CronConfig,
    }

    workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, temporal.ParentWorkflow)
    if err != nil {
      log.Fatalln("Unable to execute workflow", err)
    }
    log.Println("Started workflow",
      "WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())

    // Synchronously wait for the Workflow Execution to complete.
    // Behind the scenes the SDK performs a long poll operation.
    // If you need to wait for the Workflow Execution to complete from another process use
    // Client.GetWorkflow API to get an instance of the WorkflowRun.
    /*var result string
      fmt.Println("waiting")
      err = workflowRun.Get(context.Background(), &result)
      if err != nil {
      	log.Fatalln("Failure getting workflow result", err)
      }
      log.Printf("Workflow result: %v", result)*/
  }()
  go func() {
    defer wg.Done()
    // This Workflow ID can be a user supplied business logic identifier.
    workflowIDB := "parent-b-workflow_" + uuid.New()
    workflowOptionsB := client.StartWorkflowOptions{
      ID:           workflowIDB,
      TaskQueue:    temporal.TaskQueueNameB,
      CronSchedule: temporal.CronConfigB,
    }

    workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptionsB, temporal.ParentWorkflowB)
    if err != nil {
      log.Fatalln("Unable to execute workflow", err)
    }
    log.Println("Started workflow",
      "WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())

    // Synchronously wait for the Workflow Execution to complete.
    // Behind the scenes the SDK performs a long poll operation.
    // If you need to wait for the Workflow Execution to complete from another process use
    // Client.GetWorkflow API to get an instance of the WorkflowRun.
    /*var result string
      fmt.Println("waiting b")
      err = workflowRun.Get(context.Background(), &result)
      if err != nil {
      	log.Fatalln("Failure getting workflow result", err)
      }
      log.Printf("Workflow result: %v", result)*/
  }()
  wg.Wait()
}

// @@@SNIPEND
