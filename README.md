# Moesif AWS Lambda Example for Go

[Moesif](https://www.moesif.com) is an API analytics platform.
[moesif-aws-lambda-go](https://github.com/Moesif/moesif-aws-lambda-go)
is a middleware that logs API calls to Moesif for AWS Lambda.

This example is a Go application with Moesif's API analytics and monitoring integrated.

## How to run this example.

Make sure `moesif-aws-lambda-go` is installed. If not, you could install with `go get github.com/moesif/moesif-aws-lambda-go`

Create a new AWS Lambda function that is trigged by AWS API Gateway

Create a zip using following steps:
1. GOOS=linux go build main.go
2. zip function.zip main

Upload this zip, when prompted for handler, enter `main`

You will also want to add an environment vairable `MOESIF_APPLICATION_ID` with the value being your 
application id from your Moesif account.

Your Moesif Application Id can be found in the [_Moesif Portal_](https://www.moesif.com/).
After signing up for a Moesif account, your Moesif Application Id will be displayed during the onboarding steps. 

You can always find your Moesif Application Id at any time by logging 
into the [_Moesif Portal_](https://www.moesif.com/), click on the top right menu,
 and then clicking _Installation_.

Go to the URL for the API gateway such as https://XXXXXX.execute-api.us-west-2.amazonaws.com/default/my-test-function

The API Calls should show up in Moesif.