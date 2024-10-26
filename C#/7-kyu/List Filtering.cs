using System.Collections;
using System.Collections.Generic;

public class ListFilterer
{
   public static IEnumerable<int> GetIntegersFromList(List<object> listOfItems)
   {
       List<int> integersOnly = new List<int>();

       foreach (var item in listOfItems)
       {
           if (item is int)
           {
               integersOnly.Add((int)item);
           }
       }

       return integersOnly;
   }
}
