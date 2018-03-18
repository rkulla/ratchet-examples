package main

import (
	"database/sql"

	"github.com/dailyburn/ratchet"
	"github.com/dailyburn/ratchet/logger"
	"github.com/dailyburn/ratchet/processors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rkulla/ratchet-examples/example4/packages"
)

// Builds on example3
func main() {
	inputDB := setupDB("mysql", "root:@tcp(127.0.0.1:3306)/srcDB")
	extractDP := processors.NewSQLReader(inputDB, mypkg.Query1())
	extractDP2 := processors.NewSQLReader(inputDB, mypkg.Query2())

	transformDP := mypkg.NewMyTransformer()

	outputDB := setupDB("mysql", "root@tcp(127.0.0.1:3306)/dstDB")
	outputTable := "users2"
	loadDP := processors.NewSQLWriter(outputDB, outputTable)

	layout, err := ratchet.NewPipelineLayout(
		ratchet.NewPipelineStage( // Stage 1
			ratchet.Do(extractDP).Outputs(transformDP),
			ratchet.Do(extractDP2).Outputs(transformDP),
		),
		ratchet.NewPipelineStage( // Stage 2
			ratchet.Do(transformDP).Outputs(loadDP),
		),
		ratchet.NewPipelineStage( // Stage 3
			ratchet.Do(loadDP),
		),
	)

	pipeline := ratchet.NewBranchingPipeline(layout)
	pipeline.Name = "My Pipeline"

	err = <-pipeline.Run()
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
