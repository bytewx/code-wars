using System.Linq;

public class Kata
{
    public static bool BetterThanAverage(int[] ClassPoints, int YourPoints)
    {
        double classAverage = (double)ClassPoints.Sum() / ClassPoints.Length;
        
        return YourPoints > classAverage;
    }
}
