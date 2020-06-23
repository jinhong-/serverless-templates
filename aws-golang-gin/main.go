package main

import (
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	awsExecutionEnv := os.Getenv("AWS_EXECUTION_ENV")

	// set server mode
	g := gin.New()
	if port == "" {
		port = "3000"
	}
	addr := ":" + port
	if awsExecutionEnv == "" { //Running outside of lambda
		http.ListenAndServe(addr, g())
	} else { //Running in lambda
		gateway.ListenAndServe(addr, g())
	}
}