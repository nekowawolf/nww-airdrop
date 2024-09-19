package url

import (
	"github.com/nekowawolf/nww-airdrop/controller"

	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

	page.Get("/checkip", controller.Homepage)
	page.Get("/freeairdrop", controller.GetAirdropFreeHandler)
	page.Get("/paidairdrop", controller.GetAirdropPaidHandler)
	page.Get("/allairdrop", controller.GetAllAirdropHandler)
	page.Get("/freeairdrop/:id", controller.GetAirdropFreeByIDHandler)
	page.Get("/paidairdrop/:id", controller.GetAirdropPaidByIDHandler)
	page.Get("/freeairdrop/search/:name", controller.GetAirdropFreeByNameHandler)
	page.Get("/paidairdrop/search/:name", controller.GetAirdropPaidByNameHandler)
	page.Post("/freeairdrop", controller.InsertAirdropFreeHandler)
	page.Post("/paidairdrop", controller.InsertAirdropPaidHandler)
	page.Put("/freeairdrop/:id", controller.UpdateAirdropFreeByIDHandler)
    page.Put("/paidairdrop/:id", controller.UpdateAirdropPaidByIDHandler)
	page.Delete("/freeairdrop/:id", controller.DeleteAirdropFreeByIDHandler)
    page.Delete("/paidairdrop/:id", controller.DeleteAirdropPaidByIDHandler)
	
}
