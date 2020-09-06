import React from "react";

export default class ContactForm extends React.Component {
	constructor(props) {
		super(props)
		this.state = {
			subject: "",
			body: "",
			api_resp: null
		}
	}
	onSubjectChange(e) {
		this.setState(Object.assign({}, this.state, {subject: e.target.value}));
	}
	onBodyChange(e) {
		this.setState(Object.assign({}, this.state, {body: e.target.value}));
	}
	sendButtonClicked() {
		this.setState(Object.assign({}, this.state, {api_resp: null}));
		fetch("http://localhost/api/anon_mails", {
			method: "post",
			body: JSON.stringify({Subject:this.state.subject,Body:this.state.body})
		}).then(res => res.json()).then(json => {
			console.log(json)
			this.setState(Object.assign({}, this.state, {api_resp: json}));
		})
	}
	render() {
		var maxLengthSubject = 128;
		var maxLengthBody = 1024;
		return (
			<div class="contact-form-root">
				<div class="flex-col-container">

					<div class="input-text-group flex-col-container">
						<label for="contact-subject" class="text-label">Subject</label>
						<input id="contact-subject" class="text-box" type="text" 
						placeholder="Subject" maxlength={maxLengthSubject} 
						onChange={this.onSubjectChange.bind(this)} value={this.state.subject}/>
						<div class="remaining-chars">{maxLengthSubject - this.state.subject.length}</div>
					</div>

					<div class="input-text-group flex-col-container">
						<label for="contact-body" class="text-label">Body</label>
						<textarea id="contact-body" class="text-box" 
						placeholder="Elaborate your query here..." maxlength={maxLengthBody} 
						onChange={this.onBodyChange.bind(this)} value={this.state.body} rows="5"/>
						<div class="remaining-chars">{maxLengthBody - this.state.body.length}</div>
					</div>

					<div class="input-text-group">
						<div class="text-label">Note:</div>
						<ul>
							<li>This contact form will generate an anonymous mail and send it to me on my gmail account.</li>
							<li class="removable-screen-375-lesser">The service is throttled to allow just 3 anonymous mails per user per device per 48 hours, for security reasons.</li>
						</ul>
					</div>

					<div>
						<div class="send-button generic-content-box-hovering-emboss-border" 
							onClick={this.sendButtonClicked.bind(this)}>
							Send
						</div>
					</div>

					{this.state.api_resp != null ?
					(<div class={this.state.api_resp.Success ? "success-msg" : "error-msg"}>
						{this.state.api_resp.Message}
					</div>) : "" }

				</div>
			</div>
		);
	}
}