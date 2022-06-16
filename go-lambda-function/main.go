package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct{
	Name string `json:"what is your name?"`
	Age int `json:"How old are you"`
}

type Response struct{
	Message string `json:"Answer"`
}
func HandleLambdaEvent(event Event)(Response, error){
	return Response{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main(){
	lambda.Start(HandleLambdaEvent)
}

