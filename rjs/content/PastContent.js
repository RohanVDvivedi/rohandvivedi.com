import React from "react";

import ApiComponent from "../utility/ApiComponent";

class TimedEvent extends React.Component {
	render() {
		return (<div class="flex-row-container" style={{justifyContent: "space-between"}}>
					<div style={{fontSize:"17px", fontWeight: "550"}}>{this.props.revent}</div>
					<div style={{fontStyle:"italic"}}>{this.props.rtime}</div>
				</div>);
	}
}

export default class PastContent extends ApiComponent {
	apiPath() {
        return "/api/owner?get_pasts=true";
    }
    bodyDataBeforeApiFirstResponds() {
    	return {Fname:"Firstname",Lname:"Lastname",Email:"loading email id","Socials":null,"Pasts":[]};
    }
    render() {
    	var owner = this.state.api_response_body;
        return (
            <div class="content-root-background content-screen-widthed content-screen-heighted flex-col-container"
                style={{justifyContent: "center",
                        alignItems: "center",}}>

                    <div class="set_sub_content_background_color generic-content-box-border"
                    	style={{ width: "60%", padding: "1.5%",
                    		color: "var(--font_color_lighter)",
                    		}}>

                        <div style={{
                        	textAlign: "center",
							fontFamily: "lato, sans-serif",
							fontSize: "25px",
							fontWeight: "700",
							fontStyle: "italic",
							paddingBottom: "10px"
                        }}>
                            Past
                        </div>

                        <div>
	                        <div style={{marginTop: "0px"}}>
	                            <div style={{fontSize:"20px",fontWeight: "600"}}>SDE1 (Software Development Engineer I)</div>
	                            <div  style={{marginLeft:"10px"}}>
		                        	<a href="https://www.oyorooms.com/" target="_blank" style={{fontSize:"18px"}}>OYO</a>
		                            <div>
			                            <TimedEvent revent="OYO Vacation Homes, Amsterdam, Netherlands"
			                            			rtime="Jul’19-Feb’20" />
			                            <TimedEvent revent="Finance Tech. Team, Gurgaon, India"
			                            			rtime="Dec’18-Jun’19" />
			                            <TimedEvent revent="Supply Tech. Team, Gurgaon, India"
			                            			rtime="Aug’18-Dec’18" />
			                        </div>
		                        </div>
	                        </div>

	                        <div style={{marginTop: "15px"}}>
	                            <div style={{fontSize:"20px",fontWeight: "600"}}>Thesis</div>
	                            <div  style={{marginLeft:"10px"}}>
	                            	<TimedEvent revent="Flexible Processor Architecture Design"
			                            			rtime="Jul’17-Dec’17" />
		                            <div>DOI: <a href="https://ieeexplore.ieee.org/document/9008052" target="_blank">10.1109/DISCOVER47552.2019.9008052</a></div>
		                            <div>Authors: D. R. Vipulkumar, P. V. Bhanu and J. Soumya</div>
		                        </div>
	                        </div>

	                        <div style={{marginTop: "15px"}}>
	                            <div style={{fontSize:"20px",fontWeight: "600"}}>Education</div>
		                        <div  style={{marginLeft:"5px"}}>
		                            <a href="https://www.bits-pilani.ac.in/hyderabad/" target="_blank">BITS Pilani</a>
		                            <TimedEvent revent="B.E. (Hons.) in Mechanical Engineering"
			                            			rtime="Jul’14-Jul’18" />
		                        </div>
	                        </div>
	                    </div>

                    </div>

            </div>
        );
    }
}

    /*
    Past
Work Experience
OYO
Position: SDE1 (Software Development Engineer I)
OYO Vacation Homes, Amsterdam, Netherlands
Jul’19-Feb’20
Finance Tech. Team, Gurgaon, India
Dec’18-Jun’19
Supply Tech. Team, Gurgaon, India
Aug’18-Dec’18
Job Description: SDE1, Backend Developer contributing in java (spring boot) and ruby on rails techstack to Finance and Supply Technology teams at OYO. I also contributed in website revamp and feature development for OYO Vacation Homes' subsidiary entity Belvilla B.V..
Research Experience
Flexible Processor Architecture Design
Jul’17-Dec’17
DOI: 10.1109/DISCOVER47552.2019.9008052
Authors : D. R. Vipulkumar, P. V. Bhanu and J. Soumya
Abstract problem statement: To Design a processor that can execute any instruction set (custom designed or commercial). To try and encourage programmers to come up with their own custom instructions to target higher efficiencies for their applications.
Undergraduate Education
BITS Pilani
B.E. (Hons.) in Mechanical Engineering
Jul’14-Jul’18
Description: I graduated from BITS Pilani, Hyderabad Campus, majoring in Mechanical Engineering. My exposure to Computer Science, during my undergraduate education, gave me confidence to explore the field further which led me to learn to analyze problems and code their way out.
*/