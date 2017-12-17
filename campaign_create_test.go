package sendinblue

import (
	"fmt"
	"testing"

	"github.com/vijay1811/sendinblue/protocol/sendinblueobj"
)

const (
	testData = `<!DOCTYPE html>
	<html>
	<body>
	
	<h1>My First Heading</h1>
	<p>My first paragraph.</p>
	
	</body>
	</html>`
)

func TestCreateCampaign(t *testing.T) {
	key := "your key"
	client := NewSIBClient(key)
	campaign := &sendinblueobj.CampaignCreate{
		Name:        "TestCreateCampaign",
		HTMLContent: testData,

		//ScheduledDate: time.Now().AddDate(0, 0, 1),
		ListID:    []uint32{2},
		Subject:   "Test Subject",
		FromEmail: "vijay18111.kumar@gmail.com",
		//	FromName:  "vijay kumar",
	}

	resp, err := client.CreateCampaign(campaign)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
