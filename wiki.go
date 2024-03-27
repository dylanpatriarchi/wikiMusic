package main

import (
    "math/rand"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

var artists = map[string][]string{
    "rock":      {"Queen", "The Beatles", "Led Zeppelin", "Pink Floyd", "The Rolling Stones", "Nirvana", "AC/DC", "Radiohead", "Guns N' Roses", "Metallica"},
    "pop":       {"Michael Jackson", "Madonna", "Elton John", "Beyoncé", "Justin Bieber", "Taylor Swift", "Ariana Grande", "Ed Sheeran", "Mariah Carey", "Lady Gaga"},
    "hip-hop":   {"Eminem", "Jay-Z", "Drake", "Kanye West", "Nicki Minaj", "Tupac Shakur", "Notorious B.I.G.", "Snoop Dogg", "Nas", "Kendrick Lamar"},
    "jazz":      {"Louis Armstrong", "Miles Davis", "John Coltrane", "Duke Ellington", "Charlie Parker", "Ella Fitzgerald", "Billie Holiday", "Thelonious Monk", "Dave Brubeck", "Chick Corea"},
    "reggae":    {"Bob Marley", "Peter Tosh", "Burning Spear", "Jimmy Cliff", "Toots and the Maytals", "Lee 'Scratch' Perry", "Bunny Wailer", "Gregory Isaacs", "Damian Marley", "Buju Banton"},
    "classical": {"Ludwig van Beethoven", "Johann Sebastian Bach", "Wolfgang Amadeus Mozart", "Franz Schubert", "Antonio Vivaldi", "Pyotr Ilyich Tchaikovsky", "Gustav Mahler", "Frederic Chopin", "Igor Stravinsky", "Claude Debussy"},
    "country":   {"Johnny Cash", "Willie Nelson", "Dolly Parton", "Hank Williams", "Garth Brooks", "Shania Twain", "Johnny Paycheck", "George Strait", "Alan Jackson", "Kenny Rogers"},
    "blues":     {"Robert Johnson", "B.B. King", "Muddy Waters", "John Lee Hooker", "Howlin' Wolf", "Albert King", "Stevie Ray Vaughan", "Etta James", "Buddy Guy", "Lead Belly"},
    "metal":     {"Black Sabbath", "Iron Maiden", "Judas Priest", "Metallica", "Slayer", "Megadeth", "Pantera", "Tool", "Opeth", "Dream Theater"},
    "electronic": {"Daft Punk", "The Chemical Brothers", "Aphex Twin", "Kraftwerk", "Deadmau5", "Skrillex", "Jean-Michel Jarre", "Fatboy Slim", "Tiesto", "Avicii"},
    "folk":      {"Bob Dylan", "Joni Mitchell", "Woody Guthrie", "Pete Seeger", "Simon & Garfunkel", "Neil Young", "James Taylor", "Bon Iver", "Nick Drake", "Joan Baez"},
    "reggaeton": {"Daddy Yankee", "Don Omar", "Wisin & Yandel", "Nicky Jam", "Ozuna", "J Balvin", "Maluma", "Anuel AA", "Karol G", "Bad Bunny"},
    "indie":     {"Arctic Monkeys", "The Strokes", "Vampire Weekend", "Arcade Fire", "Tame Impala", "The Killers", "Fleet Foxes", "Interpol", "The National", "MGMT"},
}

func getRandomArtists(genres []string) []string {
    var recommendedArtists []string
    for _, genre := range genres {
        genre = strings.ToLower(genre)
        if artistsList, ok := artists[genre]; ok {
            artist := artistsList[rand.Intn(len(artistsList))]
            recommendedArtists = append(recommendedArtists, artist)
        }
    }
    return recommendedArtists
}

func main() {
    router := gin.Default()

    router.Static("/static", "./static")

    router.LoadHTMLFiles("index.html")

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    router.POST("/recommend", func(c *gin.Context) {
        var requestBody struct {
            Genres []string `json:"genres"`
        }
        if err := c.ShouldBindJSON(&requestBody); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        recommendedArtists := getRandomArtists(requestBody.Genres)

        c.JSON(http.StatusOK, gin.H{"artists": recommendedArtists})
    })

    router.Run(":8080")
}
