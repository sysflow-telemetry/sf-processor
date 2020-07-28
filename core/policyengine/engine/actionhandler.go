//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
package engine

// ActionHandler type
type ActionHandler struct {
	conf Config
	//fileHasher *FileHasher
}

// NewActionHandler creates a new handler.
func NewActionHandler(conf Config) ActionHandler {
	// var fh *FileHasher
	// if c, ok := conf[EnrichmentConfigKey]; ok {
	// 	fh = NewFileHasher()
	// 	fh.Init(c)
	// }
	// return ActionHandler{conf, fh}
	return ActionHandler{conf}
}

// HandleActionAsync handles actions defined in rule.
func (s ActionHandler) HandleActionAsync(rule Rule, r *Record, out func(r *Record)) {
	for _, a := range rule.Actions {
		switch a {
		case Hash:
			// h, _ := s.fileHasher.ProcessSync(r)
			// if hs, ok := h.(HashSet); ok {
			// 	r.Ctx.SetHashes(hs)
			// }
			fallthrough
		case Tag:
			fallthrough
		case Alert:
			fallthrough
		default:
			r.Ctx.AddRule(rule)
			out(r)
		}
	}
}

// HandleAction handles actions defined in rule.
func (s ActionHandler) HandleAction(rule Rule, r *Record) {
	for _, a := range rule.Actions {
		switch a {
		case Hash:
			// h, _ := s.fileHasher.ProcessSync(r)
			// if hs, ok := h.(HashSet); ok {
			// 	r.Ctx.SetHashes(hs)
			// }
			fallthrough
		case Tag:
			fallthrough
		case Alert:
			fallthrough
		default:
			r.Ctx.AddRule(rule)
		}
	}
}

func (s ActionHandler) computeHashes(r *Record) HashSet {
	var hs HashSet = HashSet{}
	rtype := Mapper.MapStr(SF_TYPE)(r)
	if rtype == TyFE || rtype == TyFF {
		//path := Mapper.MapStr("sf.file.path")(r)
		//contID := Mapper.MapStr("sf.container.id")(r)
		// if contID != sfgo.Zeros.String {
		// 	// md5Hash, sha1Hash, sha256Hash, err := s.getDockerHashes(path, contID)
		// 	s.fileHasher.ProcessAsync(r)
		// }
	}
	return hs
}

// func (s ActionHandler) getDockerHashes(path string, contID string) ([]byte, []byte, []byte, error) {
// 	command := []string{"cat", path}
// 	execConfig := types.ExecConfig{Tty: false, AttachStdout: true, AttachStderr: false, Cmd: command}
// 	respIDExecCreate, err := s.clt.ContainerExecCreate(context.Background(), contID, execConfig)
// 	if err != nil {
// 		logger.Error.Println(err)
// 		return nil, nil, nil, err
// 	}
// 	respID, err := s.clt.ContainerExecAttach(context.Background(), respIDExecCreate.ID, types.ExecStartCheck{})
// 	if err != nil {
// 		return nil, nil, nil, err
// 	}
// 	defer respID.Close()
// 	scanner := bufio.NewScanner(respID.Reader)
// 	i := 0
// 	md5Hash := md5.New()
// 	sha1Hash := sha1.New()
// 	sha256Hash := sha256.New()
// 	for scanner.Scan() {
// 		buff := scanner.Bytes()
// 		if i == 0 {
// 			md5Hash.Write(buff[1:])
// 			sha1Hash.Write(buff[1:])
// 			sha256Hash.Write(buff[1:])
// 		} else {
// 			md5Hash.Write(buff)
// 			sha1Hash.Write(buff)
// 			sha256Hash.Write(buff)
// 		}
// 		i++
// 	}
// 	return md5Hash.Sum(nil), sha1Hash.Sum(nil), sha256Hash.Sum(nil), nil

// }

// func (s ActionHandler) getDockerHashesCmd(cmd, path string, contID string) (string, error) {
// 	command := []string{cmd, path}
// 	execConfig := types.ExecConfig{Tty: false, AttachStdout: true, AttachStderr: false, Cmd: command}
// 	respIDExecCreate, err := s.clt.ContainerExecCreate(context.Background(), contID, execConfig)
// 	if err != nil {
// 		logger.Error.Println(err)
// 		return "", err
// 	}
// 	respID, err := s.clt.ContainerExecAttach(context.Background(), respIDExecCreate.ID, types.ExecStartCheck{})
// 	if err != nil {
// 		return "", err
// 	}
// 	defer respID.Close()
// 	var outBuf, errBuf bytes.Buffer
// 	_, err = stdcopy.StdCopy(&outBuf, &errBuf, respID.Reader)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return strings.Split(outBuf.String(), SPACE)[0], nil
// }

// func (s ActionHandler) computeHashesOnDocker(r *Record) HashSet {
// 	var hs HashSet = HashSet{}
// 	rtype := Mapper.MapStr("sf.type")(r)
// 	if rtype == "FE" || rtype == "FF" {
// 		path := Mapper.MapStr("sf.file.path")(r)
// 		//logger.Trace.Println("Computing hash for ", path)
// 		contID := Mapper.MapStr("sf.container.id")(r)
// 		if contID != sfgo.Zeros.String {
// 			// md5Hash, sha1Hash, sha256Hash, err := s.getDockerHashes(path, contID)
// 			md5Hash, _ := s.getDockerHashesCmd("md5sum", path, contID)
// 			sha1Hash, _ := s.getDockerHashesCmd("sha1sum", path, contID)
// 			sha256Hash, err := s.getDockerHashesCmd("sha256sum", path, contID)
// 			if err != nil {
// 				logger.Error.Println(err)
// 			} else {
// 				hs.MD5 = md5Hash
// 				hs.SHA1 = sha1Hash
// 				hs.SHA256 = sha256Hash
// 			}
// 		}
// 	}
// 	return hs
// }
