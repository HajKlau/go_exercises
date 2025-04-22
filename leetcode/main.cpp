#include "palindrome_number.h"
#include<iostream>
#include<vector>

using namespace std;

int main()
{
    int myNumber = -121;
    bool result = isPalindrome(myNumber);

    if (result)
    {
        cout << "Number: " << myNumber << " is palindrome" << endl;
    } else {
        cout << "Number: " << myNumber << " is not palindrome" << endl;
    }
    return 0;
}