<?php

// quickSort - быстрая сортировка, является рекурсивным алгоритмом.
// Базовый случай рекурсии - массив имееи один эелемент.
// Сначала происходит подсчет точки деления $p и сортировка относительно этой точки
// (подробнее в описании partition).
// Далее quickSort вызывается для полученных частей массива.

// $partitionType = false - Hoar partition, true - Lomuto partition
function quickSort(array &$arr, int $firstIndex, int $lastIndex, bool $partitionType = false) :void
{
	if ($firstIndex < $lastIndex)
	{
		$p = $partitionType ? partitionLomuto($arr, $firstIndex, $lastIndex)
			: partitionHoar($arr, $firstIndex, $lastIndex);
		$p = $partitionType ? $p - 1 : $p;

		quickSort($arr, $firstIndex, $p);
		quickSort($arr, $p+1, $lastIndex);
	}
}

// Существуют разные способы выбора "точки разделения":
// первый, средний, последний элемент массива, среднее арифметическое,
// средний по значению из трех, случайным образом...
// В данной реализации берется средний элемент массива.
// Основная идея зключается в том, чтобы в выбранном массиве (участке массива)
// после выполнения функции слева относительно точки разделения находились
// элементы, которые меньше по значению, справа все остальные.
// Существуют два способа прохода по массиву: разбиение Хоара и Ломута.
// В данной реализации используется разбиение Хоара:
// два индекса, указывающие на начало и коней массива идут навстречу другу.
// Как только находится элемент слева больший $pivot и справа меньший $pivot,
// данные элементы меняются местами. Работа функции заканчивается, как только
// индексы совпадут или значение индекса, идущего слева, будет больше идущего справа.
function partitionHoar(array &$arr, int $firstIndex, int $lastIndex) :int
{
	$pivot = $arr[($firstIndex + $lastIndex)/2];
	$i = $firstIndex;
	$j = $lastIndex;
	while ($i < $j)
	{
		while ($arr[$i] < $pivot)
		{
			$i++;
		}
		while ($arr[$j] > $pivot)
		{
			$j--;
		}
		if ($i >= $j)
		{
			break;
		}
		swap($arr[$i++], $arr[$j--]);
	}
	return $j;
}

function partitionLomuto(array &$arr, int $firstIndex, int $lastIndex) :int
{
	$pivot = $arr[$lastIndex];
	$i = $firstIndex-1;
		for ($j = $firstIndex; $j <= $lastIndex; $j++)
		{
			if ($arr[$j] <= $pivot)
			{
				swap($arr[++$i], $arr[$j]);
			}
		}
		return $i;
}