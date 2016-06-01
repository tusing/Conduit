// Project contains the tasks the client wants to run
type Project struct{ 
    jobs []Job
}

// Returns the Jobs of the Project in an array
func (p Project) GetJobs() []Job { }

// Adds a job to the Project with name jobName
func (p Project) AddJob(jobName String, job Job) { }

// Removes a job from the Project specified by jobName
func (p Project) RemoveJob(jobName String) { }
