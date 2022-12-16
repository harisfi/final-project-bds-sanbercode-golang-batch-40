package routes

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/controllers"
	m "github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	jwtMiddleware := middleware.JWT([]byte(os.Getenv("JWT_SECRET")))

	v1 := e.Group("/api/v1")

	subjectsV1 := v1.Group("/subjects")
	subjectsV1Auth := v1.Group("/subjects", jwtMiddleware)

	subjectsV1.GET("", controllers.GetAllSubject)
	subjectsV1Auth.POST("", controllers.CreateSubject)
	subjectsV1Auth.PUT("/:id", controllers.UpdateSubjectById)
	subjectsV1Auth.DELETE("/:id", controllers.DeleteSubjectById)

	publishersV1 := v1.Group("/publishers")
	publishersV1Auth := v1.Group("/publishers", jwtMiddleware)

	publishersV1.GET("", controllers.GetAllPublisher)
	publishersV1Auth.POST("", controllers.CreatePublisher)
	publishersV1Auth.PUT("/:id", controllers.UpdatePublisherById)
	publishersV1Auth.DELETE("/:id", controllers.DeletePublisherById)

	usersV1 := v1.Group("/users")
	usersV1Auth := v1.Group("/users", jwtMiddleware)

	usersV1.GET("", controllers.GetAllUser)
	usersV1.POST("", controllers.RegisterUser)
	usersV1.POST("/login", controllers.LoginUser)
	usersV1Auth.PUT("/:id", controllers.UpdateUserById)
	usersV1Auth.DELETE("/:id", controllers.DeleteUserById)

	booksV1 := v1.Group("/books")
	booksV1Auth := v1.Group("/books", jwtMiddleware)

	booksV1.GET("", controllers.GetAllBook)
	booksV1Auth.POST("", controllers.CreateBook)
	booksV1Auth.GET("/:id/borrow", controllers.BorrowBook)
	booksV1Auth.GET("/:id/return", controllers.ReturnBook)
	booksV1Auth.PUT("/:id", controllers.UpdateBookById)
	booksV1Auth.DELETE("/:id", controllers.DeleteBookById)

	e.Validator = &m.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = m.ErrorHandler

	return e
}
