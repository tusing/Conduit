type Provider struct{
  provider_id := Conduit.getID()
  rm = Conduit.getRegionalMaster()
  lm = rm.getLocalMaster()
}

// TODO: loop this function
func (p Provider) heart_beat() {
  lm.registerBeat(p.provider_id)
}
  
func (p Provider) doJob(jr JobRequest, j Job) {
  rm.jobStart(jr.request_id, j.hashJob())
  currently_handling = append(jr.request_id)
  // TODO: impleme job handling 
  currently_handling = // TODO: figure out how to delete by ID from a slice
  rm.jobFinish(jr.request_id)
}