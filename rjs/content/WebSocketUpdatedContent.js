import React from "react";

import WebSocketComponent from "../utility/websocketcomponent/WebSocketComponent.js";

export default class WebSocketUpdatedContent extends WebSocketComponent {
    socketPath() {
        return "/soc";
    }

    renderOnMessage() {
        console.log(this.state);
        return (
            <div>
                time on server : {this.state.socket_data_body.time_on_server}
            </div>
        );
    }

    onConnectionOpenResponse() {
        return {
            message: "hello"
        };
    }

    onMessageReceivedResponse(message) {
        return {
            message: "Time was received"
        };
    }
}
