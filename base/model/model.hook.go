package model

type Hook struct {
	ID              string      `json:"id"`
	ResponseMessage string      `json:"response-message"`
	Trigger         TriggerRule `json:"trigger-rule"`
}

type TriggerRule struct {
	And []struct {
		Match struct {
			Type      string `json:"type"`
			Secret    string `json:"secret"`
			Parameter struct {
				Source string `json:"source"`
				Name   string `json:"name"`
			} `json:"parameter"`
		} `json:"match"`
	} `json:"and"`
}
