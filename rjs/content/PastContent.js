import React from "react";

export default class PastContent extends React.Component {
    render() {
        return (
            <div class="content-root-background content-screen-widthed content-screen-heighted flex-col-container"
                style={{justifyContent: "center",
                        alignItems: "center",}}>

                    <div class="set_sub_content_background_color generic-content-box-border"
                    	style={{ width: "60%", padding: "1.5%",}}>

                        <div style={{
                        	textAlign: "center",
							fontFamily: "lato, sans-serif",
							fontSize: "25px",
							color: "rgb(50,50,50)",
							fontWeight: "600",
							fontStyle: "italic",
                        }}>
                            Past
                        </div>
                        <div>
                            <div>Work Experience</div>
                            <a href={"https://www.oyorooms.com/"} target="_blank">OYO</a>

                            <div>
                            	<div>Job Description: Fullstack Developer contributing in php, java and Progress to Belvilla (now OYO Vacation Homes).</div>
                            	<div>OYO Vacation Homes Team, Amsterdam, Netherlands</div>
                            	<div>Jul’19-Feb’20</div>
                            </div>

                            <div>
                            	<div>Job Description: Backend Developer contributing in java and ruby tachstack to Finance Technology teams in generating reconciliation summary (accounting for commission, payments and taxations).</div>
                            	<div>Team : Finance Tech. Team, Gurgaon, India</div>
                            	<div>Jan’19-Jun’19</div>
                            </div>
                            <div>
                            	<div>Job Description: Backend Developer contributing in java and ruby tachstack to Supply Technology teams.</div>
                            	<div>Team : Supply Tech. Team, Gurgaon, India</div>
                            	<div>Aug’18-Dec’18</div>
                            </div>
                        </div>

                        <div>
                            <div>Research Experience</div>
                            <a href="https://ieeexplore.ieee.org/document/9008052" target="_blank">Flexible Processor Architecture Design</a>
                            <div>Abstract: Designing a processor that can execute any instruction set. Allowing the programmers to come up with their own custom instruction sets targetting the needs of their application.</div>
                            <div>Jul’17-Dec’17</div>
                        </div>

                        <div>
                            <div>Undergraduate Education</div>
                            <a href="https://www.bits-pilani.ac.in/hyderabad/" target="_blank">BITS Pilani</a>
                            <div>B.E. (Hons.) in Mechanical Engineering</div>
                            <div>Jul’14-Jul’18</div>
                        </div>

                    </div>

            </div>
        );
    }
}