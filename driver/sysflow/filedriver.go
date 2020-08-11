package sysflow

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"

	"github.com/linkedin/goavro"
	"github.com/sysflow-telemetry/sf-apis/go/converter"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.ibm.com/sysflow/goutils/logger"
	"github.ibm.com/sysflow/sf-processor/driver/driver"
	"github.ibm.com/sysflow/sf-processor/driver/pipeline"
)

func getFiles(filename string) ([]string, error) {
	var fls []string
	if fi, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	} else if fi.IsDir() {
		logger.Trace.Println("File is a directory")
		var files []os.FileInfo
		var err error
		if files, err = ioutil.ReadDir(filename); err != nil {
			return nil, err
		}
		for _, file := range files {
			f := filename + "/" + file.Name()
			logger.Trace.Println("File in Directory: " + f)
			fls = append(fls, f)
		}
		if len(fls) == 0 {
			return nil, errors.New("No files present in directory: " + filename)
		}

	} else {
		fls = append(fls, filename)
	}
	logger.Trace.Printf("Number of files in list: %d\n", len(fls))
	return fls, nil
}

// FileDriver represents reading a sysflow file from source
type FileDriver struct {
	pipeline *pipeline.Pipeline
}

// NewFileDriver creates a new file driver object
func NewFileDriver() driver.Driver {
	return &FileDriver{}
}

// Init initializes the file driver with the pipeline
func (f *FileDriver) Init(pipeline *pipeline.Pipeline) error {
	f.pipeline = pipeline
	return nil
}

// Run runs the file driver
func (f *FileDriver) Run(path string, running *bool) error {
	channel := f.pipeline.GetRootChannel()
	sfChannel := channel.(*plugins.SFChannel)

	records := sfChannel.In

	logger.Trace.Println("Loading file: ", path)

	sfobjcvter := converter.NewSFObjectConverter()

	files, err := getFiles(path)
	if err != nil {
		logger.Error.Println("files error: ", err)
		return err
	}
	for _, fn := range files {
		logger.Trace.Println("Loading file: " + fn)
		f, err := os.Open(fn)
		if err != nil {
			logger.Error.Println("file open error: ", err)
			return err
		}
		reader := bufio.NewReader(f)
		sreader, err := goavro.NewOCFReader(reader)
		if err != nil {
			logger.Error.Println("reader error: ", err)
			return err
		}

		for sreader.Scan() {
			datum, err := sreader.Read()
			if err != nil {
				logger.Error.Println("datum reading error: ", err)
				return err
			}
			if !*running {
				break
			}

			records <- sfobjcvter.ConvertToSysFlow(datum)
		}
		f.Close()
	}
	logger.Trace.Println("Closing main channel")
	close(records)
	f.pipeline.Wait()
	return nil
}
