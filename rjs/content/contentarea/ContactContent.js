import React from "react";

import Icon from "../../utility/Icon";
import CopyToClipboard from "../../utility/Clipboard";

export default class ContactContent extends React.Component {
    render() {
        return (
        	<div class="content-root-background content-screen-widthed content-screen-heighted flex-col-container"
        		style={{justifyContent: "center",
                        alignItems: "center",}}>
                    <div class="flex-row-container">
                    	<a href="https://github.com/RohanVDvivedi" target="_blank" style={{display: "block",}}>
	                        <Icon height="35px" width="35px" iconPath="/icon/github.png"/>
	                    </a>
	                    <a href="https://www.linkedin.com/in/rohan-dvivedi-ab3014128/" target="_blank" style={{display: "block",}}>
	                        <Icon height="35px" width="35px" iconPath="/icon/linkedin.png"/>
	                    </a>
	                </div>
	                <div class="flex-row-container">
	                    <a href="https://www.youtube.com/channel/UCgn_REjbUH2Dm8CaOXvajJg?view_as=subscriber" target="_blank" style={{display: "block",}}>
	                        <Icon height="35px" width="35px" iconPath="/icon/youtube.png"/>
	                    </a>
	                </div>
	                <div class="flex-row-container">
	                    <a href="https://www.facebook.com/rohan.dvivedi.961" target="_blank" style={{display: "block",}}>
	                        <Icon height="35px" width="35px" iconPath="/icon/facebook.png"/>
	                    </a>
			        </div>
                    <div class="flex-row-container">
			            <a href="https://drive.google.com/file/d/12hE5q84en4QAsGkIlOcPEjlFL4kzgHxw/view?usp=sharing" target="_blank" style={{display: "block",}}>
	                        <Icon height="35px" width="35px" iconPath="/icon/download.png"/>
	                    </a>
	                    <a href="https://drive.google.com/file/d/12hE5q84en4QAsGkIlOcPEjlFL4kzgHxw/view?usp=sharing" target="_blank" style={{display: "block",}}>
	                        <Icon height="35px" width="35px" iconPath="/icon/pdf.png"/>
	                    </a>
			        </div>
			        <div class="flex-row-container">
			        	<div>rohandvivedi.gamil.com</div>
			        	<a onClick={()=>{CopyToClipboard("rohandvivedi@gmail.com")}} style={{display: "block",}}>
	                        <Icon height="35px" width="35px" iconPath="/icon/mail.png"/>
	                    </a>
			        	<a href="https://mail.google.com/mail/u/0/?view=cm&fs=1&tf=1&to=rohandvivedi@gmail.com&su=Hi&body=Dear%20Rohan,%0D%0A%0D%0ALooking%20forward%20to%20hearing%20from%20you,%0D%0AYours%20sincerely." target="_blank" style={{display: "block",}}>
	                        <Icon height="35px" width="35px" iconPath="/icon/gmail.png"/>
	                    </a>
			        </div>
		            <div class="flex-row-container">
		            	<div>+91-6352496059</div>
		            </div>
	        </div>
        );
    }
}