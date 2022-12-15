<?php

require_once ROOT . '/boot.php';

// вспомогательная функция для heapsort,
// сортирует маленькие кучки из трех элементов
// и рекурсивно кучку справа, если наибольший
// элемент был в правом узле и слева, если
// соответствующий элемент был слева
function heapify(int $currentIndex, array &$arr, int $arrSize)
{
	// индекс левого узла при представлении кучи в виде бинарного дерева
	$indexLeft = 2 * $currentIndex + 1;
	// с помощью следующих двух ветвей if отыскивается индекс наибольшего
	// элемента среди данного элемента, узла справа и слева.
	$theBiggestNumIndex = $currentIndex;
	if ($indexLeft < $arrSize && $arr[$theBiggestNumIndex] < $arr[$indexLeft])
	{
		$theBiggestNumIndex = $indexLeft;
	}

	// индекс правого узла при представлении кучи в виде бинарного дерева
	$indexRight = $indexLeft + 1;
	if ($indexRight < $arrSize && $arr[$theBiggestNumIndex] < $arr[$indexRight])
	{
		$theBiggestNumIndex = $indexRight;
	}

	// если наибольший элемент не сверху, то он поднимается наверх
	if ($theBiggestNumIndex !== $currentIndex)
	{
		swap($arr[$theBiggestNumIndex],$arr[$currentIndex]);
		heapify($theBiggestNumIndex, $arr, $arrSize);
	}
}

// Пирамидальная сортировка
function heapSort(&$arr, $arrSize)
{
	// функция построения кучек вызывается для всех элементов
	// снизу вверх, кроме листьев (при представлении кучи бинарным деревом),
	// так как листья уже находятся в правильном месте относительно
	// null-потомков. Строится куча.
	for ($i = $arrSize / 2 - 1; $i >= 0; $i--)
	{
		heapify($i, $arr, $arrSize);
	}
	// Наверху кучи - самый большой элемент.
	// Данный цикл занимается тем, что меняет местами корень
	// и элемент с последним индексом и восстанавливает кучу
	// уже без элемента с последним индексом. Таким в конце начинает
	// образовываться отсортированный массив.
	for ($i = $arrSize - 1; $i >= 0; $i--)
	{
		swap($arr[$i], $arr[0]);
		heapify(0, $arr, $i);
	}
}