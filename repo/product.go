package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(product Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

// Constructor function
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productID {
			// utils.SendData(w, product, 200)
			return product, nil
		}
	}

	return nil, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Delete(productID int) error {
	var tempList []*Product

	for _, p := range r.productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}

	r.productList = tempList
	return nil
}

func (r *productRepo) Update(product Product) (*Product, error) {
	for index, p := range r.productList {
		if p.ID == product.ID {
			r.productList[index] = &product
		}
	}

	return &product, nil
}

func generateInitialProducts(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is good for health",
		Price:       100,
		ImgUrl:      "https://plus.unsplash.com/premium_photo-1670512181061-e24282f7ee78?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
	}
	prd2 := &Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is good for health",
		Price:       200,
		ImgUrl:      "https://plus.unsplash.com/premium_photo-1724249990837-f6dfcb7f3eaa?q=80&w=687&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
	}
	prd3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is good for health",
		Price:       300,
		ImgUrl:      "https://plus.unsplash.com/premium_photo-1667926862695-629f15968976?q=80&w=687&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
	}
	prd4 := &Product{
		ID:          4,
		Title:       "Grapes",
		Description: "Grapes is good for health",
		Price:       400,
		ImgUrl:      "https://plus.unsplash.com/premium_photo-1666270423730-9af384b9cb0f?q=80&w=687&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
	}
	prd5 := &Product{
		ID:          5,
		Title:       "Pineapple",
		Description: "Pineapple is good for health",
		Price:       500,
		ImgUrl:      "https://plus.unsplash.com/premium_photo-1724255994628-dceb76a829e8?q=80&w=687&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
	}
	prd6 := &Product{
		ID:          6,
		Title:       "Mango",
		Description: "Mango is good for health",
		Price:       600,
		ImgUrl:      "https://plus.unsplash.com/premium_photo-1724255863045-2ad716767715?q=80&w=687&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
	}
	r.productList = append(r.productList, prd1)
	r.productList = append(r.productList, prd2)
	r.productList = append(r.productList, prd1)
	r.productList = append(r.productList, prd3)
	r.productList = append(r.productList, prd4)
	r.productList = append(r.productList, prd5)
	r.productList = append(r.productList, prd6)
}
