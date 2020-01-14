import React from "react";
import ReactDOM from "react-dom";

import AlwaysStayBar from "./content/AlwaysStayBar.js";
import ScrollAwayBar from "./content/ScrollAwayBar.js";
import ContentNavigationSideBar from "./content/ContentNavigationSideBar.js";
import ContentArea from "./content/ContentArea.js";

class Root extends React.Component {
    render() {
        return (
            <div>
                <AlwaysStayBar />
                <ScrollAwayBar />
                <ContentNavigationSideBar />
                <ContentArea />
            </div>
        );
    }
}

// ================================= >>>>

ReactDOM.render(<Root />, document.getElementById("root"));