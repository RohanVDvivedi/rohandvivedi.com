import React from "react";
import Websocket from 'react-websocket';
import Loading from "../loadingcomponent/Loading.js";

export default class WebSocketComponent extends React.Component {
    constructor(props) {
        super(props);

        // we donoe initialize the socket variable here, 
        // because the sockkect connection is made only when we make a socket, 
        // and that should be after the component is mounted
        this.socket = null;
        this.socket_connection_timeout = 250;

        // initialize the state variables
        this.state = {
            // specifies if the connection is active right now
            socket_connection_live: false,

            // the json data the server has sent us
            socket_data_body: null,
        }
    }

    // define the path where your socket server is listening 
    socketPath() {
        throw new Error('Implementing socketPath method is mandatory for sub class of WebSocketComponent');
    }

    // once the component is mounted, make socket connection to the server
    componentDidMount() {
        this.connect();
    }

    connect() {
        // create a socket connection
        this.socket = new WebSocket(window.location.origin.toString().replace("http", "ws") + this.socketPath());

        var timeoutVar = null;

        // call this is when the socket connection is opened
        this.socket.onopen = () => {
            // clear any timeout event, so we do not try to reconnect
            if(timeoutVar != null) {
                clearTimeout(timeoutVar);
                timeoutVar = null;
            }

            // and once connected reset the timeout to its default value
            this.socket_connection_timeout = 250;

            // to render, when the socket connection is live
            this.setState({
                socket_connection_live: true,
                socket_data_body: null,
            });

            // that response you have to send on connection open, and send it
            this.sendMEssage(this.onConnectionOpenResponse(message));
        }

        // we update state on receving message from server, so the render method gets called to update the component
        this.socket.onmessage = msg => {
            // parse the message as json, and set the state, this would intern call the render method
            var message = JSON.parse(msg.data);
            this.setState({
                socket_connection_live: true,
                socket_data_body: message,
            });

            // that response you have to send on this message, and send it
            this.sendMessage(this.onMessageReceivedResponse(message));
        }

        // we update state, when the socket connection breaks or is closed
        this.socket.onclose = e => {
            // to call the render function, and render when the connection is closed, for now
            this.setState({
                socket_connection_live: false,
                socket_data_body: null,
            });

            // we will wait for twice the time we waited earlier, but lesser than or equal to 10000 milliseconds (10 seconds)
            this.socket_connection_timeout = Math.min(this.socket_connection_timeout * 2, 10000);

            // log the connection closing, and updated timeout
            console.log("Socket is closed. Reconnect will be attempted in" +  Math.min(10000 / 1000, (this.socket_connection_timeout) / 1000 ) + " millisecond." + e.reason);

            // set timeout event, to try to reconnect, when required
            timeoutVar = setTimeout(function(){
                if(this.socket == null || this.socket.readyState == WebSocket.CLOSED) {
                    this.connect();
                }
            }, this.socket_connection_timeout);
        }

        // websocket on error handler, just print the error and close socket
        this.socket.onerror = err => {
            console.error("Socket encountered error: " + err.message + "Closing socket");
            this.socket.close();
            this.setState({
                socket_connection_live: false,
                socket_data_body: null,
            });
        };
    }

    // this is the json data, we send when the connection has just been opened
    onConnectionOpenResponse() {
        return null;
    }

    // this is the json data, we send when a message has been received, byt he client
    onMessageReceivedResponse(message) {
        return null;
    }

    // call this function to send a message to server, please donot access the send method of socket directly
    sendMessage(message) {
        try {
            if(message != null) {
                this.socket.send(message);
            }
        } catch(error) {
            console.log(error);
        }
    }

    render() {
        if(this.state.socket_connection_live == true && this.state.socket_data_body != null) {
            return this.renderOnMessage();
        } else {
            return (
                <Loading />
            );
        }
    }

    renderOnMessage() {
        throw new Error('Implementing renderOnMessage method is mandatory for sub class of WebSocketComponent');
    }
}