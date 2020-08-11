package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	sparta "github.com/mweagle/Sparta"
	spartaCF "github.com/mweagle/Sparta/aws/cloudformation"
	"github.com/sirupsen/logrus"
	"os"
)

func helloWorld(ctx context.Context) (string, error) {
	logger, loggerOk := ctx.Value(sparta.ContextKeyLogger).(*logrus.Logger)
	if loggerOk {
		logger.Info("Accessing structured logger!")
	}
	return "Hello World. Welcome to AWS Lambda!", nil
}

func main() {
	lambdaFn, _ := sparta.NewAWSLambda("Hello World", helloWorld, sparta.IAMRoleDefinition{})

	sess := session.Must(session.NewSession())
	awsName, awsNameErr := spartaCF.UserAccountScopedStackName("MyHelloWorldStack", sess)
	if awsNameErr != nil {
		fmt.Println("Failed to create stack name")
		os.Exit(1)
	}

	var lambdaFunctions []*sparta.LambdaAWSInfo
	lambdaFunctions = append(lambdaFunctions, lambdaFn)

	err := sparta.Main(awsName,
		"Simple Sparta HelloWorld application",
		lambdaFunctions,
		nil,
		nil,
		)
	if err != nil {
		os.Exit(1)
	}
}
