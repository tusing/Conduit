// Project contains the tasks the client wants to run
type Job struct{ 
    dockerFile String
    jobFile String
}

// Returns the channel by which Actions are sent
func (j Job) JobConnection() { }

// Returns the docker file of the job as a string
func (j Job) GetDockerFile() {
    return dockerFile
}

// Can run job and stop job as specified by the string action
func (j Job) Act(action String) { }

// Compiles the Docker File and runs the jobFile in that instance and returns necessary output as a string
func (j Job) RunJob(jobName String) { }
