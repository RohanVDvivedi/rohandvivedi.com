package api

import (
	"net/http"
	"encoding/json"
	"rohandvivedi.com/src/data"
)

func FindProject(w http.ResponseWriter, r *http.Request) {
	project_name := r.URL.Query().Get("name");
	//project_type := r.URL.Query().Get("type");

	project := data.GetProjectByName(project_name);

	if(project != nil) {
		json, _ := json.Marshal(*data.GetOwner());
		w.Write(json);
	} else {
		w.Write([]byte(`{
					        "Name": "project-name",
					        "ProjectField": "the field : embedded, fpga, electronics, backend, frontend, system programming etc",
					        "ProjectDescriptionShort": "project description short",
					        "ProjectDescriptionLong": "project description at length, this will be shown on the inner when the user opens the project",
					        "ProjectImage" : "/img/pcb.jpeg",
					        "GithubLink": "https://github.com/RohanVDvivedi/rohandvivedi.com",
					        "YoutubeLink": "https://www.youtube.com/watch?v=fjIN9XPqp_A",
					        "YoutubeLinks": ["link1", "link2"],
					        "Explanation" : [
					            {
					                "ExplanationText": "this must be a valid html consisting only of inline contents",
					                "ExplanationImage": "link to the image that we show next to the given paragraph"
					            },
					            {
					                "ExplanationText": "this must be a valid html consisting only of inline contents",
					                "ExplanationImage": "link to the image that we show next to the given paragraph"
					            }
					        ],
					        "SetupInstructions": [
					            {
					                "InstructionType" : "to be followed or terminal instruction",
					                "InstructionStatement": "The statement of the instruction or the command to write in terminal",
					                "InstructionReason": "this is the reason, or the ouput of performing the command",
					                "InstructionImage": "there can be an instruction inage showning how to assemble the product"
					            },
					            {
					                "InstructionType" : "to be followed or terminal instruction",
					                "InstructionStatement": "The statement of the instruction or the command to write in terminal",
					                "InstructionReason": "this is the reason, or the ouput of performing the command",
					                "InstructionImage": "there can be an instruction inage showning how to assemble the product"
					            }
					        ],
					        "RunInstruction" : [
					            {
					                "InstructionType" : "to be followed or terminal instruction",
					                "InstructionStatement": "The statement of the instruction or the command to write in terminal",
					                "InstructionReason": "this is the reason, or the ouput of performing the command",
					                "InstructionImage": "there can be an instruction inage showning how to assemble the product"
					            },
					            {
					                "InstructionType" : "to be followed or terminal instruction",
					                "InstructionStatement": "The statement of the instruction or the command to write in terminal",
					                "InstructionReason": "this is the reason, or the ouput of performing the command",
					                "InstructionImage": "there can be an instruction inage showning how to assemble the product"
					            }
					        ],
					        "UsageInstruction" : [
					            {
					                "InstructionType" : "to be followed or terminal instruction",
					                "InstructionStatement": "The statement of the instruction or the command to write in terminal",
					                "InstructionReason": "this is the reason, or the ouput of performing the command",
					                "InstructionImage": "there can be an instruction inage showning how to assemble the product"
					            },
					            {
					                "InstructionType" : "to be followed or terminal instruction",
					                "InstructionStatement": "The statement of the instruction or the command to write in terminal",
					                "InstructionReason": "this is the reason, or the ouput of performing the command",
					                "InstructionImage": "there can be an instruction inage showning how to assemble the product"
					            }
					        ]
					    }`));
	}
}