package sendinblue

import (
	"encoding/json"
	"fmt"

	"github.com/vijay1811/sendinblue/protocol"
	"github.com/vijay1811/sendinblue/protocol/sendinblueobj"
)

const campaignCreateFilePath = "sendinblue/campaign_create.go:"

func (c *SIBClient) CreateCampaign(campaign *sendinblueobj.CampaignCreate) (*sendinblueobj.CampaignCreateResponse, error) {
	data, err := json.Marshal(campaign)
	if err != nil {
		return nil, fmt.Errorf("%s error while marshalling campaign: %s", campaignCreateFilePath, err)
	}

	resp, err := c.sendRequest("campaign", data)
	if err != nil {
		return nil, err
	}

	switch resp.Code {
	case protocol.ResponseCode_Success:
		var createResp *sendinblueobj.CampaignCreateResponse
		if err := json.Unmarshal(resp.Data, &createResp); err != nil {
			return nil, fmt.Errorf("%s error while Un marshaling campaign create response: %s", campaignCreateFilePath, err)
		}
		return createResp, nil
	case protocol.ResponseCode_Failure:
		return nil, fmt.Errorf("%s failed to create campaign: %s", campaignCreateFilePath, resp.Message)
	default:
		return nil, fmt.Errorf("%s failed to create campaign unknown response code: %s, message: %v", campaignCreateFilePath, resp.Code, resp.Message)
	}
}
