package productHttpDelivery

// ResponseError represent the reseponse error struct
/*type ResponseError struct {
	Message string `json:"message"`
}

// BookHandler  represent the httphandler for product
type BookHandler struct {
	BUsercase domain.BookUsecase
}

// New will initialize the product resources endpoint
func New(e *echo.Echo, bu domain.BookUsecase) {
	handler := &BookHandler{
		BUsercase: bu,
	}
	e.GET("/books", handler.Fetch)
	e.POST("/books", handler.StoreMany)
	e.PUT("/product", handler.Update)
	e.DELETE("/product/:id", handler.Delete)
}

// Fetch will fetch the product based on given params
func (b *BookHandler) Fetch(c echo.Context) error {
	limitS := c.QueryParam("limit")
	limit, _ := strconv.Atoi(limitS)

	listBook, err := b.BUsercase.ListBooks(int64(limit), reflect.TypeOf(domain.Book{}))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listBook)
}

// StoreMany will store the books by given request body
func (b *BookHandler) StoreMany(c echo.Context) error {
	books := new([]domain.Book)
	err := c.Bind(&books)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValidSlice(books); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := b.BUsercase.InsertBooks(books)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

// Update will update product by given param
func (b *BookHandler) Update(c echo.Context) error {
	product := new(domain.Book)
	err := c.Bind(&product)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	_, err = b.BUsercase.UpdateBook(*product)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusOK)
}

// Delete will delete product by given param
func (b *BookHandler) Delete(c echo.Context) error {
	_, err := b.BUsercase.DeleteBook(c.Param("id"))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusOK)
}

func isRequestValid(product *domain.Book) (bool, error) {
	return true, nil
}

func isRequestValidSlice(books *[]domain.Book) (bool, error) {
	for _, product := range *books {
		return isRequestValid(&product)
	}

	return true, nil
}*/
