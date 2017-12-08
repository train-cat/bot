package wording

import "fmt"

// List of keys available for wording
const (
	HelloOne            = "hello_one"
	HelloTwo            = "hello_two"
	HelloTwoReplies     = "hello_quick_reply"
	Cancel              = "cancel"
	StartCreateAlert    = "start_create_alert"
	AskListSchedule     = "ask_list_schedule"
	AskOrigin           = "ask_origin"
	AskDestination      = "ask_destination"
	AskSchedule         = "ask_schedule"
	ConfirmationAlert   = "confirmation_alert"
	ChoiceOutOfRange    = "choice_out_of_range"
	ReAskSelectSchedule = "re_ask_select_schedule"
	Retry               = "retry"
	IssueOne            = "issue_one"
	IssueTwo            = "issue_two"
	HasFail             = "has_fail"
)

var (
	wordings = map[string][]string{
		HelloOne:            {"Bonjour ! Je suis Simone de la SNCF"},
		HelloTwo:            {"Je suis là pour te prévenir si ton train est pertubé et surtout pour te trouver une solution 😉"},
		HelloTwoReplies:     {"Créer mon alerte"},
		StartCreateAlert:    {"😊 C'est parti !"},
		AskOrigin:           {"Je vais avoir besoin de ta gare de départ !"},
		AskDestination:      {"Départ de %s c'est noté ! Peux-tu me dire ta gare de destination ?"},
		AskSchedule:         {"%s -> %s, à quelle heure ?"},
		AskListSchedule:     {"J'ai une liste de train, lequel correspond à ton train ?"},
		ConfirmationAlert:   {"👌 C'est bon pour moi ! Ton alerte est bien enregistrée ! Et comme on dit, \"Pas de nouvelles, bonne nouvelles 😉\""},
		ChoiceOutOfRange:    {"Oulà 😵 Tu vas trop loin, je n'ai que %d choix"},
		ReAskSelectSchedule: {"Désolé mais, quel trajet correspond à ton train ?"},
		Retry:               {"Très bien, recommençons depuis le début"},
		Cancel:              {"Très bien, j'arrête !"},
		IssueOne:            {"Hey, malheuresement ton train au départ de %s à %s est %s. ✊✊"},
		IssueTwo:            {"Je ne suis pas encore capable de t'aider 😢. Mais promis je travaille dur pour y arriver ! Bon courage 😊"},
		HasFail:             {"Désolé, il semble que quelque chose n'est pas fonctionné correctement, peux-tu réessayer plus tard ?"},
	}
)

// Get random message for the key
func Get(key string, a ...interface{}) string {
	str, ok := wordings[key]

	if !ok {
		return ""
	}

	return fmt.Sprintf(str[0], a...)
}
