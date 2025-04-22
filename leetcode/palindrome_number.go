#include "palindrome_number.h"
#include<iostream>
#include<vector>

using namespace std;

bool isPalindrome(int number)
{
    if (number < 0)
    {
        return false;
    }

    vector<int> fromFrontNumber;
    vector<int> fromBackNumber;
    string numberToString = to_string(number);

    for (int i = 0; i < numberToString.size(); i++) 
    {
        fromFrontNumber.push_back(numberToString[i] - '0');
    }

    for(int j = numberToString.size() - 1; j >= 0; j--)
    {
        fromBackNumber.push_back(numberToString[j] - '0');
    }

    return fromFrontNumber == fromBackNumber;
}
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

package main

import (
    "fmt"  
)

func isPalindrome(number int) bool {
    
    if number < 0 {
        return false
    }

    numberToString := strconv.Itoa(number)

    fromFrontNumber := []int{}
    fromBackNumber := []int{}

    for _, ch := range numberToString {
        fromFrontNumber = append(fromFrontNumber, int(ch-'o'))
    }

    for i := len(numberToString) - 1; i >= 0; i-- {
        fromBackNumber = append(fromBackNumber, int(numberToString[i]-'o'))
    }
}