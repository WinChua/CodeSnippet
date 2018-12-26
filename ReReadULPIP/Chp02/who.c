#include <utmp.h>
#include <stdio.h>
#include <fcntl.h>
#include <time.h>

#define UTMP_FILE_T "/var/run/utmp"

void show_time(time_t second) {
    char *cp;
    cp = ctime(&second);
    printf("% 12.12s", cp+4);
}

void show_user(struct utmp * user) {
    if( user->ut_type == USER_PROCESS ) {
        printf("%-8.8s", user->ut_user);
        printf(" ");
        show_time(user->ut_tv.tv_sec);
        printf("\n");
    }
}

int main() {
    printf("The sizeof utmp is %d\n", sizeof(struct utmp));
    int fd = open(UTMP_FILE_T, O_RDONLY);
    if( -1 == fd ) {
        perror(UTMP_FILE_T);
    }
    struct utmp user;
    while((read(fd, &user, sizeof(struct utmp))) == sizeof(struct utmp)) {
        show_user(&user);
    }
    close(fd);
    return 0; 
}
