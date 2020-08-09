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
							paddingBottom: "10px"
                        }}>
                            Past
                        </div>

                        <div>
	                        <div>
	                            <div>Work Experience</div>
	                            <a href={"https://www.oyorooms.com/"} target="_blank">OYO</a>
	                            <div>Position: SDE1 (Software Development Engineer I)</div>

	                            <div class="flex-row-container" style={{justifyContent: "space-between"}}>
	                           		<div>OYO Vacation Homes, Amsterdam, Netherlands</div>
	                           		<div style={{fontStyle:"italic"}}>Jul’19-Feb’20</div>
	                           	</div>
	                           	<div class="flex-row-container" style={{justifyContent: "space-between"}}>
	                           		<div>Finance Tech. Team, Gurgaon, India</div>
	                           		<div style={{fontStyle:"italic"}}>Dec’18-Jun’19</div>
	                           	</div>
	                           	<div class="flex-row-container" style={{justifyContent: "space-between"}}>
	                           		<div>Supply Tech. Team, Gurgaon, India</div>
	                           		<div style={{fontStyle:"italic"}}>Aug’18-Dec’18</div>
	                           	</div>
	                            <div>
	                            Job Description: 
	                            SDE1, Backend Developer contributing in java (spring boot) and ruby on rails techstack to Finance and Supply Technology teams at OYO.
	                            I also contributed in website revamp and feature development for <a href="https://oyovacationhomes.com/" target="_blank">OYO Vacation Homes'</a> subsidiary entity <a href="https://www.belvilla.com/" target="_blank">Belvilla B.V.</a>.
	                            </div>
	                        </div>

	                        <div style={{marginTop: "5px"}}>
	                            <div>Research Experience</div>
	                            <div class="flex-row-container" style={{justifyContent: "space-between"}}>
	                            	<div>Flexible Processor Architecture Design</div>
	                            	<div style={{fontStyle:"italic"}}>Jul’17-Dec’17</div>
	                            </div>
	                            <div>DOI: <a href="https://ieeexplore.ieee.org/document/9008052" target="_blank">10.1109/DISCOVER47552.2019.9008052</a></div>
	                            <div>Authors : D. R. Vipulkumar, P. V. Bhanu and J. Soumya</div>
	                            <div>Abstract problem statement: To Design a processor that can execute any instruction set (custom designed or commercial).
	                            To try and encourage programmers to come up with their own custom instructions to target higher efficiencies for their applications.</div>
	                        </div>

	                        <div style={{marginTop: "5px"}}>
	                            <div>Undergraduate Education</div>

	                            <a href="https://www.bits-pilani.ac.in/hyderabad/" target="_blank">BITS Pilani</a>

	                            <div class="flex-row-container" style={{justifyContent: "space-between"}}>
	                            	<div>B.E. (Hons.) in Mechanical Engineering</div>
	                            	<div style={{fontStyle:"italic"}}>Jul’14-Jul’18</div>
	                            </div>

	                            <div>
	                            Description: I graduated from BITS Pilani, Hyderabad Campus, majoring in Mechanical Engineering.
	                            My exposure to Computer Science, during my undergraduate education, gave me confidence to explore the field further which led me to learn to analyze problems and code their way out.
	                            </div>
	                        </div>
	                    </div>

                    </div>

            </div>
        );
    }
}