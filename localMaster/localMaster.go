package main

import (
	"github.com/tusing/Conduit/common"
)

type LocalMaster struct {
	providers     map[string]message // {providerID: last heartbeat}
	currentJobs   map[string]string  // {jobID: providerID}
	completedJobs map[string]error   // {jobID: error}
}

func (lm LocalMaster) registerMessage(provider_id string, m message) {
	/*
	   Track messages (heartbeats and job start/stops) from providers.
	   Utilize the fact that every job only sends 2 messages (start/stop) to log
	   job completion. Keep track of provider uptime with every message.

	   Args:
	       provider_id: The ID of the provider being tracked.
	       m: The message being logged.
	*/

	if m.jobID == nil {
		lm.providers[provider_id] = message
	} else {
		lm.providers[provider_id] = message{time: m.time}
		_, present = currentJobs[m.jobID]
		if !present {
			currentJobs[m.jobID] = m.err
		} else {
			completedJobs[m.jobID] = m.err
			delete(currentJobs, m.jobID)
		}
	}
}
