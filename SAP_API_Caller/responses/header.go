package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			MasterRecipeGroup           string `json:"MasterRecipeGroup"`
			MasterRecipe                string `json:"MasterRecipe"`
			MasterRecipeInternalVersion string `json:"MasterRecipeInternalVersion"`
			IsMarkedForDeletion         bool   `json:"IsMarkedForDeletion"`
			BillOfOperationsDesc        string `json:"BillOfOperationsDesc"`
			LongTextLanguageCode        string `json:"LongTextLanguageCode"`
			PlainLongText               string `json:"PlainLongText"`
			Plant                       string `json:"Plant"`
			BillOfOperationsStatus      string `json:"BillOfOperationsStatus"`
			BillOfOperationsUsage       string `json:"BillOfOperationsUsage"`
			ResponsiblePlannerGroup     string `json:"ResponsiblePlannerGroup"`
			BillOfOperationsProfile     string `json:"BillOfOperationsProfile"`
			MinimumLotSizeQuantity      string `json:"MinimumLotSizeQuantity"`
			MaximumLotSizeQuantity      string `json:"MaximumLotSizeQuantity"`
			BillOfOperationsUnit        string `json:"BillOfOperationsUnit"`
			CreationDate                string `json:"CreationDate"`
			LastChangeDate              string `json:"LastChangeDate"`
			ValidityStartDate           string `json:"ValidityStartDate"`
			ValidityEndDate             string `json:"ValidityEndDate"`
			ChangeNumber                string `json:"ChangeNumber"`
			ChangedDateTime             string `json:"ChangedDateTime"`
			ToMaterialAssignment        struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_MatlAssgmt"`
			ToOperation struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Operation"`
		} `json:"results"`
	} `json:"d"`
}
