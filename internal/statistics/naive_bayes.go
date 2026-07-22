package statistics

// p(A | B) = Probabilité que A est vrai en sachant que B est vrai
// Theoreme bayes
//p(A|B) = ( ( p(B|A) * p(A) ) / p(B)
// Soit
// Posterior = Prior * Likelihood / evidence
// B = ensemble de classe = { Classe1, Classe2, Classe3 }
// A = Autre ensemble = { Ensemble1, Ensemble2, Ensemble3 }
// Pour être plus concret
// A = { Pomme, Poire, Banane }
// B = { Rouge, Jaune, Vert }
// Donc
// La probabilité sachant que c'est Rouge, quelle est la proba que ce soit une Pomme
// P(A1 | B1)
// est égale à
// La probabilité qu'une Pomme soit Rouge * La probabilité que ça soit une Pomme
// divisé par la probabilité que ca soit rouge
// soit
// P(B1 | A1) * P(A1) / P(B1)
// P(B1|A1) est ici la vraisemblance (likelihood), P(A1) le prior,
// P(B1) l'evidence, et le résultat P(A1|B1) est le posterior

// dataset généré :

// 20 observations au total, réparties comme suit :
//
//              Rouge   Vert   Jaune   | Total
// Pomme          7       2      1    |  10
// Poire          1       4      1    |   6
// Banane         0       0      4    |   4
// -----------------------------------------
// Total          8       6      6    |  20

import (
	"fmt"
)

type Cases struct {
	Fruit string
	Color string
}

var Dataset = []Cases{
	// Pommes (10) : 7 rouges, 2 vertes, 1 jaune
	{"Pomme", "Rouge"}, {"Pomme", "Rouge"}, {"Pomme", "Rouge"}, {"Pomme", "Rouge"},
	{"Pomme", "Rouge"}, {"Pomme", "Rouge"}, {"Pomme", "Rouge"},
	{"Pomme", "Vert"}, {"Pomme", "Vert"},
	{"Pomme", "Jaune"},

	// Poires (6) : 1 rouge, 4 vertes, 1 jaune
	{"Poire", "Rouge"},
	{"Poire", "Vert"}, {"Poire", "Vert"}, {"Poire", "Vert"}, {"Poire", "Vert"},
	{"Poire", "Jaune"},

	// Bananes (4) : 4 jaunes
	{"Banane", "Jaune"}, {"Banane", "Jaune"}, {"Banane", "Jaune"}, {"Banane", "Jaune"},
}

type Bayes struct {
	Dataset    []Cases
	TotalCases float32
}

func NewBayes(Dataset []Cases) *Bayes {
	return &Bayes{
		Dataset:    Dataset,
		TotalCases: TotalCases(Dataset),
	}
}

func BayesImplementation() {
	B := NewBayes(Dataset)
	answer := B.Posterior("Pomme", "Rouge")
	answer2 := B.Posterior("Banane", "Jaune")
	answer3 := B.Posterior("Banane", "Vert")
	answer4 := B.Posterior("Vert", "Banane")

	fmt.Println(answer)
	fmt.Println(answer2)
	fmt.Println(answer3)
	fmt.Println(answer4)
}

func (b *Bayes) Posterior(A, B string) float32 {
	return b.Likelihood(B, A) * b.Prior(A) / b.Prior(B)
}

func (b *Bayes) Likelihood(A, B string) float32 {
	return b.BothOccurrences(A, B) / b.Occurrences(B)
}

func (b *Bayes) Prior(feature string) float32 {
	return b.Occurrences(feature) / b.TotalCases
}

func (b *Bayes) Occurrences(feature string) float32 {
	var occ float32
	for i := range b.Dataset {
		if feature == b.Dataset[i].Color || feature == b.Dataset[i].Fruit {
			occ += 1
		}
	}
	return occ
}

func (b *Bayes) BothOccurrences(A, B string) float32 {
	var occ float32
	for i := range b.Dataset {
		if (b.Dataset[i].Fruit == A && b.Dataset[i].Color == B) ||
			(b.Dataset[i].Fruit == B && b.Dataset[i].Color == A) {
			occ++
		}
	}
	return occ
}

func TotalCases(dataset []Cases) float32 {
	return float32(len(dataset))
}
