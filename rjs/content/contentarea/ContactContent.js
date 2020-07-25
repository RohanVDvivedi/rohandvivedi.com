import React from "react";

export default class ContactContent extends React.Component {
    render() {
        return (
            <div id={this.props.selected.toLowerCase() + "-content"} class="content-component">
            	<div>
	                <div>
	                    <div>Formal address</div>
	                </div>
	                <div>
	                    <div>Formal phone number</div>
	                </div>
	                <div>
	                    <div>Linked In Icon</div>
	                    <div>Github</div>
	                    <div>Download curriculum vitae</div>
	                </div>
	                <div>
	                    <div>Email Icon
	                        <div>Send Email now Icon, </div>
	                        <div>also Copies Email To Clipboard</div>
	                    </div>
	                    <div>rohandvivedi@gmail.com</div>
	                </div>
	            </div>
            </div>
        );
    }
}