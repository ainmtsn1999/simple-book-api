package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{}

func GetBook(ctx *gin.Context) {
	bookId := ctx.Param("bookId")
	stat := false
	var bookData Book

	for i, book := range BookDatas {
		id, _ := strconv.Atoi(bookId)
		if id == book.Id {
			stat = true
			bookData = BookDatas[i]
			break
		}
	}

	if !stat {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Book Data Not Found!",
		})
		return
	}

	ctx.JSON(http.StatusOK, bookData)
}

func GetAllBooks(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, BookDatas)
}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newBook.Id = len(BookDatas) + 1
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Book Data Created!",
	})
}

func UpdateBook(ctx *gin.Context) {
	bookId := ctx.Param("bookId")
	stat := false
	var updateBook Book

	if err := ctx.ShouldBindJSON(&updateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		id, _ := strconv.Atoi(bookId)
		if id == book.Id {
			stat = true
			updateBook.Id = id
			BookDatas[i] = updateBook
			break
		}
	}

	if !stat {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Book Data Not Found!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book Data Updated!",
	})
}

func DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("bookId")
	stat := false
	var bookIndex int

	for i, book := range BookDatas {
		id, _ := strconv.Atoi(bookId)
		if id == book.Id {
			stat = true
			bookIndex = i
			break
		}
	}

	if !stat {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Book Data Not Found!",
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book Data Deleted!",
	})
}
