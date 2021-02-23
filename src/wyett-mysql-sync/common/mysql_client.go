package common

import "database/sql"

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/23 9:11
 * @description: TODO
 */

type HostPortPair struct {
	Host string
	port uint32
}

type MySQLUrl struct {
	HpPairs  []HostPortPair
	Protocal string
	Username string
	Password string
	Charset  string

	//CollectorSourceMode             string   `config:collector.source.mode`
	//CollectorSourceHost             []string `config:collector.source.host`
	//CollectorSourceUsername         string   `config:collector.source.user`
	//CollectorSourcePassword         string   `config:collector.source.password`
	//CollectorSourceDbname           []string `config:collector.source.dbname`
	//CollectorSourceCharset          []string `config:collector.source.charset`
	//CollectorSourcePoolMaxIdleConns uint16   `config:collector.source.pool.maxIdleConns`
	//CollectorSourcePoolMaxLifeTime  uint32   `config:collector.source.pool.maxLifeTime`
	//CollectorSourcePoolMaxIdleTime  uint32   `config:collector.source.pool.maxIdleTime`
	//CollectorSourcePoolMaxOpenConns uint16   `config:collector.source.pool.maxOpenConns`
	//CollectorSourceProxyMethod      string   `config:collector.source.proxy.method`

}

type MySQLConn struct {
	conn *sql.Conn
}
