package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const projectDirName = "mongo-crud-go" // change to relevant project name

var (
	dbHost string
	dbUser string
	dbPass string
	dbPort string
	dbName string
)

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	fmt.Println(string(rootPath))

	err := godotenv.Load(string(rootPath) + `/docker/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbHost = os.Getenv("DB_HOST")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASSWD")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")

}

func GetCollection(collection string) *mongo.Collection {
	loadEnv()

	fmt.Printf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(dbName).Collection(collection)
}
