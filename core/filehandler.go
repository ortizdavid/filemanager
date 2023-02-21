package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"github.com/ortizdavid/filemanager/helpers"
)

type FileHandler struct {}

func (fm *FileHandler) GetFileExt(path string) string {
	return filepath.Ext(path)
}

func (fm *FileHandler) GetFileType(ext string) string {
	var imageExts = []string{".png", ".gif", ".jpg", ".jiff"}
	var documentExts = []string{".txt", ".pdf", ".docx", ".ppt", ".pptx"}
	var audioExts = []string{".mp3", ".aac", ".wav", ".flac"}
	var videoExts = []string{".mp4", ".mkv", ".avi", ".flv"}
	var fileType string = ""

	if helpers.Contains(imageExts, ext) {
		fileType = "Image"
	} else if helpers.Contains(documentExts, ext) {
		fileType = "Document"
	} else if helpers.Contains(audioExts, ext) {
		fileType = "Audio"
	} else if helpers.Contains(documentExts, ext) {
		fileType = "Document"
	} else if helpers.Contains(videoExts, ext) {
		fileType = "Video"
	} else {
		fileType = "Other"
	}
	return fileType
}

func (fm *FileHandler) GetFileInfo(fileName string) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nFile Name:", fileInfo.Name())  
	fmt.Println("Extension:", fm.GetFileExt(fileName))          
	fmt.Println("Size:", fileInfo.Size(), " bytes")  
	fmt.Println("Type:", fm.GetFileType(fm.GetFileExt(fileName)))              
	fmt.Println("Last Modified:", fileInfo.ModTime()) 
	fmt.Println("Permissions:", fileInfo.Mode())     
}

func (fm *FileHandler) ListFiles(dir string)  {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nAll Files in '%s':\n", dir)
	helpers.PrintChar("-", 125)
	fmt.Println("NAME\t\t\t\tTYPE\t\t\tSIZE(Bytes)\t\t\tLAST MODIFIED")
	helpers.PrintChar("-", 125)
	for _, file := range files {
		var ext = fm.GetFileExt((file.Name()))
		fmt.Printf("%s\t\t\t%s\t\t\t%d\t\t\t%s\n", file.Name(), fm.GetFileType(ext), file.Size(), file.ModTime())
	}
}

func (fm *FileHandler) CreateFolder(folderName string) {
	folder := os.Mkdir(folderName, os.FileMode(1))
	if folder != nil {
		fmt.Println(folder.Error())
	} else {
		fmt.Printf(helpers.FOLDER_CREATED, folderName)
	}
}

func (fm *FileHandler) CreateFolderAll(folderName string) {
	folder := os.MkdirAll(folderName, os.FileMode(1))
	if folder != nil {
		fmt.Println(folder.Error())
	} else {
		fmt.Printf(helpers.FOLDER_CREATED, folderName)
	}
}

func (fm *FileHandler) CreateFile(dirName string, fileName string) {
	file, err := os.Create(dirName +"/"+ fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf(helpers.FILE_CREATED, file.Name())
	}
}

func (fm *FileHandler) CreateFileAll(dirName string, files ...string) {
	for _, file := range files {
		fm.CreateFile(dirName, file)
	}
}

func (fm *FileHandler) MoveFile(fileName string, origin string, destination string) {
	err := os.Rename(origin +"/"+ fileName, destination +"/"+ fileName)
	if err != nil {
		fmt.Println(err)
	}
}

func (fm *FileHandler) WriteFile(folderName string, fileName string, content string) {
	file, err := os.OpenFile(folderName +"/"+ fileName, os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		file.WriteString(content)
	}
	file.Close()
}

