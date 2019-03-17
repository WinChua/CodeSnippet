#ifndef H_PC
#define H_PC
void * produce(void *arg) {
    for(;;) {
        sem_wait(&shared.nempty);
        sem_wait(&shared.mutex);
        if(shared.nput >= nitems) {
            sem_post(&shared.nempty);
            sem_post(&shared.mutex);
            return(NULL);
        }
        shared.buff[shared.nput % NBUFF] = shared.nputval;
        shared.nput++;
        shared.nputval++;
        sem_post(&shared.mutex);
        sem_post(&shared.nstored);
        *((int*) arg) += 1;
    }
}

void *consume(void *arg) {
    int i;
    for(i = 0; i< nitems; i++) {
        sem_wait(&shared.nstored);
        sem_wait(&shared.mutex);

        if(shared.buff[i % NBUFF] != i) {
            printf("error: buff[%d] = %d\n", i, shared.buff[i % NBUFF]);
        }
        sem_post(&shared.mutex);
        sem_post(&shared.nempty);
    }
    return (NULL);
}
#endif
