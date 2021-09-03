package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover() // mengecek apakah terjadi error atau tidak
	if err != nil {  // jika terjadi error
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(err)
	} else { // jika tidak terjadi error
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
