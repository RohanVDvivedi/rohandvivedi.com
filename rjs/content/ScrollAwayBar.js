import React from "react";

export default class ScrollAwayBar extends React.Component {
    smoothScrollTo(target) {
        var curtop = window.pageYOffset || document.documentElement.scrollTop;
        var finaltop = curtop + document.getElementById(target).getBoundingClientRect().top;
        var timeDelta = 2;
        var stepDelta = 5;
        var animateScroll = function(){       
            var curtop = window.pageYOffset || document.documentElement.scrollTop;
            var sign = +1;
            if(finaltop<curtop){
                sign = -1;
            }
            var newtop = curtop + Math.min(stepDelta, Math.abs(finaltop-curtop)) * sign;
            window.scrollTo(0, newtop);
            if(Math.abs(newtop-finaltop)>0) {
                setTimeout(animateScroll, timeDelta);
            }
        };
        animateScroll();
    }
    render() {
        var container_style = {
            justifyContent: "flex-end",

            position: "fixed",
            zIndex: "1",
            top: "0px",
            left: "0px",
            width: "100%",
        };
        var component_style = {
            width: "10%",
        };
        return (
            <div id="container" class="container_style flex-row-container" style={container_style}>
                <a id="about" class="component_style" style={component_style} onClick={this.smoothScrollTo.bind(this,'AboutContent')}>About</a>
                <a id="projects" class="component_style" style={component_style} onClick={this.smoothScrollTo.bind(this,'ProjectsContent')}>Projects</a>
                <a id="contact" class="component_style" style={component_style} onClick={this.smoothScrollTo.bind(this,'ContactContent')}>Contact</a>
                <a id="social" class="component_style" style={component_style} onClick={this.smoothScrollTo.bind(this,'SocialContent')}>Social</a>
            </div>
        );
    }
}