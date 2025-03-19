package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/beharc/go-example-apps/pkg/common/health"
	"github.com/beharc/go-example-apps/pkg/common/logger"
	"github.com/beharc/go-example-apps/pkg/common/state_machine"
)

const (
	TransferCreated    = "Created"
	TransferProcessing = "Processing"
	TransferCompleted  = "Completed"
	TransferFailed     = "Failed"
)

var transferStateTransitions = map[string][]string{
	TransferCreated:    {TransferProcessing},
	TransferProcessing: {TransferCompleted, TransferFailed},
	TransferCompleted:  {},
	TransferFailed:     {},
}

type TransferWorkItem struct {
	ID    string
	State string
}

func main() {
	log := logger.New()
	log.Info("Starting transfer service")

	jobs := make(chan TransferWorkItem, 100)

	go transferProcessor(jobs, log)

	mux := http.NewServeMux()
	health.AddHealthCheck(mux)
	mux.HandleFunc("/transfer", transferHandler(jobs))

	log.Info("Starting transfer service on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func transferHandler(jobs chan<- TransferWorkItem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transfer := TransferWorkItem{
			ID:    generateID(),
			State: TransferCreated,
		}

		select {
		case jobs <- transfer:
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("Transfer queued for processing"))
		default:
			http.Error(w, "Transfer queue is full", http.StatusServiceUnavailable)
		}
	}
}

func generateID() string {
	return fmt.Sprintf("TRF-%d", time.Now().UnixNano())
}

func processTransfer(transfer TransferWorkItem, log *logger.Logger) {
	sm := state_machine.NewStateMachine(transfer.State, transferStateTransitions)

	for {
		currentState := sm.GetState()
		var nextState string

		switch currentState {
		case TransferCreated:
			time.Sleep(time.Second)
			nextState = TransferProcessing
		case TransferProcessing:
			time.Sleep(20 * time.Millisecond)
			if rand.Float32() < 0.8 {
				nextState = TransferCompleted
			} else {
				nextState = TransferFailed
			}
		case TransferCompleted, TransferFailed:
			return
		}

		err := sm.Transition(nextState)
		if err != nil {
			log.Errorf("Failed to transition transfer %s: %v", transfer.ID, err)
			return
		}
		log.Infof("Transfer %s transitioned from %s to %s", transfer.ID, currentState, sm.GetState())
	}
}

func transferProcessor(jobs <-chan TransferWorkItem, log *logger.Logger) {
	for job := range jobs {
		go processTransfer(job, log)
	}
}
