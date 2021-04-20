<?php
class Solution {

    /**
     * @param Integer[] $g  小朋友
     * @param Integer[] $s   饼干
     * @return Integer
     */
    function findContentChildren($g, $s) {
        sort($g);
        sort($s);
        $joffset = 0;
        $li = count($s);
        $lj = count($g);
        for($i=0;$i<$li;$i++){
            for($j=$joffset;$j<$lj;$j++){
                if($s[$i]>=$g[$j]){
                    $i++;
                    $joffset = $j+1;
                }
            }
        }
        return $joffset;
    }
    
}
