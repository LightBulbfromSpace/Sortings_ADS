<?php

require_once ROOT . '/boot.php';

// mergeSort - сортировка слиянием
// $type = false - mergeSortRecursive, $type = true - mergeSortIterative
function mergeSort(array &$arr, bool $type = true) : void
{
	if ($type)
	{
		mergeSortIterative($arr);
	}
	else
	{
		mergeSortRecursive($arr, 0, count($arr));
	}
}


// Массив рекурсивно делится пополам,
// отсортированные половины сливаются функцией merge,
// создающей отсортированный массив из двух отсортированных массивов.
// $left - первый индекс массива,
// $right - кол-во элементов в массиве (последний индекс + 1).
function mergeSortRecursive(array &$arr, int $left, int $right)
{
	if ($right - $left <= 1)
	{
		return;
	}
	$mid = ceil(($left + $right) / 2);
	mergeSortRecursive($arr, $left, $mid);
	mergeSortRecursive($arr, $mid, $right);
	merge($arr, $left, $mid, $right);
}

// функция merge сливает два отсортированных массива в один
// (оба массива находятся в $arr, $mid указвает на границу разделения).
// Первый цикл выполняется, пока элементы есть в обоих массивах.
// Когда элемент устанавливливается в результирующий массив,
// соответствующая итерационная переменная ($it1, $it2) увеличивается на один.
// Второй и третий цикл устанавливают в результирующий массив значения из
// массива, который не был пройден до конца после выполнения первого цикла.
function merge(array &$arr, int $left, int $mid, int $right)
{
	$result = [];
	$it1 = 0;
	$it2 = 0;
	while ($left + $it1 < $mid && $mid + $it2 < $right)
	{
		if ($arr[$left+$it1] <= $arr[$mid+$it2])
		{
			$result[$it1+$it2] = $arr[$left+$it1];
			$it1++;
		}
		else
		{
			$result[$it1+$it2] = $arr[$mid+$it2];
			$it2++;
		}
	}

	while ($left + $it1 < $mid)
	{
		$result[$it1+$it2] = $arr[$left+$it1];
		$it1++;
	}
	while ($mid + $it2 < $right)
	{
		$result[$it1+$it2] = $arr[$mid+$it2];
		$it2++;
	}
	for ($i = 0; $i < $it1 + $it2; $i++)
	{
		$arr[$left+$i] = $result[$i];
	}

}

function mergeSortIterative(array &$arr) : array
{
	$mainLen = count($arr);
	for ($subArrLen = 1; $subArrLen <= $mainLen; $subArrLen *= 2)
	{
		$it = 0;
		for (; $it + $subArrLen <= $mainLen; $it += $subArrLen)
		{
			merge($arr, $it, ceil((2*$it + $subArrLen) / 2), $it+$subArrLen);
		}
		if ($it < $mainLen)
		{
			merge($arr, $it - $subArrLen, $it, $mainLen);
		}
	}
	return $arr;
}