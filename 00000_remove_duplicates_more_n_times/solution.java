package com.company;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class Main {

    public static void main(String[] args) {
        int[] dat = {1,2,2,3,3,3,4,5,5};
        int[] res = solution(dat, 2);
        System.out.println(Arrays.toString(res));
    }

    public static int[] solution(int[] data, int n) {
        int[] res = new int[]{};

        Map<Integer, Integer> dbl = new HashMap<>();

        for (int datum : data) {
            if (!dbl.containsKey(datum)) {
                dbl.put(datum, 1);
            } else {
                Integer cnt = dbl.get(datum);
                dbl.put(datum, cnt + 1);
            }
        }

        for (Map.Entry<Integer, Integer> entry: dbl.entrySet()) {
            if (entry.getValue() <= n) {
                res = Arrays.copyOf(res, res.length + 1);
                res[res.length - 1] = entry.getKey();
            }
        }

        return res;
    }
}