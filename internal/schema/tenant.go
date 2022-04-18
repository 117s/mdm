package schema

// Tenant is the minimal data isolation layer
// FIXME: at the very early development, we just have a quite simple tenant model
type Tenant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
