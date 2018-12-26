#include <time.h>
#include <sys/time.h>
#include <stdio.h>

int main(int argc, char * argv[]) {
    time_t now;
    now = time(NULL);
    printf("the time is %d, the ctime is %s", now, ctime(&now));
    struct tm *locp = gmtime(&now);
    struct tm loc = *locp;
    printf("the asctime of gmtime is %s", asctime(locp));
    locp = localtime(&now);
    printf("the asctime of localtime is %s", asctime(locp));
    char outstr[26];
    strftime(outstr, 25, "%Y%m%d", locp);
    printf("Ymd is %s\n", outstr);
    return 0;
}
