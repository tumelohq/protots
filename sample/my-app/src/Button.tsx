import React, {Component} from 'react';
import {TestService} from "./idl/sample/service";

type Props = {}

class Button extends Component<Props, { number: string }> {
  constructor(props: Props) {
    super(props);
    this.state = {number: "Before click"}
  }

  onClick = () => {
    console.log("clicked");
    const client = new TestService("http://localhost:3000");
    client.TestEndpointGet({id: "1"}).then(({message}) => {
        this.setState({number: message.int64type})
      }
    ).catch(
      (e) => {
        console.log(e)
      }
    )
  }

  render() {
    return (
      <div onClick={this.onClick}>
        {this.state.number}
      </div>
    )
  }
}

export default Button;