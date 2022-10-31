# Remove duplicated number from sorted array

## Question: 

- Input will be a sorted array which contains integer numbers. Some number will disappear more than once. For example:`[1, 2, 3, 3, 3]`
- Try to remove the duplicated numbers to make every number show up only once. For example, remove last two `3`, then the correct array will be `[1,2,3]`

## Thinking

- Input array is sorted, it's means the duplicated numbers will stay together one by one.

    `[1,2,2,3,4]`  OR `[3,2,2,1]`
- If number[i] equal number[i+1], we should remove one of them.
- Their will be three kinds of input we should handle carefully. 
  1. Duplicated numbers placed at beginning of array
     `[1,1,1,2,3,4,5]`
  2. Duplicated numbers placed at the middle of array
     `[1,2,3,3,3,4,5]`
  3. Duplicated numbers placed at the end of array 
     `[1,2,3,4,4,4,4]`
- If number `n` equal to number `n+1`, 
