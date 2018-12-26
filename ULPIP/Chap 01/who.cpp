#include <utmp.h>
#include <fcntl.h>
#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define SHOWHOST
void show_info(struct utmp * record) {
    printf("%-8.8s", record->ut_user);
    printf(" ");
    printf("%-8.8s", record->ut_line);
    printf(" ");
    printf("%12.12s", ctime((time_t*)&(record->ut_tv.tv_sec))+4);
    printf(" ");
    #ifdef SHOWHOST
    printf("( %s )", record->ut_host);
    #endif
    printf("\n");
    return;
}
int main() {
    struct utmp current_record;
    int utmpfd;
    int reclen = sizeof(current_record);
    if( ( utmpfd = open(UTMP_FILE, O_RDONLY) ) == -1 ) {
        perror( UTMP_FILE );
        exit(1);
    }
    while ( read( utmpfd, &current_record, reclen ) == reclen )
    {
        if(current_record.ut_type == 7)
            show_info(&current_record);
    }
    close(utmpfd);
    return 0;
}


