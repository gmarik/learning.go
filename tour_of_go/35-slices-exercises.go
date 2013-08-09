// Exercise: Slices
// 
// Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.
// 
// The choice of image is up to you. Interesting functions include x^y, (x+y)/2, and x*y.
// 
// (You need to use a loop to allocate each []uint8 inside the [][]uint8.)
// 
// (Use uint8(intValue) to convert between types.)

// MAKING THIS WORK
// $ export GOPATH=~/golang
// $ mkdir $GOPATH
// $ brew install mercurial # go get depends on this
// $ go get "code.google.com/p/go-tour/pic"
// $ go run 35-slices-exercises.go

// put into a brower url bar: data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQAAAAEACAIAAADTED8xAAAFK0lEQVR4nOzcYWpbZxCFYTnkR35mCXcpWlpWHlfGRW1adMlwI/Rx32do35wchuHDYI7NZPT1cnm/KFWtL7f/394uiE2+fSbAl49vhMvPnxeaTmkJgGnefge4vL9//Herrx9/+6f4/NP7EgDT/DsBPn8eoumalgCY5i+/A/y7vn37r3Mv/fpP0y8BMM1ffgf4/Nb5/89JfP5ZfQmAadoD8NO+BMA07QHotJYAmObDPcCj+v79dzvvZb75y86XAJjmwz3AI61f/5n6JQCmaQ/AT/sSANO0B6DTWgJgmu4B9Kf7JQCm6R6An/YlAKZpD8BP+xIA07QHoNNaAmCa43uAaW3bsybfy/t3yvv3SwJgmuN7gKk23/yV50sATNMegJ/2JQCmaQ9Ap7UEwDTdA+hP90sATNM9AD/tSwBM0x6An/YlAKZpD0CntQTANMf3AKt9vrv55h+ZLwEwzfE9gH79Z+qXAJimPQA/7UsATNMegE5rCYBpugfQn+6XAJimewB+2pcAmKY9AD/tSwBM0x6ATmsJgGmO7wFWq+v11S84XL7+LywJgGmO7wFW097v/Ue0BMA07QH4aV8CYJr2AHRaSwBM0z2A/nS/BMA03QPw074EwDTtAfhpXwJgmvYAdFpLAExzfA+w2ue7m2/+kfkSANMc3wPo13+mfgmAadoD8NO+BMA07QHotJYAmKZ7AP3pfgmAaboH4Kd9CYBp2gPw074EwDTtAei0lgCY5vgeYFrb9qzJ9/L+nfL+/ZIAmOb4HmCqzTd/5fkSANO0B+CnfQmAadoD0GktATBN9wD60/0SANN0D8BP+xIA07QH4Kd9CYBp2gPQaS0BMM3xPcBqn+9uvvlH5ksATHN8D6Bf/5n6JQCmaQ/AT/sSANO0B6DTWgJgmu4B9Kf7JQCm6R6An/YlAKZpD8BP+xIA07QHoNNaAmCatz9/+59aqyfUjx+vfkG7JACmOb4HoP+s9vV/rZYAmKY9AD/tSwBM0x6ATmsJgGm6B9Cf7pcAmKZ7AH7alwCYpj0AP+1LAEzTHoBOawmAaT7cAzyq1T7f3Xzzj8yXAJjm+B5Av/4z9UsATNMegJ/2JQCmaQ9Ap7UEwDTdA+hP90sATNM9AD/tSwBM0x6An/YlAKZpD0CntQTANMf3ANPatmdNvpf375T375cEwDTH9wBTbb75K8+XAJimPQA/7UsATNMegE5rCYBpugfQn+6XAJimewB+2pcAmKY9AD/tSwBM0x6ATmsJgGmO7wFW+3x3880/Ml8CYJrjewD9+s/ULwEwTXsAftqXAJimPQCd1hIA03QPoD/dLwEwTfcA/LQvATBNewB+2pcAmKY9AJ3WEgDTHN8DrFbX66tfcLh8/V9YEgDTHN8DrKa93/uPaAmAadoD8NO+BMA07QHotJYAmKZ7AP3pfgmAaboH4Kd9CYBp2gPw074EwDTtAei0lgCY5vgeYLXPdzff/CPzJQCmOb4H0K//TP0SANO0B+CnfQmAadoD0GktATBN9wD60/0SANN0D8BP+xIA07QH4Kd9CYBp2gPQaS0BMM3xPcC0tu1Zk+/l/Tvl/fslATDN8T3AVJtv/srzJQCmaQ/AT/sSANO0B6DTWgJgmu4B9Kf7JQCm6R6An/YlAKZpD8BP+xIA07QHoNNaAmCa43uA1T7f3Xzzj8yXAJjm+B5Av/4z9UsATNMegJ/2JQCmaQ9Ap7UEwDTdA+hP90sATNM9AD/tSwBM0x6An/YlAKZpD0CntQTANP8KAAD//9NVQsA0jnjyAAAAAElFTkSuQmCC

package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
 var xs = make([]([]uint8), dy)
    for y, _ := range xs {
        xs[y] = make([]uint8, dx)
        for x, _ := range xs[y] {
            xs[y][x] = uint8((x & y) & (x & y))
        }
    }
    
  return xs  

}

func main() {
    pic.Show(Pic)
}
