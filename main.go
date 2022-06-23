package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django"
	"github.com/qeaml/fib/config"
	"github.com/qeaml/fib/gallery"
)

const DiagScript = "/public/frontend/diag/diag.js"
const DiagStyle = "/public/frontend/diag/diag.css"

func Error(code int, err string, c *fiber.Ctx) error {
	return c.Status(code).Render("error", fiber.Map{
		"_title":    fmt.Sprintf("Error :: %d", code),
		"_escripts": []string{DiagScript},
		"_estyles":  []string{DiagStyle},
		"error":     err,
	})
}

func User(c *fiber.Ctx) error {
	id := c.Params("id")
	if user, ok := gallery.Users[id]; ok {
		return c.Render("user", fiber.Map{
			"user":    user,
			"_title":  fmt.Sprintf("User :: %s", user.Name),
			"_styles": []string{"user"},
		})
	}
	return Error(http.StatusNotFound, "User not found", c)
}

func Image(c *fiber.Ctx) error {
	idRaw := c.Params("id")
	id, err := strconv.ParseUint(idRaw, 10, 32)
	if err != nil {
		return Error(http.StatusBadRequest, "Invalid image id", c)
	}
	if image, ok := gallery.Images[uint32(id)]; ok {
		return c.Render("image", fiber.Map{
			"image":   image,
			"_title":  fmt.Sprintf("Image :: %s", image.Title),
			"_styles": []string{"image"},
		})
	}
	return Error(http.StatusNotFound, "Image not found", c)
}

func ImageFile(c *fiber.Ctx) error {
	idRaw := c.Params("id")
	id, err := strconv.ParseUint(idRaw, 10, 32)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	if _, ok := gallery.Images[uint32(id)]; ok {
		return c.SendFile(fmt.Sprintf("data/img/%d", id), true)
	}
	return c.SendStatus(http.StatusNotFound)
}

//go:embed templates
var templates embed.FS

func main() {
	log.Println("Loading config...")
	if err := config.LoadConfig(); err != nil {
		log.Fatalln(err)
	}

	log.Println("Preparing directories...")
	if err := os.MkdirAll("data/img", 0755); err != nil {
		log.Fatalln(err)
	}

	log.Println("Loading users...")
	if err := gallery.LoadUsers(); err != nil {
		log.Fatalln(err)
	}
	log.Printf("Loaded %d users\n", len(gallery.Users))

	if sys, ok := gallery.Users["sys"]; !ok {
		log.Println("Creating sys user...")
		user := &gallery.User{
			ID:         "sys",
			Flags:      gallery.UserFlagAdmin,
			Name:       "System",
			Bio:        "System user",
			Avatar:     0,
			Registered: time.Now(),
			LastLogin:  time.Now(),
		}
		gallery.Users["sys"] = user
		if err := gallery.SaveUsers(); err != nil {
			log.Fatalln(err)
		}
	} else {
		sys.LastLogin = time.Now()
		gallery.Users["sys"] = sys
		if err := gallery.SaveUsers(); err != nil {
			log.Fatalln(err)
		}
	}

	log.Println("Loading images...")
	if err := gallery.LoadImages(); err != nil {
		log.Fatalln(err)
	}
	log.Printf("Loaded %d images\n", len(gallery.Images))

	if sample, ok := gallery.Images[0]; !ok {
		log.Println("Creating sample image...")
		image := &gallery.Image{
			ID:         0,
			Flags:      gallery.ImageFlagNone,
			Title:      "Sample image",
			Desc:       "This is a sample image",
			Tags:       []string{"sample", "image"},
			Uploader:   "sys",
			UploadedAt: time.Now(),
			UpdatedAt:  time.Now(),
		}
		gallery.Images[0] = image
		if err := gallery.SaveImages(); err != nil {
			log.Fatalln(err)
		}
	} else {
		sample.UpdatedAt = time.Now()
		gallery.Images[0] = sample
		if err := gallery.SaveImages(); err != nil {
			log.Fatalln(err)
		}
	}

	log.Println("Starting server...")
	subfs, err := fs.Sub(templates, "templates")
	if err != nil {
		log.Fatalln(err)
	}
	views := django.NewFileSystem(http.FS(subfs), ".html")
	views.Debug(true)
	app := fiber.New(fiber.Config{
		AppName:      "fib",
		ServerHeader: "fib",
		Views:        views,
		ViewsLayout:  "main",
	})
	app.Use(logger.New())
	app.Get("/u/:id", User)
	app.Get("/i/:id/file", ImageFile)
	app.Get("/i/:id", Image)
	app.Static("/public", "public")

	app.Use(func(c *fiber.Ctx) error {
		return Error(http.StatusNotFound, "Page not found", c)
	})

	log.Fatalln(app.Listen(fmt.Sprintf(":%d", config.GlobalConfig.Port)))
}
