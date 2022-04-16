package main

import (
	"ZetaDB/container"
	"ZetaDB/execution"
	"ZetaDB/network"
	parser "ZetaDB/parser"
	"ZetaDB/storage"
	"ZetaDB/utility"
	"os"
)

func main() {
	//get Parser, rewriter, executionEngine, transaction
	Parser := parser.GetParser()
	rewriter := execution.GetRewriter()
	executionEngine := execution.GetExecutionEngine()
	transaction := storage.GetTransaction()

	//create a channel to buffer all requests received from socket
	requestChannel := make(chan container.Session, utility.DEFAULT_REQUEST_CHANNEL_CAPACITY)

	//open socket and listen to request
	//push all request into requestChannel
	go network.Listen(requestChannel)

	//main loop
	for {
		//fetch a request if channel is not empty
		currentRequest := <-requestChannel

		//get userSql and executeSql string from this request
		//userSql := currentRequest.UserSql
		executeSql := currentRequest.Sql

		//parse this sql and get an AST, if sql syntax invalid, reply immediately
		sqlAstNode, parseErr := Parser.ParseSql(executeSql)
		if parseErr != nil {
			network.Reply(currentRequest.Connection, "error: sql syntax invalid", -1)
			continue
		}

		//TODO unfinished, change userId
		//generate an executionPlan from current userId, AST and sql string
		executionPlan, rewriteErr := rewriter.ASTNodeToExecutionPlan(currentRequest.UserId, sqlAstNode, executeSql)
		if rewriteErr != nil {
			network.Reply(currentRequest.Connection, rewriteErr.Error(), -1)
			continue
		}

		//TODO debug
		if executionPlan == nil {
			network.Reply(currentRequest.Connection, "error: not supported currently", -1)
			continue
		}

		//use executionEngine to execute this executionPlan, get a result string for reply
		executionResult := executionEngine.Execute(executionPlan)

		//update all modification into disk
		transaction.PushTransactionIntoDisk()

		//reply
		if len(executionResult) > 8 && executionResult[0:8] == "userId: " {
			//return logged userId
			network.Reply(currentRequest.Connection, executionResult[8:], 1)
		} else if len(executionResult) > 7 && executionResult[0:7] == "error: " {
			network.Reply(currentRequest.Connection, executionResult, -1)
		} else if executionResult == "Execute OK, system halt" {
			network.Reply(currentRequest.Connection, executionResult, -2)
		} else {
			network.Reply(currentRequest.Connection, executionResult, 0)
		}

		//halt if required
		if executionResult == "Execute OK, system halt" {
			os.Exit(0)
		}
	}
}
