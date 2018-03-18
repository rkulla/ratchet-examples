package main

import (
	"database/sql"

	"github.com/dailyburn/ratchet"
	"github.com/dailyburn/ratchet/logger"
	"github.com/dailyburn/ratchet/processors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rkulla/ratchet-examples/example2/packages"
)

// Uses a basic 3 stage pipeline, including a transformer.
func main() {
	inputDB := setupDB("mysql", "root:@tcp(127.0.0.1:3306)/srcDB")
	extractDP := processors.NewSQLReader(inputDB, mypkg.Query(5))

	transformDP := mypkg.NewMyTransformer()

	outputDB := setupDB("mysql", "root@tcp(127.0.0.1:3306)/dstDB")
	outputTable := "users2"
	loadDP := processors.NewSQLWriter(outputDB, outputTable)

	pipeline := ratchet.NewPipeline(extractDP, transformDP, loadDP)
	pipeline.Name = "My Pipeline"

	err := <-pipeline.Run()
	if err != nil {
		logger.ErrorWithoutTrace(pipeline.Name, ":", err)
		logger.ErrorWithoutTrace(pipeline.Stats())
	} else {
		logger.Info(pipeline.Name, ": Completed successfully.")
	}
}

func setupDB(driver, conn string) *sql.DB {
	db, err := sql.Open(driver, conn)
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	return db
}
