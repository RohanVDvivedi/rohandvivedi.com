import React from "react";
import Loading from "../loadingcomponent/Loading.js";

export default class WebSocketComponent extends React.Component {
    constructor(props) {
        super(props);
        this.socket = new WebSocket(window.location.origin.toString() + this.socketPath());
        this.socket_connection_timeout = 250;
        this.state = {
            socket_connection_live: false,
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
        this.socket.onopen = () => {
            // once connected reset the timeout to its default value
            this.socket_connection_timeout = 250;

            // to render, when the socket connection is live
            this.setState({
                socket_connection_live: true,
                socket_data_body: null,
            });

            // clear any timeout event, so we do not try to reconnect
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
            response = this.respondOnMessage(message);
            if(response != null) {
                this.socket.send(message);
            }
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
        }

        // websocket onerror event listener
        this.socket.onerror = err => {
            console.error("Socket encountered error: " + err.message + "Closing socket");
            this.socket.close();
        };
    }

    respondOnMessage(message) {
        return null;
    }

    render() {
        if(this.state.socket_connection_live == true) {
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