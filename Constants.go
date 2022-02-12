package lovelettercreator

func genSetOne() []string {
 return []string{
      "You complete me.",
      "My heart always skips a beat when I think of you.",
      "You are my reason to be.",
      "I want to take you far far away.",
      "You are my favorite distraction.",
      "Your love makes my world go round.",
    }
}

func genSetTwo() []string {
    return []string{
      "You give meaning to my life.",
      "There was never a love like ours and there never will be again.",
      "My life is complete with you.",
      "You are my sunny day in spring.",
      "You are my warmth in winter.",
      "You are my breeze in summer.",
    }
}

func heSays() []string {
    return []string{
      "The way you look at me makes me feel strong and confident.",
      "You are the beauty of autumn.",
      "I crown you the Queen of my heart.",
      "Every time I see you I fall in love all over again.",
      "You bring more to me than I ever wanted.",
      "You are the girl of my dreams.",
    }
}
func sheSays() []string{
    return []string {
      "When I am in your arms, I feel safe, I feel loved and I feel like I am going to be protected from all the stones life can throw.",
      "Thank you for making me feel like the most beautiful person in the world.",
      "You should teach other men how to love their partners and make them feel special because you have it down to an art.",
      "You are my knight in shining armor. ",
      "You are the man of my dreams.",
    }
}

func altSalFemale() []string {
    return []string{
      "dame",
      "sweetheart",
      "delicacy",
      "darling",
      "angel",
      "missy",
      "baby",
    }
}

func altSalMale() []string{
    return []string{
      "sir",
      "sweetheart",
      "darling",
      "suitor",
      "sire",
      "gentleman",
      "monsieur",
      "baby",
    }
}

func verbs() []string  {
    return []string{
      "desires",
      "yearns",
      "ravish",
      "craves",
      "covets",
      "fancies",
      "caress",
      "lusts",
      "cares",
      "adores",
      "pines",
      "thirsts",
      "longs",
    }
}

func adverbs() []string{
    return []string {
      "passionately",
      "patiently",
      "fondly",
      "avidly",
      "intensely",
      "earnestly",
      "dearly",
      "eagerly",
      "fiercely",
      "lovingly",
      "amorously",
      "tenderly",
    }
}

func senderAdj()[]string {
    return []string{
      "burning",
      "throbbing",
      "fiery",
      "edgy",
      "sensual",
      "affectionate",
      "romantic",
      "carnal",
      "deep",
      "fierce",
      "avid",
      "ardent",
      "sympathetic",
      "erotic",
      "unsatisfied",
    }
}

func recipAdj() []string {
    return []string{
      "lovely",
      "charming",
      "little",
      "sweet",
      "amazing",
      "teasing",
      "magical",
      "seductive",
      "lovable",
      "dearest",
      "fond",
      "tender",
    }
}

func senderNouns() []string {
    return []string{
      "eagerness",
      "longing",
      "devotion",
      "infatuation",
      "adoration",
      "hunger",
      "wish",
      "ambition",
      "desire",
      "passion",
      "obsession",
    }
}

func recipNouns() []string {
    list := []string{
      "tenderness",
      "love",
      "kindness",
      "hotness",
      "loveliness",
      "allure",
      "gorgeousness",
      "sexiness",
      "fancy",
      "rapture",
      "fervour",
      "ardour",
    }
    return list
}

func bodyParts(gender bool) []string {
    list:= []string{
      "eyes",
      "nose",
      "mouth",
      "hair",
      "hands",
      "arms",
      "rear",
      "thighs",
      "legs",
      "lips",
      "voice",
      "touch",
      "face",
      "smile",
    }
    if gender {
        list=append(list,"bosom") //female
    }else{
        list=append(list,"chest") //male
    }
    return list
}

func closeWords() []string {
    return []string{
        "Always",
        "Cheers",
        "Hugs",
        "Later",
        "Love",
        "See ya",
        "Reply soon",
        "Take care",
        "XOXO",
    }
}
