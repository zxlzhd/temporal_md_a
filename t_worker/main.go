package t_worker

import (
	"github.com/zxlzhd/temporal_md_a/temporal"
	"go.temporal.io/sdk/worker"
	"log"
	"sync"

	"go.temporal.io/sdk/client"
)

// @@@SNIPSTART samples-go-child-workflow-example-worker-starter
func MainW() {
	// The client is a heavyweight object that should be created only once per process.
	c, err := client.Dial(client.Options{
		//HostPort: client.DefaultHostPort,
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
		w := worker.New(c, temporal.TaskQueueName, worker.Options{})
		w.RegisterWorkflow(temporal.ParentWorkflow)
		w.RegisterWorkflow(temporal.SampleChildWorkflow)
		w.RegisterActivity(temporal.WriteMessage)
		err = w.Run(worker.InterruptCh())
	}()
	go func() {
		defer wg.Done()
		w := worker.New(c, temporal.TaskQueueNameB, worker.Options{})
		w.RegisterWorkflow(temporal.ParentWorkflowB)
		w.RegisterWorkflow(temporal.SampleChildWorkflowB)
		w.RegisterActivity(temporal.ReadMessage)
		err = w.Run(worker.InterruptCh())
	}()
	wg.Wait()
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}

// @@@SNIPEND
