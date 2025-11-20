package customer

import (
    "strconv"

    "github.com/gofiber/fiber/v2"
)

type Handler struct{ svc Service }

func NewCustomerHandler(svc Service) *Handler { return &Handler{svc: svc} }

func (h *Handler) RegisterRoutes(app *fiber.App) {
    r := app.Group("/customers")
    r.Post("/", h.create)
    r.Get(":id", h.getByID)
}

func (h *Handler) create(c *fiber.Ctx) error {
    cu := new(Customer)
    if err := c.BodyParser(cu); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
    }
    if err := h.svc.Create(c.Context(), cu); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(fiber.StatusCreated).JSON(cu)
}

func (h *Handler) getByID(c *fiber.Ctx) error {
    idStr := c.Params("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
    }
    cu, err := h.svc.Get(c.Context(), uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
    }
    return c.JSON(cu)
}
