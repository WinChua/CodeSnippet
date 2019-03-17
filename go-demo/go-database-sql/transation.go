package main

import (
    "database/sql"
    "log"
)

func ModifyDataWithExecAndQuery(db *sql.DB) {
//  Use Exec(), preferably with a prepared statement, to accomplish an INSERT, DELETE
//    ,UPDATE, or other statement that doesn't return rows. The following example sh-
//  ows, how to insert a row and inspect metadata about operation
//  executing the stmt will return a sql.Result that gives access to stmt metadata:
//  the last inserted ID and the number of rows affected.

//  Never use db.Query("DELETE from users"), for the reason that sql.Rows returns by
//  Query will reserve a database connection util sql.Rows.Close() called. The gc will
//  eventually close the underlying net.Conn
	stmt, err := db.Prepare("INSERT INTO users(id, name) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(0, "Dolly")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

// In Go, a transaction indicates that a database connection is reserved by an object
// which let you do all of the operations on the same connections.
// You begin a transaction with a call to db.Beging(), and close it with db.Commit() or Rollback()
// method on the resulting Tx variable. Under the covers, the Tx gets a connection from the pool,
// and reserves it for use only with that transaction. The methods on the Tx map one-for-one to
// methods you can call on the database itself, such as Query() and so forth.

func TransactionExample(db *sql.DB) {

}
