#include <unistd.h>
#include <fcntl.h>
#include <time.h>

int main() {
    int fd = open("hello", O_RDONLY);
    struct flock lock;
    lock.l_type = F_RDLCK;
    lock.l_start = 0;
    lock.l_whence = SEEK_SET;
    lock.l_len = 1;
    fcntl(fd, F_SETLK, &lock);
    sleep(1);
    close(fd);
    return 0;
}
