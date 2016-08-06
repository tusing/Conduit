package main

import (
	"fmt"
	"github.com/lukedmor/conduit_tiny/common"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/exec"
	"path"
	"strconva"
	"time"
)

type provider struct {
	server      *rpc.Client
	addr        string
	provider_id string
	rm          RegionalMaster
	lm          LocalMaster
}

// TODO: Implement error handling
func (p provider) heartbeat(s string) {
	lm.registerBeat(p.provider_id, s)
	return err // True if ping unsuccessful
}

type Executor struct{}

// GOROUTINE FOR HEARTBEATS
// TODO: Implement heartbeat on start/end
// TODO: Implement hash checking
func (p *provider) Execute(x *common.Executable, r *common.ExecutionReply) error {
	heartbeat("Starting executable...")
	dir := os.TempDir()
	fileName := path.Join(dir, "content")
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Errorf("couldn't create file in temp dir: %s", err)
		return err
	}

	_, err = file.Write(x.Content)
	if err != nil {
		fmt.Errorf("couldn't write content to temp file: %s", err)
		return err
	}
	file.Close()

	// should validate x.Interpreter first, this is hacky
	// and must fit paradigm like `python file.py`
	cmd := exec.Command(x.Interpreter, fileName)
	r.Output, err = cmd.CombinedOutput()
	return err
}

// DONE: Implement provider_id, rm, lm
func newProvider(c Conduit) *provider {
	p := new(provider)
	serverConn, err := rpc.Dial("tcp", common.ProviderListenerAddr)
	if err != nil {
		log.Fatalf("Couldn't connect to client: %s", err)
	}
	p.server = serverConn
	p.addr = ":" + strconv.Itoa(port)
	p.provider_id = c.getID()
	rm = c.requestRegionalMaster()
	lm = rm.requestLocalMaster()
	p.heartbeat("Provider created.")
	return p
}

func (p *provider) terminate() {
	p.heartbeat("Terminated job.")
	args := common.ProviderJoinLeaveArgs{p.addr}
	err := p.server.Call("Provider.Leave", &args, nil)
	if err != nil {
		fmt.Errorf("error leaving conduit: %s", err)
	}
}

func (p *provider) join() {
	args := common.ProviderJoinLeaveArgs{p.addr}
	err := p.server.Call("Provider.Join", &args, nil)
	if err != nil {
		log.Fatalf("couldn't join conduit: %s", err)
	}
}

func (p *provider) listen() {
	rpc.Register(new(Executor))
	l, err := net.Listen("tcp", p.addr)
	if err != nil {
		log.Fatal(err)
	}
	go p.join()
	rpc.Accept(l)
}

func main() {
	p := newProvider()
	defer p.terminate()
	
	go func() {
		p.listen()
	}()

	err := foo()
	log.Fatal(err)
}

func foo() (err error) {
	for err == nil {
		err = p.heartbeat("")
		time.Sleep(100 * time.Millisecond)
	}
	sum(1, 2, 3, 4, 5)
	nums := []int{1, 2, 3, 4, 5}
	sum(-1, -2, nums...)
	return
}

func sum(nums ...int) (total int) {
  for _, num := range nums {
    total += num
  }
  return
}


//func (cs *conduitServer) run() {
//	go cs.listenForProviders()
//	go cs.listenForClients()
//	log.Fatal(<-cs.fail)
//}