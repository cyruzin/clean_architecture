package routeentity

// Route is struct used for route type.
type Route struct {
	Departure   string `json:"departure"`
	Destination string `json:"destination"`
	Price       int    `json:"price"`
}
