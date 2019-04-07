import React, { Component } from 'react';
import { Route, BrowserRouter } from "react-router-dom";
import Home from './Home'

export default class Router extends Component {
    render() {
        return (
            <BrowserRouter>
                <Route exact path='/' component={Home}></Route>
            </BrowserRouter>
        )
    }
}