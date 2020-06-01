package engine

import (
	"github.ibm.com/sysflow/goutils/config"
	"github.ibm.com/sysflow/goutils/container/async"
	"github.ibm.com/sysflow/goutils/logger"
)

type FileHasher struct {
	fileHashTable *async.FileHashTable
}

func NewFileHasher() *FileHasher {
	return new(FileHasher)
}

func fileHashCallback(fhi *async.FileHashInfo) {
	logger.Trace.Println("CallBACK: " + fhi.File.Path + " " + fhi.Md5)
}

func (s *FileHasher) Init(confPath string) error {
	conf, err := config.GetConfig(confPath)
	if err != nil {
		return err
	}
	s.fileHashTable, err = async.NewFileHashTable(conf, nil)
	if err != nil {
		return err
	}
	s.fileHashTable.Init()
	return nil
}

func (s *FileHasher) ProcessSync(r *Record) (interface{}, error) {
	rtype := Mapper.MapStr("sf.type")(r)
	contID := Mapper.MapStr("sf.container.id")(r)
	if rtype == "FE" || rtype == "FF" {
		filePath := Mapper.MapStr("sf.file.path")(r)
		return s.process(filePath, contID), nil
	} else if rtype == "PE" {
		filePath := Mapper.MapStr("sf.proc.name")(r)
		return s.process(filePath, contID), nil
	}
	return nil, nil
}

func (s *FileHasher) ProcessAsync(r *Record, callback func(o interface{})) error {
	return nil
}

func (s *FileHasher) Cleanup() error {
	s.fileHashTable.Close()
	return nil
}

func (s *FileHasher) process(filePath string, contID string) HashSet {
	hif := s.fileHashTable.Get(contID, filePath)
	if hif != nil {
		return HashSet{MD5: hif.Md5, SHA1: hif.Sha1, SHA256: hif.Sha256, Size: hif.Size, UpdateTs: hif.LastUpdate.Unix()}
	}
	fp := &async.FilePath{Path: filePath, ContainerId: contID}
	s.fileHashTable.StartHash(fp)
	return HashSet{}
}
