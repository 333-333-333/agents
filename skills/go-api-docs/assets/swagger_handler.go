// internal/booking/infrastructure/handler/http.go

// Create godoc
// @Summary Create a new booking
// @Description Creates a booking between a pet owner and a caregiver
// @Tags bookings
// @Accept json
// @Produce json
// @Param request body CreateBookingRequest true "Booking details"
// @Success 201 {object} server.Response{data=application.CreateBookingOutput}
// @Failure 400 {object} server.Response{error=server.Error}
// @Failure 401 {object} server.Response{error=server.Error}
// @Failure 422 {object} server.Response{error=server.Error}
// @Security BearerAuth
// @Router /bookings [post]
func (h *BookingHandler) Create(c *gin.Context) { ... }

// GetByID godoc
// @Summary Get booking by ID
// @Description Retrieves a booking by its unique identifier
// @Tags bookings
// @Produce json
// @Param id path string true "Booking ID" format(uuid)
// @Success 200 {object} server.Response{data=domain.Booking}
// @Failure 404 {object} server.Response{error=server.Error}
// @Security BearerAuth
// @Router /bookings/{id} [get]
func (h *BookingHandler) GetByID(c *gin.Context) { ... }
