//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/join-operations#query-expression-syntax-examples

type Product struct {
	Name       string
	CategoryId int
}

type Category struct {
	Id           int
	CategoryName string
}

func main() {
	products := go2linq.NewEnSlice(
		Product{Name: "Cola", CategoryId: 0},
		Product{Name: "Tea", CategoryId: 0},
		Product{Name: "Apple", CategoryId: 1},
		Product{Name: "Kiwi", CategoryId: 1},
		Product{Name: "Carrot", CategoryId: 2},
	)
	categories := go2linq.NewEnSlice(
		Category{Id: 0, CategoryName: "Beverage"},
		Category{Id: 1, CategoryName: "Fruit"},
		Category{Id: 2, CategoryName: "Vegetable"},
	)

	// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/join-operations#join
	// Join products and categories based on CategoryId
	query := go2linq.JoinMust(products, categories,
		func(product Product) int { return product.CategoryId },
		func(category Category) int { return category.Id },
		func(product Product, category Category) string {
			return fmt.Sprintf("%s - %s", product.Name, category.CategoryName)
		},
	)
	enrJoin := query.GetEnumerator()
	for enrJoin.MoveNext() {
		item := enrJoin.Current()
		fmt.Println(item)
	}

	fmt.Println()

	// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/join-operations#groupjoin
	// Join categories and product based on CategoryId and grouping result
	productGroups := go2linq.GroupJoinMust(categories, products,
		func(category Category) int { return category.Id },
		func(product Product) int { return product.CategoryId },
		func(category Category, products go2linq.Enumerable[Product]) go2linq.Enumerable[Product] {
			return products
		},
	)
	enrGroupJoin := productGroups.GetEnumerator()
	for enrGroupJoin.MoveNext() {
		fmt.Println("Group")
		productGroup := enrGroupJoin.Current()
		enrProductGroup := productGroup.GetEnumerator()
		for enrProductGroup.MoveNext() {
			product := enrProductGroup.Current()
			fmt.Printf("%8s\n", product.Name)
		}
	}
}
