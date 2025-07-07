long pillars(int num_of_pillars, int distance, int width) {
    if (num_of_pillars <= 1)
        return 0;

    int gaps = num_of_pillars - 1;
    int inner_pillars = (num_of_pillars > 2) ? num_of_pillars - 2 : 0;

    return (gaps * distance * 100) + (inner_pillars * width);
}
