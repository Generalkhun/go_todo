# Goreact todoapp 
Toy Todo app using [Go](https://golang.org/), [MongoDB](https://www.mongodb.com/) where a logged in user can access their todo list and use simple [React app](https://reactjs.org/) to display user interface. With [Heroku](https://dashboard.heroku.com/) server, Deployed project is [here](https://goreactmongo-todo.herokuapp.com/loginformR)

### Login page
![alt text](https://github.com/Generalkhun/go_todo/blob/master/loginPage.png?raw=true)
### Create account page
![alt text](https://github.com/Generalkhun/go_todo/blob/master/registerPage.png?raw=true)
### Tasks page
![alt text](https://github.com/Generalkhun/go_todo/blob/master/tasksPage.png?raw=true)

## server
Use [Go](https://golang.org/) with [Gin](https://github.com/gin-gonic/gin) to be serverside. After recieved requests from client, server will perform tasks 
 1. create account => add username and password to mongoDB
 2. login => check username and password via matching JWT token
 3. create/delete/undo tasks => perform CRUD operations on MongoDB
 4. refresh token for users that interact with app and their token is not expired yet

## client 
Use [React app](https://reactjs.org/) to connect API from server and display user's todolist data. 
After created an account and login, User can interact with their todolist database on [MongoDB](https://www.mongodb.com/) including: Add a todo item, toggle a todo item as finished (or unfinished), Delete a todo item and Delete whole todolist 

# Deployment
Use [npm build](https://docs.npmjs.com/cli/v6/commands/npm-build) with React app and [Dockerize](https://www.docker.com/) it with /server. then push the app container to [Heroku](https://dashboard.heroku.com/). Following a guide from [here](https://medium.com/jaysonmulwa/deploying-a-go-fiber-go-react-app-to-heroku-using-docker-7379ed47e0fc)
