import React, { Component, useEffect, useState } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon, Button } from "semantic-ui-react";
import "semantic-ui-css/semantic.min.css";
let endpoint = "http://127.0.0.1:8080";

function Tasks(props) {
  // state variable
  const [task, setTask] = useState("");
  const [item, setItem] = useState([]);
  const [username, setUsername] = useState("");

  //initiate tasks from the user
  GetTask();

  //Handler functions
  function onChangeHandlers(event) {
    setTask(event.target.value);
  }
  
  //task functions
  // gettask: get all task from db
  function GetTask() {
    console.log("hey");
    useEffect(() => {
      axios
        .get(endpoint + "/task/getTasks", { withCredentials: true })
        .then((response) => {
          var responseData = response.data;
          // set username to display as welcome message

          setUsername(
            responseData.slice(
              responseData.indexOf("?") + 1,
              responseData.length - 1
            )
          );
          console.log(username);
          var todoArray = responseData.slice(0, responseData.indexOf("?") - 1);

          todoArray = JSON.parse(todoArray);
          if (todoArray.length > 0) {
            var mappedItem = todoArray.map((item) => {
              let color = "yellow";
              if (item.status) {
                color = "green";
              }
              return (
                <Card key={item._id} color={color} fluid>
                  <Card.Content>
                    <Card.Header textAlign="left">
                      <div style={{ wordWrap: "break-word" }}>{item.task}</div>
                    </Card.Header>

                    <Card.Meta textAlign="right">
                      <Icon
                        name="check circle"
                        color="green"
                        onClick={() => completeTask(item._id)}
                      />
                      <span style={{ paddingRight: 10 }}>Done</span>
                      <Icon
                        name="undo"
                        color="yellow"
                        onClick={() => undoTask(item._id)}
                      />
                      <span style={{ paddingRight: 10 }}>Undo</span>
                      <Icon
                        name="delete"
                        color="red"
                        onClick={() => deleteTask(item._id)}
                      />
                      <span style={{ paddingRight: 10 }}>Delete</span>
                    </Card.Meta>
                  </Card.Content>
                </Card>
              );
            });
            setItem(mappedItem);
          } else {
            setItem([]);
          }
        });
    },[]);
  }

  function updateTask(){
    axios
    .get(endpoint + "/task/getTasks", { withCredentials: true })
    .then((response) => {
      var responseData = response.data;
      // set username to display as welcome message

      setUsername(
        responseData.slice(
          responseData.indexOf("?") + 1,
          responseData.length - 1
        )
      );
      console.log(username);
      var todoArray = responseData.slice(0, responseData.indexOf("?") - 1);

      todoArray = JSON.parse(todoArray);
      if (todoArray.length > 0) {
        var mappedItem = todoArray.map((item) => {
          let color = "yellow";
          if (item.status) {
            color = "green";
          }
          return (
            <Card key={item._id} color={color} fluid>
              <Card.Content>
                <Card.Header textAlign="left">
                  <div style={{ wordWrap: "break-word" }}>{item.task}</div>
                </Card.Header>

                <Card.Meta textAlign="right">
                  <Icon
                    name="check circle"
                    color="green"
                    onClick={() => completeTask(item._id)}
                  />
                  <span style={{ paddingRight: 10 }}>Done</span>
                  <Icon
                    name="undo"
                    color="yellow"
                    onClick={() => undoTask(item._id)}
                  />
                  <span style={{ paddingRight: 10 }}>Undo</span>
                  <Icon
                    name="delete"
                    color="red"
                    onClick={() => deleteTask(item._id)}
                  />
                  <span style={{ paddingRight: 10 }}>Delete</span>
                </Card.Meta>
              </Card.Content>
            </Card>
          );
        });
        setItem(mappedItem);
      } else {
        setItem([]);
      }
    })

  }

  function undoTask(id) {
    axios
      .put(endpoint + "/task/undoTask/" + id, {}, { withCredentials: true })
      .then((res) => {
        console.log(res);
        updateTask();
      });
  }

  function completeTask(id) {
    console.log("code is here");
    axios
      .put(endpoint + "/task/completeTask/" + id, {}, { withCredentials: true })
      .then((res) => {
        console.log("code is her2");
        console.log(res);
        updateTask();
      });
  }

  function deleteTask(id) {
    //Becareful, axios.delete has different structure: header is on second argument
    axios
      .delete(endpoint + "/task/deleteTask/" + id, { withCredentials: true })
      .then((res) => {
        console.log(res);
        updateTask();
      });
  }

  function createTask(event) {
    if(task !== ""){
      axios
        .post(endpoint + "/task/createTask",{task},{withCredentials:true})
        .then((res)=>{
          console.log("Task Added")
          updateTask()
        })
    }
  }



  //render
  return (
    <div>
      <div className="row">
        <Header className="header" as="h2">
          Hi {username}, want you want to do next?
        </Header>
      </div>
      <div className="row">
        <Form>
          <Input
            type="text"
            name="task"
            onChange={onChangeHandlers}
            value={task}
            fluid
            placeholder="Create Task"
          />
          <Button onClick={createTask}>Create Task</Button>
        </Form>
      </div>
      <div className="row">
        <Card.Group>{item}</Card.Group>
      </div>
    </div>
  );
}

export default Tasks;
