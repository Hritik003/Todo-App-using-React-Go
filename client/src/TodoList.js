import React, { Component } from "react";
import axios from "axios";
import { Card, Form, Icon, Header, Input } from "semantic-ui-react";

let endpoint = "http://backend-url:8080";

class ToDoList extends Component {
  constructor(props) {
    super(props);

    this.state = {
      task: "",
      items: [],
    };
  }

  componentDidMount() {
    console.log("Component mounted, fetching tasks...");
    this.getTask();
  }

  onChange = (event) => {
    this.setState({
      [event.target.name]: event.target.value,
    });
  };

  onSubmit = () => {
    let { task } = this.state;

    if (task) {
      axios
        .post(
          `${endpoint}/api/tasks`,
          { task },
          {
            headers: {
              "Content-Type": "application/json",
            },
          }
        )
        .then((res) => {
          this.getTask();
          this.setState({
            task: "",
          });
          console.log(res);
        });
    }
  };

  getTask = () => {
    axios.get(`${endpoint}/api/tasks`).then((res) => {
      console.log(res.data);
      if (res.data) {
        this.setState({
          items: res.data.map((item) => {
            console.log(item);
            let color = "yellow";
            let style = {
              wordWrap: "break-word",
            };

            if (item.status) {
              color = "green";
              style["textDecorationLine"] = "line-through";
            }

            return (
              <Card key={item._id} color={color} fluid className="rough">
                <Card.Content>
                  <Card.Header textAlign="left">
                    <div style={style}>
                      {typeof item.task === "string"
                        ? item.task
                        : JSON.stringify(item.task)}
                    </div>
                  </Card.Header>
                  <Card.Meta textAlign="right">
                    <Icon
                      name="check circle"
                      color="blue"
                      onClick={() => this.undoTask(item._id)}
                    />
                    <Icon
                      name="delete"
                      color="red"
                      onClick={() => this.deleteTask(item._id)}
                    />
                    <span style={{ paddingRight: 10 }}>Delete</span>
                  </Card.Meta>
                </Card.Content>
              </Card>
            );
          }),
        });
      } else {
        this.setState({
          items: [],
        });
      }
    });
  };

  updateTask = (id) => {
    axios
      .put(`${endpoint}/api/tasks/${id}`, {
        headers: {
          "Content-Type": "application/json",
        },
      })
      .then((res) => {
        console.log(res);
        this.getTask();
      });
  };

  undoTask = (id) => {
    axios
      .put(`${endpoint}/api/undoTask/${id}`, {
        headers: {
          "Content-Type": "application/json",
        },
      })
      .then((res) => {
        console.log(res);
        this.getTask();
      });
  };

  deleteTask = (id) => {
    axios
      .delete(`${endpoint}/api/deleteTask/${id}`, {
        headers: {
          "Content-Type": "application/json",
        },
      })
      .then((res) => {
        console.log(res);
        this.getTask();
      });
  };

  render() {
    return (
      <div>
        <div className="row">
          <Header className="header" as="h2" color="yellow">
            To DO LIST
          </Header>
        </div>
        <div className="row">
          <Form onSubmit={this.onSubmit}>
            <Input
              type="text"
              name="task"
              onChange={this.onChange}
              value={this.state.task}
              fluid
              placeholder="Create Task"
            />
          </Form>
        </div>
        <div className="row">
          <Card.Group>{this.state.items}</Card.Group>
        </div>
      </div>
    );
  }
}

export default ToDoList;