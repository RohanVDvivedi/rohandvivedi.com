import React from "react";

import Icon from "../../utility/Icon";
import CopyToClipboard from "../../utility/Clipboard";

class ContactSubContent extends React.Component {
	render() {
		var styl = {
				padding: "10px",
			};
		if(this.props.last == null || this.props.last != "true")
		{
			styl["borderBottom"] = "4px solid #c3c3c3";
		}
		return (
			<div style={styl}>
				<div style={{
					fontFamily: "Capriola, Helvetica, sans-serif",
				}}>
					{this.props.title}
				</div>
				<div class="flex-row-container"
					style={{
						justifyContent: "space-evenly",
						alignItems: "center",
						padding: "10px",
					}}>
			    	{this.props.children}
				</div>
			</div>);
	}
}

export default class ContactContent extends React.Component {
    render() {
        return (
        	<div class="content-root-background content-screen-widthed content-screen-heighted flex-col-container"
        		style={{justifyContent: "center",
                        alignItems: "center",}}>

                    <div class="set_sub_content_background_color flex-col-container"
                    style={{
                    	padding: "10px",
                    }}>
                    	<ContactSubContent title="CV">
		                    <a href="https://drive.google.com/file/d/12hE5q84en4QAsGkIlOcPEjlFL4kzgHxw/view?usp=sharing" target="_blank" style={{display: "block",}}>
		                        <Icon height="35px" width="35px" iconPath="/icon/pdf.png"/>
		                    </a>
				        </ContactSubContent>
	                    <ContactSubContent title="Online Presence">
	                    	<a href="https://github.com/RohanVDvivedi" target="_blank" style={{display: "block",}}>
		                        <Icon height="35px" width="35px" iconPath="/icon/github.png"/>
		                        <div>Github</div>
		                    </a>
		                    <a href="https://www.youtube.com/channel/UCgn_REjbUH2Dm8CaOXvajJg?view_as=subscriber" target="_blank" style={{display: "block",}}>
		                        <Icon height="35px" width="35px" iconPath="/icon/youtube.png"/>
		                        <div>Youtube</div>
		                    </a>
		                    <a href="https://www.linkedin.com/in/rohan-dvivedi-ab3014128/" target="_blank" style={{display: "block",}}>
		                        <Icon height="35px" width="35px" iconPath="/icon/linkedin.png"/>
		                        <div>Linkedin</div>
		                    </a>
		                    <a href="https://www.facebook.com/rohan.dvivedi.961" target="_blank" style={{display: "block",}}>
		                        <Icon height="35px" width="35px" iconPath="/icon/facebook.png"/>
		                        <div>Facebook</div>
		                    </a>
				        </ContactSubContent>
				        <ContactSubContent title="Email" last="true">
				        	<a href="#" onClick={()=>{CopyToClipboard("rohandvivedi@gmail.com")}} style={{display: "block",}}>
		                        <div style={{display:"inline-block"}}>rohandvivedi@gmail.com</div>
		                    </a>
		                    <div> OR </div>
				        	<a href="mailto:rohandvivedi@gmail.com" style={{display: "block",}}>
		                        <Icon height="35px" width="35px" iconPath="/icon/mail.png"/>
		                    </a>
				        </ContactSubContent>

			        </div>
	        </div>
        );
    }
}