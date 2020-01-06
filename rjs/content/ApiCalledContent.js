import React from "react";

import ApiComponent from "../utility/apicomponent/ApiComponent.js";

export default class ApiCalledComponent extends ApiComponent {
    apiPath() {
        return "/api";
    }
    renderAfterApiSuccess() {
        console.log(this.state);
        return (
            <div>
                {this.state.api_response_body.Name} is a skilled {this.state.api_response_body.Skill}
            </div>
        );
    }
}
