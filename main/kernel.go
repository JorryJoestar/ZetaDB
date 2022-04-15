package main

import (
	"ZetaDB/container"
	"ZetaDB/execution"
	"ZetaDB/network"
	"ZetaDB/parser"
	"ZetaDB/storage"
	"ZetaDB/utility"
)

func main() {
	//get parser, rewriter, executionEngine, transaction
	parser := parser.GetParser()
	rewriter := execution.GetRewriter()
	executionEngine := execution.GetExecutionEngine()
	transaction := storage.GetTransaction()

	//create a channel to buffer all requests received from socket
	requestChannel := make(chan container.Request, utility.DEFAULT_REQUEST_CHANNEL_CAPACITY)

	//open socket and listen to request
	//push all request into requestChannel
	go network.Listen(requestChannel)

	//main loop
	for {
		//fetch a request if channel is not empty
		currentRequest := <-requestChannel

		//get sql string from this request
		currentSql := currentRequest.Sql

		//parse this sql and get an AST, if sql syntax invalid, reply immediately
		astNode, parseErr := parser.ParseSql(currentSql)
		if parseErr != nil {
			network.Reply(currentRequest.Connection, "Error: sql syntax invalid")
			continue
		}

		//TODO unfinished, change userId
		//generate an executionPlan from current userId, AST and sql string
		executionPlan, _ := rewriter.ASTNodeToExecutionPlan(1, astNode, currentSql)

		//use executionEngine to execute this executionPlan, get a result string for reply
		executionResult := executionEngine.Execute(executionPlan)

		//update all modification into disk
		transaction.PushTransactionIntoDisk()

		//reply
		network.Reply(currentRequest.Connection, executionResult)

	}

	//ktm := execution.GetKeytableManager()

}

//tm := execution.GetTableManipulator()
//ktm.InitializeSystem()

//sql := "create table student(id int, name varchar(20));"
//sql := "drop table student;"
//sql := "insert into student values (976, 'Alex');"
//sql := "delete from student where id = 976;"
/* 	sql := "select * from student;"
   	astNode, _ := Parse.ParseSql(sql)
   	pp, _ := rewriter.ASTNodeToExecutionPlan(1, astNode, sql)
   	result := ee.Execute(pp)
   	fmt.Println(result)

   	transaction.PushTransactionIntoDisk() */

//PrintTable(2)
//PrintTable(8)
//PrintTable(9)
//PrintTable(15)
