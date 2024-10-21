package database

import (
	"context"
	"fmt"
	"log"
	"os"
)

func DbInstance() *mongo.Client {
	MongoDb := os.Getenv("MONGO_URI")
}
