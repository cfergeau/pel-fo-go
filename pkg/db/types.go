package db

type Db struct {
	Teams      []Team      `json:"Teams"`
	Games      []Game      `json:"Games"`
	Boxes      []Box       `json:"Boxes"`
	Volunteers []Volunteer `json:"Volunteers"`

	GamesById      map[int]*Game
	TeamsById      map[int]*Team
	VolunteersById map[int]*Volunteer
	BoxesByName    map[string][]*Box
}

/* Dummy struct to unnest the DB from the json */
type rawDb struct {
	States struct {
		Db
	}
	Types struct {
		/* This seems to be some kind of schema, but the field names
		 * are translated while the field names used in States are not
		 */
	}
}

type Team struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Min         int    `json:"min"`
	Max         int    `json:"max"`
	Description string `json:"description"`
	Avant       string `json:"avant"`
	Pendant     string `json:"pendant"`
	Apres       string `json:"après"`
	Statut      string `json:"statut"`
	Ordre       string `json:"ordre"`
}

type Announcement struct {
	Id                  int    `json:"id"`
	Type                string `json:"type"`
	Titre               string `json:"titre"`
	Url                 string `json:"url"`
	InformeAvecUneNotif bool   `json:"informéAvecUneNotif"`
}

type Postulant struct {
	Id                       int    `json:"id"`
	Prenom                   string `json:"prenom"`
	Nom                      string `json:"nom"`
	Email                    string `json:"email"`
	Telephone                string `json:"telephone"`
	CommentContacter         string `json:"commentContacter"`
	Potentiel                bool   `json:"potentiel"`
	DejaVenu                 bool   `json:"déjàVenu"`
	DateRencontre            string `json:"dateRencontre"`
	CommentaireDateRencontre string `json:"commentaireDateRencontre"`
	Commentaire              string `json:"commentaire"`
}

type Wish struct {
	Id         int      `json:"id"`
	Domaine    string   `json:"domaine"`
	Precisions string   `json:"precisions"`
	Equipes    []string `json:"equipes"`
	DateAjout  string   `json:"dateAjout"` // type: date
}

type Misc struct {
	Id                int    `json:"id"`
	RencontreId       string `json:"rencontreId"`
	RencontreTitre    string `json:"rencontreTitre"`
	RencontreUrl      string `json:"rencontreUrl"`
	InvitationDiscord string `json:"invitationDiscord"`
}

type Game struct {
	Id         int    `json:"id"`
	Titre      string `json:"title"` // TODO: check all field names
	MinJoueurs int    `json:"minJoueurs"`
	MaxJoueurs int    `json:"maxJoueurs"`
	Duree      int    `json:"duree"`
	Type       string `json:"type"`
	Poufpaf    string `json:"poufpaf"`
	BggId      int    `json:"bggId"`
	Ean        int    `json:"ean"`
	BggPhoto   string `json:"bggPhoto"`

	Boxes []int

	Ok   []int
	Bof  []int
	Niet []int
}

type Box struct {
	Id                int    `json:"id"`
	GameId            int    `json:"gameId"`
	Container         string `json:"container"`
	Injouable         bool   `json:"injouable"`
	EanSpecifique     int    `json:"eanSpécifique"`
	PartiesManquantes string `json:"partiesManquantes"`
	Verifie           string `json:"verifié"` // type: date
}

type Volunteer struct {
	Id                      int      `json:"id"`
	Prenom                  string   `json:"prenom"`
	Nom                     string   `json:"nom"`
	Mail                    string   `json:"mail"`
	Telephone               string   `json:"telephone"`
	Photo                   string   `json:"photo"`
	Majeur                  int      `json:"majeur"`
	Roles                   []string `json:"roles"`
	Actif                   string   `json:"actif"`
	DiscordId               string   `json:"discordId"`
	EnviesJours             []string `json:"enviesJours"`
	CommentaireEnviesJours  string   `json:"commentaireEnviesJours"`
	NbDeTshirts             int      `json:"nbDeTshirts"`
	TailleDeTshirts         string   `json:"tailleDeTshirts"`
	Alimentation            string   `json:"alimentation"`
	Equipe                  int      `json:"team"`
	EnviesEquipe            []int    `json:"enviesEquipe"`
	CommentaireEnviesEquipe string   `json:"commentaireEnviesEquipe"`
	CommentContacter        string   `json:"commentContacter"`
	AideEnAmont             string   `json:"aideEnAmont"`
	MembrePel               bool     `json:"membrePel"`
	QuestionsCachees        []int    `json:"questionsCachees"`
	Creation                string   `json:"creation"` // type: date
	Passe1                  string   `json:"passe1"`
	Passe2                  string   `json:"passe2"`
	PushNotifSubscription   string   `json:"pushNotifSubscription"`
	AccepteLesNotifs        string   `json:"accepteLesNotifs"`
	Ok                      []int    `json:"OK"`
	Bof                     []int    `json:"Bof"`
	Niet                    []int    `json:"Niet"`
	BesoinHebergement       bool     `json:"besoinHébergement"`
	NombreHeberges          int      `json:"nombreHébergés"`
	DistanceAuFestival      int      `json:"distanceAuFestival"`
	CommentaireHebergement  string   `json:"commentaireHébergement"`
	Repas                   []string `json:"repas"`
}
