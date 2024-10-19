public class Kata 
{
    public static int FindSmallestInt(int[] args) 
    {
        int smallest = args[0];

        foreach (int num in args) 
        {
            if (num < smallest) 
            {
                smallest = num;
            }
        }

        return smallest;
    }
}
