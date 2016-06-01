import (
    "net/http"
)

// Represents a single provider that runs a job (Jason, Ryan)
type Provider struct {
    var id int // Should be initialized by request to server to identify provider
    var serverID = "conduit.com/server" // not sure what the server URL is
}

// Retrieves new Job from Server 
// Checks to see if there are any already open Docker instances to run 
// Waits if there are no jobs to run
func (p Provider) GetNewJob() {
    resp, err := http.NewRequest("GET",p.serverID)
    
}

// Create and maintain docker instance
func (p Provider) StartJob(Job j) {}

// For remote termination
// Needs to stop currently running job
// May want to communicate termination and any finished return values
func (p Provider) KillJob(Job j) {}
