<?php

function quickSort(array &$arr, int $firstIndex, int $lastIndex) :void
{
	if ($firstIndex < $lastIndex)
	{
		$p = partition($arr, $firstIndex, $lastIndex);
		quickSort($arr, $firstIndex, $p);
		quickSort($arr, $p+1, $lastIndex);
	}
}

function partition(array &$arr, int $firstIndex, int $lastIndex) :int
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