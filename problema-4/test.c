// test.c
#include <stdio.h>

extern int globalVar;
static const int MAX_COUNT = 100;

enum State {
    IDLE,
    RUNNING,
    STOPPED
};

void calculate(int a, float b) {
    double result = (double)a * b;
    if (result > 0) {
        for (int i = 0; i < MAX_COUNT; i++) {
            if (i % 10 == 0) {
                continue;
            }
            // Some operation
        }
    } else {
        while (1) {
            // infinite loop, just for keyword
            break;
        }
    }
    return;
}

int main() {
    int x = 10;
    if (x > 5) {
        return 1;
    } else {
        char myChar = 'A';
        switch (myChar) {
            case 'A':
                printf("It's A\n");
                break;
            default:
                printf("It's something else\n");
        }
        do {
            x--;
        } while (x > 0);
        return 0;
    }
}
