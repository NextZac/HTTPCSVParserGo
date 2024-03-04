package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func search(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	name := query.Get("name")
	//Check if MemoryStore is empty and return "No Data" if it is
	if len(memoryStores) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No Data"})
		return
	}
	fmt.Println(name)
	result := SearchViaName(name)
	ctx.JSON(http.StatusOK, result)
}

func pageination(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	page := query.Get("page")
	// Parse the page to int and pass it to the Pagination function
	if len(memoryStores) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No Data"})
		return
	}
	pageInt := 0
	if page != "" {
		pageInt, _ = strconv.Atoi(page)
	}
	result := Pagination(pageInt)
	ctx.JSON(http.StatusOK, result)

}

func homePage(ctx *gin.Context) {
	ctx.String(http.StatusOK, fmt.Sprintf("Hello"))
}

func textUpload(ctx *gin.Context) {
	files, _ := os.ReadDir("public/uploads")
	var WasUploadedBefore = false
	if len(files) > 0 {
		WasUploadedBefore = true
		os.RemoveAll("public/uploads")
		os.MkdirAll("public/uploads", 0777)
	}

	file, header, _ := ctx.Request.FormFile("file")
	fileExt := filepath.Ext(header.Filename)
	orgFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(orgFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	out, err := os.Create("public/uploads/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	// Adding the read file to the In Memory Store
	readParse("public/uploads/" + filename)

	if WasUploadedBefore {
		ctx.JSON(http.StatusOK, gin.H{"filename": filename, "message": "File was Replaced"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"filename": filename})
	}
}

func main() {
	files, _ := os.ReadDir("public/uploads")
	if len(files) > 0 {
		for _, file := range files {
			readParse("public/uploads/" + file.Name())
		}
	}
	r := gin.Default()
	r.StaticFS("/uploads", http.Dir("public"))
	r.GET("/", homePage)
	r.GET("/search", search)
	r.GET("/pagination", pageination)
	r.POST("/upload", textUpload)
	r.Run()
}
