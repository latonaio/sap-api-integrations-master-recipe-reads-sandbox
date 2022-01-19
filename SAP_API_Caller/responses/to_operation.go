package responses

type ToOperation struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			MasterRecipeGroup             string `json:"MasterRecipeGroup"`
			MasterRecipe                  string `json:"MasterRecipe"`
			MasterRecipeOperationIntID    string `json:"MasterRecipeOperationIntID"`
			MstrRcpOperationIntVersion    string `json:"MstrRcpOperationIntVersion"`
			Operation                     string `json:"Operation"`
			OperationText                 string `json:"OperationText"`
			LongTextLanguageCode          string `json:"LongTextLanguageCode"`
			PlainLongText                 string `json:"PlainLongText"`
			WorkCenterTypeCode            string `json:"WorkCenterTypeCode"`
			WorkCenterInternalID          string `json:"WorkCenterInternalID"`
			OperationStandardTextCode     string `json:"OperationStandardTextCode"`
			Plant                         string `json:"Plant"`
			OperationControlProfile       string `json:"OperationControlProfile"`
			OperationReferenceQuantity    string `json:"OperationReferenceQuantity"`
			OperationUnit                 string `json:"OperationUnit"`
			OpQtyToBaseQtyNmrtr           string `json:"OpQtyToBaseQtyNmrtr"`
			OpQtyToBaseQtyDnmntr          string `json:"OpQtyToBaseQtyDnmntr"`
			NumberOfTimeTickets           string `json:"NumberOfTimeTickets"`
			NumberOfConfirmationSlips     string `json:"NumberOfConfirmationSlips"`
			OperationCostingRelevancyType string `json:"OperationCostingRelevancyType"`
			BusinessProcess               string `json:"BusinessProcess"`
			OperationSetupType            string `json:"OperationSetupType"`
			OperationSetupGroupCategory   string `json:"OperationSetupGroupCategory"`
			OperationSetupGroup           string `json:"OperationSetupGroup"`
			CapacityCategoryCode          string `json:"CapacityCategoryCode"`
			CreationDate                  string `json:"CreationDate"`
			LastChangeDate                string `json:"LastChangeDate"`
			ValidityStartDate             string `json:"ValidityStartDate"`
			ValidityEndDate               string `json:"ValidityEndDate"`
			ChangeNumber                  string `json:"ChangeNumber"`
			ChangedDateTime               string `json:"ChangedDateTime"`
			ToPhase                       struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Phase"`
		} `json:"results"`
	} `json:"d"`
}
