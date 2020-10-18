package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"../models"
)

//CORSMiddleware (Cross-Origin Resource Sharing) middleware that used to handle Response Header
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

// GetAllTask get all the task route
func GetAllTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		cur, err := collection.Find(context.Background(), bson.D{{}})
		if err != nil {
			log.Fatal(err)
		}

		var results []primitive.M
		for cur.Next(context.Background()) {
			var result bson.M
			e := cur.Decode(&result)
			if e != nil {
				log.Fatal(e)
			}
			// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
			results = append(results, result)

		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		cur.Close(context.Background())
		w := json.NewEncoder(c.Writer).Encode(results)
		c.JSON(http.StatusOK, w)

	}

}

//DB connect

const connectionString = "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
const dbName = "todo"
const collectionName = "todoTasks"

var collection *mongo.Collection

func init() {
	// Set the client options, specified database location by using ApplyURI
	clientOptions := options.Client().ApplyURI(connectionString)

	//Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance created!")

}

// CreateTask create task route
func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.ToDoList
		_ = json.NewDecoder(c.Request.Body).Decode(&task)
		fmt.Println(task, c.Request.Body)
		insertOneTask(task)
		w := json.NewEncoder(c.Writer).Encode(task)
		c.JSON(http.StatusOK, w)
	}
}

// Insert one task in the DB
func insertOneTask(task models.ToDoList) {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}
