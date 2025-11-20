package product

import (
    "strconv"

    "github.com/gofiber/fiber/v2"
)

type Handler struct {
    svc Service
}

func NewProductHandler(svc Service) *Handler { return &Handler{svc: svc} }

func (h *Handler) RegisterRoutes(app *fiber.App) {
    r := app.Group("/products")
    r.Post("/", h.create)
    r.Get("/", h.list)
    r.Get(":id", h.getByID)
}

func (h *Handler) create(c *fiber.Ctx) error {
    p := new(Product)
    if err := c.BodyParser(p); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
    }
    if err := h.svc.Create(c.Context(), p); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(fiber.StatusCreated).JSON(p)
}

func (h *Handler) list(c *fiber.Ctx) error {
    out, err := h.svc.List(c.Context())
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(out)
}

func (h *Handler) getByID(c *fiber.Ctx) error {
    idStr := c.Params("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
    }
    p, err := h.svc.Get(c.Context(), uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
    }
    return c.JSON(p)
}
