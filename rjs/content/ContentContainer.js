import React from "react";

import ContentHash from "../ContentHash";

function capitalizeFL(string) {
	return string.charAt(0).toUpperCase() + string.slice(1)
}

export default class ContentArea extends React.Component {
    render() {
    	var ContentComp = ContentHash[this.props.selected];
        return <ContentComp selected={this.props.selected}/>;
    }
}