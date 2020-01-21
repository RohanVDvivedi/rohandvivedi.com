import React from "react";

import AboutContent from "./contentarea/AboutContent";
import ProjectsContent from "./contentarea/ProjectsContent";
import ContactContent from "./contentarea/ContactContent";
import SocialContent from "./contentarea/SocialContent";

export default class ContentArea extends React.Component {
    render() {
        return (
            <div id="content-container">
                <AboutContent />
                <ProjectsContent />
                <ContactContent />
                <SocialContent />
            </div>
        );
    }
}