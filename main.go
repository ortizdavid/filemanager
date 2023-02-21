package main

import (
	"github.com/ortizdavid/filemanager/core"
)

func main() {

	var fileHandler *core.FileHandler

	// Folders
	rootFolder := "Files"

	mediaFolder := rootFolder+"/Media"
	musicFolder := mediaFolder+"/Music"
	imagesFolder := mediaFolder+"/Images"

	documentsFolder := rootFolder+"/Documents"
	docxDocumentsFolder := documentsFolder+"/docx"
	txtDocumentsFolder := documentsFolder+"/txt"

	// FILES
	txtFile1 := "file1.txt"
	txtFile2 := "file2.txt"
	txtFile3 := "file3.txt"
	docxFile := "file1.docx"

	// Create sigle Folders
	fileHandler.CreateFolder(rootFolder)

	// Create many folders
	fileHandler.CreateFolder(mediaFolder)
	fileHandler.CreateFolder(documentsFolder)
	fileHandler.CreateFolderAll(musicFolder)
	fileHandler.CreateFolderAll(imagesFolder)
	fileHandler.CreateFolderAll(docxDocumentsFolder)
	fileHandler.CreateFolderAll(txtDocumentsFolder)

	// Create many Files
	fileHandler.CreateFileAll(txtDocumentsFolder, txtFile1, txtFile2, txtFile3)

	// Create single File
	fileHandler.CreateFile(docxDocumentsFolder, docxFile)


	// Write in files
	fileHandler.WriteFile(txtDocumentsFolder, txtFile1, "Hello\nWorld")
	fileHandler.WriteFile(txtDocumentsFolder, txtFile2, "I am Here")
	fileHandler.WriteFile(txtDocumentsFolder, txtFile3, "123456789")

	//List Contents in some folder
	fileHandler.ListFiles(txtDocumentsFolder)

	//Get File Info
	fileHandler.GetFileInfo(txtDocumentsFolder+"/"+txtFile1)	
	
}

