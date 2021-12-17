import axios from 'axios'
import React, { Component } from 'react'

class CreateUser extends Component {

    constructor(props) {
        super(props)
    
        this.state = {
             fname:'',
             lname:'',
             age: '',
             mobileno: '',
             dob: '',
             email: '',
             address: ''
        }
    }

    changeHandler = (event) => {
        let value = event.target.value
        if (event.target.name === "age"){
            value = parseInt(event.target.value,10)
        }
        this.setState({ [event.target.name]: value })
    }

    sendData() {
        const url = "http://localhost:8080/createuser";

        // const { user } = this.state
        console.log("Testing:", this.state)
        
        axios.post( url, this.state)
        .then((resp) => {
            console.log("Response:", resp)
            console.log("Data:", resp.data)
        })
        .then((data) => {
            console.log("from data:", data)
        })
        // .then((data) => {
        //     console.log(data);
        // })
        .catch((error) => {
            console.log(error)
        })
    }

    submitHandler = (event) => {
        event.preventDefault()
        //console.log(this.state)
        this.sendData();
    }

    render () {
        const { fname, lname, age, mobileno, dob, email, address }  = this.state

        return (
            <div className="container">
                <h1 className="text-center text-primary">Create User</h1>
                <form onSubmit={this.submitHandler} className="was-validated"> 
                    <div className="mb-3">
                        <label htmlFor="fName" className="form-label">First name</label>
                        <input type="text" name="fname" value={fname} required className="form-control" id="fname" onChange={this.changeHandler} />
                        <div class="valid-feedback">looks good</div>
                        <div class="invalid-feedback">Please fill out this field</div>
                    </div>
                    <div className="mb-3">
                        <label htmlFor="lName" className="form-label">Last name</label>
                        <input type="text" name="lname" value={lname} required className="form-control" id="age" onChange={this.changeHandler}/>
                        <div class="valid-feedback">looks good</div>
                        <div class="invalid-feedback">Please fill out this field</div>
                    </div>
                    <div className="mb-3">
                        <label htmlFor="age" className="form-label">Age</label>
                        <input type="number" name="age" value={age} required className="form-control" id="age" onChange={this.changeHandler}/>
                        <div class="valid-feedback">looks good</div>
                        <div class="invalid-feedback">Please fill out this field</div>
                    </div>
                    <div className="mb-3">
                        <label htmlFor="MobileNo" className="form-label">Phone number</label>
                        <input type="text" name="mobileno" value={mobileno} required className="form-control" id="mobileno" onChange={this.changeHandler}/>
                        <div class="valid-feedback">looks good</div>
                        <div class="invalid-feedback">Please fill out this field</div>
                    </div>
                    <div className="mb-3">
                        <label htmlFor="DOB" className="form-label">Date of birth</label>
                        <input type="text" name="dob" value={dob} required className="form-control" id="dob" onChange={this.changeHandler}/>
                        <div class="valid-feedback">looks good</div>
                        <div class="invalid-feedback">Please fill out this field</div>
                    </div>
                    <div className="mb-3">
                        <label htmlFor="Email" className="form-label">Email address</label>
                        <input type="email" name="email" value={email} required className="form-control" id="email" onChange={this.changeHandler}/>
                        <div class="valid-feedback">looks good</div>
                        <div class="invalid-feedback">Please fill out this field</div>
                    </div>
                    <div className="mb-3">
                        <label htmlFor="Address" className="form-label">Address</label>
                        <input type="text" name="address" value={address} required className="form-control" id="address" onChange={this.changeHandler}/>
                        <div class="valid-feedback">looks good</div>
                        <div class="invalid-feedback">Please fill out this field</div>
                    </div>
                    <button type="submit" className="btn btn-primary">Submit</button>
                </form>
            </div>
        )
    }
}

export default CreateUser