class Cube
{
public:
  int GetSide() const
  {
    return Side;
  }
  
  void SetSide(int num)
  {
    Side = num;  
  }
  
private:
  int Side = 0;
};
