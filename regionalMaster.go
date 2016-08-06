import {
  "time"
  "math/rand"
  "github.com/patrickmn/go-cache"
  "net"
  "fmt"
  "bufio"
  "strings" // only needed below for sample processing
}

func main() {
  newRegionalMaster().run()
}

type RegionalMaster struct{
  regional_master_id string
  activeLocalMasters cache
}

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func newRegionalMaster() RegionalMaster {
  rm := new(RegionalMaster)
  // Active Local Masters is a cache where masters are removed after 30 seconds if no ping
  rm.activeLocalMasters = cache.New(5*time.Minute, 30*time.Second)u0
  return rm
}

func (r RegionalMaster) run() {
    r.listenForComm()
}

func (r RegionalMaster) listenForComm() {
  // Make new server
  // Listen for communications (control data + other info)
  // If control data is ping
    // add request info
  // If beat
    // registerBeat()
  // If login request
    // login()
}

// Appends the JobRequest to the log and returns the requestID, request_time, and a set of local masters
func (r RegionalMaster) makeNewRequest(provider_id) (int, string, []string) {
  requestTime := time.Now()
  requestID := random(0, 2147483647)
  // Todo: check if collision with another request_id
  newRequest := JobRequest{request_time: requestTime, requestID: requestID}
  r.appendNewInfo(new_request)
  masters = r.getLocalMasters()
  // send requestID, request_time and local masters to client
}
  
// Appends start time to Request
func (r RegionalMaster) appendStartTime(requestID string, time Time) {
}

// Appends end time to Request
func (r RegionalMaster) appendEndTime(requestID string, time Time) {
}
  
// Gets pings from providers (appends the job info to the log) and local masters (and updates the current local master list)
func (r RegionalMaster) registerBeats() {
  // Get pings
  id = lm.registerBeat()
  err := nil
  r.activeLocalMasters.Set(id, err, cache.DefaultExpiration)
  // TODO: Add local master to active local masters
}

// Login provider and gives ID
func (r RegionalMaster) login() {
  
}
  
// Returns request info from requestID
func (r RegionalMaster) getRequest(requestID int) (JobRequest) {
  
}