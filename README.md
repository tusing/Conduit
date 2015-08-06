**Client-side Server**

*Used to keep track of live connections with slaves; Use P2P for establishing connections and asking for resources.*

* * [ ] **```send_packet(?type{parent=Packet})```**

    * * [ ] **```keep_alive()```**
    *Generates keep alive packet, saves, and passes packet to sending channel*

    * * [ ] **```verify_integrity()```**
    *Ensure slaves are up to date and have same checksums*

    * * [ ] **```verify_identity()```**
    *Ensure slave identity is secure / security breach issues*

    * * [ ] **```send_resource()```**
    *Send packet with resource/memory requested by slave*

    * * [ ] **```assign_job()```**
    *Gives a slave a job, with timeout to ensure no overloads*

* * [ ] **```read_packet(?type{parent=Channel})```**
*Reads packets from this channel*

* * [ ] **```compile_results()```**
*Compile all results together somehow.*

    * *If too slow, split between channels.*

    * *If not consistent, ask for more slaves.*

* * [ ] **```distribute_among_slaves()```**
*Creates required channels and distribute workload [not sure how to implement]
This in itself is it’s own program*

* * [ ] **```get_master()```**
*Finds the master server host, or nothing (null)``` if this is already the master*

* * [ ] **```get_living_slaves()```**
*Test which slave connections still stable, and retrieve*

* * [ ] **```find_master()```**
*Ensure healthy connection with master. If master has been disconnected, use last backup, assign new master, and revert to previous state*

* * [ ] **```find_slaves()```**
*Establish connections in the form of channels, ensure living slaves*

* * [ ] **```die()```**
*Sends death signal to master, warning possible shutdown or disconnect
If master is dying, assign new master, and tell all slaves*

* * [ ] **```backup()```**
*Create backup of resources from master*

* * [ ] **```finish()```**
*Notify slaves to die
If master server, tell primary server this run is complete, and wait for primary server response
Verify integrity of connection, and security
Send packets of data, pausing of necessary*

**Client-side VM Manager**

*Used to manage the virtual machine created by Conduit.*

* * [ ] **```create_vm()```**
*Creates a virtual machine.*

* * [ ] **```ssh_vm(keyfile)```**
*Will use a keyfile received by the server to ssh into the internal VM.*

* * [ ] **```compute_switch(true|false)```**
*Will start or stop the running VM.*

* * [ ] **```accept_job(user,job_id)```**
*Will accept a Conduit job from a user with a job id. (Routed through main server)```*

* * [ ] **```check_status()```**
*Will  check the current VM state such as available memory, etc.*

* * [ ] **```allocate_vm_dockers(num)```**
*Will allocate a number of Docker instances depending on certain factors - available memory, etc. If more Docker instances than requested are running at the same time, the jobs being run will finish before closing.*

**Client-side VM Implementation****
***Used to manage Docker instances*.

* * [ ] **```ssh_keygen(user_details)```**
Will use relevant user details to contact the server and agree upon a key.

* * [ ] **```docker_make(instance_name)```**
Creates a Docker instance with a reference name or ID.

* * [ ] **```docker_close(instance_name)```**
Closes a docker instance.
