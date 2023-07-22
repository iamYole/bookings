package models

// Template Date holds data
type TemplateData struct {
	StringMap             map[string]string
	IntMap                map[string]int
	FloatMap              map[string]float32
	Data                  map[string]interface{}
	CSRFToken             string
	Flash, Warning, Error string
}
