<?php

require_once ROOT . '/lab_10_merge_sort/merge_sort.php';

function BucketSort(&$arr)
{
	$maxValue = max($arr);
	$mainLen = count($arr);

	// difference between first and last element value in bucket
	$bucketRange = ceil($maxValue / $mainLen);

	$buckets = [];
	foreach ($arr as $item)
	{
		$indexOfBucket = $item / $bucketRange;
		if (!isset($buckets[$indexOfBucket]))
		{
			$buckets[$indexOfBucket] = [];
		}
		$buckets[$indexOfBucket][] = $item;
	}
	$bucketsLen = count($buckets);
	for($i = 0; $i < $bucketsLen; $i++)
	{
		if(!isset($buckets[$i]))
		{
			continue;
		}
		mergeSortIterative($buckets[$i]);
	}
	ksort($buckets);
	$i = 0;
	foreach ($buckets as $bucket)
	{
		foreach ($bucket as $item)
		{
			$arr[$i] = $item;
			$i++;
		}
	}
}