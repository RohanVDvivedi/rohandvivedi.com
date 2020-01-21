import React from "react";

export default class AbstractNav extends React.Component {
    constructor(props) {
        super(props);
    }
    getNavId() {
        return this.name + "-nav"
    }
    getContentId() {
        return this.name + "-content"
    }
    smoothScrollToContent() {
        // to animate and slide to the correct position in the website page
        var curtop = window.pageYOffset || document.documentElement.scrollTop;
        var finaltop = curtop + document.getElementById(this.getContentId()).getBoundingClientRect().top;
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
    activateNavIfContentIsInFocus() {
        var curtop = window.pageYOffset || document.documentElement.scrollTop;
        var navElement = document.getElementById(this.getNavId());
        var contentElement = document.getElementById(this.getContentId());
        var contentElementtop = curtop + contentElement.getBoundingClientRect().top;
        var contentElementbot = curtop + contentElement.getBoundingClientRect().bottom;
        if(window.scrollY >= contentElementtop && window.scrollY < contentElementbot) {
            navElement.classList.add("active");
        } else {
            navElement.classList.remove("active");
        }
    }
    componentDidMount() {
        window.addEventListener('scroll', this.activateNavIfContentIsInFocus.bind(this));
    }
    render() {
        return (
            <a id={this.getNavId()} class="nav-button-style"
            onClick={this.smoothScrollToContent.bind(this)}>
                {this.navTitle}
            </a>
        );
    }
}