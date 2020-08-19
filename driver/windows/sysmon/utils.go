package sysmon

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"path/filepath"
	"strings"
	"time"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

// GetTimestamp returns a ns timestamp from a timestamp string.
// The format of the string must be: 2006-01-02 15:04:05.000 and is
// assumed to be in UTC.
func GetTimestamp(tm string) int64 {
	ts, err := time.Parse(cTimeFormat, tm)
	if err != nil {
		fmt.Printf("Error parsing timestamp: %v\n", err)
		return 0
	}
	return ts.UnixNano()
}

// GetExeAndArgs returns a canonicalized path, and a set of arguments given a command line.
func GetExeAndArgs(cmdline string) (string, string) {
	if len(cmdline) > 1 && cmdline[0] == '"' {
		index := strings.Index(cmdline[1:], "\"")
		if index != -1 {
			p := cmdline[0 : index+2]
			if len(cmdline) == index+2 {
				return filepath.Clean(p), ""
			} else {
				return filepath.Clean(p), strings.TrimLeft(cmdline[index+3:], " ")
			}
		}

	}

	fields := strings.Fields(cmdline)
	if len(fields) == 1 {
		return filepath.Clean(fields[0]), ""
	} else if len(fields) > 1 {
		return filepath.Clean(fields[0]), strings.Join(fields[1:], " ")
	}
	return "", ""
}

func createOID(oid *sfgo.OID) *sfgo.OID {
	return &sfgo.OID{Hpid: oid.Hpid, CreateTS: oid.CreateTS}
}

func createPOID(oid *sfgo.OID) *sfgo.UnionNullOID {
	poid := sfgo.NewUnionNullOID()
	poid.OID = createOID(oid)
	poid.UnionType = sfgo.UnionNullOIDTypeEnumOID
	return poid
}

func createContainerID(contID *sfgo.UnionNullString) *sfgo.UnionNullString {
	return &sfgo.UnionNullString{UnionType: contID.UnionType,
		String: contID.String}
}

func createSFProcess(proc *sfgo.Process) *sfgo.SysFlow {
	sf := sfgo.NewSysFlow()
	sf.Rec = sfgo.NewUnionSFHeaderContainerProcessFileProcessEventNetworkFlowFileFlowFileEventNetworkEventProcessFlow()
	process := &sfgo.Process{State: proc.State, Ts: proc.Ts, Exe: proc.Exe,
		ExeArgs: proc.ExeArgs, Uid: proc.Uid, UserName: proc.UserName,
		Gid: proc.Gid, GroupName: proc.GroupName, Tty: proc.Tty,
		Entry: proc.Entry}
	process.Oid = createOID(proc.Oid)
	if proc.Poid != nil && proc.Poid.UnionType == sfgo.UnionNullOIDTypeEnumOID {
		process.Poid = createPOID(proc.Poid.OID)
	}
	if proc.ContainerId != nil && proc.ContainerId.UnionType == sfgo.UnionNullStringTypeEnumString {
		process.ContainerId = createContainerID(proc.ContainerId)
	}
	sf.Rec.UnionType = sfgo.SF_PROCESS
	sf.Rec.Process = process
	return sf
}

func ip2Int(ipAddr string) (uint32, error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return 0, errors.New("wrong ipAddr format")
	}
	ip = ip.To4()
	return binary.LittleEndian.Uint32(ip), nil
}
