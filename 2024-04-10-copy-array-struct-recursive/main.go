package main

import "fmt"

/*
	- Làm sao để copy GIÁ TRỊ của một mảng trong đó mỗi phần tử lại chứa 1 mảng struct khác?
	- Có một số cách có thể nghĩ đến:
		+ Sử dụng phép gán (a := b) -> cách này thì a và b sẽ trỏ cùng địa chỉ -> thay đổi a thì b cũng bị thay đổi
		+ Sử dụng hàm copy trong Golang -> cách này sẽ chỉ phần tử đầu tiên được copy, còn mảng chứa struct sẽ vẫn
		   trỏ đến cùng để chỉ
    - Sau đây sẽ là code cho 2 cách trên và giải pháp cho bài toán này
*/

func main() {
	// Sử dụng phép gán
	aTree := []Tree{
		{
			ID:   "1",
			Name: "a",
			Children: []Tree{
				{
					ID:   "1",
					Name: "a",
					Children: []Tree{
						{
							ID:   "1",
							Name: "a",
						},
					},
				},
			},
		},
	}

	bTree := aTree
	bTree[0].Name = "b"
	bTree[0].Children[0].Name = "b"
	bTree[0].Children[0].Children[0].Name = "b"

	fmt.Println("aTree", aTree, "bTree", bTree) // khi bTree thay đổi thì aTree cũng bị thay đổi theo

	// Sử dụng hàm copy
	cTree := []Tree{
		{
			ID:   "1",
			Name: "c",
			Children: []Tree{
				{
					ID:   "1",
					Name: "c",
					Children: []Tree{
						{
							ID:   "1",
							Name: "c",
						},
					},
				},
			},
		},
	}

	dTree := make([]Tree, len(cTree))
	copy(dTree, cTree)

	dTree[0].Name = "d"
	dTree[0].Children[0].Name = "d"
	dTree[0].Children[0].Children[0].Name = "d"

	fmt.Println("cTree", cTree, "dTree", dTree) // chỉ có Name của phần tử đầu tiên của cTree không bị thay đổi theo dTree

	// Giải pháp
	eTree := []Tree{
		{
			ID:   "1",
			Name: "e",
			Children: []Tree{
				{
					ID:   "1",
					Name: "e",
					Children: []Tree{
						{
							ID:   "1",
							Name: "e",
						},
					},
				},
			},
		},
	}

	fTree := copyTrees(eTree)

	fTree[0].Name = "f"
	fTree[0].Children[0].Name = "f"
	fTree[0].Children[0].Children[0].Name = "f"

	fmt.Println("eTree", eTree, "fTree", fTree)
	/*
		- fTree thay đổi không làm ảnh hưởng đến eTree
		- Giải pháp là return ra một struct mới, lặp lại với children của nó -> kết hợp thêm đệ quy
	*/
}

type Tree struct {
	ID       string
	Name     string
	Children []Tree
}

func copyTrees(trees []Tree) []Tree {
	result := make([]Tree, 0)

	for _, t := range trees {
		result = append(result, copyTree(t))
	}

	return result
}

func copyTree(tree Tree) Tree {
	children := make([]Tree, 0)

	for _, t := range tree.Children {
		children = append(children, copyTree(t))
	}

	return Tree{
		ID:       tree.ID,
		Name:     tree.Name,
		Children: children,
	}
}
