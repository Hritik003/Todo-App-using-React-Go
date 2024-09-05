import React from 'react';
import './App.css';
import { Container }  from 'semantic-ui-react';  
import TodoList from './TodoList';


function App(){
  return (
    <div>
      <Container>
        <TodoList/>
      </Container>
    </div>
  );
}

export default App;