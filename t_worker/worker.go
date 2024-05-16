package t_worker

import (
	"github.com/zxlzhd/temporal_md_a/temporal"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

// @@@SNIPSTART samples-go-child-workflow-example-worker-starter
func WorkerRun() {
	// The client is a heavyweight object that should be created only once per process.
	c, err := client.Dial(client.Options{
		//HostPort: client.DefaultHostPort,
		HostPort: temporal.TemporalPath,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	//defer c.Close()
	err = writeWorker(c)
	if err != nil {
		log.Fatalln("Unable to start writeworker", err)
	}
	err = readWorker(c)
	if err != nil {
		log.Fatalln("Unable to start readworker", err)
	}
}
func writeWorker(c client.Client) error {
	w := worker.New(c, temporal.TaskQueueName, worker.Options{})
	w.RegisterWorkflow(temporal.ParentWorkflow)
	w.RegisterWorkflow(temporal.SampleChildWorkflow)
	w.RegisterActivity(temporal.WriteMessage)
	//err := w.Run(worker.InterruptCh())
	err := w.Start()
	if err != nil {
		log.Fatalln("Unable to create writeworker", err)
	}
	return err
}
func readWorker(c client.Client) error {
	w := worker.New(c, temporal.TaskQueueNameB, worker.Options{})
	w.RegisterWorkflow(temporal.ParentWorkflowB)
	w.RegisterWorkflow(temporal.SampleChildWorkflowB)
	w.RegisterActivity(temporal.ReadMessage)
	//err := w.Run(worker.InterruptCh())
	err := w.Start()
	if err != nil {
		log.Fatalln("Unable to create readworker", err)
	}
	return err
}

// @@@SNIPEND
