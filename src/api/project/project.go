package project

import(
	"io/ioutil"
	"encoding/json"
)

type Instruction struct {
	InstructionType string
	InstructionStatement string
	InstructionReason string
	InstructionImage string
};

type Paragraph struct {
	ExplanationText string
    ExplanationImage string
};

type Project struct {
	Name string
	ProjectField string
    ProjectDescriptionShort string
    ProjectDescriptionLong string
    ProjectImage string
    GithubLink string
    YoutubeLink string
    Explanation []Paragraph
    SetupInstructions []Instruction
    RunInstruction []Instruction
    UsageInstruction []Instruction
};

var projects []Project = nil;

func GetProjectByName(Name string) (*Project) {
	if(projects == nil) {
		data, err := ioutil.ReadFile("./projects.json")
		if(err != nil) {
			return nil;
		}
		err = json.Unmarshal(data, &projects);
		if(err != nil){
			projects = nil
			return nil
		}
	}
	for _, project := range projects {
		if(project.Name == Name) {
			return &project;
		}
	}
	return nil
}