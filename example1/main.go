package main

import (
	"database/sql"

	"github.com/dailyburn/ratchet"
	"github.com/dailyburn/ratchet/logger"
	"github.com/dailyburn/ratchet/processors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rkulla/ratchet-examples/example1/packages"
)

// The simplest Ratchet program. Consists of a pipeline with just
// Extract and Load stages, no transformer.
func main() {
	inputDB := setupDB("mysql", "root:@tcp(127.0.0.1:3306)/srcDB")
	extractDP := processors.NewSQLReader(inputDB, mypkg.Query(5))

	outputDB := setupDB("mysql", "root@tcp(127.0.0.1:3306)/dstDB")
	outputTable := "users2"
	loadDP := processors.NewSQLWriter(outputDB, outputTable)

	pipeline := ratchet.NewPipeline(extractDP, loadDP)
	pipeline.Name = "My Pipeline"

	// To see debugging output, uncomment the following lines:
	//   logger.LogLevel = logger.LevelDebug
	//   pipeline.PrintData = true

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
