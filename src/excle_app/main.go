package main

import (
    "fmt"
    "os"
	"excelize"
	"strconv"
)

func main() {
//	read
    xlsx, err := excelize.OpenFile("./Workbook.xlsx")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    // Get value from cell by given sheet index and axis.
    // cell := xlsx.GetCellValue("Sheet1", "B2")
    // fmt.Println(cell)

	// Get sheet index.
    index_1 := xlsx.GetSheetIndex("Sheet1")
    // Get all the rows in a sheet.
    rows_1 := xlsx.GetRows("sheet" + strconv.Itoa(index_1))
    for _, row_1 := range rows_1 {
        for _, colCell_1 := range row_1 {
            fmt.Print(colCell_1, "\t")
        }
        fmt.Println()
    }

	// Get sheet index.
    index_2 := xlsx.GetSheetIndex("Sheet2")
    // Get all the rows in a sheet.
    rows_2 := xlsx.GetRows("sheet" + strconv.Itoa(index_2))
    for _, row_2 := range rows_2 {
        for _, colCell_2 := range row_2 {
            fmt.Print(colCell_2, "\t")
        }
        fmt.Println()
    }
}

	// Here is a minimal example usage that will create XLSX file.
    // xlsx := excelize.NewFile()
    // // Create a new sheet.
    // xlsx.NewSheet(2, "Sheet2")
    // // Set value of a cell.
    // xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
    // xlsx.SetCellValue("Sheet1", "B2", 100)
    // // Set active sheet of the workbook.
    // xlsx.SetActiveSheet(2)
    // // Save xlsx file by the given path.
    // err := xlsx.SaveAs("./Workbook.xlsx")
    // if err != nil {
    //     fmt.Println(err)
    //     os.Exit(1)
    // }

	// read
    // xlsx, err := excelize.OpenFile("./Workbook.xlsx")
    // if err != nil {
    //     fmt.Println(err)
    //     os.Exit(1)
    // }
    // // Get value from cell by given sheet index and axis.
    // cell := xlsx.GetCellValue("Sheet1", "B2")
    // fmt.Println(cell)
    // // Get sheet index.
    // index := xlsx.GetSheetIndex("Sheet2")
    // // Get all the rows in a sheet.
    // rows := xlsx.GetRows("sheet" + strconv.Itoa(index))
    // for _, row := range rows {
    //     for _, colCell := range row {
    //         fmt.Print(colCell, "\t")
    //     }
    //     fmt.Println()
    // }
