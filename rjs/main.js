import React from "react";
import ReactDOM from "react-dom";

import ScrollAwayBar from "./content/ScrollAwayBar.js";
import ContentArea from "./content/ContentArea.js";

class Root extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            pageInView: "about",
        }
    }
    setNewPageInView(pageName) {
        this.setState({
            pageInView: pageName,
        });
    }
    render() {
        return (
            <div>
                <ScrollAwayBar pageUpdaterCallback={this.setNewPageInView.bind(this)}/>
                <ContentArea navigateTo={this.state.pageInView}/>
            </div>
        );
    }
}

// ================================= >>>>

ReactDOM.render(<Root />, document.getElementById("root"));