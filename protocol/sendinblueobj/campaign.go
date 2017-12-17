package sendinblueobj

import (
	"time"
)

type InlineImage uint8

const (
	InlineImage_NotAllowed InlineImage = 0
	InlineImage_Allowed    InlineImage = 1
)

type MirrorActive uint8

const (
	MirrorActive_NotActive MirrorActive = 0
	MirrorActive_Active    MirrorActive = 1
)

type SendNow uint8

const (
	SentNow_NotAllowed SendNow = 0
	SendNow_Allowed    SendNow = 1
)

type CampaignCreate struct {
	Category      string       `json:"category,omitempty"`
	FromName      string       `json:"from_name,omitempty"`
	Name          string       `json:"name"`
	Bat           string       `json:"bat,omitempty"`
	HTMLContent   string       `json:"html_content"`
	HTMLUrl       string       `json:"html_url"`
	ListID        []uint32     `json:"listid"`
	ScheduledDate *time.Time   `json:"scheduled_date,omitempty"`
	Subject       string       `json:"subject"`
	FromEmail     string       `json:"from_email"`
	ReplyTo       string       `json:"reply_to,omitempty"`
	ToField       string       `json:"to_field,omitempty"` //[PRENOM][NOM] to specify the first name and last name of the recipient
	ExcludeList   []uint32     `json:"exclude_list,omitempty"`
	AttachmentURL string       `json:"attachment_url,omitempty"`
	InlineImage   InlineImage  `json:"inline_image,omitempty"`
	MirrorActive  MirrorActive `json:"mirror_active,omitempty"`
	SendNow       SendNow      `json:"send_now,omitempty"`
	UTMCampaign   rune         `json:"utm_campaign,omitempty"`
}

type CampaignCreateResponse struct {
	ID uint32 `json:"id"`
}
