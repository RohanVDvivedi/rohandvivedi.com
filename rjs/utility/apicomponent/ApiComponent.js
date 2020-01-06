import React from "react";
import Loading from "../loadingcomponent/Loading.js";

export default class ApiComponent extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            api_response_received: false,
            api_response_body: null,
        }
    }
    componentDidMount() {
        fetch(window.location.origin.toString() + this.apiPath(), {
            method: this.apiMethod(),
            headers: this.apiHeaders(),
            body: this.apiBody()
        }).then(res => res.json()).then(json => {
            this.setState({
                api_response_received: true,
                api_response_body: json,
            });
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
    render() {
        if(this.state.api_response_received) {
            return this.renderAfterApiSuccess();
        } else {
            return (
                <Loading />
            );
        }
    }
    renderAfterApiSuccess() {
        throw new Error('Implement renderAfterApiSuccess method is mandatory for sub class of ApiComponent');
    }
}