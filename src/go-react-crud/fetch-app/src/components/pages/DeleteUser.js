import axios from 'axios'
import React, { Component } from 'react'

class DeleteUser extends Component {
    constructor(props) {
        super(props)
    
        this.state = {
             email: ''
        }
    }
    
    changeHandler = (event) => {
        this.setState({ [event.target.name]: event.target.value })
    }

    updateData() {
        const { email } = this.state
        const url = `http://localhost:8080/deleteuser/${email}`

        axios.delete(url)
        .then((resp) => {
            console.log("Response", resp)
        })
        .catch((error) => {
            console.log(error)
        })  
    }

    submitHandler = (event) => {
        event.preventDefault()
        this.updateData();
    }

    render () {
        const { email } = this.state

        return (
            <div>
                <h1>Delete user</h1>
                <form onSubmit={this.submitHandler}>
                    <div>
                        <p>
                            Enter Email to Delete User
                        </p>
                        <input type="email" name="email" value={email} onChange={this.changeHandler}/>
                    </div>
                    <button type="submit">Delete</button>
                </form>
            </div>
        )
    }
}

export default DeleteUser
