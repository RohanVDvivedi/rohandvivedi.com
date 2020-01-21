import React from "react";

export default class AbstractNav extends React.Component {
    constructor(props) {
        super(props);
    }
    smoothScrollToContent() {
        // to animate and slide to the correct position in the website page
        var curtop = window.pageYOffset || document.documentElement.scrollTop;
        var finaltop = curtop + document.getElementById(this.name + "-content").getBoundingClientRect().top;
        var timeDelta = 2;
        var stepDelta = 10;
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
    activateIfMyContentIsInFocus() {
        var curtop = window.pageYOffset || document.documentElement.scrollTop;
        var navElement = document.getElementById(this.name + "-nav");
        var contentElement = document.getElementById(this.name + "-content");
        var contentElementtop = curtop + contentElement.getBoundingClientRect().top;
        var contentElementbot = curtop + contentElement.getBoundingClientRect().bottom;
        if(window.scrollY >= contentElementtop && window.scrollY < contentElementbot) {
            navElement.classList.add("active");
        } else {
            navElement.classList.remove("active");
        }
    }
    componentDidMount() {
        window.addEventListener('scroll', this.activateIfMyContentIsInFocus.bind(this));
    }
    render() {
        return (
            <a id={this.name + "-nav"} class="nav-button-style"
            onClick={this.smoothScrollToContent.bind(this)}>
                {this.navTitle}
            </a>
        );
    }
}