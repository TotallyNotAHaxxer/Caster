package CastHunter

import "fmt"

// Controller actions
func Up(Hostname string)    { NewPostNoData(fmt.Sprintf(Keys["up"], Hostname), false) }    // Toggles Up
func Down(Hostname string)  { NewPostNoData(fmt.Sprintf(Keys["down"], Hostname), false) }  // Toggles Down
func Left(Hostname string)  { NewPostNoData(fmt.Sprintf(Keys["left"], Hostname), false) }  // Toggles Left
func Right(Hostname string) { NewPostNoData(fmt.Sprintf(Keys["right"], Hostname), false) } // Toggles Right
func Click(Hostname string) { NewPostNoData(fmt.Sprintf(Keys["OK"], Hostname), false) }    // Toggles Click
func Back(Hostname string)  { NewPostNoData(fmt.Sprintf(Keys["back"], Hostname), false) }  // Toggles Back
func Home(Hostname string)  { NewPostNoData(fmt.Sprintf(Keys["home"], Hostname), false) }  // Toggles Home

// sub functions

func SApp(Hostname string) {
	NewPostNoData(fmt.Sprintf(Keys["launch"], Hostname, ApplicationIDROKU), false)
}
