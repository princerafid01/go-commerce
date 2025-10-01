package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productID int) *Product {

	for _, product := range productList {
		if product.ID == productID {
			// utils.SendData(w, product, 200)
			return &product
		}
	}

	return nil
}

func Update(product Product) {
	for index, p := range productList {
		if p.ID == product.ID {
			productList[index] = product
		}
	}

}

func Delete(productID int) {
	var tempList []Product

	for _, p := range productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}

	productList = tempList
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is good for health",
		Price:       100,
		ImgUrl:      "https://plus.unsplash.com/premium_photo-1670512181061-e24282f7ee78?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is good for health",
		Price:       200,
		ImgUrl:      "https://plus.unsplash.com/premium_photo-1724249990837-f6dfcb7f3eaa?q=80&w=687&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is good for health",
		Price:       300,
		ImgUrl:      "https://plus.unsplash.com/premium_photo-1667926862695-629f15968976?q=80&w=687&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Grapes",
		Description: "Grapes is good for health",
		Price:       400,
		ImgUrl:      "https://plus.unsplash.com/premium_photo-1666270423730-9af384b9cb0f?q=80&w=687&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Pineapple",
		Description: "Pineapple is good for health",
		Price:       500,
		ImgUrl:      "https://plus.unsplash.com/premium_photo-1724255994628-dceb76a829e8?q=80&w=687&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
	}
	prd6 := Product{
		ID:          6,
		Title:       "Mango",
		Description: "Mango is good for health",
		Price:       600,
		ImgUrl:      "https://plus.unsplash.com/premium_photo-1724255863045-2ad716767715?q=80&w=687&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
	}
	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd1)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
	productList = append(productList, prd6)
}
