package sendinblue

import (
	"github.com/vijay1811/sendinblue/protocol/sendinblueobj"
)

type Campaigner interface {
	CreateCampaign(c *sendinblueobj.CampaignCreate) (*sendinblueobj.CampaignCreateResponse, error)
}
