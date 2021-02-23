package config

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/23 9:22
 * @description: TODO
 */

type CollectorConfig struct {
	// 1. version
	CollectorVersion string `config:collector.version`

	// 2. common
	CollectorMasterQuorum string `config:master_quorum`
	FullSyncHttpPort      uint32 `config:full_sync.http_port`
	IncrSyncHttpPort      uint32 `config:incr_sync.http_port`
	SystemProfilePort     uint32 `config:system_profile_port`
	LogLevel              string `config:log.level`
	LogDir                string `config:log.dir`
	LogFile               string `config:log.file`
	LogFlush              bool   `config:log.flush`
	CollectorMode         string `config:collector.mode`

	// 3.source
	CollectorSourceMode             string   `config:collector.source.mode`
	CollectorSourceHost             []string `config:collector.source.host`
	CollectorSourceUsername         string   `config:collector.source.user`
	CollectorSourcePassword         string   `config:collector.source.password`
	CollectorSourceDbname           []string `config:collector.source.dbname`
	CollectorSourceCharset          []string `config:collector.source.charset`
	CollectorSourcePoolMaxIdleConns uint16   `config:collector.source.pool.maxIdleConns`
	CollectorSourcePoolMaxLifeTime  uint32   `config:collector.source.pool.maxLifeTime`
	CollectorSourcePoolMaxIdleTime  uint32   `config:collector.source.pool.maxIdleTime`
	CollectorSourcePoolMaxOpenConns uint16   `config:collector.source.pool.maxOpenConns`
	CollectorSourceProxyMethod      string   `config:collector.source.proxy.method`

	// 4. sync
	CollectorSyncMode      string   `config:collector.sync.mode`
	CollectorSyncWhiteList []string `config:collector.sync.whitelist`
	CollectorSyncBlackList []string `config:collector.sync.blacklist`
	CollectorFilterType    string   `config:collector.filter.type`

	// 5.destination
	CollectorDestMode             string   `config:collector.destination.mode`
	CollectorDestMySQLHost        []string `config:collector.destination.mysql.host`
	CollectorDestMySQLUsername    string   `config:collector.destination.mysql.user`
	CollectorDestMySQLPassword    string   `config:collector.destination.mysql.password`
	CollectorDestMySQLDbname      string   `config:collector.destination.mysql.dbname`
	CollectorDestMySQLCharset     string   `config:collector.destination.mysql.charset`
	CollectorDestPoolMaxIdleConns uint16   `config:collector.destination.pool.maxIdleConns`
	CollectorDestPoolMaxLifeTime  uint32   `config:collector.destination.pool.maxLifeTime`
	CollectorDestPoolMaxIdleTime  uint32   `config:collector.destination.pool.maxIdleTime`
	CollectorDestPoolMaxOpenConns uint16   `config:collector.destination.pool.maxOpenConns`
	CollectorDestProxyMethod      string   `config:collector.destination.proxy.method`

	// 6.mq
	CollectorDestMQServers         string `config:collector.destination.mq.servers`
	CollectorDestMQRetries         uint8  `config:collector.destination.mq.retries`
	CollectorDestMQBatchSize       uint32 `config:collector.destination.mq.batchSize`
	CollectorDestMQMaxRequestSize  uint32 `config:collector.destination.mq.maxRequestSize`
	CollectorDestMQLingerMs        uint32 `config:collector.destination.mq.lingerMs`
	CollectorDestMQBufferMemory    uint32 `config:collector.destination.mq.bufferMemory`
	CollectorDestMQCanalBatchSize  uint32 `config:collector.destination.mq.canalBatchSize`
	CollectorDestMQCanalGetTimeOut uint32 `config:collector.destination.mq.canalGetTimeout`
	CollectorDestMQFlatMessage     bool   `config:collector.destination.mq.flatMessage`
	CollectorDestMQCompressionType string `config:collector.destination.mq.compressionType`
	CollectorDestMQAcks            uint8  `config:collector.destination.mq.acks`
	CollectorDestMQProducerGroup   string `config:collector.destination.mq.producerGroup`
	CollectorDestMQAccessChannel   string `config:collector.destination.mq.accessChannel`
	CollectorDestMQTransaction     string `config:collector.destination.mq.transaction`

	// 6.checkpoint
	CheckPointStorageUrl        string            `config:checkpoint.storage.url`
	CheckPointStorageDb         string            `config:checkpoint.storage.db`
	CheckPointStorageCollection string            `config:checkpoint.storage.collection`
	CheckPointStartPosition     uint32            `config:checkpoint.start_position`
	TransaformNamespace         map[string]string `config:transform.namespace`

	// 7.full sync
	FullSyncReaderTableParallel       uint8  `config:full_sync.reader.table_parallel`
	FullSyncReaderWriteRowParallel    uint8  `config:full_sync.reader.write_row_parallel`
	FullSyncReaderRowBatchSize        uint8  `config:full_sync.reader.row_batch_size`
	FullSyncReaderReadRowCount        uint8  `config:full_sync.reader.read_row_count`
	FullSyncTableExistDrop            bool   `config:full_sync.table_exist_drop`
	FullSyncCreateIndex               string `config:full_sync.create_index`
	FullSyncExecutorInsertOnDupUpdate bool   `config:full_sync.executor.insert_on_dup_update`
}

var option *CollectorConfig
