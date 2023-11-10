package main

import (
	"fmt"
	"os"
	"strings"
)

func getFileInfo(){

	var path string
	
	fmt.Println("What path do u want organize (type in /home/user/...):   ")
	fmt.Scan(&path)


	files,err:=os.ReadDir(path)

	if err!=nil{
		fmt.Println("could not read files ",err)
	}
	allFiles:=[]string{}
	var newFile string
	fileExtensions:=[]string{}
	for _,file:=range files{
		newFile=strings.Replace(file.Name()," ","-",-1)
		err = os.Rename(path+file.Name(),path+newFile)
		if err!=nil{
			fmt.Println(err)
		}
		numOfDots:=strings.Count(newFile,".")
		if numOfDots<=0{
			fmt.Println("no file extension");
		}else if strings.Count(newFile,".")>1{
		
			slicedFiles:=strings.SplitN(newFile,".",numOfDots+1)
				
			fileExtensions = append(fileExtensions,slicedFiles[numOfDots])
		}else{


			allFiles=append(allFiles,newFile)
			//filePath:="/home/bonk/Downloads/" + file.Name()
			//fileInfo,err:=os.Stat(filePath)
			if err!=nil{
				fmt.Println("could not read file",err)
			}
			slicedFiles:=strings.SplitN(newFile,".",2)
			fileExtensions = append(fileExtensions,slicedFiles[1])
		
		}	
	}

	var docPath string

	fmt.Println(fileExtensions)
	fmt.Println(allFiles)
	i:=0
	for i=0;i!=len(fileExtensions);i++{
		//imgs
		
		if fileExtensions[i] == "jpg" || fileExtensions[i]=="png"{
			fmt.Println(i)


			err = os.Rename(path+allFiles[i],"/home/bonk/Pictures/"+allFiles[i])
			if err!=nil{
				fmt.Println(err)
			}
		}else if fileExtensions[i] == "docx" || fileExtensions[i] == "pdf" || fileExtensions[i]=="txt"{
			fmt.Println(i)
			
			
	
			fmt.Println("What path do u want organize your .docx , .pdf in(type in /home/user/...):   ")
			fmt.Scan(&docPath)
			err = os.Rename("/home/bonk/Downloads/"+allFiles[i],docPath+allFiles[i])
			if err!=nil{
				fmt.Println(err)
			}
		}else if fileExtensions[i] == "mp3"{ 
			fmt.Println(i)


			err = os.Rename(path+allFiles[i],"/home/bonk/Music/"+allFiles[i])
			if err!=nil{
				fmt.Println(err)
			}
		}
		


		
	}
}
func setPrefix(){	
	//var path string
	
	//fmt.Println("What path do u want organize (type in /home/user/...):   ")
	//fmt.Scan(&path)

	path:="/home/bonk/Downloads/"
	files,err:=os.ReadDir("/home/bonk/Downloads/")
	allFiles:=[]string{}
	if err!=nil{
		fmt.Println(err)
	}
	var i int
	for _,file:=range files{
		allFiles=append(allFiles,file.Name())
		suffix:=string(file.Name()[0])

		if suffix =="w"{
			err = os.Rename(path+allFiles[i],"/home/bonk/Documents/"+allFiles[i])
			if err!=nil{
				fmt.Println(err)
			}
		}
		i++
	}

}

func main(){
	getFileInfo()
	setPrefix()
}
