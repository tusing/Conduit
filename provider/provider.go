// CURRENT TODO:
// * Implement taking care of unsent messages
// * Implement provider job start/error/end in localMaster
// * Implement searching for new RM/LM if one fails
// * Complete handling of unsent messages
// * Complete handling of heartbeat loop

// QUESTIONS:
// * When to use capital letter for struct name?
// * In-depth look at how handleExecutable works

package main

import (
	"crypto/md5"
	"fmt"
	"github.com/tusing/Conduit/common"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/exec"
	"path"
	"strconva"
	"time"
)

type Executor struct{}

type provider struct {
	server     *rpc.Client
	addr       string
	providerID string
	rm         RegionalMaster
	lm         LocalMaster
	jobs       map[string]string // [jobID]md5
	messages   []message         // FIFO for unsent messages
}

type job struct {
	jobID        string // job ID = client ID + time (?)
	clientID     string // client's ID
	expected_md5 string // client's md5 of job
	executable   *common.Executable
	reply        *common.ExecutionReply
}

type message struct {
	time  Time // Time message was created
	jobID string
	err   error // Error relating to job
}

func (p provider) sendMessage(m message) {
	/*
	   Send a message from the provider to the local master.
	   Save unsent messages.

	   Args:
	       p: The provider sending the message.
	       m: The relavent message to send.
	*/

	err = p.lm.registerMessage(p.providerID, m)
	if err != nil {
		p.messages = append(p.messages, m)
	}
}

func (p provider) heartbeat() {
	/*
	   Send a heartbeat to the relevant local master.

	   Args:
	       p: The provider sending the heartbeat.
	*/
	p.sendMessage(message{time: time.Now().UTC()})
}

func (p provider) doJob(j job) error {
	/*
	   Have the provider execute the given job. Notify the local master
	   on a job start or stop. Check for MD5 hash agreement, and log any
	   errors.

	   Args:
	       j: The given job.
	*/

	actual_md5, expected_md5 := md5.Sum(*j.executable), j.expected_md5
	if actual_md5 != expected_md5 {
		err = fmt.Errorf("md5 mismatch! \n %q, %q", actual_md5, expected_md5)
		p.sendMessage(message{time: time.Now().UTC(), jobID: j.jobID, err: err})
		return
	}

	p.sendMessage(message{time: time.Now().UTC(), jobID: job.jobID}) // Message on job start
	output, err := handleExecutable(j.executable, j.reply)
	p.sendMessage(message{time: time.Now().UTC(), jobID: job.jobID, err: err})
}

func handleExecutable(x *common.Executable, r *common.ExecutionReply) {
	/*
	   Execute the given executable.

	   Args:
	       x: The executable to execute.
	       r: Executable output.
	   Returns:
	       r.Output: Executable output.
	       error: Errors that might have arisen.
	*/

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
	return cmd.CombinedOutput() // returns output, err
}

func newProvider(c Conduit) *provider {
	/*
	   Create a new provider.

	   Args:
	       c: The main Conduit server.
	       returns: provider.
	   Returns:
	       p: The created provider.
	*/

	p := new(provider)
	serverConn, err := rpc.Dial("tcp", common.ProviderListenerAddr)
	if err != nil {
		log.Fatalf("Couldn't connect to client: %s", err)
	}
	p.server = serverConn
	p.addr = ":" + strconv.Itoa(port)
	p.providerID = c.getID()
	rm = c.requestRegionalMaster()
	lm = rm.requestLocalMaster()
	p.heartbeat("Provider created.")
	return p
}

func (p *provider) terminate() {
	/*
	   Terminate the provider. Leave the Conduit.

	   Args:
	       p: The provider to terminate.
	*/

	p.heartbeat("Terminated job.")
	args := common.ProviderJoinLeaveArgs{p.addr}
	err := p.server.Call("Provider.Leave", &args, nil)
	if err != nil {
		fmt.Errorf("error leaving conduit: %s", err)
	}
}

func (p *provider) join() {
	/*
	   Join the Conduit.

	   Args:
	       p: The provider to join to the Conduit.
	*/

	args := common.ProviderJoinLeaveArgs{p.addr}
	err := p.server.Call("Provider.Join", &args, nil)
	if err != nil {
		log.Fatalf("couldn't join conduit: %s", err)
	}
}

func (p *provider) listen() {
	/*
	   Listen for job requests.

	   Args:
	       p: The provider to listen for requests.
	*/
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

	// err := foo()
	log.Fatal(err)
}

// EXAMPLE CODE FOR FUTURE REFERENCE
// func foo() (err error) {
//  for err == nil {
//      err = p.heartbeat("")
//      time.Sleep(100 * time.Millisecond)
//  }
//  sum(1, 2, 3, 4, 5)
//  nums := []int{1, 2, 3, 4, 5}
//  sum(-1, -2, nums...)
//  return
// }

// func sum(nums ...int) (total int) {
//  for _, num := range nums {
//      total += num
//  }
//  return
// }

//func (cs *conduitServer) run() {
//  go cs.listenForProviders()
//  go cs.listenForClients()
//  log.Fatal(<-cs.fail)
//}
