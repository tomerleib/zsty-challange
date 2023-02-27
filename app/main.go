package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	awsRegion = "us-east-1"
	endpointAddress = "http://localhost:8000"
	tableName = "devops-challenge"
	keyName = "codeName"
	secretKey = "secretCode"
)
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/secret", getSecret)
	r.GET("/health", getHealth)		
	return r
}

func getHealth(c *gin.Context){
	data := map[string]interface{}{
		"status": "Healthy!",
		"container":  "tomerleib/goapp:latest",
	}
	c.IndentedJSON(http.StatusOK, data)
}

func getSecret(c *gin.Context){
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = awsRegion
		return nil
	})
	if err != nil {
			panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options){
		o.EndpointResolver = dynamodb.EndpointResolverFromURL(endpointAddress)
	})
	out, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
				keyName: &types.AttributeValueMemberS{Value: "theDoctor"},
		},
})

	if err != nil {
			log.Print(err)
			c.AbortWithStatus(500)
			return
	}
	var secretValue string
	err = attributevalue.Unmarshal(out.Item[secretKey], &secretValue)
	if err != nil {
		log.Print(err)
		c.AbortWithStatus(500)
		return
	}
	data := map[string]interface{}{
		keyName: "theDoctor",
		secretKey:  secretValue,
	}
	c.IndentedJSON(http.StatusOK, data)
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}