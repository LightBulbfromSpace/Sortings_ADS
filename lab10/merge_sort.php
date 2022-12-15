<?php

require_once ROOT . '/boot.php';

function mergeSort(array &$arr, int $left, int $right)
{
	if ($left + 1 >= $right)
	{
		return;
	}
	$mid = ceil(($left + $right) / 2);
	mergeSort($arr, $left, $mid);
	mergeSort($arr, $mid + 1, $right);
	merge($arr, $left, $mid, $right);
}

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