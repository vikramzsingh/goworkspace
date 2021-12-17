import React, { Component } from "react";

class FetchUser extends Component {
  state = {
    loading: false,
    person: null,
  };

  callData() {
    this.setState({ loading: true });
    const url = "http://localhost:8080";

    //const response = 
      fetch(url)
      .then((resp) => resp.json())
      .then((data) => {
        this.setState({ users: data });
      })
      .catch((e) => {
        console.error(e);
      })
      .finally(() => {
        this.setState({ loading: false });
      });
  }

  componentDidMount() {
    this.callData();
  }

  render() {
    const { loading, users } = this.state;

    if (loading) {
      return <div>Loading ...</div>;
    }

    return (
      <div>
        {users &&
          users.map((user, key) => {
            return (
              <div key={key}>
                <div>{user.name}</div>
                <div>{user.age}</div>
                <div>{user.address}</div>
              </div>
            );
          })}
      </div>
    );
  }
}

export default FetchUser;
