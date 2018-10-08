package main

import (
	"os"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)

type Echo struct {
	GitSHA string `json:"gitSha"`
	GitBranch string `json:"gitBranch"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	echo := Echo{
		GitSHA: os.Getenv("gitSha"),
		GitBranch: os.Getenv("gitBranch"),
	}
	echoMarshalled, err := json.Marshal(echo)
	if err != nil {
		panic(err)
	}
	return events.APIGatewayProxyResponse{Body: string(echoMarshalled), StatusCode: 200}, nil
}


func main() {
	lambda.Start(Handler)
}
