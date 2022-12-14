<?php

require_once 'boot.php';

function main()
{
	$arrs = [];
	$arrs[] = [0, 1, 3, 9, 4, 5, 33, 12, 6, 55, 77, 88, 66, 100, 99];
	$arrs[] = [200, 33, 12, 6, 55, 77, 88, 66, 100, 99, 10, 9, 2, 1];
	foreach ($arrs as $arr)
	{
		heapSort($arr, count($arr));
		//mergeSort($arr, 0, count($arr));
		//quickSort($arr, 0, count($arr) - 1);
		echo'<pre>';
		var_dump($arr);
		echo'</pre>';
	}
}

main();