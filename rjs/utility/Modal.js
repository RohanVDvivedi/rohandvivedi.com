import React from "react";

export default class Modal extends React.Component {
	constructor(props) {
		super(props)
		this.state = {
			modal_hidden: true
		}
	}
    render() {
        return (
        	<div>
        		<div class="modal-background" style = {{
        			display: (this.state.modal_hidden ? "none" : "block"),
        		}}>
        			<div class="modal">
        				<div class="modal-header">
	        				<div class="modal-title">
	        					{this.props.title}
	        				</div>
	        				<div class="close-icon">
	        				</div>
	        			<div>
	        			<div class="modal-content">
        					{this.props.children}
        				</div>
        			</div>
        		</div>
        	</div>
        );
    }
}