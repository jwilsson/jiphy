package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type StackProps struct {
	awscdk.StackProps
}

func NewJiphyStack(scope constructs.Construct, id string, props *StackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, &id, &sprops)

	// Setup Lambda function
	function := awslambda.NewFunction(stack, jsii.String("JiphyFunction"), &awslambda.FunctionProps{
		Architecture: awslambda.Architecture_ARM_64(),
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("./app/build/"), nil),
		Environment: &map[string]*string{
			"DYNAMO_TABLE_NAME":    jsii.String(os.Getenv("DYNAMO_TABLE_NAME")),
			"SLACK_SIGNING_SECRET": jsii.String(os.Getenv("SLACK_SIGNING_SECRET")),
		},
	})

	functionUrl := function.AddFunctionUrl(&awslambda.FunctionUrlOptions{
		AuthType: awslambda.FunctionUrlAuthType_NONE,
	})

	awscdk.NewCfnOutput(stack, jsii.String("jiphyFunctionUrlOutput"), &awscdk.CfnOutputProps{
		Value: functionUrl.Url(),
	})

	// Setup DynamoDB Table
	awsdynamodb.NewTable(stack, jsii.String(os.Getenv("DYNAMO_TABLE_NAME")), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("image_name"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewJiphyStack(app, "JiphyStack", &StackProps{})

	app.Synth(nil)
}
