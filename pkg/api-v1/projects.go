package api_v1

type GetProjectsAndBoardsOutput struct {
	Projects []struct {
		Name   string `json:"name"`
		Id     string `json:"id"`
		Boards []struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"boards"`
	} `json:"projects"`
}

func (c *Client) GetProjectsAndBoards() GetProjectsAndBoardsOutput {
	getProjectsAndBoardsOutput := GetProjectsAndBoardsOutput{}

	c.doRequest("get_projects_and_boards", struct {}{}, &getProjectsAndBoardsOutput)

	return getProjectsAndBoardsOutput
}
