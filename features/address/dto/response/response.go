package response

type AddressResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	City       string `json:"city"`
	CityId     string `json:"city_id"`
	Subdistric string `json:"subdistric"`
	ZipCode    string `json:"zip_code"`
	Phone      string `json:"phone"`
}