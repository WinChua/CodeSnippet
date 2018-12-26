#include <sys/stat.h>
#include <time.h>


void show_stat(char * filename) {
    struct stat buf;
    int result = stat(filename, &buf);
    if( -1 == result ) {
        perror(filename);
        return;
    }
    printf("file: %s\n\tlink: %d\n\tsize: %d\n\tmode: %o\n", filename, buf.st_nlink, buf.st_size, buf.st_mode);
    printf("\tmtime: %s", ctime(&buf.st_mtime));
    printf("\tatime: %s", ctime(&buf.st_atime));
    printf("\tctime: %s", ctime(&buf.st_ctime));
    show_filetype(buf.st_mode);
    return;
}

void show_filetype(mode_t mode) {
    mode_t type = mode & S_IFMT;
    mode_t types[] = {
        S_IFREG,
        S_IFDIR,
        S_IFBLK,
        S_IFCHR,
        S_IFIFO,
        S_IFLNK,
        S_IFSOCK
    };
    char * msgs[] = {
        "reg file",
        "dir file",
        "block file",
        "character special",
        "fifo file",
        "symbolic link",
        "socket file",
    };
    int i;
    for ( i = 0; i < 7; ++i ) {
        if( type == types[i] ) {
            printf("filetype:\t%s\n", msgs[i]);
            break;
        }
    }
    return;
}

int main(int argc, char * argv[]) {
    int i;
    for ( i = 1; i < argc; ++i ) {
        show_stat(argv[i]);
    }
}
