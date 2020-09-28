package eliza

var (
	pre = map[string]string{
		"no":      "don't",
		"no puedo":      "can't",
		"no quiero":      "won't",
		"recordar": "remember",
		"soñe":    "dreamed",
		"sueño":    "dream",
		"quizas":     "perhaps",
		"como":       "que",
		"cuando":      "what",
		"sefuramente": "si",
		"maquina":   "computadora",
		"computadoras": "computadora",
		"fue":      "fue",
		"tu eres":    "tu eres",
		"yo soy":       "yo soy",
		"igual":      "parecido",
	}

	post = map[string]string{
		"soy":    "somos",
		"tuyo es":   "mio",
		"yo":     "tu",
		"yo soy": "tu eres",
		"para ti": "para mi",
		"yoo":      "tu",
		"tuyo":    "mio",
		"mio":     "tuyo0",
		"yo soy el ":  "tu eres",
	}

	synonyms = map[string][]string{
		"creer":   {"creer", "sentir", "pensar", "creer", "desear"},
		"familia":   {"familia", "madre", "mama", "papa", "papi", "hermana", "hermano", "esposa", "hijos", "hijo"},
		"deseo":   {"deseo", "quiero", "necesito"},
		"triste":      {"triste", "infeliz", "des=presivo", "enfermo"},
		"feliz":    {"felix", "elated", "glad", "mejor"},
		"no puedo":   {"no puedo", "no puedo"},
		"cada uno": {"todos", "en todo el mundo", "nadie", "ninguno"},
		"ser":       {"soy", "yo", "es", "somos", "fue"},
	}

	quit = []string{"bye","adios", "te veo luego"}
)

type keyword struct {
	Weight         uint8 // Importance of the keyword - will be sorted descending
	Decompositions []*decomp
}

type decomp struct {
	Pattern      string
	Assemblies   []string
	AssemblyNext uint8
}

var keywordMap = map[string]keyword{

	"xnone": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"No creo estar entendiendo.",
					"Por favor , continua.",
					"Como eso te hace sentir ?",
					"Realmente quieres discutir este tema ?",
				},
			},
		},
	},

	"sorry": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"No necesitas excusarte.",
					"Las excusas no son necesarias.",
					"Te repito, no es necesario excusarte.",
				},
			},
		},
	},

	"apologise": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"goto sorry",
				},
			},
		},
	},

	"remember": {
		Weight: 5,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?i remember ?(.*)",
				Assemblies: []string{
					"que tanto piensas de  (2) ?",
					"Does thinking of (2) bring anything else to mind ?",
					"que mas puedes recolectar ?",
					"Why do you recollect (2) just now ?",
					"esta situacion personal que podemos discutir (2) ?",
					"cual es la conexion entre yo y (2) ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?do you remember ?(.*)",
				Assemblies: []string{
					"Did you think I would forget (2) ?",
					"Why do you think I should recall (2) now ?",
					"What about (2) ?",
					"goto what",
					"You mentioned (2) ?",
				},
			},
		},
	},

	"if": {
		Weight: 3,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?if ?(.*)",
				Assemblies: []string{
					"Do you think its likely that (2) ?",
					"Do you wish that (2) ?",
					"What do you know about (2) ?",
					"Really, if (2) ?",
				},
			},
		},
	},

	"dreamed": {
		Weight: 4,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?i dreamed ?(.*)",
				Assemblies: []string{
					"Really, (2) ?",
					"Have you ever fantasized (2) while you were awake ?",
					"Have you ever dreamed (2) before ?",
					"goto dream",
				},
			},
		},
	},

	"dream": {
		Weight: 3,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"What does that dream suggest to you ?",
					"Do you dream often ?",
					"What persons appear in your dreams ?",
					"Do you believe that dreams have something to do with your problems ?",
				},
			},
		},
	},

	"perhaps": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"You don't seem quite certain.",
					"Why the uncertain tone ?",
					"Can't you be more positive ?",
					"You aren't sure ?",
					"Don't you know ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"I am not interested in names.",
					"I've told you before, I don't care about names -",
				},
			},
		},
	},

	"deutsch": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"goto xforeign",
					"I told you before, I don't understand German.",
				},
			},
		},
	},

	"francais": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"goto xforeign",
					"I told you before, I don't understand French.",
				},
			},
		},
	},

	"italiano": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"goto xforeign",
					"I told you before, I don't understand Italian.",
				},
			},
		},
	},

	"espanol": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"goto xforeign",
					"I told you before, I don't understand Spanish.",
				},
			},
		},
	},

	"xforeign": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"I speak only English.",
				},
			},
		},
	},

	"hello": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"How do you do.	Please state your problem.",
					"Hi.	What seems to be your problem ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"Do computers worry you ?",
					"Why do you mention computers ?",
					"What do you think machines have to do with your problem ?",
					"Don't you think computers can help people ?",
					"What about machines worrys you ?",
					"What do you think about machines ?",
				},
			},
		},
	},

	"am": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?am i ?(.*)",
				Assemblies: []string{
					"Do you believe you are (2) ?",
					"Would you want to be (2) ?",
					"Do you wish I would tell you you are (2) ?",
					"What would it mean if you were (2) ?",
					"goto what",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"Why do you say 'am' ?",
					"I don't understand that.",
				},
			},
		},
	},

	"are": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?are you ?(.*)",
				Assemblies: []string{
					"Why are you interested in whether I am (2) or not ?",
					"Would you prefer if I weren't (2) ?",
					"Perhaps I am (2) in your fantasies.",
					"Do you sometimes think I am (2) ?",
					"goto what",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?are ?(.*)",
				Assemblies: []string{
					"Did you think they might not be (2) ?",
					"Would you like it if they were not (2) ?",
					"What if they were not (2) ?",
					"Possibly they are (2).",
				},
			},
		},
	},

	"your": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?your ?(.*)",
				Assemblies: []string{
					"Why are you concerned over my (2) ?",
					"What about your own (2) ?",
					"Are you worried about someone else's (2) ?",
					"Really, my (2) ?",
				},
			},
		},
	},

	"was": {
		Weight: 2,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?was i ?(.*)",
				Assemblies: []string{
					"What if you were (2) ?",
					"Do you think you were (2) ?",
					"Were you (2) ?",
					"What would it mean if you were (2) ?",
					"What does (2) suggest to you ?",
					"goto what",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?i was ?(.*)",
				Assemblies: []string{
					"Were you really ?",
					"Why do you tell me you were (2) now ?",
					"Perhaps I already know you were (2).",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?was you ?(.*)",
				Assemblies: []string{
					"Would you like to believe I was (2) ?",
					"What suggests that I was (2) ?",
					"What do you think ?",
					"Perhaps I was (2).",
					"What if I had been (2) ?",
				},
			},
		},
	},

	"i": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?i @desire ?(.*)",
				Assemblies: []string{
					"What would it mean to you if you got (2) ?",
					"Why do you want (2) ?",
					"Suppose you got (2) soon ?",
					"What if you never got (2) ?",
					"What would getting (2) mean to you ?",
					"What does wanting (2) have to do with this discussion ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?i am (.*)@sad ?(.*)",
				Assemblies: []string{
					"I am sorry to hear that you are (3).",
					"Do you think that coming here will help you not to be (3) ?",
					"I'm sure it's not pleasant to be (3).",
					"Can you explain what made you (3) ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?i am (.*)@happy ?(.*)",
				Assemblies: []string{
					"How have I helped you to be (3) ?",
					"Has your treatment made you (3) ?",
					"What makes you (3) just now ?",
					"Can you explan why you are suddenly (3) ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?i was ?(.*)",
				Assemblies: []string{
					"goto was",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?i @belief (.*) i ?(.*)",
				Assemblies: []string{
					"Do you really think so ?",
					"But you are not sure you (3).",
					"Do you really doubt you (3) ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?i (.*)@belief (.*) you ?(.*)",
				Assemblies: []string{
					"goto you",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?i am ?(.*)",
				Assemblies: []string{
					"Is it because you are (2) that you came to me ?",
					"How long have you been (2) ?",
					"Do you believe it is normal to be (2) ?",
					"Do you enjoy being (2) ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?i @cannot ?(.*)",
				Assemblies: []string{
					"How do you think that you can't (2) ?",
					"Have you tried ?",
					"Perhaps you could (2) now.",
					"Do you really want to be able to (2) ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?i don't ?(.*)",
				Assemblies: []string{
					"Don't you really (2) ?",
					"Why don't you (2) ?",
					"Do you wish to be able to (2) ?",
					"Does that trouble you ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?do i feel ?(.*)",
				Assemblies: []string{
					"Tell me more about such feelings.",
					"Do you often feel (2) ?",
					"Do you enjoy feeling (2) ?",
					"Of what does feeling (2) remind you ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?i (.*) you ?(.*)",
				Assemblies: []string{
					"Perhaps in your fantasies we (2) each other.",
					"Do you wish to (2) me ?",
					"You seem to need to (2) me.",
					"Do you (2) anyone else ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"You say (1) ?",
					"Can you elaborate on that ?",
					"Do you say (1) for some special reason ?",
					"That's quite interesting.",
				},
			},
		},
	},

	"you": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?you remind me of ?(.*)",
				Assemblies: []string{
					"goto alike",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?you are ?(.*)",
				Assemblies: []string{
					"What makes you think I am (2) ?",
					"Does it please you to believe I am (2) ?",
					"Do you sometimes wish you were (2) ?",
					"Perhaps you would like to be (2).",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?you (.*) me ?(.*)",
				Assemblies: []string{
					"Why do you think I (2) you ?",
					"You like to think I (2) you -",
					"What makes you think I (2) you ?",
					"Really, I (2) you ?",
					"Do you wish to believe I (2) you ?",
					"Suppose I did (2) you -",
					"Does someone else believe I (2) you ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?you ?(.*)",
				Assemblies: []string{
					"We were discussing you -",
					"Oh, I (2) ?",
					"You're not really talking about me -",
					"What are your feelings now ?",
				},
			},
		},
	},

	"yes": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"You seem to be quite positive.",
					"You are sure.",
					"I see.",
					"I understand.",
				},
			},
		},
	},

	"no": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"Are you saying no just to be negative?",
					"You are being a bit negative.",
					"Why not ?",
					"Why 'no' ?",
				},
			},
		},
	},

	"my": {
		Weight: 2,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "$(.*) ?my ?(.*)",
				Assemblies: []string{
					"Lets discuss further why your (2).",
					"Earlier you said your (2).",
					"But your (2).",
					"Does that have anything to do with the fact that your (2) ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?my (.*)@family ?(.*)",
				Assemblies: []string{
					"Tell me more about your family.",
					"Who else in your family (4) ?",
					"Your (3) ?",
					"What else comes to mind when you think of your (3) ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?my ?(.*)",
				Assemblies: []string{
					"Your (2) ?",
					"Why do you say your (2) ?",
					"Does that suggest anything else which belongs to you ?",
					"Is it important that your (2) ?",
				},
			},
		},
	},

	"can": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?can you ?(.*)",
				Assemblies: []string{
					"You believe I can (2) don't you ?",
					"goto what",
					"You want me to be able to (2).",
					"Perhaps you would like to be able to (2) yourself.",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?can i ?(.*)",
				Assemblies: []string{
					"Whether or not you can (2) depends on you more than me.",
					"Do you want to be able to (2) ?",
					"Perhaps you don't want to (2).",
					"goto what",
				},
			},
		},
	},

	"what": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"Por qué preguntas ?",
					"Te interesa esa pregunta ?",
					"Qué es lo que realmente querías saber?",
					"Hay tantas preguntas en tu mente ?",
					"Qué respuesta te agradaría más ?",
					"Qué piensas ?",
					"Qué te viene a la mente cuando preguntas eso ?",
					"Ha hecho antes esas preguntas ?",
					"Le has preguntado a alguien más ?",
				},
			},
		},
	},

	"because": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"Es esa la verdadera razón ?",
					"No se me ocurren otras razones ?",
					"Esa razón parece explicar algo más ?",
					"Qué otras razones puede haber ?",
				},
			},
		},
	},

	"why": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?why don't you ?(.*)",
				Assemblies: []string{
					"Crees que yo no (2) ?",
					"Quizás lo haga (2) a su debido tiempo.",
					"Deberías (2) tú mismo ?",
					"Quieres que yo (2) ?",
					"goto what",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*) ?why can't i ?(.*)",
				Assemblies: []string{
					"Crees que deberías poder (2) ?",
					"Quieres poder (2) ?",
					"Crees que esto te ayudará a (2) ?",
					"Tienes alguna idea de por qué no puedes (2) ?",
					"goto what",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"goto what",
				},
			},
		},
	},

	"everyone": {
		Weight: 2,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)@everyone ?(.*)",
				Assemblies: []string{
					"De verdad, (2) ?",
					"Seguramente que no (2).",
					"Puedes pensar en alguien en particular ?",
					"Quien , Por ejemplo?",
					"Estas pensando en una persona muy especial ?",
					"A quién puedo preguntar ?",
					"Alguien especial quizás ?",
					"Tienes a una persona en particular en mente, ¿no? ?",
					"De quién crees que estás hablando ?",
				},
			},
		},
	},

	"everybody": {
		Weight: 2,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"goto everyone",
				},
			},
		},
	},

	"nobody": {
		Weight: 2,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"goto everyone",
				},
			},
		},
	},

	"no one": {
		Weight: 2,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"goto everyone",
				},
			},
		},
	},

	"always": {
		Weight: 1,
		Decompositions: []*decomp{
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"Puedes pensar en un ejemplo específico ?",
					"Cuando ?",
					"En que incidente estas pensando ?",
					"De verdad, siempre ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*)",
				Assemblies: []string{
					"De que manera ?",
					"Que semejanza le ves ?",
					"Y esta similitud que deja dicho ?",
					"Que otra conexion ves ?",
					"Qué crees que significa ese parecido ?",
					"Cuál es la conexión, supones ?",
					"Podría haber aquí realmente alguna conexión ?",
					"Como ?",
				},
			},
			{
				AssemblyNext: 0,
				Pattern:      "(.*)@be (.*) like ?(.*)",
				Assemblies: []string{
					"goto alike",
				},
			},
		},
	},
}
