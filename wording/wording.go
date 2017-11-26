package wording

import "fmt"

const (
	HelloOne             = "hello_one"
	HelloTwo             = "hello_two"
	HelloTwoReplies      = "hello_quick_reply"
	Cancel               = "cancel"
	StartCreateAlert     = "start_create_alert"
	AskListSchedule      = "ask_list_schedule"
	OriginOkAskSchedule  = "origin_ok_ask_schedule"
	AskOrigin            = "ask_origin"
	AskDestination       = "ask_destination"
	SelectOrigin         = "select_origin"
	SelectDestination    = "select_destination"
	OriginNotFound       = "origin_not_found"
	DestinationNotFound  = "destination_not_found"
	ScheduleOkAskOrigin  = "schedule_ok_ask_origin"
	AskConfirmationAlert = "ask_confirmation_alert"
	ConfirmationAlert    = "confirmation_alert"
	CancelAlert          = "cancel_alert"
	IssueOne             = "issue_one"
	IssueTwo             = "issue_two"
)

var (
	wordings = map[string][]string{
		HelloOne:             {"Bonjour ! Je suis Simone de la SNCF"},
		HelloTwo:             {"Je suis là pour te prévenir si ton train est pertubé et surtout pour te trouver une solution 😉"},
		HelloTwoReplies:      {"Créer mon alerte"},
		StartCreateAlert:     {"😊 C'est parti ! Quelle est ta gare de départ ?"},
		AskListSchedule:      {"J'ai une liste de train, lequel correspond à ton train ?"},
		AskOrigin:            {"Je vais avoir besoin de ta gare de départ, peux-tu me le dire ?"},
		AskDestination:       {"%s c'est noté ! Peux-tu me dire ta gare de destination ?"},
		OriginNotFound:       {"Je me pensais pourtant forte en géographie 🤓. Peux-tu me préciser ta gare ?"},
		DestinationNotFound:  {"Je me pensais pourtant forte en géographie 🤓. Peux-tu me préciser ta gare ?"},
		SelectOrigin:         {"J'ai plusieurs gares ! Peux-tu me préciser ta gare ?"},
		SelectDestination:    {"J'ai plusieurs gares ! Peux-tu me préciser ta gare ?"},
		OriginOkAskSchedule:  {"Destination %s ! A quelle heure pars-tu de la gare de %s?"},
		ScheduleOkAskOrigin:  {"Départ à %s ! Depuis quelle gare pars-tu ?"},
		AskConfirmationAlert: {"Départ %s à %s ! Veux-tu que j'enregistre ton alerte ?"},
		ConfirmationAlert:    {"👌 C'est bon pour moi ! Ton alerte est bien enregistrée ! Et comme on dit, \"Pas de nouvelles, bonne nouvelles 😉\""},
		Cancel:               {"Très bien, j'arrête !"},
		CancelAlert:          {"Ok ! J'oublie cette alerte, tu ne seras pas notifié"},
		IssueOne:             {"Hey, malheuresement ton train au départ de %s à %s est %s. ✊✊"},
		IssueTwo:             {"Je ne suis pas encore capable de t'aider 😢. Mais promis je travaille dure pour y arriver ! Bon courage 😊"},
	}
)

func Get(key string, a ...interface{}) string {
	str, ok := wordings[key]

	if !ok {
		return ""
	}

	return fmt.Sprintf(str[0], a...)
}
