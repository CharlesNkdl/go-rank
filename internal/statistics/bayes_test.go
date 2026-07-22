package statistics

// généré pour test
import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/* observations

go test -bench=. -benchmem ./...

Donc j'ai généré ce petit test pour voir la différence de performance entre les deux implémentations
mon hypothèse est simple :
le simple va être plus rapide sur des tout petit dataset, parce qu'il
n'aura pas la construction, le traitement complet et l'allocation de mémoire nécessaire pour stocker tout les
champs dans la structure comme dans l'optimisé

par contre, vu que ça va faire des calculs et des lectures ++ à chaque calcul, à partir d'une certaine tailles de dataset
cela sera beaucoup plus rentable

quelques stats :

Dataset: 50000 lignes, 10000 appels à Posterior
Naïf      : 11.542726334s
Optimisé  : 483.709µs
Speedup   : x23863.0

Dataset: 50000 lignes, 10 appels à Posterior
Naïf      : 21.562167ms
Optimisé  : 792ns
Speedup   : x27225.0

Dataset: 50 lignes, 10 appels à Posterior
Naïf      : 19.458µs
Optimisé  : 2.208µs
Speedup   : x8.8

Dataset: 10 lignes, 10 appels à Posterior
Naïf      : 4µs
Optimisé  : 1.625µs
Speedup   : x2.5

Dataset: 3 lignes, 3 appels à Posterior
Naïf      : 1.208µs
Optimisé  : 417ns
Speedup   : x2.9

Au final, même pas
L'optimisé sera toujours plus rapide même si un calcul plus couteux au début, il sera plus court que la recherche naïve


*/

func timeIt(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}

func generateLargeDataset(n int) []Classes {
	fruits := []string{"Pomme", "Poire", "Banane"}
	colors := []string{"Rouge", "Vert", "Jaune"}

	r := rand.New(rand.NewSource(42)) // seed fixe = résultats reproductibles
	dataset := make([]Classes, n)
	for i := 0; i < n; i++ {
		dataset[i] = Classes{
			"fruit": fruits[r.Intn(len(fruits))],
			"color": colors[r.Intn(len(colors))],
		}
	}
	return dataset
}

func generateLargeDataset2(n int) []Cases {
	fruits := []string{"Pomme", "Poire", "Banane"}
	colors := []string{"Rouge", "Vert", "Jaune"}

	r := rand.New(rand.NewSource(42)) // seed fixe = résultats reproductibles
	dataset := make([]Cases, n)
	for i := 0; i < n; i++ {
		dataset[i] = Cases{
			fruits[r.Intn(len(fruits))],
			colors[r.Intn(len(colors))],
		}
	}
	return dataset
}

const benchDatasetSize = 3

func BenchmarkNaiveConstruction(b *testing.B) {
	dataset := generateLargeDataset2(benchDatasetSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewBayes(dataset)
	}
}

func BenchmarkNaivePosterior(b *testing.B) {
	dataset := generateLargeDataset2(benchDatasetSize)
	nb := NewBayes(dataset)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nb.Posterior("Rouge", "Pomme")
	}
}

func BenchmarkOptimizedPosterior(b *testing.B) {
	dataset := generateLargeDataset(benchDatasetSize)
	ob := NewBayesOptimized(dataset)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ob.Posterior("color", "Rouge", "fruit", "Pomme")
	}
}

func BenchmarkOptimizedConstruction(b *testing.B) {
	dataset := generateLargeDataset(benchDatasetSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewBayesOptimized(dataset)
	}
}
func TestPerfComparison(t *testing.T) {
	dataset := generateLargeDataset(benchDatasetSize)
	dataset2 := generateLargeDataset2(benchDatasetSize)
	const calls = 3 // nombre d'appels à Posterior simulés

	nb := NewBayes(dataset2)
	naiveResult := timeIt(func() {
		for i := 0; i < calls; i++ {
			nb.Posterior("Rouge", "Pomme")
		}
	})

	ob := NewBayesOptimized(dataset)
	optimizedResult := timeIt(func() {
		for i := 0; i < calls; i++ {
			ob.Posterior("color", "Rouge", "fruit", "Pomme")
		}
	})

	fmt.Printf("Dataset: %d lignes, %d appels à Posterior\n", benchDatasetSize, calls)
	fmt.Printf("Naïf      : %v\n", naiveResult)
	fmt.Printf("Optimisé  : %v\n", optimizedResult)
	fmt.Printf("Speedup   : x%.1f\n", float64(naiveResult)/float64(optimizedResult))
}
