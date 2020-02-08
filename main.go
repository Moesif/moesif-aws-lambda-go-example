package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"context"
	"net/http"
	"log"
	moesifawslambda "github.com/moesif/moesif-aws-lambda-go"
	moesifOptions "github.com/moesif/moesif-aws-lambda-go-example/moesif_options"
)

func HandleLambdaEvent(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
	moesifawslambda.StartCaptureOutgoing(moesifOptions.MoesifOptions())

	// Outgoing API call to third parties like Github / Stripe or to your own dependencies
	_, err := http.Get("https://api.github.com")

	// Check for any errors while sending outgoing request
	if err != nil {
		log.Printf("Error while sending outgoing request : %s.\n", err.Error())
	}

	return events.APIGatewayProxyResponse{
		Body:       request.Body,
		StatusCode: 200,
		Headers: map[string] string {
			"RspHeader1":     "RspHeaderValue1",
			"Content-Type":   "application/json",
			"Content-Length": "1000",
		},
	   }, nil
}

func main() {
	lambda.Start(moesifawslambda.MoesifLogger(HandleLambdaEvent, moesifOptions.MoesifOptions()))
}