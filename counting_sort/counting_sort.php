<?php

// CountingSort good for arrays with small range of values.
function CountingSort(&$arr) : array
{
	$countArr = [];
	$result = [];
	$maxValue = 0;
	foreach ($arr as $value)
	{
		if (!isset($countArr[$value]))
		{
			$countArr[$value] = 0;
		}
		$countArr[$value]++;
		$maxValue = max($maxValue, $value);
	}

	if (!isset($countArr[0]))
	{
		$countArr[0] = 0;
	}

	for ($i = 1; $i <= $maxValue; $i++)
	{
		if (!isset($countArr[$i]))
		{
			$countArr[$i] = $countArr[$i-1];
			continue;
		}
		$countArr[$i] += $countArr[$i-1];
	}
	$mainLen = count($arr);
	for ($i = $mainLen - 1; $i >= 0; $i--)
	{
		$j = $arr[$i];
		$countArr[$j]--;
		$result[$countArr[$j]] = $arr[$i];
	}
	ksort($result);
	return $result;
}