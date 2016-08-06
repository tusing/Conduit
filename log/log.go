// Log of the jobs in progress and done
package main

// Conduit server connects with providers/clients and gives the IP of their regional master
type Logger struct {
  lg Log
}

type jobRequest struct {
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

//
func (lr *Logger) append(req jobRequest) error {
  if req.requestID != nil {
    lr.lg.requestLog[req.requestID] = req
    return nil
  }
  return &common.AppendError{"Cannot Append to Log Without ID"}
}

// Gets request by its ID
func (lr *Logger) get(a *common.GetJobRequestFromLogReply, requestID string) error {
  // Returns request
  if val, ok := dict[requestID]; ok {
    
    return nil
  }
  return &common.GetFromLogError{"Request with that ID is not in the log"}
}

import {
  "time"
  "math/rand"
  "github.com/patrickmn/go-cache"
  "log"
  "net"
  "net/rpc"
  "sync"
}

func main() {
  newLog().run()
}

func newLog() Log {
  lg := new(Log)
  // Active Local Masters is a cache where masters are removed after 30 seconds if no ping
  lg.fail = make(chan error)
  return lg
}

type log struct{
  requestLog map[string] jobRequest
}

func (lg *log) run() {
  go cs.listenForProviderClient()
  go cs.listenForRegionalMaster()
  log.Fatal(<-cs.fail)
}

func (lg *log) listenForCommunication() {
  s := rpc.NewServer()
  s.Register(&ProviderOrClient{cs})
  l, err := net.Listen("tcp", common.LoggerAddr)
  if err != nil {
    lg.fail <- err
  }
  s.Accept(l)
}