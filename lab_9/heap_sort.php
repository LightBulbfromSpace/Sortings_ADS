<?php

function heapify(int $currentIndex, array &$arr, int $arrSize)
{
	$index = 2 * $currentIndex + 1;
	$theBiggestNumIndex = $currentIndex;
	if ($index < $arrSize && $arr[$theBiggestNumIndex] < $arr[$index])
	{
		$theBiggestNumIndex = $index;
	}

	$index++;
	if ($index < $arrSize && $arr[$theBiggestNumIndex] < $arr[$index])
	{
		$theBiggestNumIndex = $index;
	}

	if ($theBiggestNumIndex !== $currentIndex)
	{
		swap($arr[$theBiggestNumIndex],$arr[$currentIndex]);
		heapify($theBiggestNumIndex, $arr, $arrSize);
	}
}


function heapSort(&$arr, $arrSize)
{
	for ($i = $arrSize / 2 - 1; $i >= 0; $i--)
	{
		heapify($i, $arr, $arrSize);
	}

	for ($i = $arrSize - 1; $i >= 0; $i--)
	{
		swap($arr[$i], $arr[0]);
		heapify(0, $arr, $i);
	}
}

function swap(&$a, &$b): void
{
	if ($a === $b) {
		return;
	}

	$tmp = $a;
	$a = $b;
	$b = $tmp;
}