package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-master-recipe-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			MasterRecipeGroup:           data.MasterRecipeGroup,
			MasterRecipe:                data.MasterRecipe,
			MasterRecipeInternalVersion: data.MasterRecipeInternalVersion,
			IsMarkedForDeletion:         data.IsMarkedForDeletion,
			BillOfOperationsDesc:        data.BillOfOperationsDesc,
			LongTextLanguageCode:        data.LongTextLanguageCode,
			PlainLongText:               data.PlainLongText,
			Plant:                       data.Plant,
			BillOfOperationsStatus:      data.BillOfOperationsStatus,
			BillOfOperationsUsage:       data.BillOfOperationsUsage,
			ResponsiblePlannerGroup:     data.ResponsiblePlannerGroup,
			BillOfOperationsProfile:     data.BillOfOperationsProfile,
			MinimumLotSizeQuantity:      data.MinimumLotSizeQuantity,
			MaximumLotSizeQuantity:      data.MaximumLotSizeQuantity,
			BillOfOperationsUnit:        data.BillOfOperationsUnit,
			CreationDate:                data.CreationDate,
			LastChangeDate:              data.LastChangeDate,
			ValidityStartDate:           data.ValidityStartDate,
			ValidityEndDate:             data.ValidityEndDate,
			ChangeNumber:                data.ChangeNumber,
			ChangedDateTime:             data.ChangedDateTime,
			ToMaterialAssignment:        data.ToMaterialAssignment.Deferred.URI,
			ToOperation:                 data.ToOperation.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToToMaterialAssignment(raw []byte, l *logger.Logger) ([]ToMaterialAssignment, error) {
	pm := &responses.ToMaterialAssignment{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToMaterialAssignment. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toMaterialAssignment := make([]ToMaterialAssignment, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toMaterialAssignment = append(toMaterialAssignment, ToMaterialAssignment{
			Product:                        data.Product,
			Plant:                          data.Plant,
			MasterRecipeGroup:              data.MasterRecipeGroup,
			MasterRecipe:                   data.MasterRecipe,
			MasterRecipeMaterialAssignment: data.MasterRecipeMaterialAssignment,
			MstrRcpMatlAssgmtIntVersion:    data.MstrRcpMatlAssgmtIntVersion,
			CreationDate:                   data.CreationDate,
			LastChangeDate:                 data.LastChangeDate,
			ValidityStartDate:              data.ValidityStartDate,
			ValidityEndDate:                data.ValidityEndDate,
			ChangeNumber:                   data.ChangeNumber,
			ChangedDateTime:                data.ChangedDateTime,
		})
	}

	return toMaterialAssignment, nil
}

func ConvertToToOperation(raw []byte, l *logger.Logger) ([]ToOperation, error) {
	pm := &responses.ToOperation{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToOperation. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toOperation := make([]ToOperation, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toOperation = append(toOperation, ToOperation{
			MasterRecipeGroup:             data.MasterRecipeGroup,
			MasterRecipe:                  data.MasterRecipe,
			MasterRecipeOperationIntID:    data.MasterRecipeOperationIntID,
			MstrRcpOperationIntVersion:    data.MstrRcpOperationIntVersion,
			Operation:                     data.Operation,
			OperationText:                 data.OperationText,
			LongTextLanguageCode:          data.LongTextLanguageCode,
			PlainLongText:                 data.PlainLongText,
			WorkCenterTypeCode:            data.WorkCenterTypeCode,
			WorkCenterInternalID:          data.WorkCenterInternalID,
			OperationStandardTextCode:     data.OperationStandardTextCode,
			Plant:                         data.Plant,
			OperationControlProfile:       data.OperationControlProfile,
			OperationReferenceQuantity:    data.OperationReferenceQuantity,
			OperationUnit:                 data.OperationUnit,
			OpQtyToBaseQtyNmrtr:           data.OpQtyToBaseQtyNmrtr,
			OpQtyToBaseQtyDnmntr:          data.OpQtyToBaseQtyDnmntr,
			NumberOfTimeTickets:           data.NumberOfTimeTickets,
			NumberOfConfirmationSlips:     data.NumberOfConfirmationSlips,
			OperationCostingRelevancyType: data.OperationCostingRelevancyType,
			BusinessProcess:               data.BusinessProcess,
			OperationSetupType:            data.OperationSetupType,
			OperationSetupGroupCategory:   data.OperationSetupGroupCategory,
			OperationSetupGroup:           data.OperationSetupGroup,
			CapacityCategoryCode:          data.CapacityCategoryCode,
			CreationDate:                  data.CreationDate,
			LastChangeDate:                data.LastChangeDate,
			ValidityStartDate:             data.ValidityStartDate,
			ValidityEndDate:               data.ValidityEndDate,
			ChangeNumber:                  data.ChangeNumber,
			ChangedDateTime:               data.ChangedDateTime,
			ToPhase:                       data.ToPhase.Deferred.URI,
		})
	}

	return toOperation, nil
}

func ConvertToToPhase(raw []byte, l *logger.Logger) ([]ToPhase, error) {
	pm := &responses.ToPhase{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToPhase. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toPhase := make([]ToPhase, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toPhase = append(toPhase, ToPhase{
            MasterRecipeGroup:              data.MasterRecipeGroup,
            MasterRecipe:                   data.MasterRecipe,
            MasterRecipeOperationIntID:     data.MasterRecipeOperationIntID,
            MstrRcpSuperiorOpIntVersion:    data.MstrRcpSuperiorOpIntVersion,
            MstrRcpOperationIntVersion:     data.MstrRcpOperationIntVersion,
            SuperiorOperationInternalID:    data.SuperiorOperationInternalID,
            Operation:                      data.Operation,
            OperationText:                  data.OperationText,
            LongTextLanguageCode:           data.LongTextLanguageCode,
            PlainLongText:                  data.PlainLongText,
            WorkCenterTypeCode:             data.WorkCenterTypeCode,
            WorkCenterInternalID:           data.WorkCenterInternalID,
            Plant:                          data.Plant,
            OperationControlProfile:        data.OperationControlProfile,
            ControlRecipeDestination:       data.ControlRecipeDestination,
            OperationStandardTextCode:      data.OperationStandardTextCode,
            OperationReferenceQuantity:     data.OperationReferenceQuantity,
            OperationUnit:                  data.OperationUnit,
            OpQtyToBaseQtyNmrtr:            data.OpQtyToBaseQtyNmrtr,
            OpQtyToBaseQtyDnmntr:           data.OpQtyToBaseQtyDnmntr,
            StandardWorkFormulaParam1:      data.StandardWorkFormulaParam1,
            StandardWorkFormulaParamName1:  data.StandardWorkFormulaParamName1,
            StandardWorkQuantity1:          data.StandardWorkQuantity1,
            StandardWorkQuantityUnit1:      data.StandardWorkQuantityUnit1,
            CostCtrActivityType1:           data.CostCtrActivityType1,
            StandardWorkFormulaParam2:      data.StandardWorkFormulaParam2,
            StandardWorkFormulaParamName2:  data.StandardWorkFormulaParamName2,
            StandardWorkQuantity2:          data.StandardWorkQuantity2,
            StandardWorkQuantityUnit2:      data.StandardWorkQuantityUnit2,
            CostCtrActivityType2:           data.CostCtrActivityType2,
            StandardWorkFormulaParam3:      data.StandardWorkFormulaParam3,
            StandardWorkFormulaParamName3:  data.StandardWorkFormulaParamName3,
            StandardWorkQuantity3:          data.StandardWorkQuantity3,
            StandardWorkQuantityUnit3:      data.StandardWorkQuantityUnit3,
            CostCtrActivityType3:           data.CostCtrActivityType3,
            StandardWorkFormulaParam4:      data.StandardWorkFormulaParam4,
            StandardWorkFormulaParamName4:  data.StandardWorkFormulaParamName4,
            StandardWorkQuantity4:          data.StandardWorkQuantity4,
            StandardWorkQuantityUnit4:      data.StandardWorkQuantityUnit4,
            CostCtrActivityType4:           data.CostCtrActivityType4,
            StandardWorkFormulaParam5:      data.StandardWorkFormulaParam5,
            StandardWorkFormulaParamName5:  data.StandardWorkFormulaParamName5,
            StandardWorkQuantity5:          data.StandardWorkQuantity5,
            StandardWorkQuantityUnit5:      data.StandardWorkQuantityUnit5,
            CostCtrActivityType5:           data.CostCtrActivityType5,
            StandardWorkFormulaParam6:      data.StandardWorkFormulaParam6,
            StandardWorkFormulaParamName6:  data.StandardWorkFormulaParamName6,
            StandardWorkQuantity6:          data.StandardWorkQuantity6,
            StandardWorkQuantityUnit6:      data.StandardWorkQuantityUnit6,
            CostCtrActivityType6:           data.CostCtrActivityType6,
            NumberOfTimeTickets:            data.NumberOfTimeTickets,
            NumberOfConfirmationSlips:      data.NumberOfConfirmationSlips,
            OperationCostingRelevancyType:  data.OperationCostingRelevancyType,
            BusinessProcess:                data.BusinessProcess,
            OperationSetupType:             data.OperationSetupType,
            OperationSetupGroupCategory:    data.OperationSetupGroupCategory,
            OperationSetupGroup:            data.OperationSetupGroup,
            CapacityCategoryCode:           data.CapacityCategoryCode,
            OpIsExtlyProcdWithSubcontrg:    data.OpIsExtlyProcdWithSubcontrg,
            InspectionLotType:              data.InspectionLotType,
            PurchasingInfoRecord:           data.PurchasingInfoRecord,
            PurchasingOrganization:         data.PurchasingOrganization,
            PurchaseContract:               data.PurchaseContract,
            PurchaseContractItem:           data.PurchaseContractItem,
            PurchasingInfoRecdAddlGrpgName: data.PurchasingInfoRecdAddlGrpgName,
            PlannedDeliveryDuration:        data.PlannedDeliveryDuration,
            MaterialGroup:                  data.MaterialGroup,
            PurchasingGroup:                data.PurchasingGroup,
            Supplier:                       data.Supplier,
            NumberOfOperationPriceUnits:    data.NumberOfOperationPriceUnits,
            CostElement:                    data.CostElement,
            OpExternalProcessingPrice:      data.OpExternalProcessingPrice,
            OpExternalProcessingCurrency:   data.OpExternalProcessingCurrency,
            CreationDate:                   data.CreationDate,
            LastChangeDate:                 data.LastChangeDate,
            ValidityStartDate:              data.ValidityStartDate,
            ValidityEndDate:                data.ValidityEndDate,
            ChangeNumber:                   data.ChangeNumber,
            ChangedDateTime:                data.ChangedDateTime,
			ToPhaseRelationship:            data.ToPhaseRelationship.Deferred.URI,
		})
	}

	return toPhase, nil
}

func ConvertToToPhaseRelationship(raw []byte, l *logger.Logger) ([]ToPhaseRelationship, error) {
	pm := &responses.ToPhaseRelationship{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToPhaseRelationship. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toPhaseRelationship := make([]ToPhaseRelationship, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toPhaseRelationship = append(toPhaseRelationship, ToPhaseRelationship{
	PrdcssrMasterRecipeGroup:      data.PrdcssrMasterRecipeGroup,
	PrdcssrMasterRecipe:           data.PrdcssrMasterRecipe,
	PrdcssrMstrRcpOpInternalID:    data.PrdcssrMstrRcpOpInternalID,
	SuccssrMasterRecipeGroup:      data.SuccssrMasterRecipeGroup,
	SuccssrMasterRecipe:           data.SuccssrMasterRecipe,
	SuccssrMstrRcpOpInternalID:    data.SuccssrMstrRcpOpInternalID,
	MasterRecipeRelationshipType:  data.MasterRecipeRelationshipType,
	MaxTimeIntvlIsUsedForSchedg:   data.MaxTimeIntvlIsUsedForSchedg,
	MstrRcpRelationshipIntVersion: data.MstrRcpRelationshipIntVersion,
	TimeIntvlBtwnRelshp:           data.TimeIntvlBtwnRelshp,
	MaxTimeIntvlBtwnRelshp:        data.MaxTimeIntvlBtwnRelshp,
	TimeIntvlBtwnRelshpUnit:       data.TimeIntvlBtwnRelshpUnit,
	FactoryCalendar:               data.FactoryCalendar,
	WorkCenterInternalID:          data.WorkCenterInternalID,
	Plant:                         data.Plant,
	CreationDate:                  data.CreationDate,
	LastChangeDate:                data.LastChangeDate,
	ValidityStartDate:             data.ValidityStartDate,
	ValidityEndDate:               data.ValidityEndDate,
	ChangeNumber:                  data.ChangeNumber,
	ChangedDateTime:               data.ChangedDateTime,
		})
	}

	return toPhaseRelationship, nil
}
