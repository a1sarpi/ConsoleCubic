package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"math"
	"os"
	"strconv"
)

func cubicRoots(a, b, c, d float64) ([]float64, int) {
	if a == 0 {
		return quadraticRoots(b, c, d)
	}

	p := (3*a*c - b*b) / (3 * a * a)
	q := (27*a*a*d - 9*a*b*c + 2*b*b*b) / (27 * a * a * a)

	r := q / 2
	D := (r * r) + (p * p * p / 27)

	var roots []float64

	if D > 0 {
		sqrtD := math.Sqrt(D)
		u := math.Cbrt(-r + sqrtD)
		v := math.Cbrt(-r - sqrtD)
		x1 := u + v

		roots = append(roots, x1-(b/(3*a)))
		return roots, 1 // Один действительный корень
	} else if D == 0 {
		x1 := math.Cbrt(-r)
		roots = append(roots, x1-(b/(3*a)))

		x2 := -x1/2 - (b / (3 * a))
		roots = append(roots, x2)

		return roots, 2 // Два действительных корня
	} else {
		radius := math.Sqrt(-p / 3)
		theta := math.Acos(-r / (radius * radius * radius))

		for k := 0; k < 3; k++ {
			xk := 2 * radius * math.Cos((theta + 2*math.Pi*float64(k)/3))

			roots = append(roots, xk-(b/(3*a)))
		}

		return roots, 3 // Три действительных корня
	}
}

func quadraticRoots(a, b, c float64) ([]float64, int) {
	if a == 0 {
		return linearRoots(b, c) // Линейное уравнение
	}

	D := b*b - 4*a*c
	var roots []float64

	if D > 0 {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)

		roots = append(roots, x1, x2)

		return roots, 2 // Два действительных корня
	} else if D == 0 {
		x := -b / (2 * a)

		roots = append(roots, x)

		return roots, 1 // Один действительный корень
	} else {
		return nil, 0 // Нет действительных корней
	}
}

func linearRoots(b, c float64) ([]float64, int) {
	if b == 0 {
		if c == 0 {
			return nil, -1 // Бесконечное количество корней
		}

		return nil, 0 // Нет корней
	}

	x := -c / b

	return []float64{x}, 1 // Один действительный корень
}

func main() {
	var aFlag string
	var bFlag string
	var cFlag string
	var dFlag string

	var rootsCmd = &cobra.Command{
		Use:   "cubic-roots",
		Short: "Вычисление корней кубического уравнения ax^3 + bx^2 + cx + d = 0",
		Long:  "Эта команда принимает коэффициенты a, b, c и d кубического уравнения и вычисляет его действительные корни. Коэффициенты передаются через флаги.",
		Run: func(cmd *cobra.Command, args []string) {
			a, errA := strconv.ParseFloat(aFlag, 64)
			b, errB := strconv.ParseFloat(bFlag, 64)
			c, errC := strconv.ParseFloat(cFlag, 64)
			d, errD := strconv.ParseFloat(dFlag, 64)

			if errA != nil || errB != nil || errC != nil || errD != nil {
				fmt.Println("Ошибка: все коэффициенты должны быть числами.")
				return
			}

			roots, count := cubicRoots(a, b, c, d)

			switch count {
			case -1:
				fmt.Println("Бесконечное количество корней")
			case 0:
				fmt.Println("Нет действительных корней")
			case 1:
				fmt.Printf("Найден один действительный корень: %v\n", fmt.Sprintf("%.2f", roots[0]))
			case 2:
				fmt.Printf("Найдены два действительных корня: %v и %v\n", fmt.Sprintf("%.2f", roots[0]),
					fmt.Sprintf("%.2f", roots[1]))
			case 3:
				fmt.Printf("Найдены три действительных корня: %v, %v и %v\n", fmt.Sprintf("%.2f", roots[0]),
					fmt.Sprintf("%.2f", roots[1]),
					fmt.Sprintf("%.2f", roots[2]))
			default:
				fmt.Println("Неизвестная ошибка")
			}
		},
	}

	// Объявление флагов как строки
	rootsCmd.Flags().StringVarP(&aFlag, "a", "a", "0", "Коэффициент a")
	rootsCmd.Flags().StringVarP(&bFlag, "b", "b", "0", "Коэффициент b")
	rootsCmd.Flags().StringVarP(&cFlag, "c", "c", "0", "Коэффициент c")
	rootsCmd.Flags().StringVarP(&dFlag, "d", "d", "0", "Коэффициент d")

	// Обязательные флаги
	rootsCmd.MarkFlagRequired("a")
	rootsCmd.MarkFlagRequired("b")
	rootsCmd.MarkFlagRequired("c")
	rootsCmd.MarkFlagRequired("d")

	if err := rootsCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
