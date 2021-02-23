package common

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/23 9:17
 * @description: TODO
 */

const (
	// log
	VarLogLevelDebug   = "debug"
	VarLogLevelInfo    = "info"
	VarLogLevelWarning = "warning"
	VarLogLevelError   = "error"

	// sync mode
	VarSyncModeAll  = "all"
	VarSyncModeIncr = "incr"
	VarSyncModeFull = "full"

	// mysql connect mode
	VarMySQLConnectModeReplicate    = "replicate"
	VarMySQLConnectModeShardingJDBC = "shardingjdbc"
	VarMySQLConnectModeProxy        = "proxy"
	VarMySQLConnectModeStandalone   = "standalone"

	// full_sync.create_index
	VarFullSyncCreateIndexNone       = "none"
	VarFullSyncCreateIndexForeground = "beforesync"
	VarFullSyncCreateIndexBackground = "aftersync"

	// incr_sync.mysql_fetch_method
	VarIncrSyncMySQLFetchMethodBinlog = "binlog"

	// incr_sync.tunnel
	VarTunnelDirect = "direct"
	VarTunnelRpc    = "rpc"
	VarTunnelFile   = "file"
	VarTunnelTcp    = "tcp"
	VarTunnelKafka  = "kafka"
	VarTunnelMock   = "mock"

	// incr_sync.tunnel.message
	VarTunnelMessageRaw  = "raw"
	VarTunnelMessageJson = "json"
	VarTunnelMessageBson = "bson"

	// incr_sync.conflict_write_to
	VarIncrSyncConflictWriteToNone = "none"
	VarIncrSyncConflictWriteToDb   = "db"
	VarIncrSyncConflictWriteToSdk  = "sdk"

	// checkpoint.storage.db
	VarCheckpointStorageDbReplicaDefault  = "mysqlshake"
	VarCheckpointStorageDbShardingDefault = "admin"
	VarCheckpointStorageCollectionDefault = "ckpt_default"

	// inner variable: checkpoint.storage
	VarCheckpointStorageApi      = "api"
	VarCheckpointStorageDatabase = "database"

	// innder variable: incr_sync.reader_debug
	VarIncrSyncReaderDebugNone    = ""
	VarIncrSyncReaderDebugDiscard = "discard" // throw all
	VarIncrSyncReaderDebugPrint   = "print"   // print
)
