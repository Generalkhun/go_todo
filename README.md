# Goreact todoapp 
Toy Todo app using (https://golang.org/ Go), (https://www.mongodb.com/ MongoDB) where a logged in user can access their todo list and use simple (https://reactjs.org/ React app) to display user interface. With (https://dashboard.heroku.com/ Heroku) server, Deployed project is (https://goreactmongo-todo.herokuapp.com/loginformR here)

### Login page
![alt text](https://github.com/Generalkhun/go_todo/blob/master/loginPage.png?raw=true)
### Create account page
![alt text](https://github.com/Generalkhun/go_todo/blob/master/registerPage.png?raw=true)
### Tasks page
![alt text](https://github.com/Generalkhun/go_todo/blob/master/tasksPage.png?raw=true)

## server
Use (https://golang.org/ Go) with (https://github.com/gin-gonic/gin Gin) to be serverside. After recieved requests from client, server will perform tasks 
 1. create account => add username and password to mongoDB
 2. login => check username and password via matching JWT token
 3. create/delete/undo tasks => perform CRUD operations on MongoDB
 4. refresh token for users that interact with app and their token is not expired yet

## client 
Use (https://reactjs.org/ React app) to connect API from server and display user's todolist data. 
After created an account and login, User can interact with their todolist database on (https://www.mongodb.com/ MongoDB) including: Add a todo item, toggle a todo item as finished (or unfinished), Delete a todo item and Delete whole todolist 

# Deployment
Use (https://docs.npmjs.com/cli/v6/commands/npm-build npm build) with React app and (https://www.docker.com/ Dockerize) it with /server. then push the app container to (https://dashboard.heroku.com/ Heroku). Following a guide from (https://medium.com/jaysonmulwa/deploying-a-go-fiber-go-react-app-to-heroku-using-docker-7379ed47e0fc here)
