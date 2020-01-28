import React from "react";
import AbstractContent from "./AbstractContent";

export default class ContactContent extends AbstractContent {
    constructor(props) {
        super(props);
        this.name = "contact";
        this.contentTitle = "Contact Details"
    }
    render() {
        return (
            <div id={this.getContentId()} class="content-component">
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
        );
    }
}