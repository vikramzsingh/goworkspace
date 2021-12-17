import React, { Component } from 'react'
import axios from 'axios'

class GetUser extends Component {

    constructor(props) {
        super(props)
    
        this.state = {
            users: []         
        }
    }

    callData() {
        const url ="http://localhost:8080/getuser"

        axios.get(url)
        .then((resp) => {
            console.log(resp)
            const persons = resp.data;
            this.setState({ users: persons });
        })
        .catch((error) => {
            console.log(error);
        })
    }

    componentDidMount() {
        this.callData();   
    }
    
    render () {
        const { users } = this.state
        return (
            <div>
                {
                    users.map((user, key) => {
                        return(
                            <div key={key}>
                                <div><h3>First Name: {user.fname}</h3></div>
                                <div><h3>Last Name: {user.lname}</h3></div>
                                <div><h3>Age: {user.age}</h3></div>
                                <div><h3>Contact No:{user.mobileno}</h3></div>
                                <div><h3>Date of birth:{user.dob}</h3></div>
                                <div><h3>Email:{user.email}</h3></div>
                                <div><h3>Address:{user.address}</h3></div>
                                <hr></hr>
                            </div>
                        )
                    })
                }
            </div>
        )
    }
}

export default GetUser
