package transaction

import (
    "strconv"

    "github.com/gofiber/fiber/v2"
)

type Handler struct{ svc Service }

func NewTransactionHandler(svc Service) *Handler { return &Handler{svc: svc} }

func (h *Handler) RegisterRoutes(app *fiber.App) {
    r := app.Group("/transactions")
    r.Post("/purchase", h.purchase)
    r.Post("/redeem", h.redeem)
}

type purchaseReq struct {
    CustomerID uint   `json:"customer_id"`
    ProductID  uint   `json:"product_id"`
    Size       string `json:"size"`
    Flavor     string `json:"flavor"`
    Quantity   int    `json:"quantity"`
    UnitPrice  int    `json:"unit_price"`
}

func (h *Handler) purchase(c *fiber.Ctx) error {
    var req purchaseReq
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
    }
    t, err := h.svc.CreatePurchase(c.Context(), req.CustomerID, req.ProductID, req.Size, req.Flavor, req.Quantity, req.UnitPrice)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(t)
}

type redeemReq struct {
    CustomerID uint   `json:"customer_id"`
    Size       string `json:"size"`
}

func (h *Handler) redeem(c *fiber.Ctx) error {
    var req redeemReq
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
    }
    t, err := h.svc.RedeemBySize(c.Context(), req.CustomerID, req.Size)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(t)
}
