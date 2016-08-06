package main

type LocalMaster struct {
	heartbeats map[Provider]string
}

func (lm LocalMaster) registerBeat(p Provider, s string) {
	heartbeats[p] = s
	// Figure out optimal heartbeat length
	if len(heartbeats) > 100 {
		rm.aggregateBeats(heartbeats)
	}
}


