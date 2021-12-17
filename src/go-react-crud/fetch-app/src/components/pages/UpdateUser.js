import axios from 'axios'
import React, { Component } from 'react'

class UpdateUser extends Component {
    constructor(props) {
        super(props)
    
        this.state = {
             email: '',
             fname: ''
        }
    }
    
    changeHandler = (event) => {
        this.setState({ [event.target.name]: event.target.value })
    }

    updateData() {
        const { email, fname } = this.state
        const url = `http://localhost:8080/updateuser/${email}`

        axios.put(url, {fname})
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
        const { email, fname } = this.state

        return (
            <div>
                <h1>Update User</h1>
                <form onSubmit={this.submitHandler}>
                    <div>
                        <p>
                            Enter Email to Update name
                        </p>
                        <input type="email" name="email" value={email} onChange={this.changeHandler}/>
                    </div>
                    <div>
                        <input type="text" name="fname" value={fname} onChange={this.changeHandler}/>
                    </div>
                    <button type="submit">Update</button>
                </form>
            </div>
        )
    }
}

export default UpdateUser
