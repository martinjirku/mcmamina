package models

type Activity struct {
	ID          string
	Title       string
	Img         string
	Description string
	Time        string
	Contact     ActivityContact
}

type ActivityContact struct {
	Email string
	Phone string
	FB    string
}

func GetActivities() []Activity {
	return []Activity{
		{
			ID:          "dielnicky",
			Title:       "Montessori hernička",
			Img:         "bg-dielnicky",
			Contact:     ActivityContact{Phone: "+421 948 523 493", FB: "https://www.facebook.com/MaterskeCentrumMamina/"},
			Description: "Zážitkový a vzdelávací program pre najmenších od 2 do 4 rokov inšpirovaný princípmi Montessori pedagogiky.",
			Time:        "Každý piatok o 9:30.",
		},
		{
			ID:          "tinka",
			Title:       "Angličtina s Tinkou",
			Img:         "bg-tinka",
			Contact:     ActivityContact{Phone: "+421 907 948 207", Email: "anglictinamcmamina@gmail.com"},
			Description: "Tinka vedie krúžok angličtiny hravou a prirodzenou cestou.",
			Time:        "Utorok o 16:30 v 3 skupinách a štvrtok o 10:00.",
		},
		{
			ID:          "happy-gym",
			Title:       "Happy gym",
			Img:         "bg-hrave-cvicenie",
			Contact:     ActivityContact{Phone: "+421 911 528 887", Email: "happygymzv@gmail.com"},
			Description: "Cvičenie pre najmenších zamerané na psychomotorický, sociálny, citový a rozumový vývoj dieťaťa.",
			Time:        "Utorok o 8:45 v troch skupinkách",
		},
		{
			ID:          "pohyb",
			Title:       "Tanečno - pohybová príprava",
			Img:         "bg-pohybova-priprava",
			Contact:     ActivityContact{Email: "tkelement@tkelement.com"},
			Description: "Cieľom u detí je získať hravou formou - správne držanie tela, hudobno-pohybové cítenie, zamerať sa na rytmus, tempo, takt, dynamiku, frázovanie a iné.",
			Time:        "Pondelok o 16:00 v dvoch skupinkách",
		},
	}
}
