import React, { Component } from 'react'
import { NavLink } from 'react-router-dom'

function Index() {
        return (
            <div className="container text-center text-primary">
                <h1>Welcome</h1>
                <NavLink className="nav-link" exact to="/">Home</NavLink>
                <NavLink className="nav-link" exact to="/createuser">Create User</NavLink>
                <NavLink className="nav-link" exact to="/getuser">Get User</NavLink>
                <NavLink className="nav-link" exact to="/updateuser">Update User</NavLink>
                <NavLink className="nav-link" exact to="/deleteuser">Delete User</NavLink>
            </div>
        )
}

export default Index