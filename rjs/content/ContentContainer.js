import React from "react";

import ContentHash from "../ContentHash";

export default class ContentArea extends React.Component {
    render() {
    	var ContentComp = ContentHash[this.props.selected]["component"];
        return <ContentComp selected={this.props.selected}/>;
    }
}