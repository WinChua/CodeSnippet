#include <utmp.h>
#include <fcntl.h>
#include <stdio.h>

void show_user(struct utmp * user) {
    if(user->ut_type != USER_PROCESS) {
        return;
    }
    printf("%s\n", user->ut_name);
    return;
}

int main(int argc, char * argv[]) {

    int sfd = open(UTMP_FILE, O_RDONLY);
    if( sfd < 0 ) {
        perror(UTMP_FILE);
        return -1;
    }

    
    struct utmp user[16]; 
    int bufsize = sizeof(user);
    while(read(sfd, user, bufsize) > 0) {
        int i=0;
        for(i=0; i<16; ++i) {
            show_user(user + i);
        }
    }
}
