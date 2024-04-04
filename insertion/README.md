Insertion-Sort(A)
{
 for j=i to A.length
     key = A[i];
     // insert A[i] into sorted sequence A[1,2,3,..,i-1]
     j= i-1;
     while (j>0 and A[j]>key)
         A[j+1] = A[j]
         j= j-1
     A[j+1] = key
}
