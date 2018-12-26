#include <unistd.h>
#include <stdio.h>

int main(int argc, char * argv[]) {
    pid_t pid, ppid, pgid, sid;
    pid = getpid();
    ppid = getppid();
    pgid = getpgid(0);
    sid = getsid(0);

    printf("the pid is %d\n", pid);
    printf("the ppid is %d\n", ppid);
    printf("the pgid is %d\n", pgid);
    printf("the sid is %d\n", sid);
    
    switch(fork()) {
        case 0:
	    pid = getpid();
	    ppid = getppid();
	    pgid = getpgid(0);
	    sid = getsid(0);

	    printf("the child pid is %d\n", pid);
	    printf("the child ppid is %d\n", ppid);
	    printf("the child pgid is %d\n", pgid);
	    printf("the child sid is %d\n", sid);
            break;
        default:
            break;
    }
    return 0;
}
