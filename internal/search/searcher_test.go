package search

import (
	"testing"
)

var resultsExample = []struct {
	in  string
	out string
}{
	{`song = {
'artist':'Patricio Rey',
'song':'Mi Perro Dinamita',
'lyrics':'-Yo no sé si a tu perro le gusta ladrar a lo gogó, \nmi perro ¡no! no quiere ¡no! \ncon el h[...]',
'url':'http://lyrics.wikia.com/Patricio_Rey_Y_Sus_Redonditos_De_Ricota:Mi_Perro_Dinamita'
}
`, "http://lyrics.wikia.com/Patricio_Rey_Y_Sus_Redonditos_De_Ricota:Mi_Perro_Dinamita"},
	{`song = {
'artist':'Patricio Rey Y ' Sus Redonditos De Ricota',
'song':'Mi Perro Dinamita',
'lyrics':'-"',
'url':'http://lyrics.wikia.com/Patricio_Rey_Y_Sus_Redonditos_De_Ricota:Mi_Perro_Dinamita'
}
`, "http://lyrics.wikia.com/Patricio_Rey_Y_Sus_Redonditos_De_Ricota:Mi_Perro_Dinamita"},
}

func TestFlagParser(t *testing.T) {
	for _, tt := range resultsExample {
		t.Run(tt.in, func(t *testing.T) {
			s := wikiaExtractSong(tt.in)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}
