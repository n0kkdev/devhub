package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"text/template"
)

//Stat structure
type Stat struct {
	Name   string
	Thing  string
	Status bool
}

//StatisticData all status
type StatisticData struct {
	PageTime  string
	StatusAll []Stat
}

//Circle struct
type Circle struct {
	X, Y, R float64
}

//Brightness func create
func (c *Circle) Brightness(x, y float64) uint8 {
	var dx, dy float64 = c.X - x, c.Y - y
	d := math.Sqrt(dx*dx+dy*dy) / c.R
	if d > 1 {
		return 0
	} else {
		return 255
	}
}

var w, h int = 400, 400
var m1 = image.NewRGBA(image.Rect(0, 0, w, h))

//CreateRect ...
func CreateRect() {
	black := color.RGBA{0, 0, 0, 255}
	draw.Draw(m1, m1.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	redrect := image.Rect(50, 50, 150, 150)
	rr := color.RGBA{255, 0, 0, 255}
	draw.Draw(m1, redrect, &image.Uniform{rr}, image.ZP, draw.Src)

	greenrect := image.Rect(250, 250, 120, 120)
	gr := color.RGBA{0, 255, 0, 255}
	draw.Draw(m1, greenrect, &image.Uniform{gr}, image.ZP, draw.Src)

	bluerect := image.Rect(350, 350, 200, 200)
	br := color.RGBA{0, 0, 255, 255}
	draw.Draw(m1, bluerect, &image.Uniform{br}, image.ZP, draw.Src)

	f2, err := os.OpenFile("goserv/rect/threerect.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f2.Close()
	png.Encode(f2, m1)
}

//CreateCircle ...
func CreateCircle() {
	var hw, hh float64 = float64(w / 2), float64(h / 2)
	r := 80.0
	s := 2 * math.Pi / 3
	cr := &Circle{hw - r*math.Sin(0), hh - r*math.Cos(0), 100}
	cg := &Circle{hw - r*math.Sin(s), hh - r*math.Cos(s), 100}
	cb := &Circle{hw - r*math.Sin(-s), hh - r*math.Cos(-s), 100}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := color.RGBA{
				cr.Brightness(float64(x), float64(y)),
				cg.Brightness(float64(x), float64(y)),
				cb.Brightness(float64(x), float64(y)),
				255,
			}
			m1.Set(x, y, c)
		}
	}

	f1, err := os.OpenFile("goserv/circle/threecircle.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f1.Close()
	png.Encode(f1, m1)
}

//Hello dyn stdout
func Hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res,
		`<doctype html>
<html>
	<head>
    	<title>Hello World!</title>
	</head>
	<body>
    	<h1>Hello GO!</h1>
	</body>
</html>`)
}

//CourseStatus ...
func CourseStatus(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatal(err)
	}

	data := StatisticData{
		PageTime: "Сourse1",
		StatusAll: []Stat{
			{Name: "Ivan", Thing: "Physics", Status: false},
			{Name: "Petr", Thing: "Maths", Status: true},
			{Name: "Zina", Thing: "Biology", Status: true},
		},
	}
	tmpl.Execute(w, data)
}

func main() {

	CreateRect()
	CreateCircle()

	fs := http.FileServer(http.Dir("goserv"))

	fmt.Println("ServerListen")
	http.Handle("/", fs)
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/coursestatus", CourseStatus)
	http.ListenAndServe(":8010", nil)
}
