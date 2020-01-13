import React from "react";
import ReactDOM from "react-dom";

import ApiCalledContent from "./content/ApiCalledContent.js";
import WebSocketUpdatedContent from "./content/WebSocketUpdatedContent.js";

class Root extends React.Component {
    render() {
        return (
            <div>
                Hello World!!
                <ApiCalledContent />
                <WebSocketUpdatedContent />
            </div>
        );
    }
}

// ================================= >>>>

ReactDOM.render(<Root />, document.getElementById("root"));