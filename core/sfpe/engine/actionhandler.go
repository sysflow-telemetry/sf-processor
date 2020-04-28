package engine

import (
	"bufio"
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/sf-processor/common/logger"
)

// ActionHandler type
type ActionHandler struct {
	conf map[string]string
	clt  *client.Client
}

// NewActionHandler creates a new handler.
func NewActionHandler(conf map[string]string) ActionHandler {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	return ActionHandler{conf, cli}
}

// HandleAction handles actions defined in rule.
func (s ActionHandler) HandleAction(rule Rule, r *Record) {
	for _, a := range rule.Actions {
		switch a {
		case Hash:
			hs := s.computeHashes(r)
			r.Ctx.SetHashes(hs)
		case Alert:
		case Tag:
		default:
			r.Ctx.AddRule(rule)
		}
	}
}

func (s ActionHandler) computeHashes(r *Record) HashSet {
	var hs HashSet
	// if v, ok := s.conf[ContRuntimeType]; ok && v == Docker {
	// 	hs = s.computeHashesOnDocker(r)
	// }
	hs = s.computeHashesOnDocker(r)
	return hs
}

func (s ActionHandler) getDockerHashes(path string, contID string) ([]byte, []byte, []byte, error) {
	command := []string{"cat", path}
	execConfig := types.ExecConfig{Tty: false, AttachStdout: true, AttachStderr: false, Cmd: command}
	respIDExecCreate, err := s.clt.ContainerExecCreate(context.Background(), contID, execConfig)
	if err != nil {
		logger.Error.Println(err)
		return nil, nil, nil, err
	}
	respID, err := s.clt.ContainerExecAttach(context.Background(), respIDExecCreate.ID, types.ExecStartCheck{})
	if err != nil {
		return nil, nil, nil, err
	}
	defer respID.Close()
	scanner := bufio.NewScanner(respID.Reader)
	i := 0
	md5Hash := md5.New()
	sha1Hash := sha1.New()
	sha256Hash := sha256.New()
	for scanner.Scan() {
		buff := scanner.Bytes()
		if i == 0 {
			md5Hash.Write(buff[1:])
			sha1Hash.Write(buff[1:])
			sha256Hash.Write(buff[1:])
		} else {
			md5Hash.Write(buff)
			sha1Hash.Write(buff)
			sha256Hash.Write(buff)
		}
		i++
	}
	return md5Hash.Sum(nil), sha1Hash.Sum(nil), sha256Hash.Sum(nil), nil

}

func (s ActionHandler) computeHashesOnDocker(r *Record) HashSet {
	var hs HashSet = HashSet{}
	rtype := Mapper.MapStr("sf.type")(r)
	if rtype == "FE" || rtype == "FF" {
		path := Mapper.MapStr("sf.file.path")(r)
		contID := Mapper.MapStr("sf.container.id")(r)
		if contID != sfgo.Zeros.String {
			md5Hash, sha1Hash, sha256Hash, err := s.getDockerHashes(path, contID)
			if err != nil {
				logger.Error.Println(err)
			} else {
				hs.MD5 = string(md5Hash)
				hs.SHA1 = string(sha1Hash)
				hs.SHA256 = string(sha256Hash)
			}
		}
	}
	return hs
}
