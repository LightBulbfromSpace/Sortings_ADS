<?php

require 'heap_sort.php';

function main()
{
	$arr = [2, 1, 3, 9, 4, 5, 33, 12, 6, 55, 77, 88, 66, 100, 99];
	heapSort($arr, count($arr));
	echo'<pre>';
	var_dump($arr);
	echo'</pre>';
}

main();
