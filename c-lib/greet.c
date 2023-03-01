#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char* greetWithName(char* name) {
    char* greeting = malloc(strlen(name) + 8); // allocate memory for the greeting message
    sprintf(greeting, "Hello, %s", name); // format the greeting message
    return greeting;
}

char* greet(void) {
    return "hello from c";
}

void hello(void) {
    printf("hello from printf\n");
    puts("hello from puts");
}

int add(int a, int b) {
    return a+b;
}


// int main() {
//     hello();
//     return 0;
// }