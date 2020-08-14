package flattener

const (
	//Extended Process Attributes
	PROC_GUID_STR                = 0
	PROC_IMAGE_STR               = 1
	PROC_CURR_DIRECTORY_STR      = 2
	PROC_LOGON_GUID_STR          = 3
	PROC_LOGON_ID_STR            = 4
	PROC_TERMINAL_SESSION_ID_STR = 5
	PROC_INTEGRITY_LEVEL_STR     = 6
	PROC_SIGNATURE_STR           = 7
	PROC_SIGNATURE_STATUS        = 8
	PROC_SHA1_HASH_STR           = 9
	PROC_MD5_HASH_STR            = 10
	PROC_SHA256_HASH_STR         = 11
	PROC_IMP_HASH_STR            = 12

	NUM_EXT_PROC_ATTRS_STR = PROC_IMP_HASH_STR + 1

	PROC_SIGNED_INT        = 0
	NUM_EXT_PROC_ATTRS_INT = PROC_SIGNED_INT + 1

	//Extended File Attributes
	FILE_SHA1_HASH_STR        = 0
	FILE_MD5_HASH_STR         = 1
	FILE_SHA256_HASH_STR      = 2
	FILE_IMP_HASH_STR         = 3
	FILE_SIGNATURE_STR        = 4
	FILE_SIGNATURE_STATUS_STR = 5
	FILE_DETAILS_STR          = 6
	NUM_EXT_FILE_STR          = FILE_DETAILS_STR + 1

	FILE_SIGNED_INT  = 0
	NUM_EXT_FILE_INT = FILE_SIGNED_INT + 1

	//Extended Network Attributes
	NET_SOURCE_HOST_NAME_STR = 0
	NET_SOURCE_PORT_NAME_STR = 1
	NET_DEST_HOST_NAME_STR   = 2
	NET_DEST_PORT_NAME_STR   = 3
	NUM_EXT_NET_STR          = NET_DEST_PORT_NAME_STR + 1

	//Indexes into enriched flat record
	SYSFLOW_IDX = 0
	PROC_IDX    = 1
	FILE_IDX    = 2
	NETWORK_IDX = 3

	//Hash indexes for the hash parsing
	SHA1_HASH_STR   = 0
	MD5_HASH_STR    = 1
	SHA256_HASH_STR = 2
	IMP_HASH_STR    = 3

	//Data Sources
	SYSFLOW_SRC = 0
	PROCESS_SRC = 1
	FILE_SRC    = 2
	NETWORK_SRC = 3
)

// EnrichedFlatRecord is an enriched flat record
type EnrichedFlatRecord struct {
	Sources []int64
	Ints    [][]int64
	Strs    [][]string
}

// EFRChannel enriched flat channel
type EFRChannel struct {
	In chan *EnrichedFlatRecord
}

// NewEnFlattenerChan creates a new channel with given capacity.
func NewEnFlattenerChan(size int) interface{} {
	return &EFRChannel{In: make(chan *EnrichedFlatRecord, size)}
}
