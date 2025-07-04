#include <vector>
#include <string>

std::vector<bool> flick_switch(const std::vector<std::string>& arr)
{
    std::vector<bool> result;
    bool toggle = true;

    for (const auto& word : arr)
    {
        if (word == "flick")
        {
            toggle = !toggle;
        }

        result.push_back(toggle);
    }

    return result;
}
