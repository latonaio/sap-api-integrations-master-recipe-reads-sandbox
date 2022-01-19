package responses

type ToPhaseRelationship struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			PrdcssrMasterRecipeGroup      string `json:"PrdcssrMasterRecipeGroup"`
			PrdcssrMasterRecipe           string `json:"PrdcssrMasterRecipe"`
			PrdcssrMstrRcpOpInternalID    string `json:"PrdcssrMstrRcpOpInternalID"`
			SuccssrMasterRecipeGroup      string `json:"SuccssrMasterRecipeGroup"`
			SuccssrMasterRecipe           string `json:"SuccssrMasterRecipe"`
			SuccssrMstrRcpOpInternalID    string `json:"SuccssrMstrRcpOpInternalID"`
			MasterRecipeRelationshipType  string `json:"MasterRecipeRelationshipType"`
			MaxTimeIntvlIsUsedForSchedg   bool   `json:"MaxTimeIntvlIsUsedForSchedg"`
			MstrRcpRelationshipIntVersion string `json:"MstrRcpRelationshipIntVersion"`
			TimeIntvlBtwnRelshp           string `json:"TimeIntvlBtwnRelshp"`
			MaxTimeIntvlBtwnRelshp        string `json:"MaxTimeIntvlBtwnRelshp"`
			TimeIntvlBtwnRelshpUnit       string `json:"TimeIntvlBtwnRelshpUnit"`
			FactoryCalendar               string `json:"FactoryCalendar"`
			WorkCenterInternalID          string `json:"WorkCenterInternalID"`
			Plant                         string `json:"Plant"`
			CreationDate                  string `json:"CreationDate"`
			LastChangeDate                string `json:"LastChangeDate"`
			ValidityStartDate             string `json:"ValidityStartDate"`
			ValidityEndDate               string `json:"ValidityEndDate"`
			ChangeNumber                  string `json:"ChangeNumber"`
			ChangedDateTime               string `json:"ChangedDateTime"`
		} `json:"results"`
	} `json:"d"`
}
