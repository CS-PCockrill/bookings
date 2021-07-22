package handlers

import (
	"encoding/json"
	"fmt"
	config2 "github.com/CS-PCockrill/bookings-app/internal/config"
	"github.com/CS-PCockrill/bookings-app/internal/forms"
	"github.com/CS-PCockrill/bookings-app/internal/helpers"
	models2 "github.com/CS-PCockrill/bookings-app/internal/models"
	render2 "github.com/CS-PCockrill/bookings-app/internal/render"
	"log"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config2.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config2.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplate(w, "home.page.tmpl", &models2.TemplateData{}, r)
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// send data to the template
	render2.RenderTemplate(w, "about.page.tmpl", &models2.TemplateData{}, r)
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models2.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render2.RenderTemplate(w, "make-reservation.page.tmpl", &models2.TemplateData{
		Form: forms.New(nil),
		Data: data,
	}, r)
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// Can/could update to render 404 template page
		helpers.ServerError(w, err)
		return
	}

	reservation := models2.Reservation {
		FirstName: r.Form.Get("first_name"),
		LastName: r.Form.Get("last_name"),
		Email: r.Form.Get("email"),
		Phone: r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 3)
	form.ValidateEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render2.RenderTemplate(w, "make-reservation.page.tmpl", &models2.TemplateData{
			Form: form,
			Data: data,
		}, r)
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplate(w, "generals.page.tmpl", &models2.TemplateData{}, r)
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplate(w, "majors.page.tmpl", &models2.TemplateData{}, r)
}

// Availability searches the room availability
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplate(w, "search-availability.page.tmpl", &models2.TemplateData{}, r)
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start_date")
	end := r.Form.Get("end_date")
	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	Ok bool `json:"ok"`
	Message string `json:"message"`
}
// AvailabilityJSON handles request for availability and send JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Ok: true,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact displays contact information
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplate(w, "contact.page.tmpl", &models2.TemplateData{}, r)
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models2.Reservation)
	if !ok {
		m.App.ErrorLog.Println("Can't get error from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session.")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation

	render2.RenderTemplate(w, "reservation-summary.page.tmpl", &models2.TemplateData{
		Data: data,
	}, r)
}


