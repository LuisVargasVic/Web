package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Image string
	Title string
	Description string
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func showLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../pages/login.html")
}

func showIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../pages/index.html")
}

func showAbout(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../pages/about.html")
}

func showCatalog(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../pages/catalog.html"))
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Image: "/img/balinese-cat.jpeg", Title: "Balinese", Description: "An easygoing breed with a clownlike personality, the Balinese adores people, though less vocal than a Siamese."},
			{Image: "/img/bengal-cat.jpg", Title: "Bengal", Description: "This docile house cat has what some would call a rambunctious personality. Bengal cats are playful and love to chase, climb, investigate and be part of the action."},
			{Image: "/img/russian-blue.jpg", Title: "Russian Blue", Description: "The Russian Blue is gentle, quiet and even shy around strangers, she’ll follow you around and even ride on your shoulder."},
			{Image: "/img/persian-cat.jpg", Title: "Persian", Description: "The docile Persian is a quiet feline who enjoys a calm and relaxing environment; however, persians are independent and selective in who they show affection to."},
			{Image: "/img/british-shorthair-cat.jpeg", Title: "British shorthair", Description: "The British Shorthair is an easygoing feline. She enjoys affection but isn’t needy and dislikes being carried. She’ll follow you from room to room, though, out of curiosity."},
			{Image: "/img/himalayan-cat.jpg", Title: "Himalayan", Description: "The Himalayan Cat is a sweet and mild-tempered feline. She’s affectionate but selective. Although she loves lying in your lap and being pet, she may be reserved around guests."},
			{Image: "/img/scottish-fold-cat.jpg", Title: "Scottish fold", Description: "The smart and friendly Scottish Fold loves playing with challenging, puzzling toys to test her intelligence. She also loves human interaction, though with moderation."},
			{Image: "/img/siamese-cat.jpg", Title: "Siamese", Description: "Siamese Cats are incredibly social, intelligent and vocal—they’ll talk to anyone who wants to listen. Although they’re active and curious, they also love curling up on their human’s lap or snuggling up next to them in bed."},
			{Image: "/img/siberian-cat.jpg", Title: "Siberian", Description: "This friendly and affectionate feline will follow you around as you go about your day, and purr in your lap as you comb her coat. Siberian Cats love their humans but aren’t shy around strangers."},
			{Image: "/img/snowshoe.jpg", Title: "Snowshoe", Description: "This sweet-natured, lively cat is inquisitive and loves being the center of attention. A devoted, affectionate feline, the Snowshoe follows her human companions everywhere, purring with delight."},
			{Image: "/img/sphynx-cat.png", Title: "Sphynx", Description: "This sweet-natured, lively cat is inquisitive and loves being the center of attention. A devoted, affectionate feline, the Sphynx follows her human companions everywhere, purring with delight."},
			{Image: "/img/toyger-cat.jpeg", Title: "Toyger", Description: "This breed has a friendly, outgoing temperament and delights in being with people, even strangers, and gets along well with other pets."},
		},
	}
	tmpl.Execute(w, data)
}

func showService(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../pages/service.html")
}

func showRegister(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../pages/register.html")
}

func showContact(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../pages/contact.html")
}

func main() {
	css := http.FileServer(http.Dir("../css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	js := http.FileServer(http.Dir("../vendor"))
	http.Handle("/vendor/", http.StripPrefix("/vendor/", js))
	im := http.FileServer(http.Dir("../img"))
	http.Handle("/img/", http.StripPrefix("/img/", im))
	pg := http.FileServer(http.Dir("../pages"))
	http.Handle("/pages/", http.StripPrefix("/pages/", pg))
	http.HandleFunc("/login", showLogin)
	http.HandleFunc("/index", showIndex)
	http.HandleFunc("/about", showAbout)
	http.HandleFunc("/catalog", showCatalog)
	http.HandleFunc("/service", showService)
	http.HandleFunc("/register", showRegister)
	http.HandleFunc("/contact", showContact)
	http.ListenAndServe(":8080", nil)

}
