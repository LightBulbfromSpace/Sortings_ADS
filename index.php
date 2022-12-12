<?php

require_once 'boot.php';

function main()
{
	$arr = [2, 1, 3, 9, 4, 5, 33, 12, 6, 55, 77, 88, 66, 100, 99];
	//heapSort($arr, count($arr));
	//mergeSort($arr, 0, count($arr));
	quickSort($arr, 0, count($arr) - 1);
	echo'<pre>';
	var_dump($arr);
	echo'</pre>';
}

main();