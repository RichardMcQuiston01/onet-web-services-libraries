package onet

import "context"

// OnetDataElement is an O*NET data attribute element (ability, skill, knowledge, etc.).
type OnetDataElement struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OnetDataElementList is the list of elements for a given O*NET data attribute.
type OnetDataElementList struct {
	Element []OnetDataElement `json:"element"`
}

// OnetDataAttributeOccupations is returned by the O*NET data detail endpoints;
// it lists occupations that have the given element.
type OnetDataAttributeOccupations struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Occupation []Occupation `json:"occupation"`
}

func (c *Client) onetDataList(ctx context.Context, attribute string) (*OnetDataElementList, error) {
	var out OnetDataElementList
	return &out, c.do(ctx, "online/onet_data/"+attribute+"/", nil, &out)
}

func (c *Client) onetDataDetail(ctx context.Context, attribute, id string) (*OnetDataAttributeOccupations, error) {
	var out OnetDataAttributeOccupations
	return &out, c.do(ctx, "online/onet_data/"+attribute+"/"+id, nil, &out)
}

// OnlineONETAbilities returns the list of all O*NET ability elements.
func (c *Client) OnlineONETAbilities(ctx context.Context) (*OnetDataElementList, error) {
	return c.onetDataList(ctx, "abilities")
}

// OnlineONETAbility returns occupations associated with the given ability element ID.
func (c *Client) OnlineONETAbility(ctx context.Context, id string) (*OnetDataAttributeOccupations, error) {
	return c.onetDataDetail(ctx, "abilities", id)
}

// OnlineONETInterests returns the list of all O*NET interest elements.
func (c *Client) OnlineONETInterests(ctx context.Context) (*OnetDataElementList, error) {
	return c.onetDataList(ctx, "interests")
}

// OnlineONETInterest returns occupations associated with the given interest element ID.
func (c *Client) OnlineONETInterest(ctx context.Context, id string) (*OnetDataAttributeOccupations, error) {
	return c.onetDataDetail(ctx, "interests", id)
}

// OnlineONETKnowledge returns the list of all O*NET knowledge elements.
func (c *Client) OnlineONETKnowledge(ctx context.Context) (*OnetDataElementList, error) {
	return c.onetDataList(ctx, "knowledge")
}

// OnlineONETKnowledgeItem returns occupations associated with the given knowledge element ID.
func (c *Client) OnlineONETKnowledgeItem(ctx context.Context, id string) (*OnetDataAttributeOccupations, error) {
	return c.onetDataDetail(ctx, "knowledge", id)
}

// OnlineONETBasicSkills returns the list of all O*NET basic skill elements.
func (c *Client) OnlineONETBasicSkills(ctx context.Context) (*OnetDataElementList, error) {
	return c.onetDataList(ctx, "skills_basic")
}

// OnlineONETBasicSkill returns occupations associated with the given basic skill element ID.
func (c *Client) OnlineONETBasicSkill(ctx context.Context, id string) (*OnetDataAttributeOccupations, error) {
	return c.onetDataDetail(ctx, "skills_basic", id)
}

// OnlineONETCrossFunctionalSkills returns the list of all O*NET cross-functional skill elements.
func (c *Client) OnlineONETCrossFunctionalSkills(ctx context.Context) (*OnetDataElementList, error) {
	return c.onetDataList(ctx, "skills_cf")
}

// OnlineONETCrossFunctionalSkill returns occupations associated with the given cross-functional skill element ID.
func (c *Client) OnlineONETCrossFunctionalSkill(ctx context.Context, id string) (*OnetDataAttributeOccupations, error) {
	return c.onetDataDetail(ctx, "skills_cf", id)
}

// OnlineONETWorkActivities returns the list of all O*NET work activity elements.
func (c *Client) OnlineONETWorkActivities(ctx context.Context) (*OnetDataElementList, error) {
	return c.onetDataList(ctx, "work_activities")
}

// OnlineONETWorkActivity returns occupations associated with the given work activity element ID.
func (c *Client) OnlineONETWorkActivity(ctx context.Context, id string) (*OnetDataAttributeOccupations, error) {
	return c.onetDataDetail(ctx, "work_activities", id)
}

// OnlineONETWorkContext returns the list of all O*NET work context elements.
func (c *Client) OnlineONETWorkContext(ctx context.Context) (*OnetDataElementList, error) {
	return c.onetDataList(ctx, "work_context")
}

// OnlineONETWorkContextItem returns occupations associated with the given work context element ID.
func (c *Client) OnlineONETWorkContextItem(ctx context.Context, id string) (*OnetDataAttributeOccupations, error) {
	return c.onetDataDetail(ctx, "work_context", id)
}

// OnlineONETWorkStyles returns the list of all O*NET work style elements.
func (c *Client) OnlineONETWorkStyles(ctx context.Context) (*OnetDataElementList, error) {
	return c.onetDataList(ctx, "work_styles")
}

// OnlineONETWorkStyle returns occupations associated with the given work style element ID.
func (c *Client) OnlineONETWorkStyle(ctx context.Context, id string) (*OnetDataAttributeOccupations, error) {
	return c.onetDataDetail(ctx, "work_styles", id)
}
