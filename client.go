func main() {
  
}

func main() {
  newClient().run()
}

func newClient() Client {
  c := new(Client)
  // Active Local Masters is a cache where masters are removed after 30 seconds if no ping
  c.serverConn, err := rpc.Dial("tcp", c.RegionalMasterAddr)
	if err != nil {
		log.Fatal(err)
	}
	c.serverConn.login()
  return c
}

func (c Client) run() {
    
}

type Client struct{
  RegionalMasterAddr = "localhost:8000"
  serverConn
  localMasterID string
  localConn
  client_id string
  rm = Conduit.getRegionalMaster()
}
  
// Formulate a job request.   
func (c Client) formJobRequest(j Job) {
  // Get regional master to create job request
  jr, local_masters := rm.makeNewRequest(client_id, j.hashJob(), j.metaJob())

  // TODO: elegant error handling
  current_master := local_masters.pop() 
  // TODO: Deal with multiple master chain
  providers, err := current_master.getProviders()
  current_provider = providers.pop()
  provider.doJob(jr, j)
}