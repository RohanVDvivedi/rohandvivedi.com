import React from "react";

export default class ApiComponent extends React.Component {
    constructor(props) {
        super(props)
        this.state = Object.assign({} ,this.state, {
            api_response_body: this.bodyDataBeforeApiFirstResponds(),
        })
    }
    componentDidMount() {
        fetch(window.location.origin.toString() + this.apiPath(), {
            method: this.apiMethod(),
            headers: this.apiHeaders(),
            body: this.apiBody()
        }).then(res => res.json()).then(json => {
            this.setState(Object.assign({},this.state,{api_response_body: json,}));
        })
    }
    apiPath() {
        throw new Error('Implementing apiPath method is mandatory for sub class of ApiComponent');
    }
    apiMethod() {
        "get"
    }
    apiHeaders() {
        return {}
    }
    apiBody() {
        return null
    }
    bodyDataBeforeApiFirstResponds() {
    	throw new Error('Implementing apiPath method is mandatory for sub class of ApiComponent');
    }
}