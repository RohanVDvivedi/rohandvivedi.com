import React from "react";

// the reponse body is maintained in the state.api_response_body of your component
export default class ApiComponent extends React.Component {
    constructor(props) {
        super(props)
        // set default api response body in the satate state
        this.state = Object.assign({} ,this.state, {
            api_response_body: this.bodyDataBeforeApiFirstResponds(),
        })
    }
    // this helps to make the first api call immediately after the component has been mounted
    componentDidMount() {
        this.makeApiCallAndReRender()
    }
    // this allows the user to call set state without, modifying the state of the parent component
    // this is necessary because the ApiComponent essentially stores the response body in the state
    updateState(newState) {
    	super.setState(Object.assign({}, this.state, newState))
    }

    // the below four methods are to be used/overridden in your component if you want to change the api
    // you make have certain class variable in your component class, 
    // that will help you return appropriate path, method, headers, body for your api call
    // make sure to update those class varibales in your event handlers before calling makeApiCallAndReRender() function
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
    // api methods above

    // this is the data that will be shown to you in the render function
    // before the first api call, as made in the componentDidMount() returns with respond
    bodyDataBeforeApiFirstResponds() {
    	throw new Error('Implementing bodyDataBeforeApiFirstResponds method is mandatory for sub class of ApiComponent');
    }

    // call this function in your event handlers to make api call
    // please please do not call this function inside the render function
    makeApiCallAndReRender() {
        fetch(window.location.origin.toString() + this.apiPath(), {
            method: this.apiMethod(),
            headers: this.apiHeaders(),
            body: this.apiBody()
        }).then(res => res.json()).then(json => {
            this.updateState({api_response_body: json,});
        })
    }
}