package engine

import (
	"github.ibm.com/sysflow/sf-processor/common/logger"
)

// ActionHandler type
type ActionHandler struct {
	conf map[string]string
}

// NewActionHandler creates a new handler.
func NewActionHandler(conf map[string]string) ActionHandler {
	return ActionHandler{conf}
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
	if v, ok := s.conf[ContRuntimeType]; ok && v == Docker {
		s.computeHashesOnDocker(r)
	}
	return hs
}

func (s ActionHandler) computeHashesOnDocker(r *Record) HashSet {
	var hs HashSet = HashSet{}
	rtype := Mapper.MapStr("sf.type")(r)
	if rtype == "FE" || rtype == "FF" {
		path := Mapper.MapStr("sf.file.path")(r)
		logger.Trace.Println(path)
		// ctx := context.Background()
		// cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		// if err != nil {
		// 	panic(err)
		// }
	}
	return hs
}
