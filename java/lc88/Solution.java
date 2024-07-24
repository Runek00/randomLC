class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int k = m-1;
        int j = n-1;
        int t;
        
        for(int i = n+m-1; i >= 0; i--){
            if(j < 0) {
                return;
            }
            if(k >= 0 && nums1[k] >= nums2[j]){

                nums1[i] = nums1[k];
                k--;





            } else {
                nums1[i] = nums2[j];
                j--;
            }
                
        }
            
    }
}