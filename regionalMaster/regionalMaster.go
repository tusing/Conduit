package main

import (
  "time"
  "math/rand"
  "github.com/patrickmn/go-cache"
  "log"
  "net"
  "net/rpc"
  "tusing/conduit_append/common"
)

type LocalMaster struct {
  r regionalMaster
}

type Provider struct {
  r regionalMaster
}

type Client struct {
  r regionalMaster
}

// Gets pings from local master and updates the current local master list
func (lm LocalMaster) registerBeats() {
  // Get pings
  id = lm.Addr
  err := nil
  // Add local master to active local masters
  lm.r.activeLocalMasters.Set(id, err, cache.DefaultExpiration)
}

// internals

func main() {
  newRegionalMaster().run()
}

type regionalMaster struct{
  addr string
  activeLocalMasters *cache.Cache
  
  providers map[string]string // key is host addr
  clients map[string]string // key is host addr

  server *rpc.ConduitServer

  fail chan error
}

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func newRegionalMaster() RegionalMaster {
  rm := new(regionalMaster)
  // Active Local Masters is a cache where masters are removed after 30 seconds if no ping
  rm.activeLocalMasters = cache.New(5*time.Minute, 30*time.Second)
  // Set of logged in providers (possibly used to keep track of jobs done)
  rm.providers = make(map[string]bool)
  rm.clients = make(map[string]bool)
  rm.fail = make(chan error)
  
  serverConn, err := rpc.Dial("tcp", common.RegionalMasterListenerAddr)
  if err != nil {
    log.Fatalf("couldn't connect to conduit: %s", err)
  }
  rm.server = serverConn
  rm.addr = ":8010"
  return rm
}

func (r regionalMaster) run() {
  go r.listenForProviders()
  go r.listenForClients()
  go r.listenForLocalMasters()
  go r.ping()
  log.Fatal(<-r.fail)
}

func (r regionalMaster) listenForLocalMasters() {
  s := rpc.NewServer()
  s.Register(&LocalMaster{r})
  l, err := net.Listen("tcp", common.LocalMasterListenerAddr)
  if err != nil {
    cs.fail <- err
  }
  s.Accept(l)
}

func (r *regionalMaster) listenForProviders() {
  // Make new server
  s := rpc.NewServer()
  s.register(&Provider{r})
  l, err := net.Listen("tcp", common.ProviderListenerAddr)
  if err != nil {
    r.fail <- err
  }
  s.Accept(l)
}

func (cs *regionalMaster) listenForClients() {
  s := rpc.NewServer()
  s.register(&Client{r})
  l, err := net.Listen("tcp", common.ClientListenerAddr)
  if err != nil {
    cs.fail <- err
  }
  s.Accept(l)
}

// Pings the conduit server so it knows that the regional master is active
func (rm *regionalMaster) ping() {
  args := common.PingArgs{rm.addr}
	err := rm.server.Call("ConduitServer.Ping", &args)
	if err != nil {
		fmt.Errorf("Unable to Ping Server: %s", err)
	}
	time.Sleep(100 * time.Millisecond)
}

// Appends the JobRequest to the log and returns the requestID, request_time, and a set of local masters
func (c Client) makeNewRequest(provider_id) (int, string, []string) {
  requestTime := time.Now()
  requestID := random(0, 2147483647)
  // Todo: check if collision with another request_id
  newRequest := JobRequest{request_time: requestTime, requestID: requestID}
  r.appendNewInfo(new_request)
  masters = r.getLocalMasters()
  // send requestID, request_time and local masters to client
}
  
// Appends start time to Request
func (p Provider) appendStartTime(requestID string, time Time) {
}

// Appends end time to Request
func (p Provider) appendEndTime(requestID string, time Time) {
}

// Registers the provider and saves
func (p Provider) register(a *common.ProviderJoinLeaveArgs, reply *common.Nothing) string {
  // Generate provider ID
  pID, err := exec.Command("uuidgen").Output()
  if err != nil {
      log.Fatal(err)
  }
  p.r.providers[a.Addr] = pID
  return pID
}

// Registers the client and gives ID
func (c Client) register(a *common.ClientJoinLeaveArgs, reply *common.Nothing, p publicKey) string {
  // Generate client ID
  a.pID, err := exec.Command("uuidgen").Output()
  if err != nil {
      log.Fatal(err)
  }
  // Saves the pID of client and publicKey
  c.r.clients[a.pID] = publicKey
  return pID
}
  
// Returns request info from requestID
func (r regionalMaster) getRequest(requestID int) (JobRequest) {
  
}
