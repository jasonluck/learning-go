package deck

import "testing"

func TestCardString(t *testing.T) {

	cases := []struct {
		in   Card
		want string
	}{
		{Card{value: 13, suit: SPADES}, "KS"},
		{Card{value: 2, suit: HEARTS}, "2H"},
		{Card{value: 12, suit: CLUBS}, "QC"},
		{Card{value: 11, suit: DIAMONDS}, "JD"},
		{Card{value: 1, suit: DIAMONDS}, "AD"},
		{Card{value: 0, suit: JOKER}, "Joker"},
	}

	for _, c := range cases {
		result := c.in.String()
		if result != c.want {
			t.Errorf("Card(%q) == %q, wanted %q", c.in, result, c.want)
		}
	}
}

func TestSuits(t *testing.T) {
	result := Suits()

	if result[0] != SPADES ||
		result[1] != DIAMONDS ||
		result[2] != CLUBS ||
		result[3] != HEARTS {
		t.Errorf("Suits were not in correct order")
	}
}
