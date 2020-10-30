package main

import (
	"Generalkhun/go-todo-server/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	// create router
	router := gin.Default()

	// fromt end router
	router.Use(static.Serve("/", static.LocalFile("./build", true)))
	router.Use(static.Serve("/registerR", static.LocalFile("./build", true)))
	router.Use(static.Serve("/loginformR", static.LocalFile("./build", true)))
	router.Use(static.Serve("/taskR", static.LocalFile("./build", true)))

	// CORS
	router.Use(middleware.CORSMiddleware())

	//home-router
	//router.GET("/", middleware.PreSignin())
	router.POST("/register", middleware.Register())

	//auth-router
	routerAuth := router.Group("/auth")
	{

		routerAuth.POST("/signin", middleware.Signin())
		routerAuth.GET("/refresh", middleware.Refresh())
		routerAuth.GET("/logout", middleware.Logout())
	}

	//task-router
	routerTask := router.Group("/task")
	routerTask.Use(middleware.AuthRequired())
	{
		routerTask.GET("/welcome", middleware.Welcome())
		routerTask.GET("/getTasks", middleware.GetAllTask())
		routerTask.POST("/createTask", middleware.CreateTask())
		routerTask.PUT("/undoTask/:id", middleware.UndoTask())
		routerTask.PUT("/completeTask/:id", middleware.CompleteTask())
		routerTask.DELETE("/deleteTask/:id", middleware.DeleteTask())
		routerTask.DELETE("/deleteAllTask", middleware.DeleteAllTask())

	}

	//Run the server
	//router.Run(":8080") ---> use when runing app locally
	router.Run(getPort())

}

//Get port to deploy app (on Heroku, on this case)
func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku" + port)
	}
	return ":" + port
}
