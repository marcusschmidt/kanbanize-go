package api_v1

type CreateNewTaskInput struct {
	BoardId string `json:"boardid"`
	Title string `json:"title"`
	Description string `json:"description"`
	Priority string `json:"priority"`
	Assignee string `json:"assignee"`
	Color string `json:"color"`
	Size string `json:"size"`
	Tags string `json:"tags"`
	Deadline string `json:"deadline"`
	Extlink string `json:"extlink"`
	Type string `json:"type"`
	Template string `json:"template"`
	Subtasks string `json:"subtasks"`
	Column string `json:"column"`
	Lane string `json:"lane"`
	Position string `json:"position"`
	Exceedingreason string `json:"exceedingreason"`
	Returntaskdetails string `json:"returntaskdetails"`
}

type CreateNewTaskOutput struct {
	Id int `json:"id"`
}

func (c *Client) CreateNewTask(input *CreateNewTaskInput) CreateNewTaskOutput {
	o := CreateNewTaskOutput{}

	c.doRequest("create_new_task", input, &o)

	return o
}
