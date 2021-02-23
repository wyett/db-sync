package common

import "database/sql"

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/23 9:11
 * @description: TODO
 */

type MySQLConn struct {
	session sql.Conn
}
