import React from "react";

class AboutParagraph extends React.Component {
    constructor(props) {
        super(props)
    }
    render() {
        return(
            <div style={{
                fontSize: this.props.size + "px",
                fontFamily: "lato,sans-serif",
				fontStyle: "italic",
                fontWeight: "500",
                color: "#323232"
            }}>
                    {this.props.children}
            </div>
        );
    }
}

class ColoredBoldWord extends React.Component {
    constructor(props) {
        super(props)
    }
    render() {
        return(
            <span style={{
                fontSize: "inherit",
                fontFamily: "inherit",
                color: this.props.color,
                fontWeight: "bold",
            }}>
                    {this.props.children}
            </span>
        );
    }
}

export default class PastContent extends React.Component {
    render() {
        return (
            <div class="content-root-background content-screen-widthed content-screen-heighted flex-col-container"
                style={{justifyContent: "center",
                        alignItems: "center",}}>

                    <div>
                        <div>
                            Past
                        </div>
                        <div>
                            Work Experience :
                            <a href={"https://www.oyorooms.com/"}>OYO</a>
                            Job Description: Fullstack developer contributing in php, java and Progress to Belvilla (now OYO Vacation Homes).
                            Team : OYO Vacation Homes Team, Amsterdam, Netherlands      Jul’19-Feb’20
                            Job Description: Backend developer contributing in java and ruby tachstack to Finance Technology teams in generating reconciliation summary (accounting for commission, payments and taxations).
                            Team : Finance Tech. Team, Gurgaon, India                   Jan’19-Jun’19
                            Team : Supply Tech. Team, Gurgaon, India                    Aug’18-Dec’18
                            Research Experience :
                            <a href={"https://ieeexplore.ieee.org/document/9008052"}>Flexible Processor Architecture Design</a>
                            Description: Designing a processor that can execute any instruction set. Allowing the programmers to come up with their own custom instruction sets targetting the needs of their application.
                        </div>
                    </div>

            </div>
        );
    }
}