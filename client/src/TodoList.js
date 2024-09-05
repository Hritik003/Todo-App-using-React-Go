import React,{Component} from "react";
import axios from 'axios';
import {Card, Header, Form, Input, Icon } from 'semantic-ui-react';

let endpoint = "http://localhost";

class TodoList extends Component{
    constructor(props){
        super(props);

        this.state={ 
            task:"",
            items:[],
        }
    }

    
    

    getTask =()=>{
        axios.get(endpoint+"/api/task").then((res)=>{
            if(res.data){
                this.setState({
                    items: res.data.map((item)=>{
                        let color="yellow";
                        let style ={
                            wordWrap: "break-word",
                        };

                        if(item.status){
                            color="green";
                            style["textDecorationLine"]="line-through";
                        }

                        return(
                            <Card key ={item._id} color={color} fluild clasName ="rough">
                                <Card.Content>
                                    <Card.Header textAlign ="left">
                                        <div style={style}>{item.task}</div>
                                    </Card.Header>

                                    <Card.Meta textAlign="right">
                                        <Icon 
                                            name="edit" 
                                            color="blue"
                                            onClick={()=>this.updateTask(item._id)}
                                        />
                                        <span style ={{paddingRight:"10px"}}>Undo</span>
                                        <Icon 
                                            name="delete" 
                                            color="red"
                                            onClick={()=>this.deleteTask(item._id)}
                                        />
                                        <span style ={{paddingRight:"10px"}}>Delete</span>
                                    </Card.Meta>
                                </Card.Content>
                            </Card> 
                        );

                    }),
                });
            }else{
                this.setState({
                    items:[],
                });
            }
        });
    };

    updateTask = (id)=>{
        
        axios.put(endpoint+"/api/task/"+id,{
            headers:{
                "content-Type":"application/json",
                },
            }).then(()=>{
                console.log("Task updated");
                this.getTask();
            });
        }

    componentDidMount(){
        this.getTask();
    }

    onChange = (event => {
        this.setState({
            [event.target.name]: event.target.value
        });
    })

    onSubmit =()=>{
        let {task} = this.state;

        if(task){
            axios.post(endpoint+"/api/tasks",
                {task,},
                {headers:{

                    "content-Type":"application/json",
                }
            }
        ).then((res)=>{
            this.getTask();
            this.setState({
                task:"",
            });
        console.log("Task saved");
        });
    }
    }
    

    undoTask = (id)=>{
        axios.put(endpoint, "/api/undotask/ " + id, {
            headers:{
                "content-Type":"application/json",
            },
        }).then(()=>{
            console.log("Task updated");
            this.getTask();
        });
    }

    deleteTask = (id)=>{
        axios.delete(endpoint + "/api/deleteTask/" +id ,{
            headers:{
                "content-Type":"application/json",
            },
        }).then(()=>{
            console.log("Task deleted");
            this.getTask();
        });
    }

    render(){
        return (
            <div>
                <div className="row">
                    <Header className="header" as="h2" color="yellow">
                        My To Do List!
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
                            placeholder="create Task"                            
                        />

                    </Form>

                </div>
                <div className="row">
                    <Card.Group>
                        {this.state.items}
                    </Card.Group>
                </div>
            </div>
           
        );
    }
} 

export default TodoList;



 