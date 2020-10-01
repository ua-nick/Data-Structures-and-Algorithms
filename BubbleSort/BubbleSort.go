package BubbleSort

func BubbleSort(array []int) {
	//pass is a variable which starts from the end of the array
	//i variable starts from starting of the array
	for(int pass=n-1;pass>=0;pass--)
	{
		for(int i=0;i<=pass;i++)
		{
			if(array[i]>array[i+1])
			swap (array[i], array[i+1]);
		}
	}
}
