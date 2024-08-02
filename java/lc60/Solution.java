class Solution {
    public String getPermutation(int n, int k) {
        return backtrack(n, new HashSet<Integer>(), new Integer[]{k});
    }

    public String backtrack(int n, Set<Integer> used, Integer[] k) {
        if (used.size() >= n) {
            k[0]--;
            return "";
        }
        for(int i = 1; i <= n; i++) {
            if (!used.contains(i)){
                used.add(i);
                String out = backtrack(n, used, k);
                if (k[0] == 0) {
                    return ""+i+out;
                }
                used.remove(i);
                }
            }
        return "";
        }


}