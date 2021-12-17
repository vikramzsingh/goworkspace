import './App.css';
import { Component } from 'react';
// import FetchUser from './components/pages/FetchUser';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import '../node_modules/bootstrap/dist/css/bootstrap.css'
import Index from './components/pages/Index';
import CreateUser from './components/pages/CreateUser';
import GetUser from './components/pages/GetUser';
import UpdateUser from './components/pages/UpdateUser';
import DeleteUser from './components/pages/DeleteUser';

class App extends Component {
  render () {
    return (
      <Router>
        <div className="App">
        {/* <FetchUser></FetchUser> */}
          <Switch>
          
            <Route exact path="/" component={Index}></Route>
            <Route exact path="/createuser" component={CreateUser}></Route>
            <Route exact path="/getuser" component={GetUser}></Route>
            <Route exact path="/updateuser" component={UpdateUser}></Route>
            <Route exact path="/deleteuser" component={DeleteUser}></Route>
          </Switch>
        </div>
      </Router>
    );
  }
}

export default App;
