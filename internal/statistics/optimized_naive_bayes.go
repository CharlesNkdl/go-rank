package statistics

import "fmt"

// reprise de la logique du code du simple
// mais en essayant d'optimiser
// donc réduire le big O , parce que jme suis pas pris la tête
// et j'ai scanné plusieurs fois la totalité du dataset
// aussi de rendre agnostique l'application de naive bayes sur n'importe quel donnée

var DatasetOptimized = []Classes{
	// Pommes (10) : 7 rouges, 2 vertes, 1 jaune
	{"fruit": "Pomme", "color": "Rouge"},
	{"fruit": "Pomme", "color": "Rouge"},
	{"fruit": "Pomme", "color": "Rouge"},
	{"fruit": "Pomme", "color": "Rouge"},
	{"fruit": "Pomme", "color": "Rouge"},
	{"fruit": "Pomme", "color": "Rouge"},
	{"fruit": "Pomme", "color": "Rouge"},
	{"fruit": "Pomme", "color": "Vert"},
	{"fruit": "Pomme", "color": "Vert"},
	{"fruit": "Pomme", "color": "Jaune"},

	// Poires (6) : 1 rouge, 4 vertes, 1 jaune
	{"fruit": "Poire", "color": "Rouge"},
	{"fruit": "Poire", "color": "Vert"},
	{"fruit": "Poire", "color": "Vert"},
	{"fruit": "Poire", "color": "Vert"},
	{"fruit": "Poire", "color": "Vert"},
	{"fruit": "Poire", "color": "Jaune"},

	// Bananes (4) : 4 jaunes
	{"fruit": "Banane", "color": "Jaune"},
	{"fruit": "Banane", "color": "Jaune"},
	{"fruit": "Banane", "color": "Jaune"},
	{"fruit": "Banane", "color": "Jaune"},
}

type Classes map[string]string

type BayesOptimized struct {
	total float64
	// on va compter les occurences une seule fois
	counts map[string]map[string]float64
	// bizarre mais techniquement c'est ok, perds de l'opti au niveau mémoire mais réduit le calcul
	// on va calculer une seule fois chaque likelihood de chaque classes
	joint map[string]map[string]map[string]map[string]float64
}

func NewBayesOptimized(dataset []Classes) *BayesOptimized {
	b := &BayesOptimized{
		total:  float64(len(dataset)),
		counts: make(map[string]map[string]float64),
		joint:  make(map[string]map[string]map[string]map[string]float64),
	}
	for _, j := range dataset {
		for A, Aval := range j {
			if b.counts[A] == nil {
				b.counts[A] = make(map[string]float64)
			}
			b.counts[A][Aval]++
			for B, Bval := range j {
				if A == B {
					continue
				}
				if b.joint[A] == nil {
					b.joint[A] = make(map[string]map[string]map[string]float64)
				}
				if b.joint[A][Aval] == nil {
					b.joint[A][Aval] = make(map[string]map[string]float64)
				}
				if b.joint[A][Aval][B] == nil {
					b.joint[A][Aval][B] = make(map[string]float64)
				}
				b.joint[A][Aval][B][Bval]++
			}
		}
	}
	return b
}

func BayesOptimizing() {
	b := NewBayesOptimized(DatasetOptimized)
	answer := b.Posterior("color", "Rouge", "fruit", "Pomme")   // P(Pomme|Rouge)
	answer2 := b.Posterior("color", "Jaune", "fruit", "Banane") // P(Banane|Jaune)
	answer3 := b.Posterior("color", "Vert", "fruit", "Banane")  // P(Banane|Vert)
	answer4 := b.Posterior("fruit", "Banane", "color", "Vert")  // P(Vert|Banane) — attention, différent de answer3 !

	fmt.Println(answer)
	fmt.Println(answer2)
	fmt.Println(answer3)
	fmt.Println(answer4)

}

func (b *BayesOptimized) Prior(field, value string) float64 {
	return b.counts[field][value] / b.total
}

func (b *BayesOptimized) Likelihood(field, value, BField, BValue string) float64 {
	classCount := b.counts[BField][BValue]
	if classCount == 0 {
		return 0
	}
	return b.joint[BField][BValue][field][value] / classCount
}

func (b *BayesOptimized) Posterior(field, value, BField, BVal string) float64 {
	evidence := b.Prior(field, value)
	if evidence == 0 {
		return 0
	}
	return b.Likelihood(field, value, BField, BVal) * b.Prior(BField, BVal) / evidence
}
