<?php

function BubbleSort(&$arr) : void
{
	$arrLen = count($arr);
	for ($step = 0; $step < $arrLen - 1; $step++)
	{
		$swapped = false;
		for ($i = 0; $i < $arrLen - $step - 1; $i++)
		{
			if ($arr[$i] > $arr[$i+1])
			{
				$swapped = true;
				swap($arr[$i], $arr[$i+1]);
			}
		}
		if (!$swapped)
		{
			break;
		}
	}
}