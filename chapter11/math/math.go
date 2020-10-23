package math

//Average find the average of a serie of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

//run go install to install it in pkg folder

//to generate documentation you can run go doc golang-book/chapter11/math Average
//go doc -http":6060"" to enable to see in the browser
//http://localhost:6060/pkg/ to see the documentation generated
