package windows

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.ibm.com/sysflow/goutils/logger"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/winlogbeat/checkpoint"
	"github.com/elastic/beats/v7/winlogbeat/eventlog"
	"github.ibm.com/sysflow/sf-processor/core/flattener"
	"github.ibm.com/sysflow/sf-processor/driver/pipeline"
	"github.ibm.com/sysflow/sf-processor/driver/windows/sysmon"
)

const (
	cName        = "name"
	cAPI         = "api"
	cChannelSize = 100000
)

func initSigTerm(running *bool) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		*running = false
	}()
}

/*func processOutput(sfChan *plugins.SFChannel, wg *sync.WaitGroup) {
	hdl := writer.NewSysFlowWriter()
	conf := map[string]string{"file": "./output.avro"}
	err := hdl.Init(conf)
	if err != nil {
		logger.Error.Println(err)
		os.Exit(1)
	}
	processor := processor.NewSysFlowProcessor(hdl)
	err = processor.Init(conf)
	if err != nil {
		logger.Error.Println(err)
		os.Exit(1)
	}
	processor.Process(sfChan, wg)
	processor.Cleanup()
}*/

// ProcessWinEvtLogs processes windows event logs and creates and exports them into the pipeline
func ProcessWinEvtLogs(pluginDir string, config string) {
	channel, pipeline, wg, channels, hdlers, err := pipeline.LoadPipeline(pluginDir, config)
	if err != nil {
		logger.Error.Println("pipeline error:", err)
		return
	}
	logger.Trace.Printf("Loaded %d stages\n", len(pipeline))
	logger.Trace.Printf("Loaded %d channels\n", len(channels))
	logger.Trace.Printf("Loaded %d hdlrs\n", len(hdlers))
	running := true
	initSigTerm(&running)
	conf := make(map[string]interface{})
	efrChannel := channel.(*flattener.EFRChannel)
	sm := sysmon.NewSMProcessor(efrChannel)
	conf[cAPI] = ""
	conf[cName] = sm.GetProvider()
	cfg, err := common.NewConfigFrom(conf)
	if err != nil {
		logger.Error.Println("winevtlog provider config error:", err)
		return
	}

	eventLog, err := eventlog.New(cfg)
	if err != nil {
		logger.Error.Println("Failed to create new event log error:", err)
		return
	}
	cp := checkpoint.EventLogState{}
	err = eventLog.Open(cp)
	if err != nil {
		fmt.Printf("failed to open windows event log: %v", err)
		return
	}
	logger.Trace.Printf("Windows Event Log '%s' opened successfully", eventLog.Name())
	// setup closing the API if either the run function is signaled asynchronously
	// to shut down or when returning after io.EOF
	/*cancelCtx, cancelFn := ctxtool.WithFunc(ctxtool.FromCanceller(ctx.Cancelation), func() {
		if err := api.Close(); err != nil {
			log.Errorf("Error while closing Windows Eventlog Access: %v", err)
		}
	})
	defer cancelFn()
	*/
	// read loop
	for running {
		records, err := eventLog.Read()
		switch err {
		case nil:
			break
		case io.EOF:
			logger.Trace.Printf("End of Winlog event stream reached: %v", err)
			return
		default:
			// only log error if we are not shutting down
			/*	if cancelCtx.Err() != nil {
				return nil
			}*/

			logger.Error.Printf("Error occured while reading from Windows Event Log '%v': %v", eventLog.Name(), err)
			return
		}

		if len(records) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}
		sm.Process(records)
	}
	close(efrChannel.In)
	wg.Wait()

}
