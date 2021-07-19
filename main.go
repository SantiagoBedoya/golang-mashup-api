package main

import (
	"golang-mashups/controller"
	router "golang-mashups/http"
	"golang-mashups/service"
)

var (
	carService    service.CarDetailsService = service.NewCarDetailsService()
	carController controller.CarController  = controller.NewCarController(carService)
	httpRouter    router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":3000"
	httpRouter.GET("/car-details", carController.GetCarDetails)
	httpRouter.SERVE(port)
}
