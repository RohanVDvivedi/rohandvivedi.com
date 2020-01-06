import React from "react";
import Loading from "../loadingcomponent/Loading.js";

export default class ApiComponent extends React.Component {
    constructor(props) {
        super(props)
        this.baseUrl = "http://rohandvivedi.com";
        this.state = {
            api_response_received: false,
            api_response_body: null,
        }
    }
    componentDidMount() {
        fetch(this.baseUrl + this.apiPath(), {
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
        throw new Error('Implement apiPath method');
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
        throw new Error('Implement renderAfterApiSuccess method');
    }
}