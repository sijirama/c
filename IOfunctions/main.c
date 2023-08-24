#include <conio.h>
#include <stdio.h>

int main() {
    // cgets, cprintf, cputs have been depracated and replaced in g++, thats the
    // compiler you sue siji. creat has been replacced with open too.

    // char cgets(char *str)
    // cgets, used to read string from stdin/the keyboard
    // char arr[30];
    // arr[0] = 20; // must set the first byte to the max length of strings to
    // save cputs("Enter a string: "); cgets(arr); // Documents and Settings

    //----------------------------------------------------------------
    // void clearerr(FILE * stream)

    //----------------------------------------------------------------
    // void close(FILE * stream)
    // used to close the file stram from further processes.

    //----------------------------------------------------------------
    // int cprintf(cont=st char * str)

    fprintf(stdout, "Hello world");
}
