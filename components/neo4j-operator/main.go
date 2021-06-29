package main

import (
	"fmt"

	neo4jdatabaseclient "github.com/MattiasPernhult/go-module-test/libs/neo4j-database-client"
)

func main() {
	fmt.Printf("Database: %+v\n", neo4jdatabaseclient.Neo4jDatabase{DbId: "neo4j-operator"})
}
