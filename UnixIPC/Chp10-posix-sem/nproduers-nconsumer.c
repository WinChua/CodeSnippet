#include <stdlib.h>
#include <stdio.h>
#include <pthread.h>
#include <unistd.h>
#include <semaphore.h>
#include <string.h>
#include <math.h>

#define NBUFF 10
#define MAXNTHREADS 100
#define min(a, b) ((a) < (b) ? (a) : (b))
struct {
    int buff[NBUFF];
    int nput;
    int nputval;
    sem_t mutex, nempty, nstored;
} shared;

void *produce(void *), *consume(void *);

int nitems, nproducers;
int main(int argc, char ** argv) {
    int i, count[MAXNTHREADS];
    pthread_t tid_produce[MAXNTHREADS], tid_consume;

    if(argc != 3) {
        printf("Usage %s <#item> <#producers>\n", argv[0]);
        exit(1);
    }
    nitems = atoi(argv[1]);
    nproducers = min(atoi(argv[2]), MAXNTHREADS);

    sem_init(&shared.mutex, 0, 1);
    sem_init(&shared.nempty, 0, NBUFF);
    sem_init(&shared.nstored, 0, 0);
    
    pthread_setconcurrency(nproducers + 1);
    for (i = 0; i < nproducers; i++) {
        count[i] = 0;
        pthread_create(&tid_produce[i], NULL, produce, &count[i]);
    }
    pthread_create(&tid_consume, NULL, consume, NULL);

    for(i = 0; i< nproducers; i++) {
        pthread_join(tid_produce[i], NULL);
        printf("count[%d] = %d\n", i, count[i]);
    }

    pthread_join(tid_consume, NULL);
    sem_destroy(&shared.mutex);
    sem_destroy(&shared.nempty);
    sem_destroy(&shared.nstored);
    exit(0);
}
#include "pc.h"
