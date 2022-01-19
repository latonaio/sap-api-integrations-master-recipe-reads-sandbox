package responses

type ToMaterialAssignment struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			Product                        string `json:"Product"`
			Plant                          string `json:"Plant"`
			MasterRecipeGroup              string `json:"MasterRecipeGroup"`
			MasterRecipe                   string `json:"MasterRecipe"`
			MasterRecipeMaterialAssignment string `json:"MasterRecipeMaterialAssignment"`
			MstrRcpMatlAssgmtIntVersion    string `json:"MstrRcpMatlAssgmtIntVersion"`
			CreationDate                   string `json:"CreationDate"`
			LastChangeDate                 string `json:"LastChangeDate"`
			ValidityStartDate              string `json:"ValidityStartDate"`
			ValidityEndDate                string `json:"ValidityEndDate"`
			ChangeNumber                   string `json:"ChangeNumber"`
			ChangedDateTime                string `json:"ChangedDateTime"`
		} `json:"results"`
	} `json:"d"`
}
