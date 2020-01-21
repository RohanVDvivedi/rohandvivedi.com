import React from "react";

import AbstractNav from "./AbstractNav";

export default class AboutNav extends AbstractNav {
    smoothScrollTo(target) {
        // turn on to show the nav container
        document.getElementById("nav-container").classList.remove("hide-nav");
        var removeNavBar = function() {
            document.getElementById("nav-container").classList.add("hide-nav");
        }
        if(this.timeoutNavContainer != null){
            clearTimeout(this.timeoutNavContainer);
            this.timeoutNavContainer = null;
        }
        this.timeoutNavContainer = setTimeout(removeNavBar, 2000);

        // to animate and slide to the correct position in the website page
        var curtop = window.pageYOffset || document.documentElement.scrollTop;
        var finaltop = curtop + document.getElementById(target).getBoundingClientRect().top;
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
    render() {
        return (
            <a id="about" class="nav-button-style"
            onClick={this.smoothScrollTo.bind(this,'about-content')}>
                About
            </a>
        );
    }
}