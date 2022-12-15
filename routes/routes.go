package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/controllers"
	m "github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/middlewares"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	v1 := e.Group("/api/v1")

	subjectsV1 := v1.Group("/subjects")
	subjectsV1.GET("", controllers.GetAllSubject)
	subjectsV1.POST("", controllers.CreateSubject)
	subjectsV1.PUT("/:id", controllers.UpdateSubjectById)
	subjectsV1.DELETE("/:id", controllers.DeleteSubjectById)

	publishersV1 := v1.Group("/publishers")
	publishersV1.GET("", controllers.GetAllPublisher)
	publishersV1.POST("", controllers.CreatePublisher)
	publishersV1.PUT("/:id", controllers.UpdatePublisherById)
	publishersV1.DELETE("/:id", controllers.DeletePublisherById)

	usersV1 := v1.Group("/users")
	usersV1.GET("", controllers.GetAllUser)
	usersV1.POST("", controllers.CreateUser)
	usersV1.PUT("/:id", controllers.UpdateUserById)
	usersV1.DELETE("/:id", controllers.DeleteUserById)

	booksV1 := v1.Group("/books")
	booksV1.GET("", controllers.GetAllBook)
	booksV1.POST("", controllers.CreateBook)
	booksV1.PUT("/:id", controllers.UpdateBookById)
	booksV1.DELETE("/:id", controllers.DeleteBookById)

	e.Validator = &m.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = m.ErrorHandler

	return e
}
