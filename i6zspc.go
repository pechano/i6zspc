package main

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/martinlindhe/inputbox"
	"github.com/sqweek/dialog"
	"github.com/xuri/excelize/v2"
)

func main() {

	filename := loadi6z()
	tempfolder := extractFiles(filename)

	regNo, ok := inputbox.InputBox("BPF familieoversigt", "Indtast registreringsnummer for BPF", "xx-x")
	if ok {
		fmt.Println("Registration number:", regNo)
	} else {
		fmt.Println("No value entered")
	}
	//	tempfolder := extractFiles("sea1.i6z") //for testing. remove this line and uncomment above for production
	SPC := extractMetaSPC(tempfolder)
	defer os.RemoveAll(tempfolder)
	fmt.Println(SPC.BPFname)
	start, err := excelize.JoinCellName("a", 1)
	if err != nil {
		fmt.Println(err)
	}

	sheet := excelize.NewFile()
	defer func() {
		if err := sheet.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.

	// Set value of a cell.
	sheet.SetCellValue("Sheet1", start, "Familieoversigt")
	maxProds := numberOfProducts(SPC)
	maxSOCs := numberOfSOCs(SPC)
	maxSOCs = 0 //SOCs are not needed in the overview file
	maxASs := numberOfASs(SPC)

	maxlines := maxProds + maxSOCs + maxASs - 1

	sheet.SetCellValue("Sheet1", "A3", "BPF registreringsnummer:")
	sheet.SetColWidth("Sheet1", "A", "Z", 25)
	sheet.SetCellValue("Sheet1", "b3", regNo)
	sheet.SetCellValue("Sheet1", "a4", "BPF navn:")
	sheet.SetCellValue("Sheet1", "b4", SPC.BPFname)

	sheet.SetCellValue("Sheet1", "a5", "Aktivstof CAS:")
	sheet.SetCellValue("Sheet1", "b5", SPC.MetaSPC[0].AS[0].ASCAS)

	sheet.SetCellValue("Sheet1", "a6", "Aktivstof:")
	sheet.SetCellValue("Sheet1", "b6", SPC.MetaSPC[0].AS[0].ASName)

	sheet.SetCellValue("Sheet1", "a8", "MetaSPC#:")
	sheet.SetCellValue("Sheet1", "a9", "MetaSPC navn:")
	sheet.SetCellValue("Sheet1", "a10", "MetaSPC reg. nr.:")
	sheet.SetCellValue("Sheet1", "a11", "Aktivstof (%w/w):")
	sheet.SetCellValue("Sheet1", "a12", "Produkter:")

	for x := 0; x < maxlines; x++ {
		row := x + 13

		sheet.SetCellValue("Sheet1", "a"+fmt.Sprint(row), regNo+"-x-"+fmt.Sprint(x+1))

	}
	check(err)

	for i := 0; i < len(SPC.MetaSPC); i++ {
		activeColumn, err := excelize.ColumnNumberToName(i + 2)
		activeRow := 8
		check(err)

		activeCell, _ := excelize.JoinCellName(activeColumn, activeRow)
		sheet.SetCellValue("Sheet1", activeCell, i+1)

		activeRow++
		activeCell, _ = excelize.JoinCellName(activeColumn, activeRow)
		sheet.SetCellValue("Sheet1", activeCell, SPC.MetaSPC[i].Name)

		activeRow++
		activeCell, _ = excelize.JoinCellName(activeColumn, activeRow)
		sheet.SetCellFormula("Sheet1", activeCell, "CONCATENATE(B3,\"-\","+activeColumn+"8"+",\"-x\")")
		// HER SKER DER NOGET KRITISK HVIS DER ER FLERE AKTIVSTOFFER
		activeRow++
		activeCell, _ = excelize.JoinCellName(activeColumn, activeRow)
		sheet.SetCellValue("Sheet1", activeCell, SPC.MetaSPC[i].AS[0].ConcL+" - "+SPC.MetaSPC[i].AS[0].ConcU)

		activeRow++
		for j := 0; j < len(SPC.MetaSPC[i].Products); j++ {

			activeRow++
			activeCell, _ = excelize.JoinCellName(activeColumn, activeRow)
			sheet.SetCellValue("Sheet1", activeCell, SPC.MetaSPC[i].Products[j].Name)
		}

	}
	sheetFolder := filepath.Dir(tempfolder)
	outputsheet := filepath.Join(sheetFolder, SPC.BPFname+".xlsx")
	if err := sheet.SaveAs(outputsheet); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Saving spreadsheet as: " + outputsheet)

	tradenamefilePath := filepath.Join(sheetFolder, SPC.BPFname+" BMD.txt")
	os.Create(tradenamefilePath)
	tradenamefile, err := os.OpenFile(tradenamefilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	check(err)
	fmt.Fprint(tradenamefile, "BP/BPF navn: ")
	fmt.Fprintln(tradenamefile, SPC.BPFname)
	fmt.Fprint(tradenamefile, "Produkter, handelsnavne og anvendelser fra SPC:"+"\n"+"\n")
	meta := 1
	for _, metaSlice := range SPC.MetaSPC {
		prod := 1
		fmt.Fprintln(tradenamefile, "MetaSPC:"+metaSlice.Name+"("+regNo+"-"+fmt.Sprint(meta)+")"+" - AS konc.: "+metaSlice.AS[0].ConcL+" - "+metaSlice.AS[0].ConcU+"%")
		fmt.Fprintln(tradenamefile, "Produkter og handelsnavne:")
		for _, prodSlice := range metaSlice.Products {
			fmt.Fprintln(tradenamefile, "    "+prodSlice.Name+"("+regNo+"-"+fmt.Sprint(meta)+"-"+fmt.Sprint(prod)+")"+" - AS konc: "+prodSlice.ASConc+"%")
			for _, tradename := range prodSlice.Tradenames {

				fmt.Fprintln(tradenamefile, "        "+tradename)
			}
			prod++
		}

		/*
			fmt.Fprintln(tradenamefile, "Anvendelser:")
			for _, useSlice := range metaSlice.Uses.Use {

				fmt.Fprintln(tradenamefile, "    "+useSlice.Name)
				fmt.Fprintln(tradenamefile, "    "+useSlice.Field+"\n")
			}

			meta++

		*/
		fmt.Fprintln(tradenamefile, "")
	}
	exit := ""
	fmt.Print("Press ENTER to exit")
	fmt.Scanln(&exit)
}

func numberOfProducts(s BPFfile) (maxProducts int) {
	var products []int
	for _, meta := range s.MetaSPC {
		productsInMetaSPC := len(meta.Products)
		products = append(products, productsInMetaSPC)
	}
	maxProducts = slices.Max(products)
	return maxProducts
}

func numberOfASs(s BPFfile) (maxAS int) {
	var ASs []int
	for _, meta := range s.MetaSPC {
		AS := len(meta.AS)
		ASs = append(ASs, AS)
	}
	maxAS = slices.Max(ASs)
	return maxAS
}
func numberOfSOCs(s BPFfile) (maxSOCs int) {
	var SOCs []int
	for _, meta := range s.MetaSPC {
		SOCsinMetaSPC := len(meta.SOC)
		SOCs = append(SOCs, SOCsinMetaSPC)
	}
	maxSOCs = slices.Max(SOCs)
	return maxSOCs
}
func extractMetaSPC(path string) (baseinfo BPFfile) {

	spcPath := filepath.Join(path, "manifest.xml")

	SPC, err := os.Open(spcPath)

	if err != nil {
		log.Println(err.Error())
	}
	defer SPC.Close()
	byteValue, _ := io.ReadAll(SPC)
	var manifest Manifest
	err = xml.Unmarshal(byteValue, &manifest)

	i := 0
	for _, doc := range manifest.ContainedDocuments.Document {
		if doc.Subtype == "ProductSummaryComposition" {
			i++
			var metaSPC MetaSPC
			docID := strings.Replace(doc.ID, "/", "_", 1)

			TopDocument, err := os.Open(filepath.Join(path, docID+".i6d"))
			if err != nil {
				log.Println(err.Error())
			}
			defer TopDocument.Close()
			topDocBytes, _ := io.ReadAll(TopDocument)
			var TopDoc Document
			err = xml.Unmarshal(topDocBytes, &TopDoc)

			if TopDoc.PlatformMetadata.OrderInSectionNo == "0" {
				baseinfo.BPFname = TopDoc.PlatformMetadata.Name
			}

			metaSPC.Name = TopDoc.PlatformMetadata.Name
			for _, compound := range TopDoc.Content.FLEXIBLESUMMARYProductSummaryComposition.ProductSummaryComposition.Composition.Entry {
				if compound.SubstanceOfConcern == "false" {
					var substance Substance

					substanceFILEname := strings.Replace(compound.Name, "/", "_", 1)
					substanceFILE, err := os.Open(filepath.Join(path, substanceFILEname+".i6d"))
					check(err)
					defer substanceFILE.Close()
					substanceBYTES, _ := io.ReadAll(substanceFILE)
					err = xml.Unmarshal(substanceBYTES, &substance)

					var AS activeSubstance

					AS.ASName = substance.Content.SUBSTANCE.ChemicalName
					AS.ConcL = compound.Concentration.LowerValue
					AS.ConcU = compound.Concentration.UpperValue

					var REFsubstance refSubstance
					refSubstanceFileName := substance.Content.SUBSTANCE.ReferenceSubstance.ReferenceSubstance
					refSubstanceFileName = strings.Replace(refSubstanceFileName, "/", "_", 1)
					refSubstanceFile, err := os.Open(filepath.Join(path, refSubstanceFileName+".i6d"))

					check(err)
					defer refSubstanceFile.Close()
					REFsubstanceBYTES, _ := io.ReadAll(refSubstanceFile)
					err = xml.Unmarshal(REFsubstanceBYTES, &REFsubstance)
					AS.ASCAS = REFsubstance.Content.REFERENCESUBSTANCE.Inventory.CASNumber

					if TopDoc.PlatformMetadata.OrderInSectionNo == "0" {
						baseinfo.BPFAS = append(baseinfo.BPFAS, AS)
						continue
					}

					metaSPC.AS = append(metaSPC.AS, AS)
				}
				if compound.SubstanceOfConcern == "true" {
					var substance refSubstance

					substanceFILEname := strings.Replace(compound.Name, "/", "_", 1)
					substanceFILE, err := os.Open(filepath.Join(path, substanceFILEname+".i6d"))
					check(err)
					defer substanceFILE.Close()
					substanceBYTES, _ := io.ReadAll(substanceFILE)
					err = xml.Unmarshal(substanceBYTES, &substance)

					var SOC soc

					SOC.SOCName = substance.Content.REFERENCESUBSTANCE.ReferenceSubstanceName
					SOC.SOCconcL = compound.Concentration.LowerValue
					SOC.SOCconcU = compound.Concentration.UpperValue
					SOC.SOCCAS = substance.Content.REFERENCESUBSTANCE.Inventory.CASNumber

					if TopDoc.PlatformMetadata.OrderInSectionNo == "0" {
						baseinfo.BPFSOC = append(baseinfo.BPFSOC, SOC)
						continue
					}

					metaSPC.SOC = append(metaSPC.SOC, SOC)
				}

			}
			if TopDoc.PlatformMetadata.OrderInSectionNo == "0" {
				continue
			}

			for _, productID := range TopDoc.Content.FLEXIBLESUMMARYProductSummaryComposition.ProductCompositions.ProductCompositions.Key {
				var product product
				var productDoc productDocument
				fmt.Println(productID)
				productFilename := strings.Replace(productID, "/", "_", 1)

				productFile, err := os.Open(filepath.Join(path, productFilename+".i6d"))
				check(err)
				defer productFile.Close()
				productBytes, _ := io.ReadAll(productFile)
				err = xml.Unmarshal(productBytes, &productDoc)
				product.Name = productDoc.PlatformMetadata.Name

				for _, tradename := range productDoc.Content.FLEXIBLERECORDMixtureComposition.GeneralInformation.TradeNames {

					product.Tradenames = append(product.Tradenames, tradename.Entry.TradeName)
				}

				for _, entry := range productDoc.Content.FLEXIBLERECORDMixtureComposition.Components.Components.Entry {
					if entry.Function.Value == "1699" {
						product.ASConc = entry.TypicalConcentration.Value
					}
				}
				metaSPC.Products = append(metaSPC.Products, product)
			}
			baseinfo.MetaSPC = append(baseinfo.MetaSPC, metaSPC)

		}

	}
	return baseinfo
}

func loadi6z() (filename string) {
	filename, err := dialog.File().Filter("SPC", "i6z").Load()
	if err != nil {
		log.Println(err.Error())
	}

	return filename
}

func extractFiles(i6zPath string) (tempfolder string) {
	folder := filepath.Dir(i6zPath)
	tempFolderPath := filepath.Join(folder, "temp")
	i6z, err := zip.OpenReader(i6zPath)

	if err != nil {
		log.Print(err.Error())
	}
	defer i6z.Close()

	err = os.Mkdir(tempFolderPath, os.ModePerm)

	for _, f := range i6z.File {
		extractPath := filepath.Join(tempFolderPath, f.Name)
		// fmt.Printf("Extracting file:: %s\n", extractPath)
		//
		outFile, err := os.OpenFile(extractPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			log.Println(err.Error())

		}
		fileInArchive, err := f.Open()
		if err != nil {
			log.Println(err.Error())

		}
		if _, err := io.Copy(outFile, fileInArchive); err != nil {
			log.Println(err.Error())

		}

		outFile.Close()
		fileInArchive.Close()
	}
	fmt.Println("Files extracted")
	return tempFolderPath
}

//BPF STRUCTS****************************************

type BPFfile struct {
	BAS     string
	BPFname string
	BPFAS   []activeSubstance
	BPFSOC  []soc
	MetaSPC []MetaSPC
}

type MetaSPC struct {
	Name     string
	AS       []activeSubstance
	SOC      []soc
	Products []product
	Uses     []use
}

type activeSubstance struct {
	ASName string
	ASCAS  string
	ConcL  string
	ConcU  string
}

type soc struct {
	SOCName  string
	SOCCAS   string
	SOCconcL string
	SOCconcU string
}

type product struct {
	Name       string
	Tradenames []string
	ASConc     string
}
type use struct {
	Name  string
	Field string
}

func check(e error) {
	if e != nil {
		log.Println(e.Error())
	}
}
