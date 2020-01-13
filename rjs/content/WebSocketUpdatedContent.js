import React from "react";

import WebSocketComponent from "../utility/websocketcomponent/WebSocketComponent.js";

export default class WebSocketUpdatedContent extends WebSocketComponent {
    socketPath() {
        return "/soc";
    }

    renderOnMessage() {
        return (
            <div>
                time on server : {this.state.socket_data_body.Time} iterated : {this.state.socket_data_body.Iterator}
            </div>
        );
    }

    onConnectionOpenResponse() {
        return {
            Message: "hello world!!"
        };
    }

    onMessageReceivedResponse(message) {
        return {
            Message: "Time was received for iteration : " + message.Iterator
        };
    }
}
