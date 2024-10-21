using System;
using System.Collections.Generic;
using System.Linq;

public class HumanTimeFormat {
  public static string formatDuration(int seconds) {        
    if (seconds == 0) {
      return "now";
    } else if (seconds == 1) {
      return "1 second";
    }

    int minutes = 0;
    int hours = 0;
    int days = 0;
    int years = 0;

    years += seconds / 31536000;
    seconds -= years * 31536000; 

    days += seconds / 86400;
    seconds -= days * 86400;

    hours += seconds / 3600;
    seconds -= hours * 3600;

    minutes += seconds / 60;
    seconds -= minutes * 60;

    var timeUnits = new List<string>();

    if (years == 1) {
      timeUnits.Add("1 year");
    } else if (years > 1) {
      timeUnits.Add(years + " years");
    }

    if (days == 1) {
      timeUnits.Add("1 day");
    } else if (days > 1) {
      timeUnits.Add(days + " days");
    }

    if (hours == 1) {
      timeUnits.Add("1 hour");
    } else if (hours > 1) {
      timeUnits.Add(hours + " hours");
    }

    if (minutes == 1) {
      timeUnits.Add("1 minute");
    } else if (minutes > 1) {
      timeUnits.Add(minutes + " minutes");
    }

    if (seconds == 1) {
      timeUnits.Add("1 second");
    } else if (seconds > 1) {
      timeUnits.Add(seconds + " seconds");
    }

    if (timeUnits.Count == 1) {
      return timeUnits[0];
    } else if (timeUnits.Count == 2) {
      return timeUnits[0] + " and " + timeUnits[1];
    } else {
      return string.Join(", ", timeUnits.Take(timeUnits.Count - 1)) + " and " + timeUnits.Last();
    }
  }
}
