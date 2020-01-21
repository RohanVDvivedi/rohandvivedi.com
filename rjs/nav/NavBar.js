import React from "react";
import AboutNav from "./navbuttons/AboutNav";
import ProjectsNav from "./navbuttons/ProjectsNav";
import ContactNav from "./navbuttons/ContactNav";
import SocialNav from "./navbuttons/SocialNav";
import ChatNav from "./navbuttons/ChatNav";

export default class NavBar extends React.Component {
    smoothScrollTo(target) {
        document.getElementById("nav-container").classList.remove("hide-nav");
        var removeNavBar = function() {
            document.getElementById("nav-container").classList.add("hide-nav");
        }
        if(this.timeoutNavContainer != null){
            clearTimeout(this.timeoutNavContainer);
            this.timeoutNavContainer = null;
        }
        this.timeoutNavContainer = setTimeout(removeNavBar, 2000);

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
    onscroll() {
        var curtop = window.pageYOffset || document.documentElement.scrollTop;
        var navElement = document.getElementById("about");
        var contentElement = document.getElementById("about-content");
        var contentElementtop = curtop + contentElement.getBoundingClientRect().top;
        var contentElementbot = curtop + contentElement.getBoundingClientRect().bottom;

        navElement = document.getElementById("projects");
        contentElement = document.getElementById("projects-content");
        contentElementtop = curtop + contentElement.getBoundingClientRect().top;
        contentElementbot = curtop + contentElement.getBoundingClientRect().bottom;
        if(window.scrollY >= contentElementtop && window.scrollY < contentElementbot) {
            navElement.classList.add("active");
        } else {
            navElement.classList.remove("active");
        }

        navElement = document.getElementById("contact");
        contentElement = document.getElementById("contact-content");
        contentElementtop = curtop + contentElement.getBoundingClientRect().top;
        contentElementbot = curtop + contentElement.getBoundingClientRect().bottom;
        if(window.scrollY >= contentElementtop && window.scrollY < contentElementbot) {
            navElement.classList.add("active");
        } else {
            navElement.classList.remove("active");
        }

        navElement = document.getElementById("social");
        contentElement = document.getElementById("social-content");
        contentElementtop = curtop + contentElement.getBoundingClientRect().top;
        contentElementbot = curtop + contentElement.getBoundingClientRect().bottom;
        if(window.scrollY >= contentElementtop && window.scrollY < contentElementbot) {
            navElement.classList.add("active");
        } else {
            navElement.classList.remove("active");
        }

        document.getElementById("nav-container").classList.remove("hide-nav");
        var removeNavBar = function() {
            document.getElementById("nav-container").classList.add("hide-nav");
        }
        if(this.timeoutNavContainer != null){
            clearTimeout(this.timeoutNavContainer);
            this.timeoutNavContainer = null;
        }
        this.timeoutNavContainer = setTimeout(removeNavBar, 2000);
    }
    showNavNowAndHideAfterTms(Tms){
        // clear the previous time out
        if(this.timeoutNavContainer != null){
            clearTimeout(this.timeoutNavContainer);
            this.timeoutNavContainer = null;
        }
        // hide the nav container currently
        document.getElementById("nav-container").classList.remove("hide-nav");

        // set a timeout, to remove the NavContainer after T milliseconds
        var hideNavBar = function() {
            document.getElementById("nav-container").classList.add("hide-nav");
        }
        this.timeoutNavContainer = setTimeout(hideNavBar, Tms);
    }
    componentDidMount() {
        this.timeoutNavContainer = null;
        window.addEventListener('scroll', this.onscroll.bind(this));
        window.addEventListener('mousemove', this.showNavNowAndHideAfterTms.bind(this, 500));
    }
    render() {
        var container_style = {
            justifyContent: "flex-end",
        };
        return (
            <div id="nav-container" class="nav-container-style flex-row-container"
            style={container_style}>
                <AboutNav />
                <ProjectsNav />
                <ContactNav />
                <SocialNav />
                <ChatNav />
            </div>
        );
    }
}