package main

import (
	"flag"
	"gerrit.opencord.org/bbsim/internal/bbsim/devices"
	log "github.com/sirupsen/logrus"
	"sync"
)

func getOpts() *CliOptions {

	olt_id := flag.Int("olt_id", 0, "Number of OLT devices to be emulated (default is 1)")
	nni := flag.Int("nni", 1, "Number of NNI ports per OLT device to be emulated (default is 1)")
	pon := flag.Int("pon", 1, "Number of PON ports per OLT device to be emulated (default is 1)")
	onu := flag.Int("onu", 1, "Number of ONU devices per PON port to be emulated (default is 1)")
	flag.Parse()

	o := new(CliOptions)

	o.OltID = int(*olt_id)
	o.NumNniPerOlt = int(*nni)
	o.NumPonPerOlt = int(*pon)
	o.NumOnuPerPon = int(*onu)

	return o
}

func init() {
	log.SetLevel(log.DebugLevel)
	//log.SetReportCaller(true)
}

func main() {

	options := getOpts()

	log.WithFields(log.Fields{
		"OltID": options.OltID,
		"NumNniPerOlt": options.NumNniPerOlt,
		"NumPonPerOlt": options.NumPonPerOlt,
		"NumOnuPerPon": options.NumOnuPerPon,
	}).Info("BroadBand Simulator is on")

	wg := sync.WaitGroup{}
	wg.Add(1)


	go devices.CreateOLT(options.OltID, options.NumNniPerOlt, options.NumPonPerOlt, options.NumOnuPerPon)
	log.Debugf("Created OLT with id: %d", options.OltID)

	wg.Wait()

	defer func() {
		log.Info("BroadBand Simulator is off")
	}()
}