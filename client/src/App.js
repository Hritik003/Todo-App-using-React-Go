import react from 'react';
import './App.css';
import {container} from 'semantic-ui-react';  
import TodoList from './TodoList';


function App(){
  return (
    <div>
      <container>
        <TodoList/>
      </container>
    </div>
  );
}

export default App;