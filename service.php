<?php

function swap(&$a, &$b): void
{
	if ($a === $b) {
		return;
	}

	$tmp = $a;
	$a = $b;
	$b = $tmp;
}