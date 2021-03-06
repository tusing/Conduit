// System Overview (important!): https://drive.google.com/folderview?id=0B-z5tuznRiNGeTFkYXBpVEd5WlU&usp=sharing&tid=0BxRqJU1Ja6DcQnMtdTVwTDdES1U

// APPENDING SYSTEM
// Use a Request data structre
// Request: holds all relevant information about a single job
// Appends request to dictionary by RequestID

// PROCESS
// PROCESS: CLIENT WANTS TO SEND A JOB
// Client gets regional master from Conduit
// Client gets request_time, list of local masters from regional master
// Client asks for provider(s) from local master
//    - Client receives list of providers
// Client sends JobRequest with Job to provider; JobRequest to regional master  (APPEND 1)
//    - Regional master gets: 
//        - traceroute, request_time, job_hash
//    - Regional master computes request_time
// Provider adds request_id to Heartbeat
// Provider does job, sends update information up the tree
//    - Provider pings regional master, which appends a start_time
//    - Provider sends: traceroute, job_hash                                (APPEND 2)
//    - Provider pings regional master, which appends a finish_time

// ADVANTAGES
// Append 1 allows the Client to tell us that it's sent a job, in case Append 2 never happens
// Append 2 allows the Provider to confirm the job using job_hash, includes relevant transactional information upon finish
// Append 1 also allows cross-checking with Heartbeat to provide real-time updates for the client
// Traceroute allows us to analyze information about the path the job has taken (i.e. if local master disconnects during job)


// {request_id: JobRequest}: what is stored
// request_id: Generated by Regional Master

// start_time: perhaps ping local master

package main
type JobRequest struct {
  // traceroute: Regional master, updated with client and local master(s) once client completes job (machine ID)
  //             perhaps provide separate client-side and provider-side traceroutes
  // request_time: When the request was sent to regional master and provider from the client; obtained from regional master
  // start_time: when the provider starts the job; obtained from regional master
  // finish_time: when the provider finishs the job; obtained from regional master
  // job_hash: a hash of the job
  // client_id
  requestID, jobHash                  int
  requestTime, startTime, finishTime  Time
  clientID                            string
}

func (j JobRequest) getID()

type Heartbeat struct {
  // Send out a heartbeat upon job start or finish.
  // requests: a list of request_ids being handled by the provider
}

// LOCATIONING: Reducing Locationing Overhead
// Something like similarity hashing:
// Formulate CLIENT_ID and PROVIDER_ID in a way that quickly determines their location.
// Formulate REQUEST_ID in a way that helps us determine the hierarchy of the request,
//    to quickly notify clients in case of regional failures.
// 
// Alternatively: Utilize the tree structure to come up with an identifier.

// HEARTBEATS: Dealing with Failures
// Handling Provider failure:
// The local master maintains a list of the latest Heartbeats from every provider.
// If a provider dies, the local master pings the regional master with a list of failed request_ids.
// The regional master then notifies the client of these failures. (Otherwise, log failures.)
// Up to client to decide what to do from here.
//
// Handling Master failure:
// Every element in the chain maintains a list of higher master it can contact in case
// of failure. Higher master then reissues a local master for this element.
//
// provider then master
// when a provider fails, other providers sending heartbeats
// then providers send heartbeats up the chain
// regional master uses this to recover from the failure to the local master
// can look at the task and see which jobs are finished, accounted for, etc
// relying on start notification
//
// alternative: more effort from regional master in event of failure
// send information both to local master and regional
//
// you have start records - utilize this

// FAILURE CASES
  // Case 1: The provider shuts down before it could notify regional master about job
    // The Regional Master deletes request with no start_time after 3 seconds and notifies client
  // Case 2: The provdier shuts down during job
    // The local master keeps track of the last heartbeat (contains every request_id its working on) 
    //from each provider and notifies the Regional master to cancel those requests and notify the client


type Job struct {
  func hashJob() {}
  func metaJob() {}
}

