package main

import "encoding/xml"

type Manifest struct {
	XMLName            xml.Name `xml:"manifest"`
	Text               string   `xml:",chardata"`
	Xmlns              string   `xml:"xmlns,attr"`
	Xlink              string   `xml:"xlink,attr"`
	GeneralInformation struct {
		Text             string `xml:",chardata"`
		Title            string `xml:"title"`
		Created          string `xml:"created"`
		Author           string `xml:"author"`
		Application      string `xml:"application"`
		SubmissionType   string `xml:"submission-type"`
		ArchiveType      string `xml:"archive-type"`
		LegislationsInfo struct {
			Text        string `xml:",chardata"`
			Legislation []struct {
				Text    string `xml:",chardata"`
				ID      string `xml:"id"`
				Version string `xml:"version"`
			} `xml:"legislation"`
		} `xml:"legislations-info"`
	} `xml:"general-information"`
	Comment            string `xml:"comment"`
	BaseDocumentUuid   string `xml:"base-document-uuid"`
	ContainedDocuments struct {
		Text     string `xml:",chardata"`
		Document []struct {
			Text    string `xml:",chardata"`
			ID      string `xml:"id,attr"`
			Type    string `xml:"type"`
			Subtype string `xml:"subtype"`
			Name    struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Href string `xml:"href,attr"`
			} `xml:"name"`
			FirstModificationDate string `xml:"first-modification-date"`
			LastModificationDate  string `xml:"last-modification-date"`
			Uuid                  string `xml:"uuid"`
			Links                 struct {
				Text string `xml:",chardata"`
				Link []struct {
					Text    string `xml:",chardata"`
					RefUuid string `xml:"ref-uuid"`
					RefType string `xml:"ref-type"`
				} `xml:"link"`
			} `xml:"links"`
			Representation struct {
				Text        string `xml:",chardata"`
				SubjectType string `xml:"subject-type"`
				LegalEntity struct {
					Text    string `xml:",chardata"`
					Name    string `xml:"name"`
					Town    string `xml:"town"`
					Country string `xml:"country"`
				} `xml:"legal-entity"`
				Name            string `xml:"name"`
				Town            string `xml:"town"`
				Country         string `xml:"country"`
				IUPACName       string `xml:"IUPAC-name"`
				CASNumber       string `xml:"CAS-number"`
				InventoryNumber string `xml:"inventory-number"`
				Parent          struct {
					Text string `xml:",chardata"`
					Type string `xml:"type"`
					Name string `xml:"name"`
				} `xml:"parent"`
				ReferenceSubstance struct {
					Text            string `xml:",chardata"`
					Name            string `xml:"name"`
					IUPACName       string `xml:"IUPAC-name"`
					CASNumber       string `xml:"CAS-number"`
					InventoryNumber string `xml:"inventory-number"`
				} `xml:"reference-substance"`
			} `xml:"representation"`
		} `xml:"document"`
	} `xml:"contained-documents"`
}
type Document struct {
	XMLName          xml.Name `xml:"Document"`
	Text             string   `xml:",chardata"`
	I6c              string   `xml:"i6c,attr"`
	Xsi              string   `xml:"xsi,attr"`
	SchemaLocation   string   `xml:"schemaLocation,attr"`
	XML              string   `xml:"xml,attr"`
	PlatformMetadata struct {
		Text                  string `xml:",chardata"`
		I6m                   string `xml:"i6m,attr"`
		SchemaLocation        string `xml:"schemaLocation,attr"`
		IuclidVersion         string `xml:"iuclidVersion"`
		DocumentKey           string `xml:"documentKey"`
		ParentDocumentKey     string `xml:"parentDocumentKey"`
		Name                  string `xml:"name"`
		DocumentType          string `xml:"documentType"`
		DocumentSubType       string `xml:"documentSubType"`
		OrderInSectionNo      string `xml:"orderInSectionNo"`
		DefinitionVersion     string `xml:"definitionVersion"`
		CreationDate          string `xml:"creationDate"`
		LastModificationDate  string `xml:"lastModificationDate"`
		SubmissionType        string `xml:"submissionType"`
		SubmissionTypeVersion string `xml:"submissionTypeVersion"`
		SubmittingLegalEntity string `xml:"submittingLegalEntity"`
		DossierSubject        string `xml:"dossierSubject"`
		I5Origin              string `xml:"i5Origin"`
		CreationTool          string `xml:"creationTool"`
		SnapshotCreationTool  string `xml:"snapshotCreationTool"`
	} `xml:"PlatformMetadata"`
	Content struct {
		Text                                     string `xml:",chardata"`
		FLEXIBLESUMMARYProductSummaryComposition struct {
			Text                      string `xml:",chardata"`
			Xmlns                     string `xml:"xmlns,attr"`
			I6                        string `xml:"i6,attr"`
			AdministrativeDataSummary struct {
				Text           string `xml:",chardata"`
				DataProtection struct {
					Text            string `xml:",chardata"`
					Confidentiality string `xml:"confidentiality"`
					Justification   string `xml:"justification"`
				} `xml:"DataProtection"`
			} `xml:"AdministrativeDataSummary"`
			ProductSummaryComposition struct {
				Text                   string `xml:",chardata"`
				RepresentsMetaSPC      string `xml:"RepresentsMetaSPC"`
				ProductFamilyGroupName string `xml:"ProductFamilyGroupName"`
				FormulationType        struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value"`
				} `xml:"FormulationType"`
				Composition struct {
					Text  string `xml:",chardata"`
					Entry []struct {
						Text     string `xml:",chardata"`
						Uuid     string `xml:"uuid,attr"`
						Function struct {
							Text  string `xml:",chardata"`
							Value string `xml:"value"`
							Other string `xml:"other"`
						} `xml:"Function"`
						Concentration struct {
							Text       string `xml:",chardata"`
							UnitCode   string `xml:"unitCode"`
							LowerValue string `xml:"lowerValue"`
							UpperValue string `xml:"upperValue"`
						} `xml:"Concentration"`
						Name               string `xml:"Name"`
						SubstanceOfConcern string `xml:"SubstanceOfConcern"`
						Gci                struct {
							Text string `xml:",chardata"`
							Nil  string `xml:"nil,attr"`
						} `xml:"Gci"`
						Icg struct {
							Text string `xml:",chardata"`
							Nil  string `xml:"nil,attr"`
						} `xml:"Icg"`
						Sfc struct {
							Text string `xml:",chardata"`
							Nil  string `xml:"nil,attr"`
						} `xml:"Sfc"`
						SubstanceGeneratedInSitu struct {
							Text string `xml:",chardata"`
							Nil  string `xml:"nil,attr"`
						} `xml:"SubstanceGeneratedInSitu"`
					} `xml:"entry"`
				} `xml:"Composition"`
			} `xml:"ProductSummaryComposition"`
			ProductCompositions struct {
				Text                string `xml:",chardata"`
				ProductCompositions struct {
					Text string   `xml:",chardata"`
					Key  []string `xml:"key"`
				} `xml:"ProductCompositions"`
			} `xml:"ProductCompositions"`
			MetaSPC struct {
				Text                       string `xml:",chardata"`
				ClassificationAndLabelling string `xml:"ClassificationAndLabelling"`
				AuthorisedUses             struct {
					Text string `xml:",chardata"`
					Key  string `xml:"key"`
				} `xml:"AuthorisedUses"`
				InstructionsForUseAndMeasuresToProtect string `xml:"InstructionsForUseAndMeasuresToProtect"`
				SummaryAndEvaluation                   string `xml:"SummaryAndEvaluation"`
			} `xml:"MetaSPC"`
		} `xml:"FLEXIBLE_SUMMARY.ProductSummaryComposition"`
	} `xml:"Content"`
	Attachments struct {
		Text           string `xml:",chardata"`
		I6a            string `xml:"i6a,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	} `xml:"Attachments"`
	ModificationHistory struct {
		Text           string `xml:",chardata"`
		I6h            string `xml:"i6h,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	} `xml:"ModificationHistory"`
}

type MixtureDocument struct {
	XMLName          xml.Name `xml:"Document"`
	Text             string   `xml:",chardata"`
	I6c              string   `xml:"i6c,attr"`
	Xsi              string   `xml:"xsi,attr"`
	SchemaLocation   string   `xml:"schemaLocation,attr"`
	XML              string   `xml:"xml,attr"`
	PlatformMetadata struct {
		Text                  string `xml:",chardata"`
		I6m                   string `xml:"i6m,attr"`
		SchemaLocation        string `xml:"schemaLocation,attr"`
		IuclidVersion         string `xml:"iuclidVersion"`
		DocumentKey           string `xml:"documentKey"`
		ParentDocumentKey     string `xml:"parentDocumentKey"`
		Name                  string `xml:"name"`
		DocumentType          string `xml:"documentType"`
		DocumentSubType       string `xml:"documentSubType"`
		OrderInSectionNo      string `xml:"orderInSectionNo"`
		DefinitionVersion     string `xml:"definitionVersion"`
		CreationDate          string `xml:"creationDate"`
		LastModificationDate  string `xml:"lastModificationDate"`
		SubmissionType        string `xml:"submissionType"`
		SubmissionTypeVersion string `xml:"submissionTypeVersion"`
		SubmittingLegalEntity string `xml:"submittingLegalEntity"`
		DossierSubject        string `xml:"dossierSubject"`
		I5Origin              string `xml:"i5Origin"`
		CreationTool          string `xml:"creationTool"`
		SnapshotCreationTool  string `xml:"snapshotCreationTool"`
	} `xml:"PlatformMetadata"`
	Content struct {
		Text                             string `xml:",chardata"`
		FLEXIBLERECORDMixtureComposition struct {
			Text               string `xml:",chardata"`
			Xmlns              string `xml:"xmlns,attr"`
			I6                 string `xml:"i6,attr"`
			AdministrativeData struct {
				Text           string `xml:",chardata"`
				DataProtection struct {
					Text            string `xml:",chardata"`
					Confidentiality string `xml:"confidentiality"`
					Justification   string `xml:"justification"`
				} `xml:"DataProtection"`
			} `xml:"AdministrativeData"`
			GeneralInformation struct {
				Text       string `xml:",chardata"`
				Name       string `xml:"Name"`
				TradeNames struct {
					Text  string `xml:",chardata"`
					Entry struct {
						Text    string `xml:",chardata"`
						Uuid    string `xml:"uuid,attr"`
						Country struct {
							Text  string `xml:",chardata"`
							Value string `xml:"value"`
						} `xml:"Country"`
						TradeName string `xml:"TradeName"`
					} `xml:"entry"`
				} `xml:"TradeNames"`
				FormulationType struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value"`
				} `xml:"FormulationType"`
			} `xml:"GeneralInformation"`
			Components struct {
				Text       string `xml:",chardata"`
				Components struct {
					Text  string `xml:",chardata"`
					Entry []struct {
						Text           string `xml:",chardata"`
						Uuid           string `xml:"uuid,attr"`
						DataProtection struct {
							Text            string `xml:",chardata"`
							Confidentiality string `xml:"confidentiality"`
							Justification   string `xml:"justification"`
						} `xml:"DataProtection"`
						Reference string `xml:"Reference"`
						Function  struct {
							Text  string `xml:",chardata"`
							Value string `xml:"value"`
							Other string `xml:"other"`
						} `xml:"Function"`
						TypicalConcentration struct {
							Text      string `xml:",chardata"`
							UnitCode  string `xml:"unitCode"`
							Qualifier string `xml:"qualifier"`
							Value     string `xml:"value"`
						} `xml:"TypicalConcentration"`
						ConcentrationRange string `xml:"ConcentrationRange"`
						SubstanceOfConcern string `xml:"SubstanceOfConcern"`
						Gci                struct {
							Text string `xml:",chardata"`
							Nil  string `xml:"nil,attr"`
						} `xml:"Gci"`
						Icg struct {
							Text string `xml:",chardata"`
							Nil  string `xml:"nil,attr"`
						} `xml:"Icg"`
						Sfc struct {
							Text string `xml:",chardata"`
							Nil  string `xml:"nil,attr"`
						} `xml:"Sfc"`
						SubstanceGeneratedInSitu string `xml:"SubstanceGeneratedInSitu"`
					} `xml:"entry"`
				} `xml:"Components"`
			} `xml:"Components"`
		} `xml:"FLEXIBLE_RECORD.MixtureComposition"`
	} `xml:"Content"`
	Attachments struct {
		Text           string `xml:",chardata"`
		I6a            string `xml:"i6a,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	} `xml:"Attachments"`
	ModificationHistory struct {
		Text           string `xml:",chardata"`
		I6h            string `xml:"i6h,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	} `xml:"ModificationHistory"`
}

type Substance struct {
	XMLName          xml.Name `xml:"Document"`
	Text             string   `xml:",chardata"`
	I6c              string   `xml:"i6c,attr"`
	Xsi              string   `xml:"xsi,attr"`
	SchemaLocation   string   `xml:"schemaLocation,attr"`
	XML              string   `xml:"xml,attr"`
	PlatformMetadata struct {
		Text                  string `xml:",chardata"`
		I6m                   string `xml:"i6m,attr"`
		SchemaLocation        string `xml:"schemaLocation,attr"`
		IuclidVersion         string `xml:"iuclidVersion"`
		DocumentKey           string `xml:"documentKey"`
		ParentDocumentKey     string `xml:"parentDocumentKey"`
		Name                  string `xml:"name"`
		DocumentType          string `xml:"documentType"`
		DocumentSubType       string `xml:"documentSubType"`
		OrderInSectionNo      string `xml:"orderInSectionNo"`
		DefinitionVersion     string `xml:"definitionVersion"`
		CreationDate          string `xml:"creationDate"`
		LastModificationDate  string `xml:"lastModificationDate"`
		SubmissionType        string `xml:"submissionType"`
		SubmissionTypeVersion string `xml:"submissionTypeVersion"`
		SubmittingLegalEntity string `xml:"submittingLegalEntity"`
		DossierSubject        string `xml:"dossierSubject"`
		I5Origin              string `xml:"i5Origin"`
		CreationTool          string `xml:"creationTool"`
		SnapshotCreationTool  string `xml:"snapshotCreationTool"`
	} `xml:"PlatformMetadata"`
	Content struct {
		Text      string `xml:",chardata"`
		SUBSTANCE struct {
			Text                       string `xml:",chardata"`
			Xmlns                      string `xml:"xmlns,attr"`
			I6                         string `xml:"i6,attr"`
			Templates                  string `xml:"Templates"`
			ChemicalName               string `xml:"ChemicalName"`
			PublicName                 string `xml:"PublicName"`
			OtherNames                 string `xml:"OtherNames"`
			OwnerLegalEntityProtection struct {
				Text            string `xml:",chardata"`
				Confidentiality string `xml:"confidentiality"`
				Justification   string `xml:"justification"`
			} `xml:"OwnerLegalEntityProtection"`
			OwnerLegalEntity     string `xml:"OwnerLegalEntity"`
			ThirdPartyProtection struct {
				Text            string `xml:",chardata"`
				Confidentiality string `xml:"confidentiality"`
				Justification   string `xml:"justification"`
			} `xml:"ThirdPartyProtection"`
			ThirdParty         string `xml:"ThirdParty"`
			ContactPersons     string `xml:"ContactPersons"`
			ReferenceSubstance struct {
				Text       string `xml:",chardata"`
				Protection struct {
					Text            string `xml:",chardata"`
					Confidentiality string `xml:"confidentiality"`
					Justification   string `xml:"justification"`
				} `xml:"Protection"`
				ReferenceSubstance string `xml:"ReferenceSubstance"`
			} `xml:"ReferenceSubstance"`
			TypeOfSubstance struct {
				Text        string `xml:",chardata"`
				Composition string `xml:"Composition"`
				Origin      string `xml:"Origin"`
			} `xml:"TypeOfSubstance"`
			RoleInSupplyChain struct {
				Text           string `xml:",chardata"`
				RoleProtection struct {
					Text            string `xml:",chardata"`
					Confidentiality string `xml:"confidentiality"`
					Justification   string `xml:"justification"`
				} `xml:"RoleProtection"`
				Manufacturer struct {
					Text string `xml:",chardata"`
					Nil  string `xml:"nil,attr"`
				} `xml:"Manufacturer"`
				Importer struct {
					Text string `xml:",chardata"`
					Nil  string `xml:"nil,attr"`
				} `xml:"Importer"`
				OnlyRepresentative struct {
					Text string `xml:",chardata"`
					Nil  string `xml:"nil,attr"`
				} `xml:"OnlyRepresentative"`
				DownstreamUser struct {
					Text string `xml:",chardata"`
					Nil  string `xml:"nil,attr"`
				} `xml:"DownstreamUser"`
			} `xml:"RoleInSupplyChain"`
		} `xml:"SUBSTANCE"`
	} `xml:"Content"`
	Attachments struct {
		Text           string `xml:",chardata"`
		I6a            string `xml:"i6a,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	} `xml:"Attachments"`
	ModificationHistory struct {
		Text           string `xml:",chardata"`
		I6h            string `xml:"i6h,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	} `xml:"ModificationHistory"`
}
type refSubstance struct {
	XMLName          xml.Name `xml:"Document"`
	Text             string   `xml:",chardata"`
	I6c              string   `xml:"i6c,attr"`
	Xsi              string   `xml:"xsi,attr"`
	SchemaLocation   string   `xml:"schemaLocation,attr"`
	XML              string   `xml:"xml,attr"`
	PlatformMetadata struct {
		Text                  string `xml:",chardata"`
		I6m                   string `xml:"i6m,attr"`
		SchemaLocation        string `xml:"schemaLocation,attr"`
		IuclidVersion         string `xml:"iuclidVersion"`
		DocumentKey           string `xml:"documentKey"`
		ParentDocumentKey     string `xml:"parentDocumentKey"`
		Name                  string `xml:"name"`
		DocumentType          string `xml:"documentType"`
		DocumentSubType       string `xml:"documentSubType"`
		OrderInSectionNo      string `xml:"orderInSectionNo"`
		DefinitionVersion     string `xml:"definitionVersion"`
		CreationDate          string `xml:"creationDate"`
		LastModificationDate  string `xml:"lastModificationDate"`
		SubmissionType        string `xml:"submissionType"`
		SubmissionTypeVersion string `xml:"submissionTypeVersion"`
		SubmittingLegalEntity string `xml:"submittingLegalEntity"`
		DossierSubject        string `xml:"dossierSubject"`
		I5Origin              string `xml:"i5Origin"`
		CreationTool          string `xml:"creationTool"`
		SnapshotCreationTool  string `xml:"snapshotCreationTool"`
	} `xml:"PlatformMetadata"`
	Content struct {
		Text               string `xml:",chardata"`
		REFERENCESUBSTANCE struct {
			Text           string `xml:",chardata"`
			Xmlns          string `xml:"xmlns,attr"`
			I6             string `xml:"i6,attr"`
			DataProtection struct {
				Text            string `xml:",chardata"`
				Confidentiality string `xml:"confidentiality"`
				Justification   string `xml:"justification"`
			} `xml:"DataProtection"`
			ReferenceSubstanceName string `xml:"ReferenceSubstanceName"`
			IupacName              string `xml:"IupacName"`
			Inventory              struct {
				Text           string `xml:",chardata"`
				InventoryEntry struct {
					Text  string `xml:",chardata"`
					Entry struct {
						Text              string `xml:",chardata"`
						InventoryCode     string `xml:"inventoryCode"`
						NumberInInventory string `xml:"numberInInventory"`
					} `xml:"entry"`
				} `xml:"InventoryEntry"`
				InventoryEntryJustification string `xml:"InventoryEntryJustification"`
				CASNumber                   string `xml:"CASNumber"`
				CASName                     string `xml:"CASName"`
			} `xml:"Inventory"`
			Synonyms struct {
				Text     string `xml:",chardata"`
				Synonyms struct {
					Text  string `xml:",chardata"`
					Entry struct {
						Text           string `xml:",chardata"`
						Uuid           string `xml:"uuid,attr"`
						DataProtection struct {
							Text            string `xml:",chardata"`
							Confidentiality string `xml:"confidentiality"`
							Justification   string `xml:"justification"`
						} `xml:"DataProtection"`
						Identifier struct {
							Text  string `xml:",chardata"`
							Value string `xml:"value"`
						} `xml:"Identifier"`
						Name string `xml:"Name"`
					} `xml:"entry"`
				} `xml:"Synonyms"`
			} `xml:"Synonyms"`
			MolecularStructuralInfo struct {
				Text           string `xml:",chardata"`
				DataProtection struct {
					Text            string `xml:",chardata"`
					Confidentiality string `xml:"confidentiality"`
					Justification   string `xml:"justification"`
				} `xml:"DataProtection"`
				MolecularFormula       string `xml:"MolecularFormula"`
				MolecularWeightRange   string `xml:"MolecularWeightRange"`
				SmilesNotation         string `xml:"SmilesNotation"`
				InChl                  string `xml:"InChl"`
				StructuralFormula      string `xml:"StructuralFormula"`
				ChemicalStructureFiles string `xml:"ChemicalStructureFiles"`
			} `xml:"MolecularStructuralInfo"`
			RelatedSubstances struct {
				Text              string `xml:",chardata"`
				RelatedSubstances string `xml:"RelatedSubstances"`
			} `xml:"RelatedSubstances"`
		} `xml:"REFERENCE_SUBSTANCE"`
	} `xml:"Content"`
	Attachments struct {
		Text           string `xml:",chardata"`
		I6a            string `xml:"i6a,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	} `xml:"Attachments"`
	ModificationHistory struct {
		Text           string `xml:",chardata"`
		I6h            string `xml:"i6h,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	} `xml:"ModificationHistory"`
}

type productDocument struct {
	XMLName          xml.Name `xml:"Document"`
	Text             string   `xml:",chardata"`
	I6c              string   `xml:"i6c,attr"`
	Xsi              string   `xml:"xsi,attr"`
	SchemaLocation   string   `xml:"schemaLocation,attr"`
	XML              string   `xml:"xml,attr"`
	PlatformMetadata struct {
		Text                  string `xml:",chardata"`
		I6m                   string `xml:"i6m,attr"`
		SchemaLocation        string `xml:"schemaLocation,attr"`
		IuclidVersion         string `xml:"iuclidVersion"`
		DocumentKey           string `xml:"documentKey"`
		ParentDocumentKey     string `xml:"parentDocumentKey"`
		Name                  string `xml:"name"`
		DocumentType          string `xml:"documentType"`
		DocumentSubType       string `xml:"documentSubType"`
		OrderInSectionNo      string `xml:"orderInSectionNo"`
		DefinitionVersion     string `xml:"definitionVersion"`
		CreationDate          string `xml:"creationDate"`
		LastModificationDate  string `xml:"lastModificationDate"`
		SubmissionType        string `xml:"submissionType"`
		SubmissionTypeVersion string `xml:"submissionTypeVersion"`
		SubmittingLegalEntity string `xml:"submittingLegalEntity"`
		DossierSubject        string `xml:"dossierSubject"`
		I5Origin              string `xml:"i5Origin"`
		CreationTool          string `xml:"creationTool"`
		SnapshotCreationTool  string `xml:"snapshotCreationTool"`
	} `xml:"PlatformMetadata"`
	Content struct {
		Text                             string `xml:",chardata"`
		FLEXIBLERECORDMixtureComposition struct {
			Text               string `xml:",chardata"`
			Xmlns              string `xml:"xmlns,attr"`
			I6                 string `xml:"i6,attr"`
			AdministrativeData struct {
				Text           string `xml:",chardata"`
				DataProtection struct {
					Text            string `xml:",chardata"`
					Confidentiality string `xml:"confidentiality"`
					Justification   string `xml:"justification"`
				} `xml:"DataProtection"`
			} `xml:"AdministrativeData"`
			GeneralInformation struct {
				Text       string `xml:",chardata"`
				Name       string `xml:"Name"`
				TradeNames []struct {
					Text  string `xml:",chardata"`
					Entry struct {
						Text    string `xml:",chardata"`
						Uuid    string `xml:"uuid,attr"`
						Country struct {
							Text  string `xml:",chardata"`
							Value string `xml:"value"`
						} `xml:"Country"`
						TradeName string `xml:"TradeName"`
					} `xml:"entry"`
				} `xml:"TradeNames"`
				FormulationType struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value"`
				} `xml:"FormulationType"`
			} `xml:"GeneralInformation"`
			Components struct {
				Text       string `xml:",chardata"`
				Components struct {
					Text  string `xml:",chardata"`
					Entry []struct {
						Text           string `xml:",chardata"`
						Uuid           string `xml:"uuid,attr"`
						DataProtection struct {
							Text            string `xml:",chardata"`
							Confidentiality string `xml:"confidentiality"`
							Justification   string `xml:"justification"`
						} `xml:"DataProtection"`
						Reference string `xml:"Reference"`
						Function  struct {
							Text  string `xml:",chardata"`
							Value string `xml:"value"`
							Other string `xml:"other"`
						} `xml:"Function"`
						TypicalConcentration struct {
							Text      string `xml:",chardata"`
							UnitCode  string `xml:"unitCode"`
							Qualifier string `xml:"qualifier"`
							Value     string `xml:"value"`
						} `xml:"TypicalConcentration"`
						ConcentrationRange string `xml:"ConcentrationRange"`
						SubstanceOfConcern string `xml:"SubstanceOfConcern"`
						Gci                struct {
							Text string `xml:",chardata"`
							Nil  string `xml:"nil,attr"`
						} `xml:"Gci"`
						Icg struct {
							Text string `xml:",chardata"`
							Nil  string `xml:"nil,attr"`
						} `xml:"Icg"`
						Sfc struct {
							Text string `xml:",chardata"`
							Nil  string `xml:"nil,attr"`
						} `xml:"Sfc"`
						SubstanceGeneratedInSitu string `xml:"SubstanceGeneratedInSitu"`
					} `xml:"entry"`
				} `xml:"Components"`
			} `xml:"Components"`
		} `xml:"FLEXIBLE_RECORD.MixtureComposition"`
	} `xml:"Content"`
	Attachments struct {
		Text           string `xml:",chardata"`
		I6a            string `xml:"i6a,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	} `xml:"Attachments"`
	ModificationHistory struct {
		Text           string `xml:",chardata"`
		I6h            string `xml:"i6h,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	} `xml:"ModificationHistory"`
}
