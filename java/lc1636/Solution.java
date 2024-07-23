class Solution {
    public int[] frequencySort(int[] nums) {
        Map<Integer, Integer> fMap = new HashMap<>();
        for (int n : nums) {
            fMap.put(n, fMap.getOrDefault(n, 0) + 1);
        }
        Map<Integer, Set<Integer>> sorter = new TreeMap<>();
        fMap.forEach((key, value) -> {
            Set<Integer> sortSet = sorter.getOrDefault(value, new TreeSet<>(Comparator.reverseOrder()));
            sortSet.add(key);
            sorter.put(value, sortSet);
        });
        return sorter.values().stream().flatMap(value -> repeatMany(value, fMap)


).mapToInt(v -> v.intValue()).toArray();
    }

Stream<Integer> repeatMany(Set<Integer> values, Map<Integer, Integer> fMap){
    return values.stream().flatMap(val -> repeater(val, fMap.get(val)));
}


Stream<Integer> repeater(int val, int times){
    List<Integer> out= new ArrayList<>();
    for (int i = 0; i < times; i++) {
        out.add(val);
    }
    return out.stream();}}