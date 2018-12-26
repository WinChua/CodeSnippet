#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

static void sigHandler(int sig) {
    printf("Ouch %d!\n", sig);
}

int main(int argc, char * argv[]) {
    int j;
    if(signal(SIGINT, sigHandler) == SIG_ERR) {
        printf("signal handler install error");
        exit(1);
    }

    for(j = 0; ; ++j) {
        printf("%d\n", j);
        sleep(3);
    }
        
}
