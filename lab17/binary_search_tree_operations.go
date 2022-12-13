package BST

import (
	"labs/structures"
)

func Search(node *structures.TreeNode, elem int) (*structures.TreeNode, bool) {
	// два базовых случая: рекурсия пришла к узлу,
	// который является указателем на ноль => нужного значения в дереве нет;
	// значение узла соответствует искомому значению => элемент найден.
	if node.Val == nil {
		return nil, false
	}
	if *node.Val == elem {
		return node, true
	}

	// используется свойство БДП: левее находятся узлы с элементами меньше данного,
	// остальные - правее.
	if elem <= *node.Val {
		if node.Left != nil {
			return Search(node.Left, elem)
		}
	} else {
		if node.Right != nil {
			return Search(node.Right, elem)
		}
	}
	return nil, false
}

func Add(node *structures.TreeNode, elem int) {
	// базовый случай: значение в узле является указателем на ноль.
	// проверяется именно значение, а не сам узел, так как узел был
	// проинициализированн перед вызовом рекурсии, если не существовал.
	// инициалиализация нового узла перед вызовом рекурсии происходит для
	// связывания нового узла с остальным девевом.
	if node.Val == nil {
		node.Val = &elem
		return
	}
	// используется свойство БДП: левее находятся узлы с элементами меньше данного,
	// остальные - правее.
	if elem <= *node.Val {
		if node.Left == nil {
			node.Left = new(structures.TreeNode)
		}
		Add(node.Left, elem)
	} else {
		if node.Right == nil {
			node.Right = new(structures.TreeNode)
		}
		Add(node.Right, elem)
	}
}

func Delete(root *structures.TreeNode, elem int) {
	node, found := Search(root, elem)
	// если элемента нет в дереве, его удаление не нужно
	if !found {
		return
	}
	// удаление листа
	// листы не удаляются до конца,
	// так как дерево в данной реализации
	// не содержит указатель на родителя.
	if node.Right == nil && node.Left == nil {
		node.Val = nil
		return
	}
	// удаление узла с одним потомком
	if node.Right.Val == nil {
		node.Val = node.Left.Val
		// адреса сохраняются в переменные, так как иначе они будут потеряны при переносе
		newRight := node.Left.Right
		newLeft := node.Left.Left
		node.Right = newRight
		node.Left = newLeft
		return
	}
	if node.Left.Val == nil {
		node.Val = node.Right.Val
		// адреса сохраняются в переменные, так как иначе они будут потеряны при переносе
		newRight := node.Right.Right
		newLeft := node.Right.Left
		node.Right = newRight
		node.Left = newLeft
		return
	}
	// удаление узла с двумя потомками
	// идем на один узел вправо и ищем самый левый
	newNode := node.Right
	for newNode.Left != nil {
		newNode = newNode.Left
	}
	// переносим значение найденного узла в удаляемый узел
	node.Val = newNode.Val
	// удаляем значение в найденном узле
	newNode.Val = nil
	// если у найденного узла есть правый потомок, делаем его перенос
	if newNode.Right != nil {
		newNode.Val = node.Right.Val
		newNode.Right = node.Right.Right
	}
}
