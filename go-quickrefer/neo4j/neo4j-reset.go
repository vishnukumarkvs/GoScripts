package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func resetNeo4j(ctx context.Context, uri, username, password string) error {
	driver, err := neo4j.NewDriverWithContext(
        uri,
        neo4j.BasicAuth(username, password, ""))
		
	if err != nil{
		log.Fatal("Neo4j connection Failed", err)
	}
    defer driver.Close(ctx)

	_, err = neo4j.ExecuteQuery(ctx, driver,
		"MATCH (n) DETACH DELETE n",
    nil, neo4j.EagerResultTransformer,
    neo4j.ExecuteQueryWithDatabase("neo4j"))

	return err
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	uri := os.Getenv("uri")
	username := os.Getenv("username")
	password := os.Getenv("password")

	err := resetNeo4j(context.Background(), uri, username, password)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
}
