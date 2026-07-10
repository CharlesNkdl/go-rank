package loop

import "fmt"

/**
A shepherd has x number of sheep. It takes him y minutes per sheep to cut the wool.
However, after every z sheep, he must take the wool to the farm then head back to the sheep. It takes him w minutes to travel the distance between the sheep and the farm.
The shepherd always start near his sheep.
Don't forget that if he still has sheep to shear, he must return to finish his job.
If the shepherd is still with the sheep when he finishes shearing all x sheep, then he must return to the farm to finish his job.
*/

func Sheep(X, Y, Z, W int) {
	var time int
	for i := 1; i < X+1; i++ {
		time += Y
		if i == X {
			time += W
		} else if i%Z == 0 {
			time += W * 2
		}
	}
	fmt.Println(time)
}
