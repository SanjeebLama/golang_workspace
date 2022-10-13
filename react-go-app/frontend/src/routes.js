import React, { Component } from "react";
import { Switch, Route, BrowserRouter } from "react-router-dom";
import App from "./App";
import Login from "./components/Login";
import Register from "./components/RegisterComponent/Register";

class Routes extends Component {
  render() {
    return (
      <BrowserRouter>
        <Switch>
          <Route path="/" exact component={App} />
          <Route path="/register" exact component={Register} />
          <Route path={"/login"} exact component={Login} />
        </Switch>
      </BrowserRouter>
    );
  }
}

export default Routes;
