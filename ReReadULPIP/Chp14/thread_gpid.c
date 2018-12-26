#include <pthread.h>
#include <stdio.h>
#include <unistd.h>

int main(int argc, char * argv[]) {
    pthread_t t1, t2;
    void* show_pid(void *);
    printf("the pid of main thread is %d\n", getpid());
    pthread_create(&t1, NULL, show_pid, NULL); 
    pthread_create(&t2, NULL, show_pid, NULL); 
    pthread_join(t1, NULL);
    pthread_join(t2, NULL);
    return 0;
}

void * show_pid(void* f) {
    printf("the pid is %d\n", getpid());
}
