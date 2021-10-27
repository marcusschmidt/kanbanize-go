package api_v1

type GetBoardActivitiesInput struct {
	BoardId        string `json:"boardid"`
	FromDate       string `json:"fromdate"`
	ToDate         string `json:"todate"`
	Page           string `json:"page"`
	Resultsperpage string `json:"resultsperpage"`
	Author         string `json:"author"`
	Eventtype      string `json:"eventtype"`
	Textformat     string `json:"textformat"`
}

type GetBoardActivitiesOutput struct {
	Allactivities int `json:"allactivities"`
	Page          int `json:"page"`
	Activities    []struct {
		Author    string `json:"author"`
		Event     string `json:"event"`
		Text      string `json:"text"`
		Date      string `json:"date"`
		Taskid    string `json:"taskid"`
		Commentid string `json:"commentid,omitempty"`
	} `json:"activities"`
}

func (c *Client) GetBoardActivities(input *GetBoardActivitiesInput) GetBoardActivitiesOutput {
	o := GetBoardActivitiesOutput{}

	c.doRequest("get_board_activities", input, &o)

	return o
}

type GetBoardStructureInput struct {
	BoardId string `json:"boardid"`
}

type GetBoardStructureOutput struct {
	Columns []struct {
		Position    string `json:"position"`
		Lcname      string `json:"lcname"`
		Description string `json:"description"`
		Tasksperrow int    `json:"tasksperrow"`
		Section     string `json:"section"`
		Path        string `json:"path"`
		Lcid        string `json:"lcid"`
		Flowtype    string `json:"flowtype"`
		Color       string `json:"color"`
	} `json:"columns"`
	Lanes []struct {
		Lcname      string `json:"lcname"`
		Color       string `json:"color"`
		Description string `json:"description"`
	} `json:"lanes"`
}

func (c *Client) GetBoardStructure(input *GetBoardStructureInput) GetBoardStructureOutput {
	b := GetBoardStructureOutput{}

	c.doRequest("get_board_structure", input, &b)

	return b
}

type GetFullBoardStructureInput struct {
	BoardId string `json:"boardid"`
}

type GetFullBoardStructureOutput struct {
	Columns []struct {
		Position    string `json:"position"`
		Lcname      string `json:"lcname"`
		Section     string `json:"section"`
		Path        string `json:"path"`
		Description string `json:"description"`
		Lcid        string `json:"lcid"`
		Flowtype    string `json:"flowtype"`
		Tasksperrow int    `json:"tasksperrow"`
	} `json:"columns"`
	Lanes []struct {
		Position    string `json:"position"`
		Lcname      string `json:"lcname"`
		Path        string `json:"path"`
		Description string `json:"description"`
		Lcid        string `json:"lcid"`
		Flowtype    string `json:"flowtype"`
		Color       string `json:"color"`
	} `json:"lanes"`
}

func (c *Client) GetFullBoardStructure(input *GetFullBoardStructureInput) GetFullBoardStructureOutput {
	b := GetFullBoardStructureOutput{}

	c.doRequest("get_full_board_structure", input, &b)

	return b
}

type GetBoardSettingsInput struct {
	BoardId string `json:"boardid"`
}

type GetBoardSettingsOutput struct {
	Avatars      []interface{} `json:"avatars"`
	Usernames    []string      `json:"usernames"`
	Templates    []string      `json:"templates"`
	Types        []string      `json:"types"`
	CustomFields []interface{} `json:"customFields"`
}

type RawOutput struct {
	data string
}

func (c *Client) GetBoardSettings(input *GetBoardSettingsInput) string {
	o := ""

	c.doRequest("get_board_settings", input, &o)

	return o
}
