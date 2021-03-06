package gherkin

// Builtin dialects for af (Afrikaans), am (Armenian), ar (Arabic), ast (Asturian), az (Azerbaijani), bg (Bulgarian), bm (Malay), bs (Bosnian), ca (Catalan), cs (Czech), cy-GB (Welsh), da (Danish), de (German), el (Greek), em (Emoji), en (English), en-Scouse (Scouse), en-au (Australian), en-lol (LOLCAT), en-old (Old English), en-pirate (Pirate), eo (Esperanto), es (Spanish), et (Estonian), fa (Persian), fi (Finnish), fr (French), ga (Irish), gj (Gujarati), gl (Galician), he (Hebrew), hi (Hindi), hr (Croatian), ht (Creole), hu (Hungarian), id (Indonesian), is (Icelandic), it (Italian), ja (Japanese), jv (Javanese), kn (Kannada), ko (Korean), lt (Lithuanian), lu (Luxemburgish), lv (Latvian), mn (Mongolian), nl (Dutch), no (Norwegian), pa (Panjabi), pl (Polish), pt (Portuguese), ro (Romanian), ru (Russian), sk (Slovak), sl (Slovenian), sr-Cyrl (Serbian), sr-Latn (Serbian (Latin)), sv (Swedish), ta (Tamil), th (Thai), tl (Telugu), tlh (Klingon), tr (Turkish), tt (Tatar), uk (Ukrainian), ur (Urdu), uz (Uzbek), vi (Vietnamese), zh-CN (Chinese simplified), zh-TW (Chinese traditional)
func GherkinDialectsBuildin() GherkinDialectProvider {
	return buildinDialects
}

const (
	feature = "feature"
	background = "background"
	scenario = "scenario"
	scenarioOutline = "scenarioOutline"
	examples = "examples"
	given = "given"
	when = "when"
	then = "then"
	and = "and"
	but = "but"
)

var buildinDialects = gherkinDialectMap{
	"af": &GherkinDialect{
		"af", "Afrikaans", "Afrikaans", map[string][]string{
			feature: []string{
				"Funksie",
				"Besigheid Behoefte",
				"Vermoë",
			},
			background: []string{
				"Agtergrond",
			},
			scenario: []string{
				"Situasie",
			},
			scenarioOutline: []string{
				"Situasie Uiteensetting",
			},
			examples: []string{
				"Voorbeelde",
			},
			given: []string{
				"* ",
				"Gegewe ",
			},
			when: []string{
				"* ",
				"Wanneer ",
			},
			then: []string{
				"* ",
				"Dan ",
			},
			and: []string{
				"* ",
				"En ",
			},
			but: []string{
				"* ",
				"Maar ",
			},
		},
	},
	"am": &GherkinDialect{
		"am", "Armenian", "հայերեն", map[string][]string{
			feature: []string{
				"Ֆունկցիոնալություն",
				"Հատկություն",
			},
			background: []string{
				"Կոնտեքստ",
			},
			scenario: []string{
				"Սցենար",
			},
			scenarioOutline: []string{
				"Սցենարի կառուցվացքը",
			},
			examples: []string{
				"Օրինակներ",
			},
			given: []string{
				"* ",
				"Դիցուք ",
			},
			when: []string{
				"* ",
				"Եթե ",
				"Երբ ",
			},
			then: []string{
				"* ",
				"Ապա ",
			},
			and: []string{
				"* ",
				"Եվ ",
			},
			but: []string{
				"* ",
				"Բայց ",
			},
		},
	},
	"ar": &GherkinDialect{
		"ar", "Arabic", "العربية", map[string][]string{
			feature: []string{
				"خاصية",
			},
			background: []string{
				"الخلفية",
			},
			scenario: []string{
				"سيناريو",
			},
			scenarioOutline: []string{
				"سيناريو مخطط",
			},
			examples: []string{
				"امثلة",
			},
			given: []string{
				"* ",
				"بفرض ",
			},
			when: []string{
				"* ",
				"متى ",
				"عندما ",
			},
			then: []string{
				"* ",
				"اذاً ",
				"ثم ",
			},
			and: []string{
				"* ",
				"و ",
			},
			but: []string{
				"* ",
				"لكن ",
			},
		},
	},
	"ast": &GherkinDialect{
		"ast", "Asturian", "asturianu", map[string][]string{
			feature: []string{
				"Carauterística",
			},
			background: []string{
				"Antecedentes",
			},
			scenario: []string{
				"Casu",
			},
			scenarioOutline: []string{
				"Esbozu del casu",
			},
			examples: []string{
				"Exemplos",
			},
			given: []string{
				"* ",
				"Dáu ",
				"Dada ",
				"Daos ",
				"Daes ",
			},
			when: []string{
				"* ",
				"Cuando ",
			},
			then: []string{
				"* ",
				"Entós ",
			},
			and: []string{
				"* ",
				"Y ",
				"Ya ",
			},
			but: []string{
				"* ",
				"Peru ",
			},
		},
	},
	"az": &GherkinDialect{
		"az", "Azerbaijani", "Azərbaycanca", map[string][]string{
			feature: []string{
				"Özəllik",
			},
			background: []string{
				"Keçmiş",
				"Kontekst",
			},
			scenario: []string{
				"Ssenari",
			},
			scenarioOutline: []string{
				"Ssenarinin strukturu",
			},
			examples: []string{
				"Nümunələr",
			},
			given: []string{
				"* ",
				"Tutaq ki ",
				"Verilir ",
			},
			when: []string{
				"* ",
				"Əgər ",
				"Nə vaxt ki ",
			},
			then: []string{
				"* ",
				"O halda ",
			},
			and: []string{
				"* ",
				"Və ",
				"Həm ",
			},
			but: []string{
				"* ",
				"Amma ",
				"Ancaq ",
			},
		},
	},
	"bg": &GherkinDialect{
		"bg", "Bulgarian", "български", map[string][]string{
			feature: []string{
				"Функционалност",
			},
			background: []string{
				"Предистория",
			},
			scenario: []string{
				"Сценарий",
			},
			scenarioOutline: []string{
				"Рамка на сценарий",
			},
			examples: []string{
				"Примери",
			},
			given: []string{
				"* ",
				"Дадено ",
			},
			when: []string{
				"* ",
				"Когато ",
			},
			then: []string{
				"* ",
				"То ",
			},
			and: []string{
				"* ",
				"И ",
			},
			but: []string{
				"* ",
				"Но ",
			},
		},
	},
	"bm": &GherkinDialect{
		"bm", "Malay", "Bahasa Melayu", map[string][]string{
			feature: []string{
				"Fungsi",
			},
			background: []string{
				"Latar Belakang",
			},
			scenario: []string{
				"Senario",
				"Situasi",
				"Keadaan",
			},
			scenarioOutline: []string{
				"Kerangka Senario",
				"Kerangka Situasi",
				"Kerangka Keadaan",
				"Garis Panduan Senario",
			},
			examples: []string{
				"Contoh",
			},
			given: []string{
				"* ",
				"Diberi ",
				"Bagi ",
			},
			when: []string{
				"* ",
				"Apabila ",
			},
			then: []string{
				"* ",
				"Maka ",
				"Kemudian ",
			},
			and: []string{
				"* ",
				"Dan ",
			},
			but: []string{
				"* ",
				"Tetapi ",
				"Tapi ",
			},
		},
	},
	"bs": &GherkinDialect{
		"bs", "Bosnian", "Bosanski", map[string][]string{
			feature: []string{
				"Karakteristika",
			},
			background: []string{
				"Pozadina",
			},
			scenario: []string{
				"Scenariju",
				"Scenario",
			},
			scenarioOutline: []string{
				"Scenariju-obris",
				"Scenario-outline",
			},
			examples: []string{
				"Primjeri",
			},
			given: []string{
				"* ",
				"Dato ",
			},
			when: []string{
				"* ",
				"Kada ",
			},
			then: []string{
				"* ",
				"Zatim ",
			},
			and: []string{
				"* ",
				"I ",
				"A ",
			},
			but: []string{
				"* ",
				"Ali ",
			},
		},
	},
	"ca": &GherkinDialect{
		"ca", "Catalan", "català", map[string][]string{
			feature: []string{
				"Característica",
				"Funcionalitat",
			},
			background: []string{
				"Rerefons",
				"Antecedents",
			},
			scenario: []string{
				"Escenari",
			},
			scenarioOutline: []string{
				"Esquema de l'escenari",
			},
			examples: []string{
				"Exemples",
			},
			given: []string{
				"* ",
				"Donat ",
				"Donada ",
				"Atès ",
				"Atesa ",
			},
			when: []string{
				"* ",
				"Quan ",
			},
			then: []string{
				"* ",
				"Aleshores ",
				"Cal ",
			},
			and: []string{
				"* ",
				"I ",
			},
			but: []string{
				"* ",
				"Però ",
			},
		},
	},
	"cs": &GherkinDialect{
		"cs", "Czech", "Česky", map[string][]string{
			feature: []string{
				"Požadavek",
			},
			background: []string{
				"Pozadí",
				"Kontext",
			},
			scenario: []string{
				"Scénář",
			},
			scenarioOutline: []string{
				"Náčrt Scénáře",
				"Osnova scénáře",
			},
			examples: []string{
				"Příklady",
			},
			given: []string{
				"* ",
				"Pokud ",
				"Za předpokladu ",
			},
			when: []string{
				"* ",
				"Když ",
			},
			then: []string{
				"* ",
				"Pak ",
			},
			and: []string{
				"* ",
				"A také ",
				"A ",
			},
			but: []string{
				"* ",
				"Ale ",
			},
		},
	},
	"cy-GB": &GherkinDialect{
		"cy-GB", "Welsh", "Cymraeg", map[string][]string{
			feature: []string{
				"Arwedd",
			},
			background: []string{
				"Cefndir",
			},
			scenario: []string{
				"Scenario",
			},
			scenarioOutline: []string{
				"Scenario Amlinellol",
			},
			examples: []string{
				"Enghreifftiau",
			},
			given: []string{
				"* ",
				"Anrhegedig a ",
			},
			when: []string{
				"* ",
				"Pryd ",
			},
			then: []string{
				"* ",
				"Yna ",
			},
			and: []string{
				"* ",
				"A ",
			},
			but: []string{
				"* ",
				"Ond ",
			},
		},
	},
	"da": &GherkinDialect{
		"da", "Danish", "dansk", map[string][]string{
			feature: []string{
				"Egenskab",
			},
			background: []string{
				"Baggrund",
			},
			scenario: []string{
				"Scenarie",
			},
			scenarioOutline: []string{
				"Abstrakt Scenario",
			},
			examples: []string{
				"Eksempler",
			},
			given: []string{
				"* ",
				"Givet ",
			},
			when: []string{
				"* ",
				"Når ",
			},
			then: []string{
				"* ",
				"Så ",
			},
			and: []string{
				"* ",
				"Og ",
			},
			but: []string{
				"* ",
				"Men ",
			},
		},
	},
	"de": &GherkinDialect{
		"de", "German", "Deutsch", map[string][]string{
			feature: []string{
				"Funktionalität",
			},
			background: []string{
				"Grundlage",
			},
			scenario: []string{
				"Szenario",
			},
			scenarioOutline: []string{
				"Szenariogrundriss",
			},
			examples: []string{
				"Beispiele",
			},
			given: []string{
				"* ",
				"Angenommen ",
				"Gegeben sei ",
				"Gegeben seien ",
			},
			when: []string{
				"* ",
				"Wenn ",
			},
			then: []string{
				"* ",
				"Dann ",
			},
			and: []string{
				"* ",
				"Und ",
			},
			but: []string{
				"* ",
				"Aber ",
			},
		},
	},
	"el": &GherkinDialect{
		"el", "Greek", "Ελληνικά", map[string][]string{
			feature: []string{
				"Δυνατότητα",
				"Λειτουργία",
			},
			background: []string{
				"Υπόβαθρο",
			},
			scenario: []string{
				"Σενάριο",
			},
			scenarioOutline: []string{
				"Περιγραφή Σεναρίου",
			},
			examples: []string{
				"Παραδείγματα",
				"Σενάρια",
			},
			given: []string{
				"* ",
				"Δεδομένου ",
			},
			when: []string{
				"* ",
				"Όταν ",
			},
			then: []string{
				"* ",
				"Τότε ",
			},
			and: []string{
				"* ",
				"Και ",
			},
			but: []string{
				"* ",
				"Αλλά ",
			},
		},
	},
	"em": &GherkinDialect{
		"em", "Emoji", "😀", map[string][]string{
			feature: []string{
				"📚",
			},
			background: []string{
				"💤",
			},
			scenario: []string{
				"📕",
			},
			scenarioOutline: []string{
				"📖",
			},
			examples: []string{
				"📓",
			},
			given: []string{
				"* ",
				"😐",
			},
			when: []string{
				"* ",
				"🎬",
			},
			then: []string{
				"* ",
				"🙏",
			},
			and: []string{
				"* ",
				"😂",
			},
			but: []string{
				"* ",
				"😔",
			},
		},
	},
	"en": &GherkinDialect{
		"en", "English", "English", map[string][]string{
			feature: []string{
				"Feature",
				"Business Need",
				"Ability",
			},
			background: []string{
				"Background",
			},
			scenario: []string{
				"Scenario",
			},
			scenarioOutline: []string{
				"Scenario Outline",
				"Scenario Template",
			},
			examples: []string{
				"Examples",
				"Scenarios",
			},
			given: []string{
				"* ",
				"Given ",
			},
			when: []string{
				"* ",
				"When ",
			},
			then: []string{
				"* ",
				"Then ",
			},
			and: []string{
				"* ",
				"And ",
			},
			but: []string{
				"* ",
				"But ",
			},
		},
	},
	"en-Scouse": &GherkinDialect{
		"en-Scouse", "Scouse", "Scouse", map[string][]string{
			feature: []string{
				"Feature",
			},
			background: []string{
				"Dis is what went down",
			},
			scenario: []string{
				"The thing of it is",
			},
			scenarioOutline: []string{
				"Wharrimean is",
			},
			examples: []string{
				"Examples",
			},
			given: []string{
				"* ",
				"Givun ",
				"Youse know when youse got ",
			},
			when: []string{
				"* ",
				"Wun ",
				"Youse know like when ",
			},
			then: []string{
				"* ",
				"Dun ",
				"Den youse gotta ",
			},
			and: []string{
				"* ",
				"An ",
			},
			but: []string{
				"* ",
				"Buh ",
			},
		},
	},
	"en-au": &GherkinDialect{
		"en-au", "Australian", "Australian", map[string][]string{
			feature: []string{
				"Pretty much",
			},
			background: []string{
				"First off",
			},
			scenario: []string{
				"Awww, look mate",
			},
			scenarioOutline: []string{
				"Reckon it's like",
			},
			examples: []string{
				"You'll wanna",
			},
			given: []string{
				"* ",
				"Y'know ",
			},
			when: []string{
				"* ",
				"It's just unbelievable ",
			},
			then: []string{
				"* ",
				"But at the end of the day I reckon ",
			},
			and: []string{
				"* ",
				"Too right ",
			},
			but: []string{
				"* ",
				"Yeah nah ",
			},
		},
	},
	"en-lol": &GherkinDialect{
		"en-lol", "LOLCAT", "LOLCAT", map[string][]string{
			feature: []string{
				"OH HAI",
			},
			background: []string{
				"B4",
			},
			scenario: []string{
				"MISHUN",
			},
			scenarioOutline: []string{
				"MISHUN SRSLY",
			},
			examples: []string{
				"EXAMPLZ",
			},
			given: []string{
				"* ",
				"I CAN HAZ ",
			},
			when: []string{
				"* ",
				"WEN ",
			},
			then: []string{
				"* ",
				"DEN ",
			},
			and: []string{
				"* ",
				"AN ",
			},
			but: []string{
				"* ",
				"BUT ",
			},
		},
	},
	"en-old": &GherkinDialect{
		"en-old", "Old English", "Englisc", map[string][]string{
			feature: []string{
				"Hwaet",
				"Hwæt",
			},
			background: []string{
				"Aer",
				"Ær",
			},
			scenario: []string{
				"Swa",
			},
			scenarioOutline: []string{
				"Swa hwaer swa",
				"Swa hwær swa",
			},
			examples: []string{
				"Se the",
				"Se þe",
				"Se ðe",
			},
			given: []string{
				"* ",
				"Thurh ",
				"Þurh ",
				"Ðurh ",
			},
			when: []string{
				"* ",
				"Tha ",
				"Þa ",
				"Ða ",
			},
			then: []string{
				"* ",
				"Tha ",
				"Þa ",
				"Ða ",
				"Tha the ",
				"Þa þe ",
				"Ða ðe ",
			},
			and: []string{
				"* ",
				"Ond ",
				"7 ",
			},
			but: []string{
				"* ",
				"Ac ",
			},
		},
	},
	"en-pirate": &GherkinDialect{
		"en-pirate", "Pirate", "Pirate", map[string][]string{
			feature: []string{
				"Ahoy matey!",
			},
			background: []string{
				"Yo-ho-ho",
			},
			scenario: []string{
				"Heave to",
			},
			scenarioOutline: []string{
				"Shiver me timbers",
			},
			examples: []string{
				"Dead men tell no tales",
			},
			given: []string{
				"* ",
				"Gangway! ",
			},
			when: []string{
				"* ",
				"Blimey! ",
			},
			then: []string{
				"* ",
				"Let go and haul ",
			},
			and: []string{
				"* ",
				"Aye ",
			},
			but: []string{
				"* ",
				"Avast! ",
			},
		},
	},
	"eo": &GherkinDialect{
		"eo", "Esperanto", "Esperanto", map[string][]string{
			feature: []string{
				"Trajto",
			},
			background: []string{
				"Fono",
			},
			scenario: []string{
				"Scenaro",
				"Kazo",
			},
			scenarioOutline: []string{
				"Konturo de la scenaro",
				"Skizo",
				"Kazo-skizo",
			},
			examples: []string{
				"Ekzemploj",
			},
			given: []string{
				"* ",
				"Donitaĵo ",
				"Komence ",
			},
			when: []string{
				"* ",
				"Se ",
			},
			then: []string{
				"* ",
				"Do ",
			},
			and: []string{
				"* ",
				"Kaj ",
			},
			but: []string{
				"* ",
				"Sed ",
			},
		},
	},
	"es": &GherkinDialect{
		"es", "Spanish", "español", map[string][]string{
			feature: []string{
				"Característica",
			},
			background: []string{
				"Antecedentes",
			},
			scenario: []string{
				"Escenario",
			},
			scenarioOutline: []string{
				"Esquema del escenario",
			},
			examples: []string{
				"Ejemplos",
			},
			given: []string{
				"* ",
				"Dado ",
				"Dada ",
				"Dados ",
				"Dadas ",
			},
			when: []string{
				"* ",
				"Cuando ",
			},
			then: []string{
				"* ",
				"Entonces ",
			},
			and: []string{
				"* ",
				"Y ",
				"E ",
			},
			but: []string{
				"* ",
				"Pero ",
			},
		},
	},
	"et": &GherkinDialect{
		"et", "Estonian", "eesti keel", map[string][]string{
			feature: []string{
				"Omadus",
			},
			background: []string{
				"Taust",
			},
			scenario: []string{
				"Stsenaarium",
			},
			scenarioOutline: []string{
				"Raamstsenaarium",
			},
			examples: []string{
				"Juhtumid",
			},
			given: []string{
				"* ",
				"Eeldades ",
			},
			when: []string{
				"* ",
				"Kui ",
			},
			then: []string{
				"* ",
				"Siis ",
			},
			and: []string{
				"* ",
				"Ja ",
			},
			but: []string{
				"* ",
				"Kuid ",
			},
		},
	},
	"fa": &GherkinDialect{
		"fa", "Persian", "فارسی", map[string][]string{
			feature: []string{
				"وِیژگی",
			},
			background: []string{
				"زمینه",
			},
			scenario: []string{
				"سناریو",
			},
			scenarioOutline: []string{
				"الگوی سناریو",
			},
			examples: []string{
				"نمونه ها",
			},
			given: []string{
				"* ",
				"با فرض ",
			},
			when: []string{
				"* ",
				"هنگامی ",
			},
			then: []string{
				"* ",
				"آنگاه ",
			},
			and: []string{
				"* ",
				"و ",
			},
			but: []string{
				"* ",
				"اما ",
			},
		},
	},
	"fi": &GherkinDialect{
		"fi", "Finnish", "suomi", map[string][]string{
			feature: []string{
				"Ominaisuus",
			},
			background: []string{
				"Tausta",
			},
			scenario: []string{
				"Tapaus",
			},
			scenarioOutline: []string{
				"Tapausaihio",
			},
			examples: []string{
				"Tapaukset",
			},
			given: []string{
				"* ",
				"Oletetaan ",
			},
			when: []string{
				"* ",
				"Kun ",
			},
			then: []string{
				"* ",
				"Niin ",
			},
			and: []string{
				"* ",
				"Ja ",
			},
			but: []string{
				"* ",
				"Mutta ",
			},
		},
	},
	"fr": &GherkinDialect{
		"fr", "French", "français", map[string][]string{
			feature: []string{
				"Fonctionnalité",
			},
			background: []string{
				"Contexte",
			},
			scenario: []string{
				"Scénario",
			},
			scenarioOutline: []string{
				"Plan du scénario",
				"Plan du Scénario",
			},
			examples: []string{
				"Exemples",
			},
			given: []string{
				"* ",
				"Soit ",
				"Etant donné que ",
				"Etant donné qu'",
				"Etant donné ",
				"Etant donnée ",
				"Etant donnés ",
				"Etant données ",
				"Étant donné que ",
				"Étant donné qu'",
				"Étant donné ",
				"Étant donnée ",
				"Étant donnés ",
				"Étant données ",
			},
			when: []string{
				"* ",
				"Quand ",
				"Lorsque ",
				"Lorsqu'",
			},
			then: []string{
				"* ",
				"Alors ",
			},
			and: []string{
				"* ",
				"Et que ",
				"Et qu'",
				"Et ",
			},
			but: []string{
				"* ",
				"Mais que ",
				"Mais qu'",
				"Mais ",
			},
		},
	},
	"ga": &GherkinDialect{
		"ga", "Irish", "Gaeilge", map[string][]string{
			feature: []string{
				"Gné",
			},
			background: []string{
				"Cúlra",
			},
			scenario: []string{
				"Cás",
			},
			scenarioOutline: []string{
				"Cás Achomair",
			},
			examples: []string{
				"Samplaí",
			},
			given: []string{
				"* ",
				"Cuir i gcás go",
				"Cuir i gcás nach",
				"Cuir i gcás gur",
				"Cuir i gcás nár",
			},
			when: []string{
				"* ",
				"Nuair a",
				"Nuair nach",
				"Nuair ba",
				"Nuair nár",
			},
			then: []string{
				"* ",
				"Ansin",
			},
			and: []string{
				"* ",
				"Agus",
			},
			but: []string{
				"* ",
				"Ach",
			},
		},
	},
	"gj": &GherkinDialect{
		"gj", "Gujarati", "ગુજરાતી", map[string][]string{
			feature: []string{
				"લક્ષણ",
				"વ્યાપાર જરૂર",
				"ક્ષમતા",
			},
			background: []string{
				"બેકગ્રાઉન્ડ",
			},
			scenario: []string{
				"સ્થિતિ",
			},
			scenarioOutline: []string{
				"પરિદ્દશ્ય રૂપરેખા",
				"પરિદ્દશ્ય ઢાંચો",
			},
			examples: []string{
				"ઉદાહરણો",
			},
			given: []string{
				"* ",
				"આપેલ છે ",
			},
			when: []string{
				"* ",
				"ક્યારે ",
			},
			then: []string{
				"* ",
				"પછી ",
			},
			and: []string{
				"* ",
				"અને ",
			},
			but: []string{
				"* ",
				"પણ ",
			},
		},
	},
	"gl": &GherkinDialect{
		"gl", "Galician", "galego", map[string][]string{
			feature: []string{
				"Característica",
			},
			background: []string{
				"Contexto",
			},
			scenario: []string{
				"Escenario",
			},
			scenarioOutline: []string{
				"Esbozo do escenario",
			},
			examples: []string{
				"Exemplos",
			},
			given: []string{
				"* ",
				"Dado ",
				"Dada ",
				"Dados ",
				"Dadas ",
			},
			when: []string{
				"* ",
				"Cando ",
			},
			then: []string{
				"* ",
				"Entón ",
				"Logo ",
			},
			and: []string{
				"* ",
				"E ",
			},
			but: []string{
				"* ",
				"Mais ",
				"Pero ",
			},
		},
	},
	"he": &GherkinDialect{
		"he", "Hebrew", "עברית", map[string][]string{
			feature: []string{
				"תכונה",
			},
			background: []string{
				"רקע",
			},
			scenario: []string{
				"תרחיש",
			},
			scenarioOutline: []string{
				"תבנית תרחיש",
			},
			examples: []string{
				"דוגמאות",
			},
			given: []string{
				"* ",
				"בהינתן ",
			},
			when: []string{
				"* ",
				"כאשר ",
			},
			then: []string{
				"* ",
				"אז ",
				"אזי ",
			},
			and: []string{
				"* ",
				"וגם ",
			},
			but: []string{
				"* ",
				"אבל ",
			},
		},
	},
	"hi": &GherkinDialect{
		"hi", "Hindi", "हिंदी", map[string][]string{
			feature: []string{
				"रूप लेख",
			},
			background: []string{
				"पृष्ठभूमि",
			},
			scenario: []string{
				"परिदृश्य",
			},
			scenarioOutline: []string{
				"परिदृश्य रूपरेखा",
			},
			examples: []string{
				"उदाहरण",
			},
			given: []string{
				"* ",
				"अगर ",
				"यदि ",
				"चूंकि ",
			},
			when: []string{
				"* ",
				"जब ",
				"कदा ",
			},
			then: []string{
				"* ",
				"तब ",
				"तदा ",
			},
			and: []string{
				"* ",
				"और ",
				"तथा ",
			},
			but: []string{
				"* ",
				"पर ",
				"परन्तु ",
				"किन्तु ",
			},
		},
	},
	"hr": &GherkinDialect{
		"hr", "Croatian", "hrvatski", map[string][]string{
			feature: []string{
				"Osobina",
				"Mogućnost",
				"Mogucnost",
			},
			background: []string{
				"Pozadina",
			},
			scenario: []string{
				"Scenarij",
			},
			scenarioOutline: []string{
				"Skica",
				"Koncept",
			},
			examples: []string{
				"Primjeri",
				"Scenariji",
			},
			given: []string{
				"* ",
				"Zadan ",
				"Zadani ",
				"Zadano ",
			},
			when: []string{
				"* ",
				"Kada ",
				"Kad ",
			},
			then: []string{
				"* ",
				"Onda ",
			},
			and: []string{
				"* ",
				"I ",
			},
			but: []string{
				"* ",
				"Ali ",
			},
		},
	},
	"ht": &GherkinDialect{
		"ht", "Creole", "kreyòl", map[string][]string{
			feature: []string{
				"Karakteristik",
				"Mak",
				"Fonksyonalite",
			},
			background: []string{
				"Kontèks",
				"Istorik",
			},
			scenario: []string{
				"Senaryo",
			},
			scenarioOutline: []string{
				"Plan senaryo",
				"Plan Senaryo",
				"Senaryo deskripsyon",
				"Senaryo Deskripsyon",
				"Dyagram senaryo",
				"Dyagram Senaryo",
			},
			examples: []string{
				"Egzanp",
			},
			given: []string{
				"* ",
				"Sipoze ",
				"Sipoze ke ",
				"Sipoze Ke ",
			},
			when: []string{
				"* ",
				"Lè ",
				"Le ",
			},
			then: []string{
				"* ",
				"Lè sa a ",
				"Le sa a ",
			},
			and: []string{
				"* ",
				"Ak ",
				"Epi ",
				"E ",
			},
			but: []string{
				"* ",
				"Men ",
			},
		},
	},
	"hu": &GherkinDialect{
		"hu", "Hungarian", "magyar", map[string][]string{
			feature: []string{
				"Jellemző",
			},
			background: []string{
				"Háttér",
			},
			scenario: []string{
				"Forgatókönyv",
			},
			scenarioOutline: []string{
				"Forgatókönyv vázlat",
			},
			examples: []string{
				"Példák",
			},
			given: []string{
				"* ",
				"Amennyiben ",
				"Adott ",
			},
			when: []string{
				"* ",
				"Majd ",
				"Ha ",
				"Amikor ",
			},
			then: []string{
				"* ",
				"Akkor ",
			},
			and: []string{
				"* ",
				"És ",
			},
			but: []string{
				"* ",
				"De ",
			},
		},
	},
	"id": &GherkinDialect{
		"id", "Indonesian", "Bahasa Indonesia", map[string][]string{
			feature: []string{
				"Fitur",
			},
			background: []string{
				"Dasar",
			},
			scenario: []string{
				"Skenario",
			},
			scenarioOutline: []string{
				"Skenario konsep",
			},
			examples: []string{
				"Contoh",
			},
			given: []string{
				"* ",
				"Dengan ",
			},
			when: []string{
				"* ",
				"Ketika ",
			},
			then: []string{
				"* ",
				"Maka ",
			},
			and: []string{
				"* ",
				"Dan ",
			},
			but: []string{
				"* ",
				"Tapi ",
			},
		},
	},
	"is": &GherkinDialect{
		"is", "Icelandic", "Íslenska", map[string][]string{
			feature: []string{
				"Eiginleiki",
			},
			background: []string{
				"Bakgrunnur",
			},
			scenario: []string{
				"Atburðarás",
			},
			scenarioOutline: []string{
				"Lýsing Atburðarásar",
				"Lýsing Dæma",
			},
			examples: []string{
				"Dæmi",
				"Atburðarásir",
			},
			given: []string{
				"* ",
				"Ef ",
			},
			when: []string{
				"* ",
				"Þegar ",
			},
			then: []string{
				"* ",
				"Þá ",
			},
			and: []string{
				"* ",
				"Og ",
			},
			but: []string{
				"* ",
				"En ",
			},
		},
	},
	"it": &GherkinDialect{
		"it", "Italian", "italiano", map[string][]string{
			feature: []string{
				"Funzionalità",
			},
			background: []string{
				"Contesto",
			},
			scenario: []string{
				"Scenario",
			},
			scenarioOutline: []string{
				"Schema dello scenario",
			},
			examples: []string{
				"Esempi",
			},
			given: []string{
				"* ",
				"Dato ",
				"Data ",
				"Dati ",
				"Date ",
			},
			when: []string{
				"* ",
				"Quando ",
			},
			then: []string{
				"* ",
				"Allora ",
			},
			and: []string{
				"* ",
				"E ",
			},
			but: []string{
				"* ",
				"Ma ",
			},
		},
	},
	"ja": &GherkinDialect{
		"ja", "Japanese", "日本語", map[string][]string{
			feature: []string{
				"フィーチャ",
				"機能",
			},
			background: []string{
				"背景",
			},
			scenario: []string{
				"シナリオ",
			},
			scenarioOutline: []string{
				"シナリオアウトライン",
				"シナリオテンプレート",
				"テンプレ",
				"シナリオテンプレ",
			},
			examples: []string{
				"例",
				"サンプル",
			},
			given: []string{
				"* ",
				"前提",
			},
			when: []string{
				"* ",
				"もし",
			},
			then: []string{
				"* ",
				"ならば",
			},
			and: []string{
				"* ",
				"かつ",
			},
			but: []string{
				"* ",
				"しかし",
				"但し",
				"ただし",
			},
		},
	},
	"jv": &GherkinDialect{
		"jv", "Javanese", "Basa Jawa", map[string][]string{
			feature: []string{
				"Fitur",
			},
			background: []string{
				"Dasar",
			},
			scenario: []string{
				"Skenario",
			},
			scenarioOutline: []string{
				"Konsep skenario",
			},
			examples: []string{
				"Conto",
				"Contone",
			},
			given: []string{
				"* ",
				"Nalika ",
				"Nalikaning ",
			},
			when: []string{
				"* ",
				"Manawa ",
				"Menawa ",
			},
			then: []string{
				"* ",
				"Njuk ",
				"Banjur ",
			},
			and: []string{
				"* ",
				"Lan ",
			},
			but: []string{
				"* ",
				"Tapi ",
				"Nanging ",
				"Ananging ",
			},
		},
	},
	"kn": &GherkinDialect{
		"kn", "Kannada", "ಕನ್ನಡ", map[string][]string{
			feature: []string{
				"ಹೆಚ್ಚಳ",
			},
			background: []string{
				"ಹಿನ್ನೆಲೆ",
			},
			scenario: []string{
				"ಕಥಾಸಾರಾಂಶ",
			},
			scenarioOutline: []string{
				"ವಿವರಣೆ",
			},
			examples: []string{
				"ಉದಾಹರಣೆಗಳು",
			},
			given: []string{
				"* ",
				"ನೀಡಿದ ",
			},
			when: []string{
				"* ",
				"ಸ್ಥಿತಿಯನ್ನು ",
			},
			then: []string{
				"* ",
				"ನಂತರ ",
			},
			and: []string{
				"* ",
				"ಮತ್ತು ",
			},
			but: []string{
				"* ",
				"ಆದರೆ ",
			},
		},
	},
	"ko": &GherkinDialect{
		"ko", "Korean", "한국어", map[string][]string{
			feature: []string{
				"기능",
			},
			background: []string{
				"배경",
			},
			scenario: []string{
				"시나리오",
			},
			scenarioOutline: []string{
				"시나리오 개요",
			},
			examples: []string{
				"예",
			},
			given: []string{
				"* ",
				"조건",
				"먼저",
			},
			when: []string{
				"* ",
				"만일",
				"만약",
			},
			then: []string{
				"* ",
				"그러면",
			},
			and: []string{
				"* ",
				"그리고",
			},
			but: []string{
				"* ",
				"하지만",
				"단",
			},
		},
	},
	"lt": &GherkinDialect{
		"lt", "Lithuanian", "lietuvių kalba", map[string][]string{
			feature: []string{
				"Savybė",
			},
			background: []string{
				"Kontekstas",
			},
			scenario: []string{
				"Scenarijus",
			},
			scenarioOutline: []string{
				"Scenarijaus šablonas",
			},
			examples: []string{
				"Pavyzdžiai",
				"Scenarijai",
				"Variantai",
			},
			given: []string{
				"* ",
				"Duota ",
			},
			when: []string{
				"* ",
				"Kai ",
			},
			then: []string{
				"* ",
				"Tada ",
			},
			and: []string{
				"* ",
				"Ir ",
			},
			but: []string{
				"* ",
				"Bet ",
			},
		},
	},
	"lu": &GherkinDialect{
		"lu", "Luxemburgish", "Lëtzebuergesch", map[string][]string{
			feature: []string{
				"Funktionalitéit",
			},
			background: []string{
				"Hannergrond",
			},
			scenario: []string{
				"Szenario",
			},
			scenarioOutline: []string{
				"Plang vum Szenario",
			},
			examples: []string{
				"Beispiller",
			},
			given: []string{
				"* ",
				"ugeholl ",
			},
			when: []string{
				"* ",
				"wann ",
			},
			then: []string{
				"* ",
				"dann ",
			},
			and: []string{
				"* ",
				"an ",
				"a ",
			},
			but: []string{
				"* ",
				"awer ",
				"mä ",
			},
		},
	},
	"lv": &GherkinDialect{
		"lv", "Latvian", "latviešu", map[string][]string{
			feature: []string{
				"Funkcionalitāte",
				"Fīča",
			},
			background: []string{
				"Konteksts",
				"Situācija",
			},
			scenario: []string{
				"Scenārijs",
			},
			scenarioOutline: []string{
				"Scenārijs pēc parauga",
			},
			examples: []string{
				"Piemēri",
				"Paraugs",
			},
			given: []string{
				"* ",
				"Kad ",
			},
			when: []string{
				"* ",
				"Ja ",
			},
			then: []string{
				"* ",
				"Tad ",
			},
			and: []string{
				"* ",
				"Un ",
			},
			but: []string{
				"* ",
				"Bet ",
			},
		},
	},
	"mn": &GherkinDialect{
		"mn", "Mongolian", "монгол", map[string][]string{
			feature: []string{
				"Функц",
				"Функционал",
			},
			background: []string{
				"Агуулга",
			},
			scenario: []string{
				"Сценар",
			},
			scenarioOutline: []string{
				"Сценарын төлөвлөгөө",
			},
			examples: []string{
				"Тухайлбал",
			},
			given: []string{
				"* ",
				"Өгөгдсөн нь ",
				"Анх ",
			},
			when: []string{
				"* ",
				"Хэрэв ",
			},
			then: []string{
				"* ",
				"Тэгэхэд ",
				"Үүний дараа ",
			},
			and: []string{
				"* ",
				"Мөн ",
				"Тэгээд ",
			},
			but: []string{
				"* ",
				"Гэхдээ ",
				"Харин ",
			},
		},
	},
	"nl": &GherkinDialect{
		"nl", "Dutch", "Nederlands", map[string][]string{
			feature: []string{
				"Functionaliteit",
			},
			background: []string{
				"Achtergrond",
			},
			scenario: []string{
				"Scenario",
			},
			scenarioOutline: []string{
				"Abstract Scenario",
			},
			examples: []string{
				"Voorbeelden",
			},
			given: []string{
				"* ",
				"Gegeven ",
				"Stel ",
			},
			when: []string{
				"* ",
				"Als ",
			},
			then: []string{
				"* ",
				"Dan ",
			},
			and: []string{
				"* ",
				"En ",
			},
			but: []string{
				"* ",
				"Maar ",
			},
		},
	},
	"no": &GherkinDialect{
		"no", "Norwegian", "norsk", map[string][]string{
			feature: []string{
				"Egenskap",
			},
			background: []string{
				"Bakgrunn",
			},
			scenario: []string{
				"Scenario",
			},
			scenarioOutline: []string{
				"Scenariomal",
				"Abstrakt Scenario",
			},
			examples: []string{
				"Eksempler",
			},
			given: []string{
				"* ",
				"Gitt ",
			},
			when: []string{
				"* ",
				"Når ",
			},
			then: []string{
				"* ",
				"Så ",
			},
			and: []string{
				"* ",
				"Og ",
			},
			but: []string{
				"* ",
				"Men ",
			},
		},
	},
	"pa": &GherkinDialect{
		"pa", "Panjabi", "ਪੰਜਾਬੀ", map[string][]string{
			feature: []string{
				"ਖਾਸੀਅਤ",
				"ਮੁਹਾਂਦਰਾ",
				"ਨਕਸ਼ ਨੁਹਾਰ",
			},
			background: []string{
				"ਪਿਛੋਕੜ",
			},
			scenario: []string{
				"ਪਟਕਥਾ",
			},
			scenarioOutline: []string{
				"ਪਟਕਥਾ ਢਾਂਚਾ",
				"ਪਟਕਥਾ ਰੂਪ ਰੇਖਾ",
			},
			examples: []string{
				"ਉਦਾਹਰਨਾਂ",
			},
			given: []string{
				"* ",
				"ਜੇਕਰ ",
				"ਜਿਵੇਂ ਕਿ ",
			},
			when: []string{
				"* ",
				"ਜਦੋਂ ",
			},
			then: []string{
				"* ",
				"ਤਦ ",
			},
			and: []string{
				"* ",
				"ਅਤੇ ",
			},
			but: []string{
				"* ",
				"ਪਰ ",
			},
		},
	},
	"pl": &GherkinDialect{
		"pl", "Polish", "polski", map[string][]string{
			feature: []string{
				"Właściwość",
				"Funkcja",
				"Aspekt",
				"Potrzeba biznesowa",
			},
			background: []string{
				"Założenia",
			},
			scenario: []string{
				"Scenariusz",
			},
			scenarioOutline: []string{
				"Szablon scenariusza",
			},
			examples: []string{
				"Przykłady",
			},
			given: []string{
				"* ",
				"Zakładając ",
				"Mając ",
				"Zakładając, że ",
			},
			when: []string{
				"* ",
				"Jeżeli ",
				"Jeśli ",
				"Gdy ",
				"Kiedy ",
			},
			then: []string{
				"* ",
				"Wtedy ",
			},
			and: []string{
				"* ",
				"Oraz ",
				"I ",
			},
			but: []string{
				"* ",
				"Ale ",
			},
		},
	},
	"pt": &GherkinDialect{
		"pt", "Portuguese", "português", map[string][]string{
			feature: []string{
				"Funcionalidade",
				"Característica",
				"Caracteristica",
			},
			background: []string{
				"Contexto",
				"Cenário de Fundo",
				"Cenario de Fundo",
				"Fundo",
			},
			scenario: []string{
				"Cenário",
				"Cenario",
			},
			scenarioOutline: []string{
				"Esquema do Cenário",
				"Esquema do Cenario",
				"Delineação do Cenário",
				"Delineacao do Cenario",
			},
			examples: []string{
				"Exemplos",
				"Cenários",
				"Cenarios",
			},
			given: []string{
				"* ",
				"Dado ",
				"Dada ",
				"Dados ",
				"Dadas ",
			},
			when: []string{
				"* ",
				"Quando ",
			},
			then: []string{
				"* ",
				"Então ",
				"Entao ",
			},
			and: []string{
				"* ",
				"E ",
			},
			but: []string{
				"* ",
				"Mas ",
			},
		},
	},
	"ro": &GherkinDialect{
		"ro", "Romanian", "română", map[string][]string{
			feature: []string{
				"Functionalitate",
				"Funcționalitate",
				"Funcţionalitate",
			},
			background: []string{
				"Context",
			},
			scenario: []string{
				"Scenariu",
			},
			scenarioOutline: []string{
				"Structura scenariu",
				"Structură scenariu",
			},
			examples: []string{
				"Exemple",
			},
			given: []string{
				"* ",
				"Date fiind ",
				"Dat fiind ",
				"Dati fiind ",
				"Dați fiind ",
				"Daţi fiind ",
			},
			when: []string{
				"* ",
				"Cand ",
				"Când ",
			},
			then: []string{
				"* ",
				"Atunci ",
			},
			and: []string{
				"* ",
				"Si ",
				"Și ",
				"Şi ",
			},
			but: []string{
				"* ",
				"Dar ",
			},
		},
	},
	"ru": &GherkinDialect{
		"ru", "Russian", "русский", map[string][]string{
			feature: []string{
				"Функция",
				"Функциональность",
				"Функционал",
				"Свойство",
			},
			background: []string{
				"Предыстория",
				"Контекст",
			},
			scenario: []string{
				"Сценарий",
			},
			scenarioOutline: []string{
				"Структура сценария",
			},
			examples: []string{
				"Примеры",
			},
			given: []string{
				"* ",
				"Допустим ",
				"Дано ",
				"Пусть ",
			},
			when: []string{
				"* ",
				"Если ",
				"Когда ",
			},
			then: []string{
				"* ",
				"То ",
				"Тогда ",
			},
			and: []string{
				"* ",
				"И ",
				"К тому же ",
				"Также ",
			},
			but: []string{
				"* ",
				"Но ",
				"А ",
			},
		},
	},
	"sk": &GherkinDialect{
		"sk", "Slovak", "Slovensky", map[string][]string{
			feature: []string{
				"Požiadavka",
				"Funkcia",
				"Vlastnosť",
			},
			background: []string{
				"Pozadie",
			},
			scenario: []string{
				"Scenár",
			},
			scenarioOutline: []string{
				"Náčrt Scenáru",
				"Náčrt Scenára",
				"Osnova Scenára",
			},
			examples: []string{
				"Príklady",
			},
			given: []string{
				"* ",
				"Pokiaľ ",
				"Za predpokladu ",
			},
			when: []string{
				"* ",
				"Keď ",
				"Ak ",
			},
			then: []string{
				"* ",
				"Tak ",
				"Potom ",
			},
			and: []string{
				"* ",
				"A ",
				"A tiež ",
				"A taktiež ",
				"A zároveň ",
			},
			but: []string{
				"* ",
				"Ale ",
			},
		},
	},
	"sl": &GherkinDialect{
		"sl", "Slovenian", "Slovenski", map[string][]string{
			feature: []string{
				"Funkcionalnost",
				"Funkcija",
				"Možnosti",
				"Moznosti",
				"Lastnost",
				"Značilnost",
			},
			background: []string{
				"Kontekst",
				"Osnova",
				"Ozadje",
			},
			scenario: []string{
				"Scenarij",
				"Primer",
			},
			scenarioOutline: []string{
				"Struktura scenarija",
				"Skica",
				"Koncept",
				"Oris scenarija",
				"Osnutek",
			},
			examples: []string{
				"Primeri",
				"Scenariji",
			},
			given: []string{
				"Dano ",
				"Podano ",
				"Zaradi ",
				"Privzeto ",
			},
			when: []string{
				"Ko ",
				"Ce ",
				"Če ",
				"Kadar ",
			},
			then: []string{
				"Nato ",
				"Potem ",
				"Takrat ",
			},
			and: []string{
				"In ",
				"Ter ",
			},
			but: []string{
				"Toda ",
				"Ampak ",
				"Vendar ",
			},
		},
	},
	"sr-Cyrl": &GherkinDialect{
		"sr-Cyrl", "Serbian", "Српски", map[string][]string{
			feature: []string{
				"Функционалност",
				"Могућност",
				"Особина",
			},
			background: []string{
				"Контекст",
				"Основа",
				"Позадина",
			},
			scenario: []string{
				"Сценарио",
				"Пример",
			},
			scenarioOutline: []string{
				"Структура сценарија",
				"Скица",
				"Концепт",
			},
			examples: []string{
				"Примери",
				"Сценарији",
			},
			given: []string{
				"* ",
				"За дато ",
				"За дате ",
				"За дати ",
			},
			when: []string{
				"* ",
				"Када ",
				"Кад ",
			},
			then: []string{
				"* ",
				"Онда ",
			},
			and: []string{
				"* ",
				"И ",
			},
			but: []string{
				"* ",
				"Али ",
			},
		},
	},
	"sr-Latn": &GherkinDialect{
		"sr-Latn", "Serbian (Latin)", "Srpski (Latinica)", map[string][]string{
			feature: []string{
				"Funkcionalnost",
				"Mogućnost",
				"Mogucnost",
				"Osobina",
			},
			background: []string{
				"Kontekst",
				"Osnova",
				"Pozadina",
			},
			scenario: []string{
				"Scenario",
				"Primer",
			},
			scenarioOutline: []string{
				"Struktura scenarija",
				"Skica",
				"Koncept",
			},
			examples: []string{
				"Primeri",
				"Scenariji",
			},
			given: []string{
				"* ",
				"Za dato ",
				"Za date ",
				"Za dati ",
			},
			when: []string{
				"* ",
				"Kada ",
				"Kad ",
			},
			then: []string{
				"* ",
				"Onda ",
			},
			and: []string{
				"* ",
				"I ",
			},
			but: []string{
				"* ",
				"Ali ",
			},
		},
	},
	"sv": &GherkinDialect{
		"sv", "Swedish", "Svenska", map[string][]string{
			feature: []string{
				"Egenskap",
			},
			background: []string{
				"Bakgrund",
			},
			scenario: []string{
				"Scenario",
			},
			scenarioOutline: []string{
				"Abstrakt Scenario",
				"Scenariomall",
			},
			examples: []string{
				"Exempel",
			},
			given: []string{
				"* ",
				"Givet ",
			},
			when: []string{
				"* ",
				"När ",
			},
			then: []string{
				"* ",
				"Så ",
			},
			and: []string{
				"* ",
				"Och ",
			},
			but: []string{
				"* ",
				"Men ",
			},
		},
	},
	"ta": &GherkinDialect{
		"ta", "Tamil", "தமிழ்", map[string][]string{
			feature: []string{
				"அம்சம்",
				"வணிக தேவை",
				"திறன்",
			},
			background: []string{
				"பின்னணி",
			},
			scenario: []string{
				"காட்சி",
			},
			scenarioOutline: []string{
				"காட்சி சுருக்கம்",
				"காட்சி வார்ப்புரு",
			},
			examples: []string{
				"எடுத்துக்காட்டுகள்",
				"காட்சிகள்",
				" நிலைமைகளில்",
			},
			given: []string{
				"* ",
				"கொடுக்கப்பட்ட ",
			},
			when: []string{
				"* ",
				"எப்போது ",
			},
			then: []string{
				"* ",
				"அப்பொழுது ",
			},
			and: []string{
				"* ",
				"மேலும்  ",
				"மற்றும் ",
			},
			but: []string{
				"* ",
				"ஆனால்  ",
			},
		},
	},
	"th": &GherkinDialect{
		"th", "Thai", "ไทย", map[string][]string{
			feature: []string{
				"โครงหลัก",
				"ความต้องการทางธุรกิจ",
				"ความสามารถ",
			},
			background: []string{
				"แนวคิด",
			},
			scenario: []string{
				"เหตุการณ์",
			},
			scenarioOutline: []string{
				"สรุปเหตุการณ์",
				"โครงสร้างของเหตุการณ์",
			},
			examples: []string{
				"ชุดของตัวอย่าง",
				"ชุดของเหตุการณ์",
			},
			given: []string{
				"* ",
				"กำหนดให้ ",
			},
			when: []string{
				"* ",
				"เมื่อ ",
			},
			then: []string{
				"* ",
				"ดังนั้น ",
			},
			and: []string{
				"* ",
				"และ ",
			},
			but: []string{
				"* ",
				"แต่ ",
			},
		},
	},
	"tl": &GherkinDialect{
		"tl", "Telugu", "తెలుగు", map[string][]string{
			feature: []string{
				"గుణము",
			},
			background: []string{
				"నేపథ్యం",
			},
			scenario: []string{
				"సన్నివేశం",
			},
			scenarioOutline: []string{
				"కథనం",
			},
			examples: []string{
				"ఉదాహరణలు",
			},
			given: []string{
				"* ",
				"చెప్పబడినది ",
			},
			when: []string{
				"* ",
				"ఈ పరిస్థితిలో ",
			},
			then: []string{
				"* ",
				"అప్పుడు ",
			},
			and: []string{
				"* ",
				"మరియు ",
			},
			but: []string{
				"* ",
				"కాని ",
			},
		},
	},
	"tlh": &GherkinDialect{
		"tlh", "Klingon", "tlhIngan", map[string][]string{
			feature: []string{
				"Qap",
				"Qu'meH 'ut",
				"perbogh",
				"poQbogh malja'",
				"laH",
			},
			background: []string{
				"mo'",
			},
			scenario: []string{
				"lut",
			},
			scenarioOutline: []string{
				"lut chovnatlh",
			},
			examples: []string{
				"ghantoH",
				"lutmey",
			},
			given: []string{
				"* ",
				"ghu' noblu' ",
				"DaH ghu' bejlu' ",
			},
			when: []string{
				"* ",
				"qaSDI' ",
			},
			then: []string{
				"* ",
				"vaj ",
			},
			and: []string{
				"* ",
				"'ej ",
				"latlh ",
			},
			but: []string{
				"* ",
				"'ach ",
				"'a ",
			},
		},
	},
	"tr": &GherkinDialect{
		"tr", "Turkish", "Türkçe", map[string][]string{
			feature: []string{
				"Özellik",
			},
			background: []string{
				"Geçmiş",
			},
			scenario: []string{
				"Senaryo",
			},
			scenarioOutline: []string{
				"Senaryo taslağı",
			},
			examples: []string{
				"Örnekler",
			},
			given: []string{
				"* ",
				"Diyelim ki ",
			},
			when: []string{
				"* ",
				"Eğer ki ",
			},
			then: []string{
				"* ",
				"O zaman ",
			},
			and: []string{
				"* ",
				"Ve ",
			},
			but: []string{
				"* ",
				"Fakat ",
				"Ama ",
			},
		},
	},
	"tt": &GherkinDialect{
		"tt", "Tatar", "Татарча", map[string][]string{
			feature: []string{
				"Мөмкинлек",
				"Үзенчәлеклелек",
			},
			background: []string{
				"Кереш",
			},
			scenario: []string{
				"Сценарий",
			},
			scenarioOutline: []string{
				"Сценарийның төзелеше",
			},
			examples: []string{
				"Үрнәкләр",
				"Мисаллар",
			},
			given: []string{
				"* ",
				"Әйтик ",
			},
			when: []string{
				"* ",
				"Әгәр ",
			},
			then: []string{
				"* ",
				"Нәтиҗәдә ",
			},
			and: []string{
				"* ",
				"Һәм ",
				"Вә ",
			},
			but: []string{
				"* ",
				"Ләкин ",
				"Әмма ",
			},
		},
	},
	"uk": &GherkinDialect{
		"uk", "Ukrainian", "Українська", map[string][]string{
			feature: []string{
				"Функціонал",
			},
			background: []string{
				"Передумова",
			},
			scenario: []string{
				"Сценарій",
			},
			scenarioOutline: []string{
				"Структура сценарію",
			},
			examples: []string{
				"Приклади",
			},
			given: []string{
				"* ",
				"Припустимо ",
				"Припустимо, що ",
				"Нехай ",
				"Дано ",
			},
			when: []string{
				"* ",
				"Якщо ",
				"Коли ",
			},
			then: []string{
				"* ",
				"То ",
				"Тоді ",
			},
			and: []string{
				"* ",
				"І ",
				"А також ",
				"Та ",
			},
			but: []string{
				"* ",
				"Але ",
			},
		},
	},
	"ur": &GherkinDialect{
		"ur", "Urdu", "اردو", map[string][]string{
			feature: []string{
				"صلاحیت",
				"کاروبار کی ضرورت",
				"خصوصیت",
			},
			background: []string{
				"پس منظر",
			},
			scenario: []string{
				"منظرنامہ",
			},
			scenarioOutline: []string{
				"منظر نامے کا خاکہ",
			},
			examples: []string{
				"مثالیں",
			},
			given: []string{
				"* ",
				"اگر ",
				"بالفرض ",
				"فرض کیا ",
			},
			when: []string{
				"* ",
				"جب ",
			},
			then: []string{
				"* ",
				"پھر ",
				"تب ",
			},
			and: []string{
				"* ",
				"اور ",
			},
			but: []string{
				"* ",
				"لیکن ",
			},
		},
	},
	"uz": &GherkinDialect{
		"uz", "Uzbek", "Узбекча", map[string][]string{
			feature: []string{
				"Функционал",
			},
			background: []string{
				"Тарих",
			},
			scenario: []string{
				"Сценарий",
			},
			scenarioOutline: []string{
				"Сценарий структураси",
			},
			examples: []string{
				"Мисоллар",
			},
			given: []string{
				"* ",
				"Агар ",
			},
			when: []string{
				"* ",
				"Агар ",
			},
			then: []string{
				"* ",
				"Унда ",
			},
			and: []string{
				"* ",
				"Ва ",
			},
			but: []string{
				"* ",
				"Лекин ",
				"Бирок ",
				"Аммо ",
			},
		},
	},
	"vi": &GherkinDialect{
		"vi", "Vietnamese", "Tiếng Việt", map[string][]string{
			feature: []string{
				"Tính năng",
			},
			background: []string{
				"Bối cảnh",
			},
			scenario: []string{
				"Tình huống",
				"Kịch bản",
			},
			scenarioOutline: []string{
				"Khung tình huống",
				"Khung kịch bản",
			},
			examples: []string{
				"Dữ liệu",
			},
			given: []string{
				"* ",
				"Biết ",
				"Cho ",
			},
			when: []string{
				"* ",
				"Khi ",
			},
			then: []string{
				"* ",
				"Thì ",
			},
			and: []string{
				"* ",
				"Và ",
			},
			but: []string{
				"* ",
				"Nhưng ",
			},
		},
	},
	"zh-CN": &GherkinDialect{
		"zh-CN", "Chinese simplified", "简体中文", map[string][]string{
			feature: []string{
				"功能",
			},
			background: []string{
				"背景",
			},
			scenario: []string{
				"场景",
				"剧本",
			},
			scenarioOutline: []string{
				"场景大纲",
				"剧本大纲",
			},
			examples: []string{
				"例子",
			},
			given: []string{
				"* ",
				"假如",
				"假设",
				"假定",
			},
			when: []string{
				"* ",
				"当",
			},
			then: []string{
				"* ",
				"那么",
			},
			and: []string{
				"* ",
				"而且",
				"并且",
				"同时",
			},
			but: []string{
				"* ",
				"但是",
			},
		},
	},
	"zh-TW": &GherkinDialect{
		"zh-TW", "Chinese traditional", "繁體中文", map[string][]string{
			feature: []string{
				"功能",
			},
			background: []string{
				"背景",
			},
			scenario: []string{
				"場景",
				"劇本",
			},
			scenarioOutline: []string{
				"場景大綱",
				"劇本大綱",
			},
			examples: []string{
				"例子",
			},
			given: []string{
				"* ",
				"假如",
				"假設",
				"假定",
			},
			when: []string{
				"* ",
				"當",
			},
			then: []string{
				"* ",
				"那麼",
			},
			and: []string{
				"* ",
				"而且",
				"並且",
				"同時",
			},
			but: []string{
				"* ",
				"但是",
			},
		},
	},
}

