package models

type TemplateData struct {

	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	AutreData map[string]interface{}

	CSRFToken string
	FlashMessage string
	WarningMessage string
	ErrorMessage string

}	
