package flattener

const (
	//Extended Process Attributes
	PROC_GUID                = 0
	PROC_IMAGE               = 1
	PROC_CURR_DIRECTORY      = 2
	PROC_LOGIN_GUID          = 3
	PROC_LOGIN_ID            = 4
	PROC_TERMINAL_SESSION_ID = 5
	PROC_INTEGRITY_LEVEL     = 6
	NUM_EXT_PROC_ATTRS       = PROC_INTEGRITY_LEVEL + 1

	//Hash Attributes
	PROC_SHA1_HASH   = 0
	PROC_MD5_HASH    = 1
	PROC_SHA256_HASH = 2
	PROC_IMP_HASH    = 3
	NUM_HASHES       = PROC_IMP_HASH + 1

	//Indexes into enriched flat record
	SYSFLOW_IDX = 0
	EXT_WIN_IDX = 1
	HASH_IDX    = 3
)

type EnrichedFlatRecord struct {
	Sources []int64
	Ints    [][]int64
	Strs    [][]string
}

type EFRChannel struct {
	In chan *EnrichedFlatRecord
}
